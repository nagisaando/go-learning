package middleware

import (
	"net/http"

	"example.com/rest-api-2/utils"
	"github.com/gin-gonic/gin"
)

// middleware = an extra function that runs in the middle of the request

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	// Gin will continue the execution even if the middleware returns (throwing error).
	// AbortWithStatusJSON => It aborts the current response and sends the response that is attached and no other request handlers thereafter will be executed
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
		return
	}

	// Context.Set() allows to attach the data to context and can be used anywhere where the context is available
	context.Set("userId", userId)
	// context.Next(): It will ensure the next request handler in line will execute correctly
	context.Next()

}
