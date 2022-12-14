package request

type CreateUserRequest struct {
	Age      int    `json:"age" validate:"required,gt=8" example:"20"`
	Email    string `json:"email" validate:"required,email" example:"@gmail.com"`
	Password string `json:"password" validate:"required,min=6" example:"hahahihilulus"`
	Username string `json:"username" validate:"required" example:""`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email" example:"@gmail.com"`
	Password string `json:"password" validate:"required,min=6" example:"hahahihilulus"`
}

type UpdateUserRequest struct {
	Email    string `json:"email" validate:"required,email" example:"@gmail.com"`
	Username string `json:"username" validate:"required" example:""`
}
