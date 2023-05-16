package controller

import (
	"airland-server/src/cross_cutting/models"
	requests "airland-server/src/domain/itinerary/commands/requests"
	handler "airland-server/src/domain/itinerary/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItineraryController struct {
	handler *handler.ItineraryHandler
}

func NewItineraryController(handler *handler.ItineraryHandler) (controller *ItineraryController) {
	return &ItineraryController{
		handler,
	}
}

func (ic *ItineraryController) Create(ctx *gin.Context) {
	userId := ctx.GetString("UserID")
	var commandRequest requests.CreateItineraryCommandRequest
	commandRequest.UserID = userId

	if err := ctx.ShouldBindJSON(&commandRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any]("", false, []string{err.Error()}, nil))
		return
	}

	ic.handler.Create(commandRequest)
}
