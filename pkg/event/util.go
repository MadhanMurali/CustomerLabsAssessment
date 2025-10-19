package event

import (
	"encoding/json"
	"strings"
)

type EventAttribute struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type EventTrait struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Event struct {
	Event           string                    `json:"event"`
	EventType       string                    `json:"event_type"`
	AppId           string                    `json:"app_id"`
	UserId          string                    `json:"user_id"`
	MessageId       string                    `json:"message_id"`
	PageTitle       string                    `json:"page_title"`
	PageURL         string                    `json:"page_url"`
	BrowserLanguage string                    `json:"browser_language"`
	ScreenSize      string                    `json:"screen_size"`
	Attributes      map[string]EventAttribute `json:"attributes"`
	Traits          map[string]EventTrait     `json:"traits"`
}

func (evt *Event) UnmarshalMinifiedJSON(data []byte) error {
	var minifiedEvent map[string]string

	if err := json.Unmarshal(data, &minifiedEvent); err != nil {
		return err
	}

	evt.Event = minifiedEvent["ev"]
	evt.EventType = minifiedEvent["et"]
	evt.AppId = minifiedEvent["id"]
	evt.UserId = minifiedEvent["uid"]
	evt.MessageId = minifiedEvent["mid"]
	evt.PageTitle = minifiedEvent["t"]
	evt.PageURL = minifiedEvent["p"]
	evt.BrowserLanguage = minifiedEvent["l"]
	evt.ScreenSize = minifiedEvent["sc"]

	for key, value := range minifiedEvent {
		if suffix, isFound := strings.CutPrefix(key, "atrk"); isFound {
			if evt.Attributes == nil {
				evt.Attributes = map[string]EventAttribute{}
			}
			evt.Attributes[value] = EventAttribute{
				Type:  minifiedEvent["atrt"+suffix],
				Value: minifiedEvent["atrv"+suffix],
			}
		} else if suffix, isFound := strings.CutPrefix(key, "uatrk"); isFound {
			if evt.Traits == nil {
				evt.Traits = map[string]EventTrait{}
			}
			evt.Traits[value] = EventTrait{
				Type:  minifiedEvent["uatrt"+suffix],
				Value: minifiedEvent["uatrv"+suffix],
			}
		}
	}

	return nil
}
