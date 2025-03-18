package roles

type Role struct {
	ID   string `gorm:"primaryKey;column:id"`
	Name string `gorm:"column:name;unique"`
}

func (r *Role) TableName() string {
	return "roles"
}

func (r *Role) MapToResponse() RoleResponse {
	return RoleResponse{
		ID:   r.ID,
		Name: r.Name,
	}
}

type CreateRoleRequest struct {
	Name string `json:"name" validate:"required"`
}

func (c *CreateRoleRequest) MapToRole() Role {
	return Role{
		Name: c.Name,
	}
}

type DeleteRoleRequest struct {
	ID string `json:"id" validate:"required,uuid"`
}

type RoleResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
