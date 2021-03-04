package dns

import (
	"github.com/miekg/dns"
	"sync"
)

type Config struct {
	nameserver string
	timeout    int
	typeName   string
}

type Client struct {
	C Config

	obj  *dns.Client
	msgs *sync.Pool
}

type Look interface {
	Query(string, *error) string
}
