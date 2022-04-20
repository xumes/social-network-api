# social-network-api

social-network-api is an API built in Go for emulate a social network, giving users the ability to register, login, create posts, follow other users to see their posts and like or unlike them.

## Installation

Fork or Clone this repository.
Rename the `.env.sample` file to `.env` and fill out the variables using your environment.
This application requires a connection to a MySQL database.
You will have to generate a `SECRET_KEY` but it can be any string you want.

## Running the API

```bash
go run main.go
```

## Endpoints

```Go
POST "/login"
POST  "/users"
GET  "/users"
GET  "/users/{id}"
PUT  "/users/{id}"
DELETE  "/users/{id}"
POST "/users/{userId}/follow"
POST "/users/{userId}/unfollow"
GET "/users/{userId}/followers"
GET "/users/{userId}/following"
POST "/users/{userId}/change-password"
POST "/posts"
GET "/posts"
GET "/posts/{postId}"
PUT "/posts/{postId}"
DELETE "/posts/{postId}"
GET "/users/{userId}/posts"
POST "/posts/{postId}/like"
POST "/posts/{postId}/unlike"
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)

## Author

Reginaldo Santos
