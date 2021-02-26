package cfg

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/savsgio/go-logger/v2"
	"github.com/yakarim/website-walid/database"
)

var jwtSignKey = []byte("TestForFasthttpWithJWT")

type userCredential struct {
	Username []byte `json:"username"`
	Password []byte `json:"password"`
	jwt.StandardClaims
}

// GenerateToken ...
func (c *Cfg) GenerateToken(username []byte, password []byte) (string, time.Time) {
	logger.Debugf("Create new token for user %s", username)

	expireAt := time.Now().Add(1 * time.Minute)

	// Embed User information to `token`
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &userCredential{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt.Unix(),
		},
	})

	// token -> string. Only server knows the secret.
	tokenString, err := newToken.SignedString(jwtSignKey)
	if err != nil {
		logger.Error(err)
	}

	return tokenString, expireAt
}

func validateToken(requestToken string) (*jwt.Token, *userCredential, error) {
	logger.Debug("Validating token...")

	user := &userCredential{}
	token, err := jwt.ParseWithClaims(requestToken, user, func(token *jwt.Token) (interface{}, error) {
		return jwtSignKey, nil
	})

	return token, user, err
}

// ValidateUser ...
func (c *Cfg) ValidateUser(email, password string) (database.User, bool, error) {

	var user database.User
	if db.Where("email = ?", email).First(&user).RecordNotFound() {
		return user, false, errors.New("EMAIL NOT FOUND")
	}

	pwd2 := c.GetPwd(password)
	pwdMatch := c.ComparePasswords(user.Password, pwd2)
	if pwdMatch == false {
		return user, pwdMatch, errors.New("Password Not Mact")
	}

	return user, pwdMatch, errors.New("Success Login")
}
