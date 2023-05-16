package ioc

import (
	itinerary_handler "airland-server/src/domain/itinerary/handlers"
	"airland-server/src/infrastructure/datastore"
	"airland-server/src/infrastructure/repository"
	"airland-server/src/presentation/controller"
	"airland-server/src/presentation/router"
	"airland-server/src/presentation/server"
)

func Bootstrap() (s *server.Server, err error) {
	data, err := datastore.NewData()

	if err != nil {
		return &server.Server{}, err
	}

	routes := router.NewRouter()
	buildEntityRoutes(routes, data)

	return server.NewServer(routes), nil
}

func buildEntityRoutes(r *router.Routes, data *datastore.Data) {
	r.SetItineraryRoutes(getItineraryController(data))
}

func getItineraryController(data *datastore.Data) (c *controller.ItineraryController) {
	repo := repository.NewItineraryRepo(data)
	handler := itinerary_handler.NewItineraryHandler(repo)
	c = controller.NewItineraryController(handler)
	return
}
