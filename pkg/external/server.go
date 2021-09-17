package external

import (
	"flag"
	"fmt"
	"log"

	"github.com/awesome-linus/go-graphql-echo-boilerplate/pkg/external/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type Server struct {
	router *echo.Echo
	DB     *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	return &Server{
		router: echo.New(),
		DB:     db,
	}
}

func (s *Server) Init(env string) {
	log.Printf("env: %s", env)
}

func StartServer() {
	var (
		port = flag.String("port", "3000", "addr to bind")
		env  = flag.String("env", "develop", "実行環境 (production, staging, develop)")
	)

	flag.Parse()

	db := mysql.Connect()

	s := NewServer(db)
	s.Init(*env)
	s.Middleware()

	s.GraphqlRouter()

	err := s.router.Start(fmt.Sprint(":", *port))
	if err != nil {
		log.Fatalln(err)
	}

}
