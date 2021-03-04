package dns

import "github.com/edunx/lua"

func LuaInjectApi(L *lua.LState , parent *lua.LTable) {
	tab := L.CreateTable(0 , 1)
	injectDnsClientApi( L , tab )

	L.SetField(parent , "dns" , tab)
}
