package request

type UserRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
