package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/software-engineering-components/lab3/server/restaurant"
)

type HttpPortNumber int

type RestaurantApiServer struct {
	Port    int
	RestaurantHandler restaurant.HttpHandlerFunc
	server  *http.Server
}

func (s *RestaurantApiServer) Start() error {
	if s.Handler == nil {
		return fmt.Errorf("channels HTTP handler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	handler.HandleFunc("/restaurant", s.RestaurantHandler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

func (s *RestaurantApiServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}
