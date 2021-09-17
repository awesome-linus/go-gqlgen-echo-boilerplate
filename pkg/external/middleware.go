package external

import (
	"github.com/labstack/echo/middleware"
)

func (s *Server) Middleware() {
	s.router.Use(middleware.Logger())
	s.router.Use(middleware.Recover())
	s.router.Use(middleware.Gzip())

	// s.router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{os.Getenv("CORS_ALLOW_ORIGIN")},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	// }))
}
