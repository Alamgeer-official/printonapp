package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
	JWT "github.com/golang-jwt/jwt"
	"githuh.com/printonapp/models"
)

const (
	SECRET string ="dfhdf8fr4y47b78b56f573857b358375b675365b8f"
)

type Claim struct {
	User models.User `json:"user"`
	JWT.StandardClaims
}

func CreateJWToken(user *models.User) (string, error) {
	//omit confidential info
	user.Password=""
	user.Phone=""
	atClaim:=Claim{
		User:           *user,
		StandardClaims: JWT.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute *240).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "printonapp",
			
		},
	}

	tokenWithClaim:=JWT.NewWithClaims(jwt.SigningMethodHS256,atClaim)
	token,err:=tokenWithClaim.SignedString([]byte(SECRET))
	if err != nil{
		return "",nil
	}

	return token,nil


}
