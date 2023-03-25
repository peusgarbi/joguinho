local logoArchive = io.open('images/logo.txt', 'r')
local logo = logoArchive:read("*all")
logoArchive:close()
print(logo)

local db = require('infra.database.sqlite.db')
db.prepareDatabase()

repeat
   io.write("Digite seu nickname: ")
   io.flush()
   answer=io.read()
until string.len(answer) > 0

print("Olá, "..answer..", tudo bem com você?")
print('O jogo ainda está em desenvolvimento, mas você pode contribuir!')
