package cfg

import (
	"fmt"
	"log"
	"time"

	jsoniter "github.com/json-iterator/go"

	"github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger"
	"golang.org/x/crypto/bcrypt"
)

// H ...
type H map[string]interface{}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// TimeIn ...
func (c *Cfg) TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

// GetPwd ...
func (c *Cfg) GetPwd(pwd string) []byte {
	_, err := fmt.Scan(&pwd)
	if err != nil {
		log.Println(err)
	}
	return []byte(pwd)
}

// HashAndSalt ...
func (c *Cfg) HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// ComparePasswords ...
func (c *Cfg) ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}

// JSON ...
func (c *Cfg) JSON(ctx *atreugo.RequestCtx, body interface{}, statusCode ...int) error {
	ctx.Response.Header.SetContentType("application/json")

	if len(statusCode) > 0 {
		ctx.Response.Header.SetStatusCode(statusCode[0])
	}

	data, err := json.Marshal(body)
	if err != nil {
		logger.Print(err)
	}

	ctx.Response.SetBody(data)

	return nil
}
