package dns

import (
	"github.com/edunx/lua"
	pub "github.com/edunx/rock-public-go"
)

const (
	MT string = "ROCK_DNS_GO_MT"
)

func injectDnsClientApi(L *lua.LState, parent *lua.LTable) {
	mt := L.NewTypeMetatable(MT)

	L.SetField(mt, "__index", L.NewFunction(Get))
	L.SetField(mt, "__newindex", L.NewFunction(Set))

	L.SetField(parent, "client", L.NewFunction(CreateDnsClientUserdata))
}

func CreateDnsClientUserdata(L *lua.LState) int {
	opt := L.CheckTable(1)
	v := &Client{
		C: Config{
			nameserver: opt.CheckSocket("nameserver", L),
			timeout:    opt.CheckInt("timeout", 5),
			typeName:   opt.CheckString("type_name", "A"),
		},
	}
	if err := v.Start(); err != nil {

		L.RaiseError("start DnsClient fail , e: %v", err)

		pub.Out.Debug("start DnsClient fail, e: %v", err)
		return 0
	}
	pub.Out.Debug("start DnsClient successful , info: %v", v.C)

	ud := L.NewUserDataByInterface(v, MT)
	L.Push(ud)
	return 1

}

func Get(L *lua.LState) int {
	self := CheckDnsClientUserDataByIdx(L, 1)
	name := L.CheckString(2)

	switch name {
	case "query":
		L.Push(L.NewFunction(func(vm *lua.LState) int {
			host := L.CheckString(1)
			if rc, size, e := self.Query(host); e != nil {
				pub.Out.Err("query fail , err: %v", e)
			} else {
				pub.Out.Info("query succeed , rc: %v , size: %d", rc, size)
			}
			return 0
		}))
	}

	return 1
}

func Set(L *lua.LState) int {
	return 0
}

func (c *Client) ToUserData(L *lua.LState) *lua.LUserData {
	return L.NewUserDataByInterface(c, MT)
}
