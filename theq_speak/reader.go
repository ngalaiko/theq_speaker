package theq_speak

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"time"
)

const (
	questionsLimit = 5
	timeout        = 5 * time.Second
	apiUrl         = "https://api.thequestion.ru/api"
)

func (t *theqSpeak) questionsLoop() {
	for {
		if err := t.fetchQuestions(questionsLimit); err != nil {
			panic(err)
		}

		time.Sleep(timeout)
	}
}

func (t *theqSpeak) fetchQuestions(limit int32) error {
	questions, err := t.getQuestions(limit)
	if err != nil {
		return err
	}

	sort.Slice(questions, func(i, j int) bool {
		return questions[i].Id > questions[j].Id
	})

	for _, question := range questions {
		if t.seen[question.Id] {
			continue
		}

		t.seen[question.Id] = true

		t.queue <- question
	}

	return nil
}

func (t *theqSpeak) getQuestions(limit int32) ([]*Question, error) {
	requestUrl := apiUrl +
		"/questions/query" +
		"?lang=%s" +
		"&sort=%s" +
		"&limit=%d"

	response := []*Question{}
	if err := t.httpGet(
		fmt.Sprintf(requestUrl, "ru", "date", limit),
		&response,
	); err != nil {
		return nil, err
	}

	return response, nil
}

func (t *theqSpeak) httpGet(url string, responsePointer interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, responsePointer); err != nil {
		return err
	}

	return nil
}
