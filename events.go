package aircall

import (
	"encoding/json"
	"fmt"
	"strings"
)

type InboundWebhook struct {
	Resource  string          `json:"resource,omitempty"`
	Event     string          `json:"event,omitempty"`
	EventName string          `json:"event_name,omitempty"`
	Timestamp int             `json:"timestamp,omitempty"`
	Token     string          `json:"token,omitempty"`
	Data      json.RawMessage `json:"data,omitempty"`
}

func (w *InboundWebhook) GetCallData() (*Call, error) {
	if w.Resource != "call" {
		return nil, fmt.Errorf("Resource is %s not a call.", w.Resource)
	}

	call := Call{}

	if err := json.Unmarshal(w.Data, &call); err != nil {
		return nil, err
	}

	return &call, nil
}

func (w *InboundWebhook) GetUserData() (*User, error) {
	if w.Resource != "user" {
		return nil, fmt.Errorf("Resource is %s not a user.", w.Resource)
	}

	user := User{}

	if err := json.Unmarshal(w.Data, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (w *InboundWebhook) GetNumberData() (*Number, error) {
	if w.Resource != "number" {
		return nil, fmt.Errorf("Resource is %s not a number.", w.Resource)
	}

	number := Number{}

	if err := json.Unmarshal(w.Data, &number); err != nil {
		return nil, err
	}

	return &number, nil
}

func (w *InboundWebhook) GetContactData() (*Contact, error) {
	if w.Resource != "contact" {
		return nil, fmt.Errorf("Resource is %s not a contact.", w.Resource)
	}

	contact := Contact{}

	if err := json.Unmarshal(w.Data, &contact); err != nil {
		return nil, err
	}

	return &contact, nil
}

func (w *InboundWebhook) GetMessageData() (*Message, error) {
	if w.Resource != "message" {
		return nil, fmt.Errorf("Resource is %s not a message.", w.Resource)
	}

	message := Message{}

	if err := json.Unmarshal(w.Data, &message); err != nil {
		return nil, err
	}

	return &message, nil
}

func (w *InboundWebhook) GetIntegrationData() (*Integration, error) {
	if w.Resource != "integration" {
		return nil, fmt.Errorf("Resource is %s not a integration.", w.Resource)
	}

	integration := Integration{}

	if err := json.Unmarshal(w.Data, &integration); err != nil {
		return nil, err
	}

	return &integration, nil
}

func (w *InboundWebhook) GetConversationIntelligenceSentimentData() (*ConversationIntelligenceSentiment, error) {
	if w.Resource != "conversation_intelligence" {
		return nil, fmt.Errorf("Resource is %s not a conversation intelligence.", w.Resource)
	}

	if eventParts := strings.Split(w.Event, "."); len(eventParts) > 0 && eventParts[0] != "sentiment" {
		return nil, fmt.Errorf("Resource is %s, but event is %s not a supported.", w.Resource, w.Event)
	}

	sentiment := ConversationIntelligenceSentiment{}

	if err := json.Unmarshal(w.Data, &sentiment); err != nil {
		return nil, err
	}

	return &sentiment, nil
}

func (w *InboundWebhook) GetConversationIntelligenceTranscriptionData() (*ConversationIntelligenceTranscription, error) {
	if w.Resource != "conversation_intelligence" {
		return nil, fmt.Errorf("Resource is %s not a conversation intelligence.", w.Resource)
	}

	if eventParts := strings.Split(w.Event, "."); len(eventParts) > 0 && eventParts[0] != "transcription" {
		return nil, fmt.Errorf("Resource is %s, but event is %s not a supported.", w.Resource, w.Event)
	}

	transcription := ConversationIntelligenceTranscription{}

	if err := json.Unmarshal(w.Data, &transcription); err != nil {
		return nil, err
	}

	return &transcription, nil
}
