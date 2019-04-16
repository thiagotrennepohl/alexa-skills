package main

import (
	"github.com/labstack/echo"
	alexaHTTPRouter "github.com/thiagotrennepohl/alexa-skills/router"
	"github.com/thiagotrennepohl/alexa-skills/service"
	"github.com/thiagotrennepohl/alexa-skills/service/farewell"
	"github.com/thiagotrennepohl/alexa-skills/service/saudation"
)

func main() {
	router := echo.New()
	intents := make(map[string]service.Intent)
	saudationService := saudation.NewSaudationService()
	farewellService := farewell.NewFarewellService()

	intents = map[string]service.Intent{
		"saudation": saudationService,
		"farewell":  farewellService,
	}

	alexaHTTPRouter.NewSaudationRouter(router, intents)
	router.Start(":8080")
}
