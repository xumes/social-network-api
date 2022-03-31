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
	{
		Uri:            "/users/{userId}/follow",
		Method:         http.MethodPost,
		Function:       controllers.FollowUserById,
		IsAuthRequired: true,
	},
	{
		Uri:            "/users/{userId}/unfollow",
		Method:         http.MethodPost,
		Function:       controllers.UnfollowUserById,
		IsAuthRequired: true,
	},
	{
		Uri:            "/users/{userId}/followers",
		Method:         http.MethodGet,
		Function:       controllers.GetFollowersById,
		IsAuthRequired: true,
	},
	{
		Uri:            "/users/{userId}/following",
		Method:         http.MethodGet,
		Function:       controllers.GetFollowingById,
		IsAuthRequired: true,
	},
}
