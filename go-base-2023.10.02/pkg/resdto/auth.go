package resdto

type SignUpSuccess struct {
	Token        string
	RefreshToken string
}

type SignInSuccess struct {
	SignUpSuccess
}
