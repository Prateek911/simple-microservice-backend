package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"simple-microservice-backend/config"
	"simple-microservice-backend/db"
	"syscall"
	"time"
)

type ServerOptions struct {
	MaxIdleConnections int
	MaxOpenConnections int
	DialTimeout        time.Duration
	Timeout            time.Duration
	ContextTimeOut     time.Duration
}

type CORSOption string

const (
	ACCESS_CONTROL_ALLOW_ORIGIN      CORSOption = "Access-Control-Allow-Origin"
	ACCESS_CONTROL_ALLOW_CREDENTIALS CORSOption = "Access-Control-Allow-Credentials"
	ACCESS_CONTROL_ALLOW_HEADERS     CORSOption = "Access-Control-Allow-Headers"
	ACCESS_CONTROL_ALLOW_METHODS     CORSOption = "Access-Control-Allow-Methods"
	DEFAULT_HEADERS                  CORSOption = "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"
)

type CORSHeaders struct {
	ContentType    string `json:"Content-Type"`
	ContentLength  string `json:"Content-Length"`
	AcceptEncoding string `json:"Accept-Encoding"`
	CSRFToken      string `json:"X-CSRF-Token"`
	Authorization  string `json:"Authorization"`
	Accept         string `json:"accept"`
	Origin         string `json:"origin"`
	CacheControl   string `json:"Cache-Control"`
	XRequestedWith string `json:"X-Requested-With"`
}

type Server struct {
	httpServer *http.Server
}

func NewServer(port string) (*Server, error) {
	mux := setupRoutes()

	httpServer := &http.Server{Addr: ":" + port, Handler: mux}

	return &Server{httpServer: httpServer}, nil
}

func (s *Server) Start() error {

	opts, err := config.NewServerConfig()
	if err != nil {
		log.Fatal("Error initializing server options:", err)
		return err
	}

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen :%s\n", err)
		}
		log.Println("Server up on port :", s.httpServer.Addr)
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	log.Println("Server shutting down --->:", <-quit)

	if err := db.Close; err != nil {
		log.Println("Error closing DB connection :", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Duration(opts.ContextTimeOut*60).Seconds()))
	defer cancel()
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shut down---->:", err)
		return err
	}

	log.Println("server exiting")
	return nil

}
