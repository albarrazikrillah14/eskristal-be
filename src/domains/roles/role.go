package roles

type Role struct {
	ID   string `gorm:"primaryKey;column:id"`
	Name string `gorm:"column:name;unique"`
}

func (r *Role) TableName() string {
	return "roles"
}
