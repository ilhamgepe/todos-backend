package helper

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ilhamgepe/todos-backend/internal/models"
)

func CreateToken(user *models.User) (*string, *string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"sub":   user.ID,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})
	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"sub":   user.ID,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(([]byte(os.Getenv("JWT_SECRET"))))
	if err != nil {
		return nil, nil, err
	}
	refreshString, err := refresh.SignedString(([]byte(os.Getenv("JWT_SECRET_REFRESH"))))
	if err != nil {
		return nil, nil, err
	}

	return &tokenString, &refreshString, nil
}

func ValidateToken(tokenString string)(*models.JWTClaims,error){
	t,err := ExtractTokenFromBearer(tokenString)
	token,err := jwt.Parse(t, func(token *jwt.Token)(interface{},error){
		if method,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil,errors.New("invalid signing method")
		}else if method != jwt.SigningMethodHS256{
			return nil,errors.New("invalid signing method")
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		// Jika terjadi kesalahan saat mem-parsing token
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return &models.JWTClaims{
		Email: claims["email"].(string),
        Sub:   int(claims["sub"].(float64)),
        Exp:   claims["exp"].(float64),
	}, nil
}

func ValidateRefreshToken(tokenString string)(*models.JWTClaims,error){
	t,err := ExtractTokenFromBearer(tokenString)
	token,err := jwt.Parse(t, func(token *jwt.Token)(interface{},error){
		if method,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil,errors.New("invalid signing method")
		}else if method != jwt.SigningMethodHS256{
			return nil,errors.New("invalid signing method")
		}

		return []byte(os.Getenv("JWT_SECRET_REFRESH")), nil
	})
	if err != nil {
		// Jika terjadi kesalahan saat mem-parsing token
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return &models.JWTClaims{
		Email: claims["email"].(string),
        Sub:   int(claims["sub"].(float64)),
        Exp:   claims["exp"].(float64),
	}, nil
}

func ExtractTokenFromBearer(bearerString string) (string,error){
	if bearerString == "" || !strings.Contains(bearerString,"Bearer"){
		return "",errors.New("invalid token")
	}

	token := strings.TrimPrefix(bearerString, "Bearer ")

	return token, nil
}