package types

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type Account struct {
	AcID    int64   `json:"ac-id"`
	Name    string  `json:"name"`
	Balance float32 `json:"balance"`
}
