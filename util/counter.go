package logutil



func NewCounter(name string, size int) *Counter {
	return &Counter{name, size, 0, "", l, NewRateLimit(Rate / 2)}
}

type Counter struct {
	name string
	size int

	cur  int
	last string
	l    Log

	limit *ratelimit
}

func (c *Counter) Set(s int, last string) {
	c.cur = s
	c.last = last
	if !c.limit.Limit() || c.cur == c.size {
		c.print()
	}
}

func (c *Counter) print() {
	c.l.Infof("%s [%3d%%] %d of %d %s", c.name, c.cur*100/c.size, c.cur, c.size, c.last)
}
