package provider

type Message struct {
	Err string `json:"msg"`
}
type User struct {
	Message
	Login    *string
	Password *string
	Nickname *string
	Avatar   *string
}
