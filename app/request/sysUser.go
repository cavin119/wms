package request

type Register struct {
	Username string   `json:"username" validate:"required,min=3,max=10" label:"用户名"`
	Password string   `json:"password" validate:"required,min=6,max=12" label:"密码"`
	RePassword string `json:"re_password" validate:"required,min=6,eqfield=Password" label:"确认密码"`
}

type Login struct {
	Username string `json:"username" validate:"required,min=3,max=10" label:"用户名"`
	Password string `json:"password" validate:"required,min=6,max=12" label:"密码"`
}
