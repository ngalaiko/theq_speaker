package reader

import (
	"fmt"
	"github.com/ngalayko/theq_speaker/types"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/exec"
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

type Reader interface {
	ReadLoop()
}

type reader struct {
	apiKey string
	queue  chan types.Text
}

func New(apiKey string, queue chan types.Text) Reader {
	return &reader{
		apiKey: apiKey,
		queue:  queue,
	}
}

func (t *reader) ReadLoop() {
	defer func() {
		recover()
	}()

	for question := range t.queue {
		if err := t.read(question.String(), question.Gender().String()); err != nil {
			panic(err)
		}
	}
}

func (t *reader) read(text string, gender string) error {
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
	defer os.Remove(fileName)

	if _, err := exec.Command("mpg123", "-q", fileName).CombinedOutput(); err != nil {
		return err
	}

	return nil
}

func (t *reader) chooseVoice(gender string) string {
	rand.Seed(time.Now().Unix())

	switch gender {
	case "female":
		i := rand.Int() % len(femaleVoices)
		return femaleVoices[i]
	default:
		i := rand.Int() % len(maleVoices)
		return maleVoices[i]
	}
}

func (t *reader) downloadFromURL(url string, fileName string) error {
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
