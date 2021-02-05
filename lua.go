package dns

import (
	"github.com/edunx/lua"
	pub "github.com/edunx/rock-public-go"
)

const (
	MT string = "ROCK_DNS_GO_MT"
)

func CheckDnsUserData(L *lua.LState, idx int) *Dns {
	ud := L.CheckUserData(idx)

	switch v := ud.Value.(type) {
	case *Dns:
		return ud.Value.(*Dns)
	default:
		L.RaiseError("expect invalid type , must be *Dns, got %T", v)
		return nil
	}

}

func CreateDnsUserdata(L *lua.LState) int {
	opt := L.CheckTable(1)
	v := &Dns{
		C: Config{
			nameserver: opt.CheckSocket("nameserver", L),
			timeout:    opt.CheckInt("timeout", 5),
		},
	}
	if err := v.Start(); err != nil {

		L.RaiseError("start Dns fail , e: %v", err)

		pub.Out.Debug("start Dns fail, e: %v", err)
		return 0
	}
	pub.Out.Debug("start Dns successful , info: %s", v.C)

	ud := L.NewUserDataByInterface(v, MT)
	L.Push(ud)
	return 1

}

func LuaInjectApi(L *lua.LState, parent *lua.LTable) {
	mt := L.NewTypeMetatable(MT)

	L.SetField(mt, "__index", L.NewFunction(Get))
	L.SetField(mt, "__newindex", L.NewFunction(Set))

	L.SetField(parent, "dns", L.NewFunction(CreateDnsUserdata))
}

func Get(L *lua.LState) int {
	return 0
}

func Set(L *lua.LState) int {
	return 0
}

func (d *Dns) ToUserData(L *lua.LState) *lua.LUserData {
	return L.NewUserDataByInterface(d, MT)
}
