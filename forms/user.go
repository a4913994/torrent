package forms

type UserSignup struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	BirthDay string `json:"birthday"`
	Gender   string `json:"gender"`
	PhotoURL string `json:"photo_url"`
}
