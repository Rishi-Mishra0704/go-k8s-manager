package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
	listenAddr string
	router     *echo.Echo
}

func NewServer(listenAddr string) *Server {
	return &Server{listenAddr: listenAddr}
}

func (server *Server) setupRouter() {
	router := echo.New()

	// Middleware
	router.POST("/hello", sayHello)
	server.router = router
}

func (server *Server) Start() error {
	server.setupRouter()
	return server.router.Start(server.listenAddr)
}

type User struct {
	Name string `json:"name"`
}

func sayHello(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, response(err, "Invalid request payload"))
	}
	return c.JSON(http.StatusOK, response(nil, "Hello "+user.Name))
}

func response(err error, message string) map[string]interface{} {
	resp := map[string]interface{}{
		"message": message,
	}
	if err != nil {
		resp["error"] = err.Error()
	}
	return resp
}
