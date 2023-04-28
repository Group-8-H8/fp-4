package helpers

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fydhfzh/fp-4/entity"
	"github.com/fydhfzh/fp-4/pkg/errs"
)

const secretKey = "keren"

var invalidTokenError = errs.NewUnauthenticatedError("invalid token")

func ClaimToken(id int, email string, role string) *jwt.MapClaims {
	return &jwt.MapClaims{
		"id": id,
		"email": email,
		"role": role,
		"exp": time.Now().Add(time.Hour * 10).Unix(),
	}

}

func SignToken(signingMethod jwt.SigningMethod, token jwt.MapClaims) string {
	claim := jwt.NewWithClaims(signingMethod, token)

	tokenString, err := claim.SignedString([]byte(secretKey))

	if err != nil {
		return ""
	}

	return tokenString
}

func GenerateToken(id int, email string, role string) (string, errs.Errs) {
	tokenClaim := ClaimToken(id, email, role)

	signedToken := SignToken(jwt.SigningMethodHS256, *tokenClaim)

	if signedToken == "" {
		return "", errs.NewInternalServerError("something went wrong")
	}

	return signedToken, nil
}

func parseToken(tokenString string) (*jwt.Token, errs.Errs) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, invalidTokenError
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, invalidTokenError
	}

	return token, nil
}

func bindTokenToUserEntity(mapClaims jwt.MapClaims) (*entity.User, errs.Errs) {
	var user entity.User

	if id, ok := mapClaims["id"].(float64); !ok {
		return nil, invalidTokenError
	} else {
		user.ID = uint(id)
	}

	if email, ok := mapClaims["email"].(string); !ok {
		return nil, invalidTokenError
	} else {
		user.Email = email
	}

	if role, ok := mapClaims["role"].(string); !ok {
		return nil, invalidTokenError
	} else {
		user.Role = role
	}

	return &user, nil
}

func GetUserData(bearerToken string) (*entity.User, errs.Errs) {
	mapClaims, err := validateToken(bearerToken)

	if err != nil {
		return nil, invalidTokenError
	}
	
	user, err := bindTokenToUserEntity(*mapClaims)
	
	if err != nil{
		return nil, invalidTokenError
	}

	return user, nil
}

func validateToken(bearerToken string) (*jwt.MapClaims, errs.Errs) {
	isBearer := strings.HasPrefix(bearerToken, "Bearer")

	if !isBearer {
		return nil, invalidTokenError
	}

	split := strings.Split(bearerToken, " ")

	if len(split) != 2 {
		return nil, invalidTokenError
	}

	tokenString := split[1]

	token, err := parseToken(tokenString)

	if err != nil {
		return nil, err
	}

	var mapClaims jwt.MapClaims

	if claims, ok := token.Claims.(jwt.MapClaims); !ok {
		return nil, invalidTokenError
	} else {
		mapClaims = claims
	}

	return &mapClaims, nil
}