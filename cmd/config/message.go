package config

type Message struct {
	Login  Login
	Logout Logout
}

type Login struct {
	StartMessage string
	EnterEmail   string
	EnterPass    string
}

type Logout struct {
	CredentialsClear string
}

var Messages = Message{
	Login{
		StartMessage: "Enter your LastBackend credentials.",
		EnterEmail:   "Enter your email or username: ",
		EnterPass:    "Enter your password (typing will be hidden): ",
	},
	Logout{
		CredentialsClear: "Local credentials cleared.",
	},
}
