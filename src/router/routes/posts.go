package routes

import (
	"api/src/controllers"
	"net/http"
)

var postsRoutes = []Route{
	{
		Uri:            "/posts",
		Method:         http.MethodPost,
		Function:       controllers.AddPosts,
		IsAuthRequired: true,
	},
	{
		Uri:            "/posts",
		Method:         http.MethodGet,
		Function:       controllers.GetPosts,
		IsAuthRequired: true,
	},
	{
		Uri:            "/posts/{postId}",
		Method:         http.MethodGet,
		Function:       controllers.GetPostById,
		IsAuthRequired: true,
	},
	{
		Uri:            "/posts/{postId}",
		Method:         http.MethodPut,
		Function:       controllers.UpdatePostById,
		IsAuthRequired: true,
	},
	{
		Uri:            "/posts/{postId}",
		Method:         http.MethodDelete,
		Function:       controllers.RemovePostById,
		IsAuthRequired: true,
	},
	{
		Uri:            "/users/{userId}/posts",
		Method:         http.MethodGet,
		Function:       controllers.GetPostsByUserId,
		IsAuthRequired: true,
	},
	{
		Uri:            "/posts/{postId}/like",
		Method:         http.MethodPost,
		Function:       controllers.LikePost,
		IsAuthRequired: true,
	},
	{
		Uri:            "/posts/{postId}/unlike",
		Method:         http.MethodPost,
		Function:       controllers.UnlikePost,
		IsAuthRequired: true,
	},
}
