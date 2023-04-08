# Joguinho

Esse é o joguinho feito pelos amigos Devs

- Divirta-se com moderação!

## Como jogar:

Há duas opções para se jogar o joguinho:

1. Clonando este repositório do GitHub;
1. Utilizando a imagem feita em Docker (**recomendado**).

### 1- Utilizando o GitHub:

- Instale o **Git** na sua máquina:
    - https://git-scm.com/downloads
- Instale a linguagem de programação **Rust** na sua máquina:
    - https://www.rust-lang.org
- Instale o **Sqlite3** sua máquina:
    - https://sqlite.org/index.html
- Faça um clone deste repositório, utilizando o comando:
    - `git clone https://github.com/peusgarbi/joguinho.git`
- Navegue até o diretório clonado do jogo (contém o arquivo `Cargo.toml`)
- Para jogar, execute o comando:
    - `cargo run`

### 2- Utilizando o Docker (**recomendado**):

- Instale o Docker na sua máquina:
    - https://www.docker.com
- Para jogar, execute o comando:
    - `docker run -v ~/joguinho/database:/app/src/database -it peusgarbi/joguinho`
- Caso queria atualizar o jogo, utilize o comando:
    - `docker pull peusgarbi/joguinho`
