package types

type ErrorMessage struct {
	Status   string      `json:"status"`
	Message  string      `json:"message"`
	Data     interface{} `json:"responseData"`
	Redirect string      `json:"redirect"`
}

type Singup struct {
	Firstname string `json:"firstName" validate:"required,min=3,max=32"`
	Lastname  string `json:"lastName" validate:"required,min=3,max=32"`
	Username  string `json:"userName" validate:"required,min=3,max=32"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=20"`
}

type Singin struct {
	Username string `json:"userName" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}
