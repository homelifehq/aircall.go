package message

type Resource struct {
	Welcome       string `json:"welcome"`
	Waiting       string `json:"waiting"`
	IVR           string `json:"ivr"`
	Voicemail     string `json:"voicemail"`
	Closed        string `json:"closed"`
	CallbackLater string `json:"callback_later"`
}
