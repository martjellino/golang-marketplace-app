package user

type UserRequest struct {			
	Username   string	`json:"username" binding:"required,min=3,max=50`
	Password   string	`json:"password" binding:"required,min=3,max=50`
	Fullname   string	`json:"fullname" binding:"required,min=3,max=50`
}
