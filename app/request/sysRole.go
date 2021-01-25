package request

type CreateRole struct {
	PId int `json:"p_id" validate:"gte=0" label:"上级ID"`
	RoleName string `json:"role_name" validate:"required,min=3,max=10" label:"角色名称"`
}
