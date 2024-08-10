package dbtype

import "database/sql"

type User struct {
	ID      sql.NullInt64
	Name    sql.NullString
	Email   sql.NullString
	Credits sql.NullInt64
}

type NotifyUsers struct {
	UserID         string	
	Name           string
	Email          string
	PreferredTime  string
	Surprise       int
	NotificationID string
	FCMID          string
}