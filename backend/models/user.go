package models

type User struct {
	Id       int
	Username string `json:"username"`
	Password string `json:"password"`
}

/*

rooms
CREATE TABLE rooms (
    roomId SERIAL PRIMARY KEY,
    roomName VARCHAR(255) NOT NULL,
);

messages

CREATE TABLE messages (
    messageId SERIAL PRIMARY KEY,
    messageText VARCHAR(1000) NOT NULL,
    roomId INT NOT NULL,
    userId INT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (roomId) REFERENCES rooms(roomId),
    FOREIGN KEY (userId) REFERENCES users(id)
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

*/
