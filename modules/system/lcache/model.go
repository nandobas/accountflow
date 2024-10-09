package lcache

import (
	"time"
)

type LocalCache_t struct {
	AccountEntries map[int64]UnitCache_t
}

type UnitCache_t struct {
	Item    interface{}
	Control ControlCache_t
}

type ControlCache_t struct {
	Updated   *time.Time
	ReadCount int
}
