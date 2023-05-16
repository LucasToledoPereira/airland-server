package router

import (
	"airland-server/src/presentation/controller"
	"fmt"
)

func (r *Routes) SetItineraryRoutes(c *controller.ItineraryController) {
	fmt.Println(c)
	ig := r.Private.Group("/itinerary")
	ig.POST("", c.Create)
}
