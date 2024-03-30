package models

type Interview struct {
	Interview_Id int    `json:"interview_id,omitempty"`
	Name         string `json:"name,omitempty" validate:"required"`
	Active       int    `json:"active,omitempty"`
	Created_On   string `json:"created_on,omitempty"`
	// Role    IdName `json:"role,omitempty"`
	// Project IdName `json:"project,omitempty"`
}

func NewInterview(interview_id int, name string, active int, created_on string) Interview {
	return Interview{interview_id, name, active, created_on}
}
