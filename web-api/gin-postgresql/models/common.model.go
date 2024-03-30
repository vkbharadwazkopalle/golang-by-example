package models

type IdName struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty" validate:"required"`
}
