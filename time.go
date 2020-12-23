package timestamps

import (
	"database/sql"
	"time"
)

const DefaultDateLayout = "2006-01-02 15:04:05"
const DefaultDateWithZoneLayout = "2006-01-02 15:04:05 Z07:00"

const DefaultFineDateLayout = "2006-01-02 15:04:05.999999999"
const DefaultFineDateWithZoneLayout = "2006-01-02 15:04:05.999999999 Z07:00"

const DefaultRFC3339DateLayout = time.RFC3339
const DefaultRFC3339NanoDateLayout = time.RFC3339Nano

func NilTime() sql.NullTime {
	return sql.NullTime{}
}

func Time(now time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  now,
		Valid: true,
	}
}

func ZeroTime() sql.NullTime {
	return sql.NullTime{
		Time:  time.Unix(0, 0),
		Valid: false,
	}
}

func Now() sql.NullTime {
	return Time(time.Now())
}

///////////////////////////////////////////////////////
///////////////////////////////////////////////////////
///////////////////////////////////////////////////////
func ParseWithLayout(layout string, date string) (sql.NullTime, error) {
	if now, err := time.ParseInLocation(layout, date, time.Local); err == nil {
		return Time(now), nil
	} else {
		return ZeroTime(), err
	}
}

func FormatWithLayout(layout string, t sql.NullTime) string {
	if t.Valid {
		return t.Time.Format(layout)
	}
	return ""
}

/////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////
func Parse(date string) (sql.NullTime, error) {
	return ParseWithLayout(DefaultDateLayout, date)
}

func Format(t sql.NullTime) string {
	return FormatWithLayout(DefaultDateLayout, t)
}

/////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////
func ParseWithZone(date string) (sql.NullTime, error) {
	return ParseWithLayout(DefaultDateWithZoneLayout, date)
}

func FormatWithZone(t sql.NullTime) string {
	return FormatWithLayout(DefaultDateWithZoneLayout, t)
}

/////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////
func ParseFine(date string) (sql.NullTime, error) {
	return ParseWithLayout(DefaultFineDateLayout, date)
}

func FormatFine(t sql.NullTime) string {
	return FormatWithLayout(DefaultFineDateLayout, t)
}

/////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////
func ParseFineWithZone(date string) (sql.NullTime, error) {
	return ParseWithLayout(DefaultFineDateWithZoneLayout, date)
}

func FormatFineWithZone(t sql.NullTime) string {
	return FormatWithLayout(DefaultFineDateWithZoneLayout, t)
}

/////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////
func ParseRFC3339(date string) (sql.NullTime, error) {
	return ParseWithLayout(DefaultRFC3339DateLayout, date)
}

func FormatRFC3339(t sql.NullTime) string {
	return FormatWithLayout(DefaultRFC3339DateLayout, t)
}

/////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////
func ParseRFC3339Nano(date string) (sql.NullTime, error) {
	return ParseWithLayout(DefaultRFC3339NanoDateLayout, date)
}

func FormatRFC3339Nano(t sql.NullTime) string {
	return FormatWithLayout(DefaultRFC3339NanoDateLayout, t)
}
