package apiserver

import "github.com/sirupsen/logrus"

// APIServer is a struct that represents the API server.
type APIServer struct {
	config *Config
	logger *logrus.Logger
}

// New creates a new instance of APIServer.
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
	}
}

// Start the API server.
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.logger.Infof("Starting server on %s", s.config.BindAddr)
	return nil
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}
