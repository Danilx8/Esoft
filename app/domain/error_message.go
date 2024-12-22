package domain

type ErrorMessage struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}
