package logutil

import (
	"io"

	"github.com/dustin/go-humanize"
)

type Printer func(string, ...interface{})

func ReadProgress(r io.Reader, printer Printer, name string, size int64) io.ReadCloser {

	var sizeHuman string

	if size > 0 {
		sizeHuman = humanize.Bytes(uint64(size))
	}
	return &ProgressBar{r, printer, name, size, 0, sizeHuman, 0, NewRateLimit(Rate)}
}

type ProgressBar struct {
	io.Reader

	printer Printer
	name    string
	size    int64

	done int64

	sizeHuman string //So we don't calcuate it in every read.
	last      int64

	limit *ratelimit
}

func (p *ProgressBar) print() {

	if p.sizeHuman == "" {
		p.printer("%s [UKN%%] %s of UKN", p.name, humanize.Bytes(uint64(p.done)))
	}
	p.printer("%s [%3d%%] %s of %s", p.name, p.done*100/p.size, humanize.Bytes(uint64(p.done)), p.sizeHuman)
}

func (p *ProgressBar) Read(b []byte) (int, error) {
	n, err := p.Reader.Read(b)
	p.done += int64(n)

	if (p.done-p.last) > (p.size/50) && !p.limit.Limit() {
		p.last = p.done
		p.print()
	}

	return n, err
}

func (p *ProgressBar) Close() error {
	p.print()
	c, ok := p.Reader.(io.Closer)
	if ok {
		return c.Close()
	}
	return nil
}
