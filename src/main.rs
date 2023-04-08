#![allow(unused)]

#[macro_use]
extern crate diesel;

mod menu;
mod db;
mod schema;
mod models;
mod ops;
mod introduction;

use std::{thread, env, fs::File, time::Duration, io::{self, Read, prelude::*}};
use console::{Term, style};
use dialoguer::{theme::ColorfulTheme, Confirm, Input};
use dotenv::dotenv;
use indicatif::ProgressBar;


fn print_logo() -> std::io::Result<()> {
    let images_dir = match env::var("IMAGES_DIR") {
        Ok(val) => val,
        Err(e) => panic!("Erro ao obter o caminho até a pasta de imagens: {}", e),
    };
    let logo_dir = format!("{}/logo.txt", images_dir);
    // Abra o arquivo
    let mut file = File::open(logo_dir)?;
    // Crie um buffer
    let mut contents = String::new();
    // Leia o conteúdo do arquivo
    file.read_to_string(&mut contents)?;
    // Imprima o conteúdo no console
    println!("{}", style(contents).color256(2));
    Ok(())
}

fn main() {
    dotenv().ok();
    ops::run_migrations();
    print_logo();
    menu::menu();
    let terminal = Term::stdout();
    let nickname_prompt = format!("{}", style("Digite seu nickname:").color256(97));
    let nickname: String = Input::with_theme(&ColorfulTheme::default())
        .with_prompt(nickname_prompt)
        .interact_text()
        .unwrap();

    let welcome_speach = format!("{}{}{}", style("Olá, ").blue(), style(nickname.trim_end()).yellow(), style(", tudo bem com você?").blue());
    println!("{}", welcome_speach);
    println!("{}", style("O jogo ainda está em desenvolvimento, mas você pode contribuir!").yellow());

    let are_you_sure_prompt = format!("{}", style("Deseja prosseguir no jogo?").color256(97));
    if Confirm::with_theme(&ColorfulTheme::default())
        .with_prompt(are_you_sure_prompt)
        .default(true)
        .show_default(false)
        .wait_for_newline(true)
        .interact()
        .unwrap()
    {
        let deps = 100;
        let pb = ProgressBar::new(deps);
        for _ in 0..deps {
            pb.inc(1);
            thread::sleep(Duration::from_millis(10));
        }
        pb.finish_and_clear();

        let username = nickname.to_string();
        ops::create_save(username);
        

        introduction::introduction();
        println!("{}", style("Por enquanto o jogo acaba por aqui...").color256(8));
        println!("{}", style("Se você contribuir no desenvolvimento, as coisas vão progredir mais rápido!").color256(8));
        println!("{}{}!", style("Bora ajudar, ").color256(8), style(nickname.trim_end()).yellow());
        menu::menu();
    } else {
        println!("{}", style("Jogo abortado, retornando ao menu principal...").red());
        println!("");
        menu::menu();
    }
}