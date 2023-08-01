package dto

type Contact struct {
	ContactId      string         `json:"contactId"`
	Email          string         `json:"email"`
	Name           string         `json:"name"`
	Ip             string         `json:"ip"`
	Origin         string         `json:"origin"`
	Href           string         `json:"href"`
	Campaign       Campaign       `json:"campaign,omitempty"`
	SourceCampaign SourceCampaign `json:"sourceCampaign,omitempty"`
	PhoneNumber    CustomField    `json:"phoneNumber,omitempty"`
}

type CustomField struct {
	FieldId string `json:"customFieldId"`
	Href    string `json:"href"`
}
