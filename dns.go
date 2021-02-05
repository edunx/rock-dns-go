package dns

import (
	"github.com/miekg/dns"
	"sync"
	"time"
)

func (d *Dns) Start() error {

	d.msgs = &sync.Pool{
		New: func() interface{} {
			return &dns.Msg{}
		},
	}

	d.client = &dns.Client{
		Timeout: time.Duration( d.C.timeout ) * time.Second,
	}

	return nil
}

func (d *Dns) Close() {

}

func (d *Dns) Query(host string, e *error) ([]string , int) {

	m := d.msgs.Get().(*dns.Msg)
	m.SetQuestion(host , dns.TypeA)

	r , _ , err := d.client.Exchange(m , d.C.nameserver)
	if err != nil {
		*e = err
		return nil , 0
	}

	rlen := len(r.Answer)
	rc := make([]string , rlen )
	size := 0

	var ans dns.RR
	for i := 0 ; i < rlen ; i++ {
		ans = r.Answer[i]

		switch d.C.TypeName  {
		case "A":
			record , t := ans.(*dns.A)
			if t {
				rc[size] = record.A.String()
				size ++
			}

		case "CNAME":
			record , t := ans.(*dns.CNAME)
			if t {
				rc[size] = record.Target
				size++
			}
		}

	}

	d.msgs.Put(m)
	return rc[:size] , size
}

