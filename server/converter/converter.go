package converter

import (
	"fmt"
	"github.com/ngalayko/theq_speaker/server/logger"
	"github.com/ngalayko/theq_speaker/server/types"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	apiURL = "https://tts.voicetech.yandex.net/generate"

	fileNameLength = 5
	fileFormat     = "mp3"
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
	queueToSend    chan types.Message
}

func New(apiKey string, queueToConvert chan types.Text, queueToSend chan types.Message) Converter {
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
	}()

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

	randomString := RandStringBytesMaskImprSrc(fileNameLength)
	fileName := fmt.Sprintf("%v.%v", randomString, fileFormat)
	if err := t.downloadFromURL(requestURL, fileName); err != nil {
		return err
	}

	t.queueToSend <- types.NewMessage(fileName)

	t.logger.Info("Text converted", logger.Fields{
		"text":     text,
		"fileName": fileName,
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

func (t *converter) downloadFromURL(url string, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if _, err := io.Copy(file, response.Body); err != nil {
		return err
	}

	return nil
}
