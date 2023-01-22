package main

import (
	"fmt"
	"log"

	endpoints "github.com/jeremyawarren15/tsg68-api/endpoints"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pocketbase.New()

	fmt.Print("Test")
	migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
		Automigrate: true,
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		endpoints.AddEventsGet(app, e)
		endpoints.AddUpdateResponse(app, e)
		endpoints.AddGetEventAttending(app, e)
		endpoints.AddGetEventResponse(app, e)

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
