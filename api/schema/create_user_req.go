package schema

type CreateUserReq struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
