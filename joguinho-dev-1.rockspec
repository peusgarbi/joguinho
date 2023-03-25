rockspec_format = "3.0"
package = "joguinho"
version = "dev-1"
source = {
   url = "git+https://github.com/peusgarbi/joguinho",
}
description = {
   summary = "Um joguinho",
   detailed = [[
      LuaRocks
   ]],
}
dependencies = {
   "lua >= 5.1",
   "lsqlite3complete",
}
