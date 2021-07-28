package service

import (
	"net/http"
	"os"
	"section8/domain"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type JwtUsecase struct {
	UserRepo domain.UserRepository
}

func NewJwtService(userRepo domain.UserRepository) JwtUsecase {
	return JwtUsecase{userRepo}
}
func (h *JwtUsecase) SetJwtAdmin(g *echo.Group) {

	secret := os.Getenv("secret")

	// validate jwt token
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte(secret),
	}))

	// validate payload related with admin type of token
	g.Use(h.validateJwtAdmin)
}

func (h *JwtUsecase) validateJwtAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		user := c.Get("user")
		token := user.(*jwt.Token)
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if claims["is_admin"] == true {
				return next(c)
			} else {
				return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
			}
		}

		return echo.NewHTTPError(http.StatusForbidden, "Invalid Token")
	}
}
