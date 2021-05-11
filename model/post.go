package model

type Post struct {
	ID         uint   `json:"id"`
	TagID      int    `json:"tag_id"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedOn  int    `json:"created_on"`
	CreatedBy  string `json:"created_by"`
	ModifiedOn int    `json:"modified_on"`
	ModifiedBy string `json:"modified_by"`
	DeletedOn  int    `json:"deleted_on"`
	State      int    `json:"state"`
}
