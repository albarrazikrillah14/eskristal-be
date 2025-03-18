package roles

type Role struct {
	ID   string `gorm:"primaryKey;column:id"`
	Name string `gorm:"column:name;unique"`
}

func (r *Role) TableName() string {
	return "roles"
}

type CreateRoleRequest struct {
	Name string `json:"name" validate:"required"`
}

func (c *CreateRoleRequest) MapToRole() Role {
	return Role{
		Name: c.Name,
	}
}
