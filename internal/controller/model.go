package controller

type EmailRequest struct {
	Sender      string               `json:"sender"`
	Receivers   []string             `json:"receivers"`
	Subject     string               `json:"subject"`
	Body        string               `json:"body"`
	Attachments []*RequestAttachment `json:"attachments"`
}

type RequestAttachment struct {
	FileName string
	FileType string
	Data     []byte
}
