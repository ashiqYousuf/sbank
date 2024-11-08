package token

import "time"

// Maker is an interface for managing tokens
// JWTMaker & PasetoMaker will implement this interface
// Implies we can easily switch b/w our JWT & PASETO tokens
type Maker interface {
	// creates a new token for a specific username & duration
	CreateToken(username string, duration time.Duration) (string, error)

	// checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
