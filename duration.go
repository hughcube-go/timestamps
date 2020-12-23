package timestamps

import (
	"database/sql"
	"time"
)

type HasDuration interface {
	GetStartedAt() time.Time
	GetEndedAt() time.Time
	SetStartedAt(now time.Time)
	SetEndedAt(now time.Time)

	SetStartedAtDate(date string) error
	SetEndedAtDate(date string) error
	GetStartedAtDate() string
	GetEndedAtDate() string

	SetStartedAtDateWithZone(date string) error
	SetEndedAtDateWithZone(date string) error
	GetStartedAtDateWithZone() string
	GetEndedAtDateWithZone() string

	SetStartedAtFineDate(date string) error
	SetEndedAtFineDate(date string) error
	GetStartedAtFineDate() string
	GetEndedAtFineDate() string

	SetStartedAtFineWithZone(date string) error
	SetEndedAtFineWithZone(date string) error
	GetStartedAtFineWithZone() string
	GetEndedAtFineWithZone() string

	GetStartedAtSqlTime() sql.NullTime
	GetEndedAtSqlTime() sql.NullTime
	SetStartedAtSqlTime(now sql.NullTime)
	SetEndedAtSqlTime(now sql.NullTime)

	SetStartedAtRFC3339Date(date string) error
	SetEndedAtRFC3339Date(date string) error
	GetStartedAtRFC3339Date() string
	GetEndedAtRFC3339Date() string

	SetStartedAtRFC3339NanoDate(date string) error
	SetEndedAtRFC3339NanoDate(date string) error
	GetStartedAtRFC3339NanoDate() string
	GetEndedAtRFC3339NanoDate() string

	SetStartedAtWithLayout(layout string, date string) error
	SetEndedAtWithLayout(layout string, date string) error
	GetStartedAtWithLayout(layout string) string
	GetEndedAtWithLayout(layout string) string

	LoadDefaultTimestamps()

	TouchStartTimestamps()
	TouchEndTimestamps()

	IsStarted() bool
	IsEnded() bool

	GetDurationLength() int64
}

type Duration struct {
	StartedAt sql.NullTime
	EndedAt   sql.NullTime
}

////////////////////////////////////////////////
////////////////////////////////////////////////
////////////////////////////////////////////////
func (t *Duration) GetStartedAt() time.Time {
	return t.StartedAt.Time
}

func (t *Duration) GetEndedAt() time.Time {
	return t.EndedAt.Time
}

func (t *Duration) SetStartedAt(now time.Time) {
	t.StartedAt = Time(now)
}

func (t *Duration) SetEndedAt(now time.Time) {
	t.EndedAt = Time(now)
}

////////////////////////////////////////////////
////////////////////////////////////////////////
////////////////////////////////////////////////
func (t *Duration) SetStartedAtDate(date string) error {
	if now, err := Parse(date); err != nil {
		return err
	} else {
		t.StartedAt = now
		return nil
	}
}

func (t *Duration) SetEndedAtDate(date string) error {
	if now, err := Parse(date); err != nil {
		return err
	} else {
		t.EndedAt = now
		return nil
	}
}

func (t *Duration) GetStartedAtDate() string {
	return Format(t.StartedAt)
}

func (t *Duration) GetEndedAtDate() string {
	return Format(t.EndedAt)
}

////////////////////////////////////////////////
////////////////////////////////////////////////
////////////////////////////////////////////////
func (t *Duration) SetStartedAtDateWithZone(date string) error {
	if now, err := ParseWithZone(date); err != nil {
		return err
	} else {
		t.StartedAt = now
		return nil
	}
}

func (t *Duration) SetEndedAtDateWithZone(date string) error {
	if now, err := ParseWithZone(date); err != nil {
		return err
	} else {
		t.StartedAt = now
		return nil
	}
}

func (t *Duration) GetStartedAtDateWithZone() string {
	return FormatWithZone(t.EndedAt)
}

func (t *Duration) GetEndedAtDateWithZone() string {
	return FormatWithZone(t.EndedAt)
}

////////////////////////////////////////////////
////////////////////////////////////////////////
////////////////////////////////////////////////
func (t *Duration) SetStartedAtFineDate(date string) error {
	if now, err := ParseFine(date); err != nil {
		return err
	} else {
		t.StartedAt = now
		return nil
	}
}

func (t *Duration) SetEndedAtFineDate(date string) error {
	if now, err := ParseFine(date); err != nil {
		return err
	} else {
		t.StartedAt = now
		return nil
	}
}

func (t *Duration) GetStartedAtFineDate() string {
	return FormatFine(t.EndedAt)
}

func (t *Duration) GetEndedAtFineDate() string {
	return FormatFine(t.EndedAt)
}

////////////////////////////////////////////////
////////////////////////////////////////////////
////////////////////////////////////////////////
func (t *Duration) SetStartedAtFineWithZone(date string) error {
	if now, err := ParseFineWithZone(date); err != nil {
		return err
	} else {
		t.StartedAt = now
		return nil
	}
}

func (t *Duration) SetEndedAtFineWithZone(date string) error {
	if now, err := ParseFineWithZone(date); err != nil {
		return err
	} else {
		t.StartedAt = now
		return nil
	}
}

func (t *Duration) GetStartedAtFineWithZone() string {
	return FormatFineWithZone(t.EndedAt)
}

func (t *Duration) GetEndedAtFineWithZone() string {
	return FormatFineWithZone(t.EndedAt)
}

////////////////////////////////////////////////
////////////////////////////////////////////////
////////////////////////////////////////////////
func (t *Duration) SetStartedAtRFC3339Date(date string) error {
	if now, err := ParseRFC3339(date); err != nil {
		return err
	} else {
		t.StartedAt = now
		return nil
	}
}

func (t *Duration) SetEndedAtRFC3339Date(date string) error {
	if now, err := ParseRFC3339(date); err != nil {
		return err
	} else {
		t.StartedAt = now
		return nil
	}
}

func (t *Duration) GetStartedAtRFC3339Date() string {
	return FormatRFC3339(t.EndedAt)
}

func (t *Duration) GetEndedAtRFC3339Date() string {
	return FormatRFC3339(t.EndedAt)
}

////////////////////////////////////////////////
////////////////////////////////////////////////
////////////////////////////////////////////////
func (t *Duration) SetStartedAtRFC3339NanoDate(date string) error {
	if now, err := ParseRFC3339Nano(date); err != nil {
		return err
	} else {
		t.StartedAt = now
		return nil
	}
}

func (t *Duration) SetEndedAtRFC3339NanoDate(date string) error {
	if now, err := ParseRFC3339Nano(date); err != nil {
		return err
	} else {
		t.StartedAt = now
		return nil
	}
}

func (t *Duration) GetStartedAtRFC3339NanoDate() string {
	return FormatRFC3339Nano(t.EndedAt)
}

func (t *Duration) GetEndedAtRFC3339NanoDate() string {
	return FormatRFC3339Nano(t.EndedAt)
}

////////////////////////////////////////////////
////////////////////////////////////////////////
////////////////////////////////////////////////
func (t *Duration) SetStartedAtWithLayout(layout string, date string) error {
	if now, err := ParseWithLayout(layout, date); err != nil {
		return err
	} else {
		t.StartedAt = now
		return nil
	}
}

func (t *Duration) SetEndedAtWithLayout(layout string, date string) error {
	if now, err := ParseWithLayout(layout, date); err != nil {
		return err
	} else {
		t.EndedAt = now
		return nil
	}
}

func (t *Duration) GetStartedAtWithLayout(layout string) string {
	return FormatWithLayout(layout, t.StartedAt)
}

func (t *Duration) GetEndedAtWithLayout(layout string) string {
	return FormatWithLayout(layout, t.EndedAt)
}

////////////////////////////////////////////////
////////////////////////////////////////////////
////////////////////////////////////////////////
func (t *Duration) GetStartedAtSqlTime() sql.NullTime {
	return t.StartedAt
}

func (t *Duration) GetEndedAtSqlTime() sql.NullTime {
	return t.EndedAt
}

func (t *Duration) SetStartedAtSqlTime(now sql.NullTime) {
	t.StartedAt = now
}

func (t *Duration) SetEndedAtSqlTime(now sql.NullTime) {
	t.EndedAt = now
}

//////////////////////////////////////////////
//////////////////////////////////////////////
//////////////////////////////////////////////
func (t *Duration) LoadDefaultTimestamps() {
	if !t.StartedAt.Valid {
		t.StartedAt = Now()
	}
}

func (t *Duration) TouchStartTimestamps() {
	t.StartedAt = Now()
}

func (t *Duration) TouchEndTimestamps() {
	t.EndedAt = Now()
}

func (t *Duration) IsStarted() bool {
	return t.StartedAt.Valid && t.StartedAt.Time.Before(time.Now())
}

func (t *Duration) IsEnded() bool {
	return t.EndedAt.Valid && t.EndedAt.Time.After(time.Now())
}

func (t *Duration) GetDurationLength() int64 {
	if !t.StartedAt.Valid || !t.EndedAt.Valid {
		return 0
	}
	return t.EndedAt.Time.UnixNano() - t.StartedAt.Time.UnixNano()
}
