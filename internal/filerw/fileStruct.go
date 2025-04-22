package filerw

import "time"

type noteInfo struct {
	Name     string
	Category string
	Time     time.Time
	Path     string
}
