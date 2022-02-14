package response

type UserInfo struct {
	UserName   string `json:"userName"`
	UserId     string `json:"userId"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Department string `json:"department"`
}
