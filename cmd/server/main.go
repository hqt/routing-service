package main

import "github.com/hqt/zendesk-assignment"

func main() {
	app, err := zendesk.NewApp()
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
