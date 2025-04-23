package apiserver

// APIServer is a struct that represents the API server.
type APIServer struct {
	config *Config
}

// New creates a new instance of APIServer.
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
	}
}

// Start the API server.
func (s *APIServer) Start() error {
	return nil
}
