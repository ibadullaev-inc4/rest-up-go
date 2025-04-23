package apiserver

import (
	"io"
	"net/http"
	"rest-up-go/internal/app/store"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer is a struct that represents the API server.
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// New creates a new instance of APIServer.
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start the API server.
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()

	if err := s.configureStrore(); err != nil {
		return err
	}

	s.logger.Infof("Starting server on %s", s.config.BindAddr)
	s.logger.Info("Router configured")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, world!")
	}
}

func (s *APIServer) configureStrore() error {
	st := store.New(&s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}
