package typings

import (
	"encoding/json"
	"time"
)

type Conversation struct {
	RequestData     *RequestData `json:"request_data"`
	LastInteraction int64        `json:"last_interaction"`
}

func (c Conversation) New() *Conversation {
	return &Conversation{
		RequestData:     RequestData{}.New(),
		LastInteraction: time.Now().Unix(),
	}
}

type RequestData struct {
	Transcript     []TranscriptEntry `json:"transcript"`
	Temperature    float64           `json:"temperature"`
	MaxTokens      int               `json:"max_tokens"`
	N              int               `json:"n"`
	HrMode         bool              `json:"hr_mode"`
	HrOverlapRatio float64           `json:"hr_overlap_ratio"`
	TopP           float64           `json:"top_p"`
}

func (t RequestData) New() *RequestData {
	return &RequestData{
		Transcript:     make([]TranscriptEntry, 0),
		Temperature:    0.5,
		MaxTokens:      512,
		N:              1,
		HrMode:         false,
		HrOverlapRatio: 0,
		TopP:           0.95,
	}
}

func (t *RequestData) AddEntry(entry TranscriptEntry) error {
	t.Transcript = append(t.Transcript, entry)
	return nil
}

func (t *RequestData) AddEntryWithText(text string) error {
	t.AddEntry(TranscriptEntry{
		Type: "text",
		Data: text,
	})
	return nil
}

func (t *RequestData) AddEntryWithImage(image string) error {
	t.AddEntry(TranscriptEntry{
		Type: "image",
		Data: image,
	})
	return nil
}

func (t *RequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(*t)
}

type TranscriptEntry struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func (t TranscriptEntry) New(transcript_type string, data string) TranscriptEntry {
	return TranscriptEntry{
		Type: transcript_type,
		Data: data,
	}
}
