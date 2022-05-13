package model

type User struct {
	Id              uint    `gorm:"column:id;AUTO_INCREMENT" json:"id"`
	CognitoUsername string  `gorm:"column:cognito_username" json:"cognito_username" binding:"required"`
	Username        string  `gorm:"column:username" json:"username"`
	Pass            *string `gorm:"column:password" json:"password"`
	Name            *string `gorm:"column:name" json:"name"`
	Market          *string `gorm:"column:market" json:"market"`
	Email           *string `gorm:"column:email" json:"email"`
	Birthday        *string `gorm:"column:birthday" json:"birthday"`
	Gender          *string `gorm:"column:gender" json:"gender"`
	Phone           *string `gorm:"column:phone_number" json:"phone_number" binding:"required"`
	UsedCount       int     `gorm:"column:used_count" json:"used_count"`
	StatusId        int     `gorm:"column:status_id" json:"status_id" binding:"required"`
	OrderingId      *int    `gorm:"column:ordering_id" json:"ordering_id"`
	Description     *string `gorm:"column:description" json:"description"`
}

type UserProfile struct {
	Id              uint    `gorm:"column:id;AUTO_INCREMENT" json:"id"`
	CognitoUsername string  `gorm:"column:cognito_username" json:"cognito_username" binding:"required"`
	Username        string  `gorm:"column:username" json:"username"`
	Pass            *string `gorm:"column:password" json:"password"`
	Name            *string `gorm:"column:name" json:"name"`
	Market          *string `gorm:"column:market" json:"market"`
	Email           *string `gorm:"column:email" json:"email"`
	Birthday        *string `gorm:"column:birthday" json:"birthday"`
	Gender          *string `gorm:"column:gender" json:"gender"`
	Phone           *string `gorm:"column:phone_number" json:"phone_number" binding:"required"`
	UsedCount       int     `gorm:"column:used_count" json:"used_count"`
	StatusId        int     `gorm:"column:status_id" json:"status_id" binding:"required"`
	OrderingId      *int    `gorm:"column:ordering_id" json:"ordering_id"`
	Description     *string `gorm:"column:description" json:"description"`
}
