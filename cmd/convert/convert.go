package convert

import (
	"database/sql"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func StringToSQLNullString(s *string) sql.NullString {
	var nullString sql.NullString
	if s != nil {
		nullString.String = *s
		nullString.Valid = true
	} else {
		nullString.Valid = false
	}
	return nullString
}

func StringToPGText(s *string) pgtype.Text {
	var pgText pgtype.Text

	if s != nil {
		pgText.String = *s
		pgText.Valid = true
	}

	if pgText.String == "" {
		pgText.Valid = false
	}

	return pgText
}

func TimeToPGTimestamp(t *time.Time) pgtype.Timestamptz {
	var pgTimestamp pgtype.Timestamptz

	if t != nil {
		pgTimestamp.Time = *t
	} else {
		pgTimestamp.Valid = false
	}

	return pgTimestamp
}
