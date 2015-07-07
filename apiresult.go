package mvc

import (
	"time"
)

type APIStatus string

const (
	APIStatus_OK  APIStatus = "OK"
	APIStatus_NOK APIStatus = "NOK"
)

type APIResult struct {
	Status   APIStatus
	Message  string
	Duration time.Duration
	Data     interface{}
}

func (a *APIResult) Run(f func(data interface{}) (interface{}, error), parm interface{}) *APIResult {
	t0 := time.Now()
	a.Status = APIStatus_OK
	a.Message = ""
	if f == nil {
		a.Data = nil
	} else {
		r, e := f(parm)
		if e != nil {
			a.Status = APIStatus_NOK
			a.Message = e.Error()
			a.Data = nil
		} else {
			a.Data = r
		}
	}
	a.Duration = time.Since(t0)
	return a
}
