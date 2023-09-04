package domain

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Login struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
}

type LoginRes struct {
	JWT string `json:"jwt"`
}
