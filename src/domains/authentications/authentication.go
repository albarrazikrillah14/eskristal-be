package authentications

type Authentication struct {
	ID    int64  `gorm:"primaryKey;autoIncremenet;column:id"`
	Token string `gorm:"column:token"`
}

func (a *Authentication) TableName() string {
	return "authentications"
}

type AuthenticationRequest struct {
	UsernameOrEmail string `json:"username_or_email" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
}
