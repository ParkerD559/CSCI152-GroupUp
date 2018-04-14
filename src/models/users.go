package models

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type user struct {
	Name           string
	WsConn         *websocket.Conn
	Online         bool
	Token          string
	CurrentGroups  []*group
	FavoriteGroups []*group
	RecentGroups   []*group
}

var userMutex = &sync.Mutex{}

// NewUser makes a new user with the given username, token, and IP
func NewUser(username string) (userToken string) {
	userToken, err := generateRandomString(32)
	if err != nil {
		log.Println(err.Error())
	}
	u := user{
		Name:  username,
		Token: userToken,
	}
	// Thread safe users modification.
	userMutex.Lock()
	users[userToken] = &u
	userMutex.Unlock()
	return
}

// RemoveUser assumes existense of user, and removes it
func RemoveUser(token string) {
	userMutex.Lock()
	delete(users, token)
	userMutex.Unlock()
	return
}

// UserExists checks if a user currently (ie connection is alive)
func UserExists(token string) bool {
	if _, exists := users[token]; exists {
		return true
	} else {
		return false
	}
}

// GetUserToken returns the token of a given user, assumes it exists
func GetUserToken(ip string) (token string) {
	token = users[ip].Token
	return
}
