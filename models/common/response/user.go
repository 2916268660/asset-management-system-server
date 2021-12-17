package response

type UserInfo struct {
	UserName string `json:"userName"`
	StuId    string `json:"stuId"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
