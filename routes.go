// Routes stuff here
package boulder

import (
	"github.com/tpageforfunzies/boulder/handlers"
	"github.com/gin-gonic/gin"
	"sync"
)

var router *gin.Engine
var once sync.Once

// GetRouter returns a singleton instance of the router
func GetRouter() *gin.Engine {
	// Singleton router instance
	once.Do(func() {
		router = gin.Default()
	})

	return router
}

func AddRoutes(group *gin.RouterGroup) {
	group.GET("/", homeHandler)
	group.POST("/user/new", handlers.CreateUser)
}