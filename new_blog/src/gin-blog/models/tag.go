package models

type Tag struct{
	Model
	Name string `json:"Name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}
