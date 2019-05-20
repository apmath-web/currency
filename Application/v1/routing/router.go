package routing

import (
	"github.com/gin-gonic/gin"
)

func GenRouter() *gin.Engine {
	router := gin.Default()
	_ = router.Group("/v1")
	return router
}
