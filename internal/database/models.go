// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package database

import (
	"time"
)

type Product struct {
	ID          int64
	Title       string
	Description string
	Created     time.Time
}
