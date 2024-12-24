package request

type UserRequest struct {
	Email    string `json:"email" binding:""`
	Password string `json:"password" binding:""`
	Name     string `json:"name" binding:"required,min=2,max=50"`
	Age      int8   `json:"age" binding:""`
}


