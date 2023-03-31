#![allow(unused)]

use std::fs::File;
use std::io::{self, Read};
use std::io::prelude::*;
use dialoguer::{theme::ColorfulTheme, Confirm};


fn print_logo() -> std::io::Result<()> {
    // Abra o arquivo
    let mut file = File::open("src/images/logo.txt")?;
    // Crie um buffer
    let mut contents = String::new();
    // Leia o conteúdo do arquivo
    file.read_to_string(&mut contents)?;
    // Imprima o conteúdo no console
    println!("{}", contents);
    Ok(())
}

fn main() {
    print_logo();
    println!("Digite seu nickname: ");
    let mut nickname = String::new();
    io::stdin()
        .read_line(&mut nickname)
        .expect("Não foi possível ler seu nickname!");
    
    if Confirm::with_theme(&ColorfulTheme::default())
        .with_prompt("Deseja prosseguir no jogo?")
        .default(true)
        .show_default(false)
        .wait_for_newline(true)
        .interact()
        .unwrap()
    {
        println!("Olá, {}, tudo bem com você?", nickname.trim_end());
        println!("O jogo ainda está em desenvolvimento, mas você pode contribuir!");
    } else {
        println!("Jogo abortado :(");
    }
}
