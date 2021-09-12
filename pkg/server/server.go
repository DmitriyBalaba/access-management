package server

import (
	"access-management/pkg/config"
	"access-management/pkg/config/server"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/rs/zerolog/log"
)

type Server struct {
	*http.Server
	config *server.Config
}

func NewServer(config *config.Config) *Server {
	fmt.Println("go to server")
	if config == nil {
		panic("cannot create new server with nil config")
	}
	fmt.Println("config exists server")
	return &Server{
		Server: &http.Server{
			Addr: ":" + strconv.Itoa(config.Server().Port),
			// TODO: set read and write timeouts as well
		},
		config: config.Server(),
	}
}

// RouteInitializer is a function which inits routes for specified router
type RouteInitializer func(r *mux.Router, ctx context.Context)

var routeInitializers []RouteInitializer

// Adds RouteInitializer to global list for NewRouter
func AddRouteInitializer(f RouteInitializer) {
	if f == nil {
		panic("cannot add nil RouteInitializer")
	}
	routeInitializers = append(routeInitializers, f)
}

func (s *Server) waitForInterrupt() {
	waitTime := time.Second * time.Duration(s.config.GracefulTimeout)

	// creating a channel and waiting for Interrupt signal sent
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	<-signalChannel

	// Interrupt signal has been sent so the server shuts down with delay
	log.Info().Msgf("Shutting down (%d seconds timeout)", s.config.GracefulTimeout)
	ctx, cancel := context.WithTimeout(context.Background(), waitTime)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal().Msg("Cannot shut down: " + err.Error())
		os.Exit(-1)
	}
	log.Info().Msg("Bye")
}

func (s *Server) Run() {
	s.Handler = s.NewRouter(context.Background())
	if s.Handler == nil {
		panic("cannot run server with nil handler: use SetHandler()")
	}

	go func() {
		log.Info().Msgf("Listen and serve on port: %d", s.config.Port)
		if err := s.Server.ListenAndServe(); err != nil {
			log.Fatal().Msg("Cannot run server: " + err.Error())
			os.Exit(-1)
		}
	}()
	s.waitForInterrupt()
}

func (s *Server) Start() {
	fmt.Println("start server")
}

type FallibleHandlerFunc func(w http.ResponseWriter, r *http.Request) (err error)

func NewRoute(r *mux.Router, path string, h FallibleHandlerFunc, constructors ...alice.Constructor) *mux.Route {
	if r == nil || h == nil {
		panic("cannot create new route with nil input")
	}
	return r.Handle(path, alice.New(constructors...).Then(exec(h)))
}

func exec(h FallibleHandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			w.WriteHeader(http.StatusForbidden)
		}
	})
}

func (s *Server) NewRouter(ctx context.Context) *mux.Router {
	if ctx == nil {
		panic("cannot create new router with nil context")
	}

	r := mux.NewRouter()
	if p := s.config.PathPrefix; p != "" {
		r = r.PathPrefix("/" + p).Subrouter()
	}

	for i := range routeInitializers {
		routeInitializers[i](r, ctx)
	}

	// reset routeInitializers
	routeInitializers = nil
	return r
}
