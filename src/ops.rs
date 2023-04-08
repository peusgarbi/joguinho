use crate::models::{NewSave, Save};
use crate::db::establish_sqlite_connection;
use diesel::prelude::*;
use diesel_migrations::run_pending_migrations;
use console::style;

pub fn create_save(username: String) {
    let connection = establish_sqlite_connection();
    let new_save = NewSave {
        username: &username,
    };

    diesel::insert_into(crate::schema::saves::dsl::saves)
        .values(&new_save)
        .execute(&connection)
        .expect("Erro ao salvar novo save.");
}

pub fn read_saves() {
    let connection = establish_sqlite_connection();
    let results = crate::schema::saves::dsl::saves
        .load::<Save>(&connection)
        .expect("Erro ao tentar carregar os saves.");

    if results.len() == 0 {
        println!("{}", style("- Nenhum jogo salvo!").yellow());
    } else if results.len() == 1 {
        println!("{} {} {}", style("- Mostrando").blue(), style(results.len()).yellow(), style("jogo salvo:").blue());
        for video in results {
            println!("> {:?}", video.username)
        };
    } else {
        println!("{} {} {}", style("- Mostrando").blue(), style(results.len()).yellow(), style("jogos salvos:").blue());
        for video in results {
            println!("[{}] {:?}", style(video.id).color256(193), style(video.username).color256(97))
        };
    }
    println!("");
}

pub fn run_migrations() {
    let connection = establish_sqlite_connection();

    match run_pending_migrations(&connection) {
        Ok(()) => println!("Migrações bem-sucedidas"),
        Err(e) => println!("Erro ao rodar migrações: {:?}", e),
    }
}