package converter

import (
	"encoding/base64"
	"github.com/ngalayko/theq_speaker/server/logger"
	"github.com/ngalayko/theq_speaker/server/types"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

const (
	apiURL = "https://tts.voicetech.yandex.net/generate"

	fileFormat = "wav"
)

var (
	maleVoices   = []string{"zahar", "ermil"}
	femaleVoices = []string{"jane", "oksana", "alyss", "omazh"}
)

type Converter interface {
	ConvertLoop()
}

type converter struct {
	logger logger.Logger

	apiKey string

	queueToConvert chan types.Text
	queueToSend    chan *types.Message
}

func New(apiKey string, queueToConvert chan types.Text, queueToSend chan *types.Message) Converter {
	return &converter{
		logger: logger.New("converter"),

		apiKey: apiKey,

		queueToConvert: queueToConvert,
		queueToSend:    queueToSend,
	}
}

func (t *converter) ConvertLoop() {
	defer func() {
		recover()

		t.logger.Info("ConvertLoop recovered", logger.Fields{})
	}()

	t.logger.Info("ConvertLoop started", logger.Fields{})

	for question := range t.queueToConvert {
		if err := t.convert(question.String(), question.Gender()); err != nil {
			t.logger.Error("ConvertLoop error", logger.Fields{
				"error": err,
			})

			panic(err)
		}
	}
}

func (t *converter) convert(text string, gender types.Gender) error {
	requestURL := apiURL +
		"?key=" + url.QueryEscape(t.apiKey) +
		"&text=" + url.QueryEscape(text) +
		"&format=" + fileFormat +
		"&quality=hi" +
		"&lang=ru" +
		"&speaker=" + t.chooseVoice(gender)

	base64String, err := t.downloadFromURL(requestURL)
	if err != nil {
		return err
	}

	t.queueToSend <- &types.Message{
		Base64: base64String,
		Text:   text,
	}

	t.logger.Info("Text converted", logger.Fields{
		"text": text,
	})

	return nil
}

func (t *converter) chooseVoice(gender types.Gender) string {
	rand.Seed(time.Now().Unix())

	switch gender.Clean() {
	case types.GenderFemale:
		i := rand.Int() % len(femaleVoices)
		return femaleVoices[i]
	default:
		i := rand.Int() % len(maleVoices)
		return maleVoices[i]
	}
}

func (t *converter) downloadFromURL(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", nil
	}

	return base64.StdEncoding.EncodeToString(bytes), nil
}
