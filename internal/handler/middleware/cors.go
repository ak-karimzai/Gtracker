// Package middleware provides various HTTP middleware functions that can be used
// to enhance and secure HTTP requests processed by the Gin framework in the Gtracker application.
package middleware

import (
	"github.com/gin-gonic/gin"
)

// Cors returns a middleware function that handles Cross-Origin Resource Sharing (CORS) settings.
// This middleware enables the server to accept requests from different origins, specifically
// for development and production environments where the client and server might not share the same origin.
//
// The CORS middleware sets several HTTP headers to manage how resources are shared between different origins:
//   - "Access-Control-Allow-Origin": specifies the allowed origin for cross-origin requests. Here, it's set to
//     allow requests from "http://localhost:4200", typically used for local development.
//   - "Access-Control-Allow-Credentials": allows credentials such as cookies, authorization headers, or TLS client
//     certificates to be sent in cross-origin requests.
//   - "Access-Control-Allow-Headers": specifies the headers allowed in the actual request.
//   - "Access-Control-Allow-Methods": specifies the methods allowed when accessing the resource in response to a preflight request.
//
// If the HTTP method is "OPTIONS", the middleware will abort the request with a 204 (No Content) status,
// indicating that the actual request can be made safely.
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:4200")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
