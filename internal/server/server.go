package server

import (
	"fmt"
	"log"

	"github.com/bypepe77/api/internal/module/posts"
	"github.com/bypepe77/api/internal/module/user"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Server struct {
	config *Config
	db     gorm.DB
	log    *logrus.Logger
	engine *gin.Engine
}

func NewServer(config *Config, db gorm.DB, log *logrus.Logger) *Server {
	return &Server{
		config: config,
		db:     db,
		log:    log,
		engine: gin.New(),
	}
}

func (s *Server) buildConnectionString() string {
	if s.config.Host == "" {
		s.config.Host = "localhost"
	}
	if s.config.Port == "" {
		s.config.Host = "localhost"
	}

	return fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
}

func (s *Server) Run() error {
	connectionString := s.buildConnectionString()
	log.Println("Server running on", s.config.Port)
	s.registerRoutes()
	return s.engine.Run(connectionString)
}

func (s *Server) registerRoutes() {
	// Initialize User routes
	userRoutes := user.NewUserRouteController(s.log, s.db, *s.engine.Group("/user"))
	userRoutes.RegisterUserRoutes()

	postRoutes := posts.NewPostRouteController(s.log, s.db, *s.engine.Group("/post"))
	postRoutes.RegisterPostRoutes()
}
