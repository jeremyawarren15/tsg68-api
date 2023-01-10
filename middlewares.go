package main

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/models"
)

func RequireValidatedRecordAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			record, _ := c.Get(apis.ContextAuthRecordKey).(*models.Record)
			if record == nil || !record.Verified() {
				return apis.NewUnauthorizedError("The request requires valid record authorization token to be set.", nil)
			}

			return next(c)
		}
	}
}
