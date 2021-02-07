package dns

import (
	"github.com/edunx/lua"
)

func CheckDnsUserDataByTable(L *lua.LState , opt *lua.LTable , key string ) *Dns {
    var obj *Dns
    var ud  *lua.LUserData
    var ok bool

	ud , ok = opt.RawGetString(key).(*lua.LUserData)
    if !ok { goto ERR }

    obj , ok = ud.Value.(*Dns)
    if !ok { goto ERR } else { return obj }

ERR:
    L.RaiseError("expect invalid type , must be *Dns")
    return nil
}

func CheckDnsUserDataByIdx(L *lua.LState, idx int) *Dns {
	obj , ok := L.CheckUserData(idx).Value.(*Dns)
    if ok { return obj }

    L.RaiseError("expect invalid type , must be *Dns")
    return nil
}

