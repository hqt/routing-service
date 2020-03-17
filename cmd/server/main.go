package main

import (
	"github.com/hqt/routing-service/src/server"
)

func main() {
	app, err := server.NewApp("etc/StationMap.csv")
	if err != nil {
		panic(err)
	}

	_ = app.Start()
	defer func() {
		err := app.Stop()
		if err != nil {
			panic(err)
		}
	}()
}
