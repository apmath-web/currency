package actions

import (
	"encoding/json"
	"net/http"

	"github.com/apmath-web/currency/Application/v1/validation"
	"github.com/apmath-web/currency/Application/v1/viewModels"
	"github.com/gin-gonic/gin"
)

func CurrencyHandler(c *gin.Context) {
	// currentCurrency, err := strconv.Atoi(c.Param("currentCurrency"))
	// wantedCurrency, err := strconv.Atoi(c.Param("wantedCurrency"))
	// amount := c.Param("amount")

	vm := viewModels.CurrencyViewModel{}
	if err := c.BindJSON(&vm); err != nil {
		validator := validation.GenValidation()
		validator.SetMessage("validation error")
		validator.AddMessage(validation.GenMessage("json", err.Error()))
		str, _ := json.Marshal(validator)
		c.String(http.StatusBadRequest, string(str))
		return
	}
	if !vm.Validate() {
		validator := vm.GetValidation()
		validator.SetMessage("validation error")
		str, _ := json.Marshal(validator)
		c.String(http.StatusBadRequest, string(str))
		return
	}

	// dm := mapper.CurrencyMapper(vm)

	c.JSON(http.StatusCreated, gin.H{"amount": 177, "currency": "EUR"})
}
