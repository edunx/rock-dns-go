package dns

import (
	"github.com/edunx/lua"
)

func CheckDnsUserDataByTable(L *lua.LState , opt *lua.LTable , key string ) *Dns {
	v := tb.RawGetString(key)

    ud , ok := v.(*lua.LUserData)
    if !ok { goto ERR }

    v , ok := ud.Value.(*Dns)
    if !ok { goto ERR } else { return v }


ERR:
    L.RaiseError("expect invalid type , must be *Dns")
    return nil
}

func CheckDnsUserDataByIdx(L *lua.LState, idx int) *Dns {
	ud := L.CheckUserData(idx)
    if _  , ok := ud.Value.(*Dns) {
        return ud.Value.(*Dns)
    }

    L.RaiseError("expect invalid type , must be *Dns, got %T", v)
    return nil
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
