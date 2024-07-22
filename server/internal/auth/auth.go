package auth

type Parameters struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Refresh  string `json:"refresh,omitempty"`
	Token    string `json:"token,omitempty"`
}

func GetToken(params Parameters) string {
	return params.Username
}
