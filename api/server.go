package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/harissucipto/xendit-task/db/sqlc"
	"github.com/harissucipto/xendit-task/util"
)

// Server servers HTTP requests for our api services
type Server struct {
	config util.Config
	store db.Store 
	router *gin.Engine
}

// NewServer creates a new HTTP Server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store: store,
	}


	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/orgs/:org-name/comments", server.listComments)
	router.POST("/orgs/:org-name/comments", server.createComment)
	router.DELETE("/orgs/:org-name/comments", server.deleteComment)

	// get  /orgs/<org-name>/members/
	router.GET("/orgs/:org-name/members", server.listMembers)
	
	server.router = router 
}

// Start runs the HTTP server on a specifc address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}