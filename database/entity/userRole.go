package entity

type UserRole struct {
	Id       *int64 `gorm:"primaryKey;column:ID;AUTO_INCREMENT" json:"id"`
	UserId   *int64 `gorm:"column:USER_ID" json:"userId"`
	RoleCode string `gorm:"column:ROLE_CODE;size:100" json:"roleCode"`
	User     User   `gorm:"foreignKey:USER_ID"`
	Role     Role   `gorm:"foreignKey:ROLE_CODE"`
}

func (UserRole) TableName() string {
	return "USER_ROLE"
}
