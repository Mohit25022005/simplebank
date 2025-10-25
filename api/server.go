package api

import (
	db "github.com/Mohit25022005/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

//Server serves HTTP requests for out banking service.
type Server struct {
	store *db.Store
	router *gin.Engine
}

//NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server{
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts",server.CreateAccount)
	router.GET("/accounts/:id",server.getAccount)
	router.GET("/accounts",server.listAccount)
	

	server.router = router
	return server
}


//Start HTTP server on a specific address
func (server *Server)Start(address string)error{
	return server.router.Run(address)
}

func errorRespose(err error)gin.H{
	return gin.H{"error": err.Error()}
}