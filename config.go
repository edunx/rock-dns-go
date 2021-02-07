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

type Dns struct {
	C Config

	client *dns.Client
	msgs   *sync.Pool
}

type Look interface {
	Query(string, *error) string
}
