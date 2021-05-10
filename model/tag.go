package model

type Tag struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	CreatedOn  int    `json:"created_on"`
	CreatedBy  string `json:"created_by"`
	ModifiedOn int    `json:"modified_on"`
	ModifiedBy string `json:"modified_by"`
	DeletedOn  int    `json:"deleted_on"`
	State      int    `json:"state"`
}
