package dns

import (
	"github.com/edunx/lua"
)

func CheckDnsClientUserDataByTable(L *lua.LState, opt *lua.LTable, key string) *Client {
	var obj *Client
	var ud *lua.LUserData
	var ok bool

	ud, ok = opt.RawGetString(key).(*lua.LUserData)
	if !ok {
		goto ERR
	}

	obj, ok = ud.Value.(*Client)
	if !ok {
		goto ERR
	} else {
		return obj
	}

ERR:
	L.RaiseError("expect invalid type , must be *DnsClient")
	return nil
}

func CheckDnsClientUserDataByIdx(L *lua.LState, idx int) *Client {
	obj, ok := L.CheckUserData(idx).Value.(*Client)
	if ok {
		return obj
	}

	L.RaiseError("expect invalid type , must be *DnsClient")
	return nil
}
