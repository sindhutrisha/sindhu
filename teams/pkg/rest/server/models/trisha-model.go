package models

type Trisha struct {
	Id int64 `json:"id,omitempty"`

	Fields string `json:"fields,omitempty"`

	Verified bool `json:"verified,omitempty"`
}
