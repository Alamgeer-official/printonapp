
package middelware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupCORSMiddleware configures and returns the CORS middleware.
func SetupCORSMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"*"} // Replace with your frontend's actual origin
	// config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}

	return cors.New(config)
}
