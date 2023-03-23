local answer
repeat
   io.write("Digite seu nickname: ")
   io.flush()
   answer=io.read()
until string.len(answer) > 0

print("Olá, "..answer..", tudo bem com você?")
print('O jogo ainda está em desenvolvimento, mas você pode contribuir!')