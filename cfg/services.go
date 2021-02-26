package cfg

type sigInInterface interface {
	SignIn(AuthDetails) (string, error)
}

type signInStruct struct{}

//let expose this interface:
var (
	Authorize sigInInterface = &signInStruct{}
)

func (si *signInStruct) SignIn(authD AuthDetails) (string, error) {
	token, err := CreateToken(authD)
	if err != nil {
		return "", err
	}
	return token, nil
}
