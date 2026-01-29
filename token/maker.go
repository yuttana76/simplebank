package token

import "time"

// Make is an interface for manage tokens
type Maker interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(username string, role string, duration time.Duration, tokenType TokenType) (string, *Payload, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string, tokenType TokenType) (*Payload, error)
}
