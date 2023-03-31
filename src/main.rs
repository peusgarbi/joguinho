#![allow(unused)]

use std::fs::File;
use std::io::{self, Read};
use std::io::prelude::*;
use std::thread;
use std::time::Duration;
use dialoguer::{theme::ColorfulTheme, Confirm, Input};
use console::{Term, style};
use indicatif::ProgressBar;


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
