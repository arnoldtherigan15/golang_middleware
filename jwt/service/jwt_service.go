package service

import (
	"fmt"
	"net/http"
	"os"
	"section8/domain"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtUsecase struct {
	UserRepo domain.UserRepository
}

func NewJwtService(userRepo domain.UserRepository) domain.JwtUsecase {
	return &JwtUsecase{userRepo}
}
func (h *JwtUsecase) SetJwtAdmin(g *echo.Group) {

	secret := os.Getenv("SECRET")
	// fmt.Println(secret, ">>>>>>> disini 24")
	// validate jwt token
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte(secret),
	}))
	// fmt.Println("masukk jwtttt<<<< 30", secret)
	// validate payload related with admin type of token
	g.Use(h.validateJwtAdmin)
}

func (h *JwtUsecase) validateJwtAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("masukkkkk>>>> 36")

	return func(c echo.Context) error {
		fmt.Println("masukkkkk>>>> 38")
		user := c.Get("user")
		token := user.(*jwt.Token)
		fmt.Println(token, ">>>>>>")
		if true {
			return next(c)
		}
		// if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// 	if claims["role"] == "admin" {
		// 		return next(c)
		// 	} else {
		// 		return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
		// 	}
		// }

		return echo.NewHTTPError(http.StatusForbidden, "Invalid Token")
	}
}
