package dns

import (
	"errors"
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

	switch d.C.typeName {
	case "A" , "CNAME":
		return nil
	default:
		return errors.New("not found dns type name")
	}
}

func (d *Dns) Close() {

}

func (d *Dns) Query(host string ) ([]string , int , error) {
	m := d.msgs.Get().(*dns.Msg)
	m.SetQuestion(host , dns.TypeA)

	r , _ , err := d.client.Exchange(m , d.C.nameserver)
	if err != nil {
		return nil , 0 , err
	}

	rlen := len(r.Answer)
	rc := make([]string , rlen )
	size := 0

	var ans dns.RR
	for i := 0 ; i < rlen ; i++ {
		ans = r.Answer[i]

		switch d.C.typeName  {
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
	return rc[:size] , size , nil
}

