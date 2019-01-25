package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	api "github.com/zrcni/go-battlenet-graphql-api"
	"github.com/zrcni/go-battlenet-graphql-api/battlenet"
	"github.com/zrcni/go-battlenet-graphql-api/config"
)

const defaultPort = "4000"

var srv = &server{}

type server struct {
	router chi.Router
}

func authenticateBattleNet(battlenetAuth *battlenet.Auth) {
	if err := battlenetAuth.Authenticate(); err != nil {
		time.Sleep(time.Second)
		authenticateBattleNet(battlenetAuth)
	}
}

func (s *server) middlewares() {
	battlenetAuth := &battlenet.Auth{}
	go func() { authenticateBattleNet(battlenetAuth) }()
	s.router.Use(battlenet.Middleware(battlenetAuth))
	s.router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	}).Handler)
}

func (s *server) routes() {
	s.router.Get("/", handler.Playground("GraphQL playground", "/query"))
	s.router.Post("/query", handler.GraphQL(api.NewExecutableSchema(api.Config{Resolvers: &api.Resolver{}})))
}

func (s *server) listen() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, s.router))
}

func (s *server) Serve() {
	s.router = chi.NewRouter()
	s.middlewares()
	s.routes()
	s.listen()
}

func init() {
	config.Setup()
}

func main() {
	srv.Serve()
}
