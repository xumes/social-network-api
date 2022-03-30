package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		Uri:            "/users",
		Method:         http.MethodPost,
		Function:       controllers.AddUser,
		IsAuthRequired: false,
	},
	{
		Uri:            "/users",
		Method:         http.MethodGet,
		Function:       controllers.GetUsers,
		IsAuthRequired: true,
	},
	{
		Uri:            "/users/{userId}",
		Method:         http.MethodGet,
		Function:       controllers.GetUserById,
		IsAuthRequired: true,
	},
	{
		Uri:            "/users/{userId}",
		Method:         http.MethodPut,
		Function:       controllers.UpdateUserById,
		IsAuthRequired: true,
	},
	{
		Uri:            "/users/{userId}",
		Method:         http.MethodDelete,
		Function:       controllers.RemoveUserById,
		IsAuthRequired: true,
	},
}
