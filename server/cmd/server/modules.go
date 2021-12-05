// +build wireinject

package main

import (
	"github.com/software-engineering-components/lab3/server/restaurant"
	"github.com/google/wire"
)

func ComposeApiServer(port int) (*RestaurantApiServer, error) {
	wire.Build(
		NewDbConnection,
		restaurant.Providers,
		wire.Struct(new(RestaurantApiServer), "Port", "RestaurantHandler"),
	)
	return nil, nil
}
