package main

import (
	"github.com/Genarodaniel/order-system/config/dependency"
	"github.com/Genarodaniel/order-system/config/env"
	"github.com/Genarodaniel/order-system/internal/infra/api/server"
	_ "github.com/lib/pq"
)

func main() {
	env.Load()

	if err := dependency.Load(); err != nil {
		panic(err)
	}
	defer dependency.DB.Close()

	s := server.Init()

	s.Run(env.Config.WebServerPort)

}
