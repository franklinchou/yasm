package main

import "time"
import (
	"./utils"
	"math/rand"
)


//*********************************************************
// Application defaults
//*********************************************************
const SessionDefaultTimeout = 120



func _createSessionId() string {
	source := rand.NewSource(time.Now().UnixNano())
	return utils.RandomString(32, source)
}

func createSession(key, token string) string {
	expiration := time.Duration(SessionDefaultTimeout)
	sessionId := _createSessionId()

	return sessionId
}