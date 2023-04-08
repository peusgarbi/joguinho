#![allow(unused)]

#[macro_use]
extern crate diesel;

mod menu;
mod db;
mod schema;
mod models;
mod ops;
mod introduction;

use std::{fs::File, io::Read};
use console::style;


fn print_logo() -> std::io::Result<()> {
    // Abra o arquivo
    let mut file = File::open("src/images/logo.txt")?;
    // Crie um buffer
    let mut contents = String::new();
    // Leia o conteúdo do arquivo
    file.read_to_string(&mut contents)?;
    // Imprima o conteúdo no console
    println!("{}", style(contents).color256(2));
    Ok(())
}

fn main() {
    ops::run_migrations();
    print_logo();
    menu::menu();
}