package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `[
			{
				"id": "_pb_users_auth_",
				"created": "2023-01-22 04:43:19.943Z",
				"updated": "2023-01-22 04:44:14.778Z",
				"name": "users",
				"type": "auth",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "users_name",
						"name": "name",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "users_avatar",
						"name": "avatar",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [
								"image/jpg",
								"image/jpeg",
								"image/png",
								"image/svg+xml",
								"image/gif"
							],
							"thumbs": null
						}
					},
					{
						"system": false,
						"id": "xmuvz4ve",
						"name": "phone_number",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": "^\\(?([0-9]{3})\\)?[-. ]?([0-9]{3})[-. ]?([0-9]{4})$"
						}
					}
				],
				"listRule": "id = @request.auth.id",
				"viewRule": "id = @request.auth.id",
				"createRule": null,
				"updateRule": "id = @request.auth.id",
				"deleteRule": "id = @request.auth.id",
				"options": {
					"allowEmailAuth": true,
					"allowOAuth2Auth": false,
					"allowUsernameAuth": true,
					"exceptEmailDomains": null,
					"manageRule": null,
					"minPasswordLength": 8,
					"onlyEmailDomains": null,
					"requireEmail": false
				}
			},
			{
				"id": "jovjo7m6f629yvk",
				"created": "2023-01-22 04:44:14.779Z",
				"updated": "2023-01-22 04:44:14.779Z",
				"name": "events",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "k2htyvqt",
						"name": "title",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "kj7bmlkj",
						"name": "slug",
						"type": "text",
						"required": true,
						"unique": true,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "u2pztjn6",
						"name": "description",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "dlvxiobi",
						"name": "body",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "dyd2royr",
						"name": "start",
						"type": "date",
						"required": true,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "n9yxt2zw",
						"name": "end",
						"type": "date",
						"required": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					}
				],
				"listRule": "@request.auth.verified = true",
				"viewRule": "@request.auth.verified = true",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "h73uugxedsps7c1",
				"created": "2023-01-22 04:44:14.781Z",
				"updated": "2023-01-22 04:44:14.781Z",
				"name": "responses",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "epqph65f",
						"name": "user",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": true
						}
					},
					{
						"system": false,
						"id": "zcrpqj2k",
						"name": "event",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"collectionId": "jovjo7m6f629yvk",
							"cascadeDelete": true
						}
					},
					{
						"system": false,
						"id": "f2sfenj6",
						"name": "response",
						"type": "select",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"pending",
								"attending",
								"declined"
							]
						}
					}
				],
				"listRule": "@request.auth.id = user.id",
				"viewRule": "@request.auth.id = user.id",
				"createRule": "@request.auth.verified = true",
				"updateRule": "@request.auth.id = user.id",
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "53cglz61bvuscz5",
				"created": "2023-01-22 04:44:14.782Z",
				"updated": "2023-01-22 04:44:14.782Z",
				"name": "updates",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "4p1ykk3d",
						"name": "title",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "fancumag",
						"name": "slug",
						"type": "text",
						"required": true,
						"unique": true,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "fi6y6wv8",
						"name": "description",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "jb5tpdmu",
						"name": "body",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "vq5g8oclfzj2e8q",
				"created": "2023-01-22 04:44:14.783Z",
				"updated": "2023-01-22 04:44:14.783Z",
				"name": "faqs",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "mafz8wvj",
						"name": "title",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "asl3sp9r",
						"name": "body",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"listRule": "",
				"viewRule": "",
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "cuhgk0tdswo9nqx",
				"created": "2023-01-22 04:44:14.784Z",
				"updated": "2023-01-22 04:44:14.784Z",
				"name": "cadets",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "1cwicko0",
						"name": "name",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "42jiyluj",
						"name": "birth_date",
						"type": "date",
						"required": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "nxcahffn",
						"name": "father",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": true
						}
					}
				],
				"listRule": null,
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		return nil
	})
}
