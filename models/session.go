package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"

	"github.com/cheeyeo/lenslocked/rand"
)

const (
	// min number of bytes to use for each session token
	MinBytesPerToken = 32
)

type Session struct {
	ID        int
	UserID    int
	Token     string
	TokenHash string
}

type SessionService struct {
	DB            *sql.DB
	BytesPerToken int
}

// Create will create a new session for the user provided. The session token will be returned as the Token field on the session type, but only hashed session token is stored in database
func (ss *SessionService) Create(userID int) (*Session, error) {
	bytesPerToken := ss.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}
	// create session token
	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("create sessionservice : %w", err)
	}
	session := Session{
		UserID:    userID,
		Token:     token,
		TokenHash: ss.hash(token),
	}
	row := ss.DB.QueryRow(`
	  UPDATE sessions
	  SET token_hash=$2
	  WHERE user_id=$1
	  RETURNING id;
	`, session.UserID, session.TokenHash)

	err = row.Scan(&session.ID)
	if err == sql.ErrNoRows {
		// no session exists so create session for user
		// Insert into DB
		row := ss.DB.QueryRow(`
		  INSERT INTO sessions(user_id, token_hash)
		  VALUES ($1, $2)
		  RETURNING id;
		`, session.UserID, session.TokenHash)
		err = row.Scan(&session.ID)
	}

	if err != nil {
		return nil, fmt.Errorf("session create: %w", err)
	}
	return &session, nil
}

func (ss *SessionService) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}

func (ss *SessionService) User(token string) (*User, error) {
	tokenHash := ss.hash(token)
	var user User
	row := ss.DB.QueryRow(`
	  SELECT user_id
	  FROM sessions
	  WHERE token_hash=$1
	`, tokenHash)

	err := row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("sessionservice user: %w", err)
	}

	row = ss.DB.QueryRow(`
	  SELECT email, password_hash
	  FROM users WHERE id=$1
	`, user.ID)
	err = row.Scan(&user.Email, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("sessionservice user: %w", err)
	}

	return &user, nil
}

func (ss *SessionService) Delete(token string) error {
	tokenHash := ss.hash(token)
	_, err := ss.DB.Exec(`
	  DELETE FROM sessions
	  WHERE token_hash=$1;
	`, tokenHash)
	if err != nil {
		return fmt.Errorf("sessionservice delete: %w", err)
	}
	return nil
}
