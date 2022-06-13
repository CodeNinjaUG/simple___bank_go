package api

import (
	db "github.com/codeninjaug/simple_bank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

//Server serves Http requests for our banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new Http Server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	//add routes to router
	router.POST("/create/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	router.POST("/create/transfer", server.createTransfer)
	router.PUT("/update/account/:id", server.UpdateAccount)
	router.DELETE("/delete/account/:id", server.DeleteAccount)
	server.router = router
	return server
}

//error reponse method for every function
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

//Start runs the http server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
