package actions

import (
	"encoding/json"
	"github.com/apmath-web/currency/Application/v1/validation"
	"github.com/apmath-web/currency/Infrastructure"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdaterHandler(c *gin.Context) {
	updater := Infrastructure.GetServiceManager().GetUpdaterService()
	if err := updater.Update(); err != nil {
		validator := validation.GenValidation()
		validator.SetMessage("Internal error")
		validator.AddMessage(validation.GenMessage("updating", err.Error()))
		str, _ := json.Marshal(validator)
		c.String(http.StatusInternalServerError, string(str))
		return
	}
}
