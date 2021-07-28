package domain

import "github.com/labstack/echo/v4"

// type Jwt struct{}

type JwtUsecase interface {
	SetJwtSuperAdmin(g *echo.Group)
	SetJwtEmployee(g *echo.Group)
	// SetJwtGeneral(g *echo.Group)
}
