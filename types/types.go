package types

type Person struct {
	Username        string `json:"username"`
	Firstname       string `json:"first"`
	Lastname        string `json:"last"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Confirmpassword string `json:"confirm"`
}

type Login struct {
	Username string `json:"name"`
	Password string `json:"password"`
}

type ChangePasswordRequest struct {
	Username    string `json:"name"`
	Oldpassword string `json:"opass"`
	Newpassword string `json:"npass"`
}
