package mappers

import "time"

type EventRequest struct {
	Name           string    `json:"name"`
	StartTime      time.Time `json:"startTime"`
	EndTime        time.Time `json:"EndTime"`
	Timezone       string    `json:"timezone"`
	Details        string    `json:"details"`
	EntryType      string    `json:"entryType"`
	EventType      string    `json:"eventType"`
	Cost           float32   `json:"cost"`
	AimedAt        string    `json:"aimedAt"`
	IsPublic       bool      `json:"isPublic"`
	ThumbnailImage []byte    `json:"thumbnailImage"`
	Images         [][]byte  `json:"images"`
	LocationID     string    `json:"locationId"`
	CategoryID     string    `json:"categoryID"`
}

type EventResponse struct {
	Name              string    `json:"name"`
	StartTime         time.Time `json:"startTime"`
	EndTime           time.Time `json:"EndTime"`
	Timezone          string    `json:"timezone"`
	Details           string    `json:"details"`
	EntryType         string    `json:"entryType"`
	EventType         string    `json:"eventType"`
	Cost              float32   `json:"cost"`
	AimedAt           string    `json:"aimedAt"`
	IsPublic          bool      `json:"isPublic"`
	ThumbnailImageURL string    `json:"thumbnailImage"`
	LocationID        string    `json:"locationId"`
	CategoryID        string    `json:"categoryID"`
}
