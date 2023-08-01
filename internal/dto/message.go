package dto

type Message struct {
	ResourceType string `json:"resourceType"`
	ResourceId   string `json:"resourceId"`
	Name         string `json:"name"`
	Subject      string `json:"subject"`
	Href         string `json:"href"`
}
