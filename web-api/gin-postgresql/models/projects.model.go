package models

type Project struct {
	Project_Id int    `json:"project_id,omitempty"`
	Name       string `json:"name,omitempty" validate:"required"`
	// Client IdName `json:"client,omitempty"`
	Active     int    `json:"active,omitempty"`
	Created_On string `json:"created_on,omitempty"`
}

func NewProject(project_id int, name string, active int, created_on string) Project {
	return Project{project_id, name, active, created_on}
}
