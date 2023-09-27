package entity

type User struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	IsDeleted   string `json:"isDeleted"`
	CreatedDate string `json:"createdDate"`
	UpdatedDate string `json:"updatedDate"`
}
