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
	apiURL         = "https://api.thequestion.ru/api"
)

func (t *theqSpeak) questionsLoop() {
	defer func() {
		recover()
	}()

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
		return questions[i].ID > questions[j].ID
	})

	for _, question := range questions {
		if t.seen[question.ID] {
			continue
		}

		t.seen[question.ID] = true

		t.queue <- question
	}

	return nil
}

func (t *theqSpeak) getQuestions(limit int32) ([]*Question, error) {
	requestURL := apiURL +
		"/questions/query" +
		"?lang=%s" +
		"&sort=%s" +
		"&limit=%d"

	response := []*Question{}
	if err := t.httpGet(
		fmt.Sprintf(requestURL, "ru", "date", limit),
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
