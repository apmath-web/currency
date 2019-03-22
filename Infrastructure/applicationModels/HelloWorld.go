package applicationModels

import "github.com/apmath-web/template-go/Domain"

type HelloWorld struct {
	Message string `json:"message"`
}

func (hw *HelloWorld) GetMessage() string {
	return hw.Message
}

func (hw *HelloWorld) SetMessage(message string) {
	hw.Message = message
}

func GenHelloWorldApplicationModel(message string) Domain.HelloWorldApplicationModel {
	hw := new(HelloWorld)
	hw.SetMessage(message)
	return hw
}
