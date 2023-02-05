package strategy

import "time"

type Context struct {
	Strategy Interface
}

func NewContext(s Interface) *Context {
	return &Context{Strategy: s}
}

func (c *Context) SetStrategy(s Interface) {
	c.Strategy = s
}

func (c *Context) CalculateETA(startAt time.Time, trip TripDetail) time.Time {
	return c.Strategy.CalculateETA(startAt, trip)
}
