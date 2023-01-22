package customEndpoints

import (
	"net/http"

	"github.com/jeremyawarren15/tsg68-api/middleware"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/samber/lo"
)

func AddEventsGet(app *pocketbase.PocketBase, e *core.ServeEvent) {
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
			middleware.RequireValidatedRecordAuth(),
		},
	})
}

func AddGetEventAttending(app *pocketbase.PocketBase, e *core.ServeEvent) {
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
			middleware.RequireValidatedRecordAuth(),
		},
	})
}

func AddGetEventResponse(app *pocketbase.PocketBase, e *core.ServeEvent) {
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
			middleware.RequireValidatedRecordAuth(),
		},
	})
}
