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
test_dependencies = {
   "lsqlite3",
}
test = {
   type = "busted",
   platforms = {
      windows = {
         flags = { "--exclude-tags=ssh,git,unix", "-Xhelper", "lua_dir=$(LUA_DIR)", "-Xhelper", "lua_interpreter=$(LUA)" }
      },
      unix = {
         flags = { "--exclude-tags=ssh,git", "-Xhelper", "lua_dir=$(LUA_DIR)", "-Xhelper", "lua_interpreter=$(LUA)" }
      }
   }
}