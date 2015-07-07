package mvc

import (
	"time"
)

type SessionMgr struct {
	Sessions map[string]*Session
}

type Session struct {
	Key        string
	Value      interface{}
	LastUpdate time.Time
	Validity   time.Duration
}

type ISessionStore interface {
	Source() string
	Set(string, interface{}, time.Time) error
	Get(string) *Session
}
