package aural

// UserID is a type alias for UserIDs provided by the Facebook API
type UserID uint64

// User is a struct for User objects provided by the Facebook API
type User struct {
	ID UserID `json:"id"`
}
