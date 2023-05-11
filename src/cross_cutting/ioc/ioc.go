package ioc

import (
	"airland-server/src/infrastructure/datastore"
	"airland-server/src/infrastructure/router"
	"airland-server/src/infrastructure/server"
)

func Bootstrap() (s *server.Server, err error) {
	_, err = datastore.NewData()

	if err != nil {
		return &server.Server{}, err
	}

	routes := router.NewRouter()

	return server.NewServer(routes), nil
}
