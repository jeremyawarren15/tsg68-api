package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/samber/lo"
)

type EventResponse struct {
	EventSlug string `json:"event"`
	Response  string `json:"response"`
}

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// Get Events by Slug
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/events/:slug",
			Handler: func(c echo.Context) error {
				event, err := app.Dao().FindFirstRecordByData("events", "slug", c.PathParam("slug"))
				if err != nil {
					return apis.NewNotFoundError("The event does not exist.", err)
				}

				apis.EnrichRecord(c, app.Dao(), event)

				return c.JSON(200, event)
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
				RequireValidatedRecordAuth(),
			},
		})

		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/events/:slug/response",
			Handler: func(c echo.Context) error {
				user, _ := c.Get(apis.ContextAuthRecordKey).(*models.Record)
				event, err := app.Dao().FindFirstRecordByData("events", "slug", c.PathParam("slug"))
				if err != nil {
					return apis.NewNotFoundError("The event does not exist.", err)
				}

				responses, err := app.Dao().FindRecordsByExpr("responses", dbx.HashExp{"event": event.Id, "user": user.Id})
				if err != nil {
					return err
				}

				if len(responses) > 0 {
					return c.JSON(200, responses[0].GetString("response"))
				}

				return c.JSON(200, "pending")
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
				RequireValidatedRecordAuth(),
			},
		})

		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/events/:slug/attending",
			Handler: func(c echo.Context) error {
				event, err := app.Dao().FindFirstRecordByData("events", "slug", c.PathParam("slug"))
				if err != nil {
					return apis.NewNotFoundError("The event does not exist.", err)
				}

				responses, err := app.Dao().FindCollectionByNameOrId("responses")
				if err != nil {
					return err
				}

				query := app.Dao().RecordQuery(responses).
					AndWhere(dbx.HashExp{"event": event.Id, "response": "attending"})

				rows := []dbx.NullStringMap{}
				if err := query.All(&rows); err != nil {
					return err
				}

				existingResponses := models.NewRecordsFromNullStringMaps(responses, rows)
				userIds := lo.Map(existingResponses, func(x *models.Record, index int) string {
					return x.GetString("user")
				})

				records, err := app.Dao().FindRecordsByIds("users", userIds)

				names := lo.Map(records, func(x *models.Record, index int) string {
					return x.GetString("name")
				})

				return c.JSON(200, names)
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
				RequireValidatedRecordAuth(),
			},
		})

		e.Router.AddRoute(echo.Route{
			Method: http.MethodPost,
			Path:   "/response",
			Handler: func(c echo.Context) error {
				user, _ := c.Get(apis.ContextAuthRecordKey).(*models.Record)

				var response EventResponse
				err := c.Bind(&response)
				if err != nil {
					return c.String(http.StatusBadRequest, "bad request")
				}

				event, err := app.Dao().FindFirstRecordByData("events", "slug", response.EventSlug)
				if err != nil {
					return apis.NewNotFoundError("The event does not exist.", err)
				}

				responses, err := app.Dao().FindCollectionByNameOrId("responses")
				if err != nil {
					return err
				}

				query := app.Dao().RecordQuery(responses).
					AndWhere(dbx.HashExp{"event": event.Id, "user": user.Id}).
					Limit(1)

				rows := []dbx.NullStringMap{}
				if err := query.All(&rows); err != nil {
					return err
				}

				existingResponses := models.NewRecordsFromNullStringMaps(responses, rows)
				if len(existingResponses) > 0 {
					record := existingResponses[0]
					record.Set("response", response.Response)

					if err := app.Dao().SaveRecord(record); err != nil {
						return err
					}
					return c.JSON(200, record)
				}

				record := models.NewRecord(responses)
				record.Set("event", event.Id)
				record.Set("user", user.Id)
				record.Set("response", response.Response)

				if err := app.Dao().SaveRecord(record); err != nil {
					return err
				}
				return c.JSON(200, record)
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
				RequireValidatedRecordAuth(),
			},
		})

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
