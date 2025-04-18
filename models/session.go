package models

import "database/sql"

type Session struct {
	ID        int
	UserID    int
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

// Create will create a new session for the user provided. The session token will be returned as the Token field on the session type, but only hashed session token is stored in database
func (ss *SessionService) Create(userID int) (*Session, error) {
	// create session token
	return nil, nil
}
