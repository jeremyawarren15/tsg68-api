package customEndpoints

import (
	"net/http"

	"github.com/jeremyawarren15/tsg68-api/middleware"
	"github.com/jeremyawarren15/tsg68-api/structs"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func AddUpdateResponse(app *pocketbase.PocketBase, e *core.ServeEvent) {
	e.Router.AddRoute(echo.Route{
		Method: http.MethodPost,
		Path:   "/events/:slug/response",
		Handler: func(c echo.Context) error {
			user, _ := c.Get(apis.ContextAuthRecordKey).(*models.Record)

			var response structs.EventResponse
			err := c.Bind(&response)
			if err != nil {
				return c.String(http.StatusBadRequest, "bad request")
			}

			event, err := app.Dao().FindFirstRecordByData("events", "slug", c.PathParam("slug"))
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
			middleware.RequireValidatedRecordAuth(),
		},
	})
}
