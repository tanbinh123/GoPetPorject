package enteties

type User struct {
	ID       string `json:"id"`
	Login    int    `json:"login"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Created  string `json:"created"`
	Modified string `json:"modified"`
}
