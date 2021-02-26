package cfg

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/savsgio/atreugo/v11"
)

// AuthDetails ...
type AuthDetails struct {
	AuthUUID string
	UserID   string
}

// UserCredential ..
type UserCredential struct {
	Username   string `json:"username"`
	Authorized bool   `json:"authorized"`
	AuthUUID   string `json:"auth_uuid"`
	UserID     string `json:"user_id"`

	jwt.StandardClaims
}

// CreateToken ...
func CreateToken(authD AuthDetails) (string, error) {
	expireAt := time.Now().Add(1 * time.Minute)
	fmt.Println(expireAt.Unix())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserCredential{
		Username:   "username",
		Authorized: true,
		UserID:     authD.UserID,
		AuthUUID:   authD.AuthUUID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt.Unix(),
		},
	})
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

// TokenValid ...
func TokenValid(ctx *atreugo.RequestCtx) error {
	token, err := VerifyToken(ctx)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {

		return err
	}
	return nil
}

// VerifyToken ...
func VerifyToken(ctx *atreugo.RequestCtx) (*jwt.Token, error) {
	tokenString := ExtractToken(ctx)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//does this token conform to "SigningMethodHMAC" ?
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// ExtractToken get the token from the request body
func ExtractToken(ctx *atreugo.RequestCtx) string {
	jwtCookie := ctx.Request.Header.Cookie("__jwt")
	if len(jwtCookie) == 0 {
		token := ctx.QueryArgs().Peek("token")
		if string(token) != "" {
			return string(token)
		}
		bearToken := ctx.Request.Header.Peek("Authorization")
		strArr := strings.Split(string(bearToken), " ")
		if len(strArr) == 2 {
			return strArr[1]
		}
		return ""
	}

	return string(jwtCookie)
}

// ExtractTokenAuth ...
func ExtractTokenAuth(ctx *atreugo.RequestCtx) (*AuthDetails, error) {
	token, err := VerifyToken(ctx)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		AuthUUID, ok := claims["auth_uuid"].(string) //convert the interface to string
		if !ok {
			return nil, err
		}
		return &AuthDetails{
			AuthUUID: AuthUUID,
			UserID:   claims["user_id"].(string),
		}, nil
	}
	return nil, err
}
