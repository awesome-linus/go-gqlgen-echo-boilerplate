package external

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/awesome-linus/go-graphql-echo-boilerplate/pkg/graph/generated"
	graph "github.com/awesome-linus/go-graphql-echo-boilerplate/pkg/graph/resolver"
	"github.com/labstack/echo"
)

func (s *Server) GraphqlRouter() {
	graphqlHandler := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{DB: s.DB}},
		),
	)
	playgroundHandler := playground.Handler("GraphQL", "/graphql")

	s.router.POST("/graphql", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	s.router.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})
}
