package fetcher

import (
	"encoding/json"
	"fmt"
	"github.com/ngalayko/theq_speaker/types"
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

type Fetcher interface {
	FetchLoop()
}

type fetcher struct {
	queue chan types.Text

	seen map[int64]bool
}

func New(queue chan types.Text) Fetcher {
	return &fetcher{
		queue: queue,

		seen: map[int64]bool{},
	}
}

func (t *fetcher) FetchLoop() {
	defer func() {
		recover()
	}()

	for {
		if err := t.fetchNext(questionsLimit); err != nil {
			panic(err)
		}

		time.Sleep(timeout)
	}
}

func (t *fetcher) fetchNext(limit int32) error {
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

func (t *fetcher) getQuestions(limit int32) ([]*types.Question, error) {
	requestURL := apiURL +
		"/questions/query" +
		"?lang=ru" +
		"&sort=date" +
		"&limit=%d"

	response := []*types.Question{}
	if err := t.httpGet(fmt.Sprintf(requestURL, limit), &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (t *fetcher) httpGet(url string, responsePointer interface{}) error {
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
