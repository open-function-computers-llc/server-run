package server

import "time"

type session struct {
	expires time.Time
}
