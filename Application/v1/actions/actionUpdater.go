package actions

import (
	"github.com/apmath-web/currency/Infrastructure"
	"github.com/gin-gonic/gin"
)

func UpdaterHandler(c *gin.Context) {
	updater := Infrastructure.GetServiceManager().GetUpdaterService()
	updater.Update()
}
