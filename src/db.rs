use diesel::prelude::*;
use dotenv::dotenv;
use std::env;

pub fn establish_sqlite_connection() -> SqliteConnection {
    let database_url = env::var("DATABASE_URL").expect("DATABASE_URL não configurada");
    SqliteConnection::establish(&database_url)
        .expect(&format!("Erro de conexão com o banco de dados: {}", database_url))
}