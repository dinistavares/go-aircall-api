package aircall

import "fmt"

// Conversation Intelligence service
type ConversationIntelligenceService service

type ConversationIntelligenceTranscriptionResponse struct {
	Transcription *ConversationIntelligenceTranscription `json:"transcription,omitempty"`
}

type ConversationIntelligenceSentimentResponse struct {
	Sentiment *ConversationIntelligenceSentiment `json:"sentiment,omitempty"`
}

type ConversationIntelligenceTopicResponse struct {
	Topic *ConversationIntelligenceTopic `json:"topic,omitempty"`
}

type ConversationIntelligenceSummaryResponse struct {
	Summary *ConversationIntelligenceSummary `json:"summary,omitempty"`
}

type ConversationIntelligenceTranscription struct {
	ID            int                                           `json:"id,omitempty"`
	CallID        int                                           `json:"call_id,omitempty"`
	CallCreatedAt string                                        `json:"call_created_at,omitempty"`
	Type          string                                        `json:"type,omitempty"`
	Content       *ConversationIntelligenceTranscriptionContent `json:"content,omitempty"`
}

type ConversationIntelligenceTranscriptionContent struct {
	Language   string                                             `json:"language,omitempty"`
	Utterances *[]ConversationIntelligenceTranscriptionUtterances `json:"utterances,omitempty"`
}

type ConversationIntelligenceTranscriptionUtterances struct {
	StartTime       float64 `json:"start_time,omitempty"`
	EndTime         float64 `json:"end_time,omitempty"`
	Text            string  `json:"text,omitempty"`
	ParticipantType string  `json:"participant_type,omitempty"`
	PhoneNumber     string  `json:"phone_number,omitempty"`
	UserID          int     `json:"user_id,omitempty"`
}

type ConversationIntelligenceSentiment struct {
	ID           int                                     `json:"id,omitempty"`
	CallID       int                                     `json:"call_id,omitempty"`
	Participants *[]ConversationIntelligenceParticipants `json:"participants,omitempty"`
}

type ConversationIntelligenceParticipants struct {
	PhoneNumber     string `json:"phone_number,omitempty"`
	Value           string `json:"value,omitempty"`
	Type            string `json:"type,omitempty"`
	UserID          string `json:"user_id,omitempty"`
	ParticipantType string `json:"participant_type,omitempty"`
}

type ConversationIntelligenceTopic struct {
	ID        int      `json:"id,omitempty"`
	CallID    int      `json:"call_id,omitempty"`
	CreatedAt string   `json:"created_at,omitempty"`
	Content   []string `json:"content,omitempty"`
}

type ConversationIntelligenceSummary struct {
	ID        int    `json:"id,omitempty"`
	CallID    int    `json:"call_id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	Content   string `json:"content,omitempty"`
}

//  ***********************************************************************************
//  GET TRANSCRIPTION
//  https://developer.aircall.io/api-references/#retrieve-a-transcription
//  ***********************************************************************************

// Get a call transcription. Reference: https://developer.aircall.io/api-references/#retrieve-a-transcription
func (service *ConversationIntelligenceService) GetTranscription(callID int) (*ConversationIntelligenceTranscriptionResponse, *Response, error) {
	_url := fmt.Sprintf("calls/%d/transcription", callID)

	responseBody := new(ConversationIntelligenceTranscriptionResponse)
	response, err := service.client.Get(_url, nil, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  GET SENTIMENTS
//  https://developer.aircall.io/api-references/#retrieve-sentiments
//  ***********************************************************************************

// Get a call sentiment. Reference: https://developer.aircall.io/api-references/#retrieve-sentiments
func (service *ConversationIntelligenceService) GetSentiment(callID int) (*ConversationIntelligenceSentimentResponse, *Response, error) {
	_url := fmt.Sprintf("calls/%d/sentiments", callID)

	responseBody := new(ConversationIntelligenceSentimentResponse)
	response, err := service.client.Get(_url, nil, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  GET TOPICS
//  https://developer.aircall.io/api-references/#retrieve-topics
//  ***********************************************************************************

// Get a call topics. Reference: https://developer.aircall.io/api-references/#retrieve-topics
func (service *ConversationIntelligenceService) GetTopics(callID int) (*ConversationIntelligenceTopicResponse, *Response, error) {
	_url := fmt.Sprintf("calls/%d/topics", callID)

	responseBody := new(ConversationIntelligenceTopicResponse)
	response, err := service.client.Get(_url, nil, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}

//  ***********************************************************************************
//  GET SUMMARY
//  https://developer.aircall.io/api-references/#retrieve-a-summary
//  ***********************************************************************************

// Get a call summary. Reference: https://developer.aircall.io/api-references/#retrieve-a-summary
func (service *ConversationIntelligenceService) GetSummary(callID int) (*ConversationIntelligenceSummaryResponse, *Response, error) {
	_url := fmt.Sprintf("calls/%d/summary", callID)

	responseBody := new(ConversationIntelligenceSummaryResponse)
	response, err := service.client.Get(_url, nil, responseBody)

	if err != nil {
		return nil, response, err
	}

	return responseBody, response, nil
}
