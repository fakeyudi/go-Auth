package server

import (
	"context"
	"fmt"
	"github.com/markbates/goth/gothic"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/", s.HelloWorldHandler)

	r.GET("/health", s.healthHandler)

	r.GET("/auth/:provider/callback", s.GetAuthCallbackFunction)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) GetAuthCallbackFunction(c *gin.Context) {
	provider := c.Param("provider")
	fmt.Println("provider", provider)
	c.Request = c.Request.WithContext(context.WithValue(context.Background(), "provider", provider))

	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		fmt.Println(err, c.Writer, c.Request)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(user)

	c.JSON(http.StatusOK, user)
}
