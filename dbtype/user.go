package dbtype

import "database/sql"

type User struct {
	ID      sql.NullInt64
	Name    sql.NullString
	Email   sql.NullString
	Credits sql.NullInt64
}

type NotifyUsers struct {
	UserID         sql.NullString
	Name           sql.NullString
	Email          sql.NullString
	PreferredTime  sql.NullString
	Surprise       sql.NullString
	NotificationID sql.NullInt64
	FCMID          sql.NullString
}