package model

type User struct {
	Id          uint   `json:"id" binding:"required"`
	StatusId    uint   `json:"status_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
	Gender      int    `json:"gender" binding:"required"`
	Description string `json:"description"`
}

type UserProfile struct {
	Id          uint   `json:"id" binding:"required"`
	UserId      uint   `json:"user_id" binding:"required"`
	JoinDate    string `json:"join_date" binding:"required"`
	ImageUrl    string `json:"image_url" binding:"required"`
	Description string `json:"description"`
}

var UserList = []User{
	{Id: 0, StatusId: 1, Name: "SEO1", Email: "SEO1@example.com", Phone: "+821012345678", Gender: 1, Description: ""},
	{Id: 1, StatusId: 0, Name: "SEO2", Email: "SEO2@example.com", Phone: "+821022222222", Gender: 2, Description: ""},
	{Id: 2, StatusId: 1, Name: "SEO3", Email: "SEO3@example.com", Phone: "+821033333333", Gender: 1, Description: ""},
	{Id: 3, StatusId: 1, Name: "SEO4", Email: "SEO4@example.com", Phone: "+821044444444", Gender: 2, Description: ""},
}

var UserProfileList = []UserProfile{
	{Id: 0, UserId: 0, JoinDate: "2022-05-10 18:03:00", ImageUrl: "111", Description: ""},
	{Id: 1, UserId: 1, JoinDate: "2022-05-11 18:03:00", ImageUrl: "222", Description: ""},
	{Id: 2, UserId: 2, JoinDate: "2022-05-12 18:03:00", ImageUrl: "333", Description: ""},
	{Id: 3, UserId: 3, JoinDate: "2022-05-13 18:03:00", ImageUrl: "444", Description: ""},
}
