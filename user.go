package aural

// UserID is a type alias for UserIDs provided by the Facebook API
type UserID string

// User is a struct for User objects provided by the Facebook API
type User struct {
	ID UserID `json:"id"`
}
