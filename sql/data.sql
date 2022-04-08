INSERT INTO users (name, nick, email, password)
VALUES
("User 1", "user_1", "user1@email.com", "$2a$10$x1nqiNp0yv9gDcFSaDCZrOHYn76FB5FRfiGHnEAthGpFxrnAnUjcO"),
("User 2", "user_2", "user2@email.com", "$2a$10$x1nqiNp0yv9gDcFSaDCZrOHYn76FB5FRfiGHnEAthGpFxrnAnUjcO"),
("User 3", "user_3", "user3@email.com", "$2a$10$x1nqiNp0yv9gDcFSaDCZrOHYn76FB5FRfiGHnEAthGpFxrnAnUjcO"),
("User 4", "user_4", "user4@email.com", "$2a$10$x1nqiNp0yv9gDcFSaDCZrOHYn76FB5FRfiGHnEAthGpFxrnAnUjcO"),
("User 5", "user_5", "user5@email.com", "$2a$10$x1nqiNp0yv9gDcFSaDCZrOHYn76FB5FRfiGHnEAthGpFxrnAnUjcO");


INSERT INTO followers (user_id, follower_id)
VALUES
(1, 2),
(1, 3),
(3, 1),
(2, 4),
(4, 2);


INSERT INTO posts (title, content, author_id)
VALUES
("user 1 post", "this is a post from user 1! Nice...", 1),
("user 2 post", "this is a post from user 2! Uhu...", 2),
("user 3 post", "this is a post from user 3! Lets go...", 3);