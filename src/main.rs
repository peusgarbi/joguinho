#![allow(unused)]

use std::fs::File;
use std::io::{self, Read};
use std::io::prelude::*;
use std::{thread, env};
use std::time::Duration;
use dotenv::dotenv;
use dialoguer::{theme::ColorfulTheme, Confirm, Input};
use console::{Term, style};
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
    println!("{}", contents);
    Ok(())
}

fn main() {
    dotenv().ok();
    print_logo();
    let terminal = Term::stdout();
    
    let nickname: String = Input::with_theme(&ColorfulTheme::default())
        .with_prompt("Digite seu nickname:")
        .interact_text()
        .unwrap();

    if Confirm::with_theme(&ColorfulTheme::default())
        .with_prompt("Deseja prosseguir no jogo?")
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

        println!("Olá, {}, tudo bem com você?", style(nickname.trim_end()).yellow());
        println!("{}", style("O jogo ainda está em desenvolvimento, mas você pode contribuir!").green());
    } else {
        println!("{}", style("Jogo abortado :(").red());
    }
}
