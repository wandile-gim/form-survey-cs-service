package domain

import "time"

type Tracker struct {
	LastOne time.Time
	Service string
}
