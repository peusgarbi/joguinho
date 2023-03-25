local sqlite3 = require("lsqlite3complete")
connection = {
    db = sqlite3.open("infra/database/sqlite/joguinho.sqlite", sqlite3.OPEN_READWRITE + sqlite3.OPEN_CREATE + sqlite3.OPEN_SHAREDCACHE)
}

function connection.prepareDatabase()
    local createDbSql = io.open('infra/database/sqlite/create.db.sql', 'r')
    if not createDbSql then return nil end
    local sqlContent = createDbSql:read("*all")
    connection.query(sqlContent, nil)
end

function connection.query(sql, params)
    stmt = connection.db:prepare(sql)
    if not stmt then return end
    if params then
        stmt:bind_names{params}
    end
    print(stmt)
    stmt:exec()
end

function connection.connect()
    if connection.db:isopen() then return nil end
    connection.db:open('joguinho.sqlite')
end

function connection.disconnect()
    if not connection.db:isopen() then return nil end
    connection.db:close()
end

return connection
