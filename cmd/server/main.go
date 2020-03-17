package main

import routingservice "github.com/hqt/routing-service"

func main() {
	app, err := routingservice.NewApp()
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
