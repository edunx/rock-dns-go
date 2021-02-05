package dns

import (
	"github.com/edunx/lua"
	pub "github.com/edunx/rock-public-go"
)

const (
	MT string = "ROCK_DNS_GO_MT"
)


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
