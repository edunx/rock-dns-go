package dns

import (
	"errors"
	"github.com/miekg/dns"
	"sync"
	"time"
)

func (c *Client) Start() error {

	c.msgs = &sync.Pool{
		New: func() interface{} {
			return &dns.Msg{}
		},
	}

	c.obj = &dns.Client{
		Timeout: time.Duration(c.C.timeout) * time.Second,
	}

	switch c.C.typeName {
	case "A", "CNAME":
		return nil
	default:
		return errors.New("not found dns type name")
	}
}

func (c *Client) Close() {

}

func (c *Client) Query(host string) ([]string, int, error) {
	m := c.msgs.Get().(*dns.Msg)
	m.SetQuestion(host, dns.TypeA)

	r, _, err := c.obj.Exchange(m, c.C.nameserver)
	if err != nil {
		return nil, 0, err
	}

	rlen := len(r.Answer)
	rc := make([]string, rlen)
	size := 0

	var ans dns.RR
	for i := 0; i < rlen; i++ {
		ans = r.Answer[i]

		switch c.C.typeName {
		case "A":
			record, t := ans.(*dns.A)
			if t {
				rc[size] = record.A.String()
				size++
			}

		case "CNAME":
			record, t := ans.(*dns.CNAME)
			if t {
				rc[size] = record.Target
				size++
			}
		}

	}

	c.msgs.Put(m)
	return rc[:size], size, nil
}
