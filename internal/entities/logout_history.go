package entities

type LogoutHistory struct {
	Model
	JTI string `gorm:"type:uuid"`
}
