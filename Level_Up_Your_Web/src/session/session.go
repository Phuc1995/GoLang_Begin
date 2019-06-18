package session

import (
	"fmt"
	"handler"
	"net/http"
	"time"
)

type Session struct {
	ID string
	UserID string
	Expiry time.Time
}

const(
	//Keep users logged in for 3 days
	SessionLength      = 24 * 3 * time.Hour
	SessionCookieName = "GophrSession"
	SessionIDLength = 20
)

func NewSession(w http.ResponseWriter) *Session {
	expiry := time.Now().Add(SessionIDLength)

	session := &Session{
		ID:     handler.GenerateID("sess", SessionIDLength),
		Expiry: expiry,
	}

	cookie := http.Cookie{
		Name: SessionCookieName,
		Value: session.ID,
		Expires: session.Expiry,
	}
	fmt.Println(cookie)
	http.SetCookie(w, &cookie)
	return session
}