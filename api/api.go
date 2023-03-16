package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/acheong08/gpt4/typings"
)

var (
	key         = os.Getenv("OPENAI_KEY")
	org         = os.Getenv("OPENAI_ORG")
	endpoint    = os.Getenv("OPENAI_ENDPOINT")
	http_client = http.DefaultClient
)

func newRequest(body io.Reader) *http.Request {
	base_request, _ := http.NewRequest("POST", endpoint, body)
	base_request.Header.Add("Authorization", "Bearer "+key)
	base_request.Header.Add("Content-Type", "application/json")
	base_request.Header.Add("OpenAI-Organization", org)
	base_request.Header.Add("Accept", "application/json")
	base_request.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36")
	return base_request
}

func Send(transcript typings.RequestData) (TextCompletion, error) {
	transcript_json, err := transcript.MarshalJSON()
	if err != nil {
		return TextCompletion{}, err
	}
	request := newRequest(bytes.NewReader(transcript_json))
	response, err := http_client.Do(request)
	if err != nil {
		return TextCompletion{}, err
	}
	defer response.Body.Close()
	// Map response to struct
	var completions TextCompletion
	err = json.NewDecoder(response.Body).Decode(&completions)
	if err != nil {
		return TextCompletion{}, err
	}
	completions.ID = ""
	return completions, nil
}
