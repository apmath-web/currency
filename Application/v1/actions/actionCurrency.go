package actions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/apmath-web/currency/Application/v1/mapper"
	"github.com/apmath-web/currency/Application/v1/validation"
	"github.com/apmath-web/currency/Application/v1/viewModels"
	"github.com/apmath-web/currency/Infrastructure"
	"github.com/gin-gonic/gin"
)

func CurrencyHandler(c *gin.Context) {
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

	dm := mapper.CurrencyMapper(vm)

	// updater := Infrastructure.GetServiceManager().GetUpdaterService()
	// updater.Update()

	exchanger := Infrastructure.GetServiceManager().GetExchangerService()
	ans := exchanger.Exchange(dm)

	fmt.Print("Answer", ans.GetAmount())
	// fetcher := applicationModels.Fetcher{}

	// rates := fetcher.FetchAll()

	// fmt.Print(rates)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	c.JSON(http.StatusCreated, gin.H{"amount": ans.GetAmount(), "currency": dm.GetWantedCurrency()})
}
