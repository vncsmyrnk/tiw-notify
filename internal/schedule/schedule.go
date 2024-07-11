// vim: noexpandtab

package schedule

import (
	"time"
)

type Job struct {
	time time.Time
	task func()
}
