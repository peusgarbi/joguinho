use std::process::exit;
use dialoguer::{console::{Term, style}, theme::ColorfulTheme, Select};

pub fn menu() -> std::io::Result<()> {
    let menu_choices = vec![
        "Novo jogo",
        "Carregar jogo",
        "Instruções",
        "Configurações",
        "Sair",
        ];

    let menu_prompt = format!("{}", style("Menu Principal").color256(8));
    let menu = Select::with_theme(&ColorfulTheme::default())
        .with_prompt(menu_prompt)
        .items(&menu_choices)
        .default(0)
        .interact_on_opt(&Term::stderr())?;

    match menu {
        Some(0) => println!(""),
        Some(1) => load_game(),
        Some(2) => instructions(),
        Some(3) => settings(),
        Some(4) => exit_game(),
        None => none(),
        _ => invalid(),
    }
    Ok(())
}

fn load_game() {
    println!("");
    println!("{}", style("- Você realmente achava que já conseguiamos salvar alguma coisa?").color256(193));
    println!("");
    menu();
}

fn instructions() {
    println!("");
    println!("{}", style("- Para pular os textos com animação de digitação, aperte a tecla ENTER.").color256(193));
    println!("");
    menu();
}

fn settings() {
    println!("");
    println!("{}", style("- Você é muito curioso, mas ainda não é possível configurar nada.").color256(193));
    println!("");
    menu();
}

fn exit_game() {
    println!("{}", style("Jogo fechado").red());
    exit(0);
}

fn none() {
    println!("");
    println!("{}", style("- Você não selecionou uma opção do menu!").color256(193));
    println!("");
    menu();
}

fn invalid() {
    println!("");
    println!("{}", style("- Opção inválida!").color256(193));
    println!("");
    menu();
}