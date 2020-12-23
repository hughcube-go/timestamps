package timestamps

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_time_DefaultDateLayout(t *testing.T) {
	a := assert.New(t)

	layouts := []string{
		DefaultDateLayout,
		DefaultDateWithZoneLayout,
		DefaultFineDateLayout,
		DefaultFineDateWithZoneLayout,
		DefaultRFC3339DateLayout,
		DefaultRFC3339NanoDateLayout,
	}

	for _, layout := range layouts {
		date := time.Now().Format(layout)
		println(date)

		now, err := time.Parse(layout, date)

		a.Nil(err)
		a.Equal(now.Format(layout), date)
	}
}
