package reqdto

type TSignInWithUsername struct {
	Username string
	Password string
}

type TSignUpWithUsername struct {
	TSignInWithUsername
	Name string
}
