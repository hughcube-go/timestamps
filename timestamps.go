package timestamps

import (
	"database/sql"
	"time"
)

type HasTimestamps interface {
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetDeletedAt() time.Time
	SetCreatedAt(now time.Time)
	SetUpdatedAt(now time.Time)
	SetDeletedAt(now time.Time)

	GetCreatedAtSqlTime() sql.NullTime
	GetUpdatedAtSqlTime() sql.NullTime
	GetDeletedAtSqlTime() sql.NullTime
	SetCreatedAtSqlTime(now sql.NullTime)
	SetUpdatedAtSqlTime(now sql.NullTime)
	SetDeletedAtSqlTime(now sql.NullTime)

	SetCreatedAtDate(date string) error
	SetUpdatedAtDate(date string) error
	SetDeletedAtDate(date string) error
	GetCreatedAtDate() string
	GetUpdatedAtDate() string
	GetDeletedAtDate() string

	SetCreatedAtDateWithZone(date string) error
	SetUpdatedAtDateWithZone(date string) error
	SetDeletedAtDateWithZone(date string) error
	GetCreatedAtDateWithZone() string
	GetUpdatedAtDateWithZone() string
	GetDeletedAtDateWithZone() string

	SetCreatedAtFineDate(date string) error
	SetUpdatedAtFineDate(date string) error
	SetDeletedAtFineDate(date string) error
	GetCreatedAtFineDate() string
	GetUpdatedAtFineDate() string
	GetDeletedAtFineDate() string

	SetCreatedAtFineDateWithZone(date string) error
	SetUpdatedAtFineDateWithZone(date string) error
	SetDeletedAtFineDateWithZone(date string) error
	GetCreatedAtFineDateWithZone() string
	GetUpdatedAtFineDateWithZone() string
	GetDeletedAtFineDateWithZone() string

	SetCreatedAtRFC3339Date(date string) error
	SetUpdatedAtRFC3339Date(date string) error
	SetDeletedAtRFC3339Date(date string) error
	GetCreatedAtRFC3339Date() string
	GetUpdatedAtRFC3339Date() string
	GetDeletedAtRFC3339Date() string

	SetCreatedAtRFC3339NanoDate(date string) error
	SetUpdatedAtRFC3339NanoDate(date string) error
	SetDeletedAtRFC3339NanoDate(date string) error
	GetCreatedAtRFC3339NanoDate() string
	GetUpdatedAtRFC3339NanoDate() string
	GetDeletedAtRFC3339NanoDate() string

	SetCreatedAtWithLayout(layout string, date string) error
	SetUpdatedAtWithLayout(layout string, date string) error
	SetDeletedAtWithLayout(layout string, date string) error
	GetCreatedAtWithLayout(layout string) string
	GetUpdatedAtWithLayout(layout string) string
	GetDeletedAtWithLayout(layout string) string

	IsDelete() bool

	LoadDefaultTimestamps()

	TouchCreateTimestamps()
	TouchUpdateTimestamps()
	TouchDeleteTimestamps()
}

type Timestamps struct {
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}

//////////////////////////////////////////////
//////////////////////////////////////////////
//////////////////////////////////////////////
func (t *Timestamps) GetCreatedAt() time.Time {
	return t.CreatedAt.Time
}

func (t *Timestamps) GetUpdatedAt() time.Time {
	return t.UpdatedAt.Time
}

func (t *Timestamps) GetDeletedAt() time.Time {
	return t.DeletedAt.Time
}

func (t *Timestamps) SetCreatedAt(now time.Time) {
	t.CreatedAt = Time(now)
}

func (t *Timestamps) SetUpdatedAt(now time.Time) {
	t.UpdatedAt = Time(now)
}

func (t *Timestamps) SetDeletedAt(now time.Time) {
	t.DeletedAt = Time(now)
}

///////////////////////////////////////////////////////
///////////////////////////////////////////////////////
///////////////////////////////////////////////////////
func (t *Timestamps) GetCreatedAtSqlTime() sql.NullTime {
	return t.CreatedAt
}

func (t *Timestamps) GetUpdatedAtSqlTime() sql.NullTime {
	return t.UpdatedAt
}

func (t *Timestamps) GetDeletedAtSqlTime() sql.NullTime {
	return t.DeletedAt
}

func (t *Timestamps) SetCreatedAtSqlTime(now sql.NullTime) {
	t.CreatedAt = now
}

func (t *Timestamps) SetUpdatedAtSqlTime(now sql.NullTime) {
	t.UpdatedAt = now
}

func (t *Timestamps) SetDeletedAtSqlTime(now sql.NullTime) {
	t.DeletedAt = now
}

///////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////
func (t *Timestamps) SetCreatedAtDate(date string) error {
	if now, err := Parse(date); err != nil {
		return err
	} else {
		t.CreatedAt = now
		return nil
	}
}

func (t *Timestamps) SetUpdatedAtDate(date string) error {
	if now, err := Parse(date); err != nil {
		return err
	} else {
		t.UpdatedAt = now
		return nil
	}
}

func (t *Timestamps) SetDeletedAtDate(date string) error {
	if now, err := Parse(date); err != nil {
		return err
	} else {
		t.DeletedAt = now
		return nil
	}
}

func (t *Timestamps) GetCreatedAtDate() string {
	return Format(t.CreatedAt)
}

func (t *Timestamps) GetUpdatedAtDate() string {
	return Format(t.UpdatedAt)
}

func (t *Timestamps) GetDeletedAtDate() string {
	return Format(t.DeletedAt)
}

///////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////
func (t *Timestamps) SetCreatedAtDateWithZone(date string) error {
	if now, err := ParseWithZone(date); err != nil {
		return err
	} else {
		t.CreatedAt = now
		return nil
	}
}

func (t *Timestamps) SetUpdatedAtDateWithZone(date string) error {
	if now, err := ParseWithZone(date); err != nil {
		return err
	} else {
		t.UpdatedAt = now
		return nil
	}
}

func (t *Timestamps) SetDeletedAtDateWithZone(date string) error {
	if now, err := ParseWithZone(date); err != nil {
		return err
	} else {
		t.DeletedAt = now
		return nil
	}
}

func (t *Timestamps) GetCreatedAtDateWithZone() string {
	return FormatWithZone(t.CreatedAt)
}

func (t *Timestamps) GetUpdatedAtDateWithZone() string {
	return FormatWithZone(t.UpdatedAt)
}

func (t *Timestamps) GetDeletedAtDateWithZone() string {
	return FormatWithZone(t.DeletedAt)
}

///////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////
func (t *Timestamps) SetCreatedAtFineDate(date string) error {
	if now, err := ParseFine(date); err != nil {
		return err
	} else {
		t.CreatedAt = now
		return nil
	}
}

func (t *Timestamps) SetUpdatedAtFineDate(date string) error {
	if now, err := ParseFine(date); err != nil {
		return err
	} else {
		t.UpdatedAt = now
		return nil
	}
}

func (t *Timestamps) SetDeletedAtFineDate(date string) error {
	if now, err := ParseFine(date); err != nil {
		return err
	} else {
		t.DeletedAt = now
		return nil
	}
}

func (t *Timestamps) GetCreatedAtFineDate() string {
	return FormatFine(t.CreatedAt)
}

func (t *Timestamps) GetUpdatedAtFineDate() string {
	return FormatFine(t.UpdatedAt)
}

func (t *Timestamps) GetDeletedAtFineDate() string {
	return FormatFine(t.DeletedAt)
}

///////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////
func (t *Timestamps) SetCreatedAtFineDateWithZone(date string) error {
	if now, err := ParseFineWithZone(date); err != nil {
		return err
	} else {
		t.CreatedAt = now
		return nil
	}
}

func (t *Timestamps) SetUpdatedAtFineDateWithZone(date string) error {
	if now, err := ParseFineWithZone(date); err != nil {
		return err
	} else {
		t.UpdatedAt = now
		return nil
	}
}

func (t *Timestamps) SetDeletedAtFineDateWithZone(date string) error {
	if now, err := ParseFineWithZone(date); err != nil {
		return err
	} else {
		t.DeletedAt = now
		return nil
	}
}

func (t *Timestamps) GetCreatedAtFineDateWithZone() string {
	return FormatFineWithZone(t.CreatedAt)
}

func (t *Timestamps) GetUpdatedAtFineDateWithZone() string {
	return FormatFineWithZone(t.UpdatedAt)
}

func (t *Timestamps) GetDeletedAtFineDateWithZone() string {
	return FormatFineWithZone(t.DeletedAt)
}

///////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////
func (t *Timestamps) SetCreatedAtRFC3339Date(date string) error {
	if now, err := ParseRFC3339(date); err != nil {
		return err
	} else {
		t.CreatedAt = now
		return nil
	}
}

func (t *Timestamps) SetUpdatedAtRFC3339Date(date string) error {
	if now, err := ParseRFC3339(date); err != nil {
		return err
	} else {
		t.UpdatedAt = now
		return nil
	}
}

func (t *Timestamps) SetDeletedAtRFC3339Date(date string) error {
	if now, err := ParseRFC3339(date); err != nil {
		return err
	} else {
		t.DeletedAt = now
		return nil
	}
}

func (t *Timestamps) GetCreatedAtRFC3339Date() string {
	return FormatRFC3339(t.CreatedAt)
}

func (t *Timestamps) GetUpdatedAtRFC3339Date() string {
	return FormatRFC3339(t.UpdatedAt)
}

func (t *Timestamps) GetDeletedAtRFC3339Date() string {
	return FormatRFC3339(t.DeletedAt)
}

///////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////
func (t *Timestamps) SetCreatedAtRFC3339NanoDate(date string) error {
	if now, err := ParseRFC3339Nano(date); err != nil {
		return err
	} else {
		t.CreatedAt = now
		return nil
	}
}

func (t *Timestamps) SetUpdatedAtRFC3339NanoDate(date string) error {
	if now, err := ParseRFC3339Nano(date); err != nil {
		return err
	} else {
		t.UpdatedAt = now
		return nil
	}
}

func (t *Timestamps) SetDeletedAtRFC3339NanoDate(date string) error {
	if now, err := ParseRFC3339Nano(date); err != nil {
		return err
	} else {
		t.DeletedAt = now
		return nil
	}
}

func (t *Timestamps) GetCreatedAtRFC3339NanoDate() string {
	return FormatRFC3339Nano(t.CreatedAt)
}

func (t *Timestamps) GetUpdatedAtRFC3339NanoDate() string {
	return FormatRFC3339Nano(t.UpdatedAt)
}

func (t *Timestamps) GetDeletedAtRFC3339NanoDate() string {
	return FormatRFC3339Nano(t.DeletedAt)
}

//////////////////////////////////////////////
//////////////////////////////////////////////
//////////////////////////////////////////////
func (t *Timestamps) SetCreatedAtWithLayout(layout string, date string) error {
	if now, err := ParseWithLayout(layout, date); err != nil {
		return err
	} else {
		t.CreatedAt = now
		return nil
	}
}
func (t *Timestamps) SetUpdatedAtWithLayout(layout string, date string) error {
	if now, err := ParseWithLayout(layout, date); err != nil {
		return err
	} else {
		t.UpdatedAt = now
		return nil
	}
}
func (t *Timestamps) SetDeletedAtWithLayout(layout string, date string) error {
	if now, err := ParseWithLayout(layout, date); err != nil {
		return err
	} else {
		t.DeletedAt = now
		return nil
	}
}
func (t *Timestamps) GetCreatedAtWithLayout(layout string) string {
	return FormatWithLayout(layout, t.CreatedAt)
}
func (t *Timestamps) GetUpdatedAtWithLayout(layout string) string {
	return FormatWithLayout(layout, t.UpdatedAt)
}
func (t *Timestamps) GetDeletedAtWithLayout(layout string) string {
	return FormatWithLayout(layout, t.DeletedAt)
}

//////////////////////////////////////////////
//////////////////////////////////////////////
//////////////////////////////////////////////
func (t *Timestamps) LoadDefaultTimestamps() {
	if !t.CreatedAt.Valid {
		t.CreatedAt = Now()
	}

	if !t.UpdatedAt.Valid {
		t.UpdatedAt = Now()
	}
}

func (t *Timestamps) IsDelete() bool {
	return t.DeletedAt.Valid
}

func (t *Timestamps) TouchCreateTimestamps() {
	t.CreatedAt = Now()
}

func (t *Timestamps) TouchUpdateTimestamps() {
	t.UpdatedAt = Now()
}

func (t *Timestamps) TouchDeleteTimestamps() {
	t.DeletedAt = Now()
}
