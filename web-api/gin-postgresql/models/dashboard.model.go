package models

type Dashboard struct {
	Id   string   `json:"id,omitempty"`
	Name string   `json:"name,omitempty" validate:"required"`
	Data Snapshot `json:"data,omitempty"`
}

type Snapshot struct {
	Complete int `json:"complete,omitempty"`
	Ongoing  int `json:"ongoing,omitempty"`
}
