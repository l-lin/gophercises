package auth

// Credentials used to authenticate the user
type Credentials struct {
	UserName, Password string
}

// Result is the result of the authentication attempt
type Result struct {
	Token string
	Error error
}

// Authenticator is used to authenticate an user
type Authenticator interface {
	Authenticate(chan *Result, *Credentials)
}
