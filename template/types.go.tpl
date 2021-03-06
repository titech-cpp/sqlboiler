package models

import (
    "time"
    "errors"
    "database/sql"
)

type timeTime = time.Time
type nullBool = sql.NullBool
type nullString = sql.NullString
type nullInt32 = sql.NullInt32
type nullInt64 = sql.NullInt64
type nullTime = sql.NullTime
var newError = errors.New