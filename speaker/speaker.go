package speaker

import (
	"fmt"
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

type Speaker interface {
	Say(text string, gender Gender) error
}

type speaker struct {
	apiKey string
}

func New(apiKey string) Speaker {
	return &speaker{
		apiKey: apiKey,
	}
}

func (t *speaker) Say(text string, gender Gender) error {
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

func (t *speaker) chooseVoice(gender Gender) string {
	rand.Seed(time.Now().Unix())

	switch gender {
	case GenderFemale:
		i := rand.Int() % len(femaleVoices)
		return femaleVoices[i]
	default:
		i := rand.Int() % len(maleVoices)
		return maleVoices[i]
	}
}

func (t *speaker) downloadFromURL(url string, fileName string) error {
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
