// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package author

import (
	"database/sql"
)

type Author struct {
	ID   int64
	Name string
	Bio  sql.NullString
}
