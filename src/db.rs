use diesel::prelude::*;

pub fn establish_sqlite_connection() -> SqliteConnection {
    let database_url = "src/database/joguinho.db";
    SqliteConnection::establish(&database_url)
        .expect(&format!("Erro de conexão com o banco de dados: {}", database_url))
}