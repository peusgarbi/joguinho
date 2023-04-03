use std::{thread, time, env, io::Write};
use dialoguer::{console::{Term, style}, theme::ColorfulTheme, Select};
use termion::{async_stdin, event::Key, input::TermRead};
    
fn explication() {
    let text = "Você é um jovem seminarista prestes a se formar Padre da Igreja Católica Apostólica Romana no auge do século XIII. Ao olhar para o seu passado, questões vem a sua cabeça se você tomou as melhores decisões...";
    let mut time = 30;
    let mut stdin = async_stdin().keys();

    for c in text.chars() {
        print!("{}", style(c).color256(208));
        std::io::stdout().flush().expect("Flushing to succeed");
        std::thread::sleep(std::time::Duration::from_millis(time));

        if let Some(_) = stdin.next() {
            println!("{}", style("Pulando texto...").black());
            println!("{}", style(text).color256(208));
            break;
        }

    }
    println!("");
}

fn choices() -> std::io::Result<()> {
    let reasons = vec![
        "Você era um pobre fudido e decidiu melhorar sua vida.",
        "Você é o escudo e a espada da igreja e acredita na instituição. Deus Vult!",
        "Você acredita em Deus, mas perdeu a fé na instituição. Ainda assim quer dedicar sua vida à missão de Cristo.",
        "Sua família te obrigou a virar um Seminarista.",
        "Você sempre teve uma natureza Inquisitiva sobre as coisas do mundo.",
        "Você só se importa com o Poder e o Status que a instituição pode lhe oferecer.",
        ];

    let selection_one_prompt = format!("{}", style("Por que você decidiu entrar na Igreja?").color256(97));
    let selection_one = Select::with_theme(&ColorfulTheme::default())
        .with_prompt(selection_one_prompt)
        .items(&reasons)
        .default(0)
        .interact_on_opt(&Term::stderr())?;

    match selection_one {
        Some(index) => env::set_var("CHOICE",reasons[index]),
        None => println!("Usuário não selecionou uma opção!")
    }

    let context = [
        "Você se dedicou muito ao estudo dos Demônios.",
        "Você se dedicou ao estudo da Liturgia."
    ];

    let selection_two_prompt = format!("{}", style("Como foram seus primeiros anos de estudo?").color256(97));
    let selection_two = Select::with_theme(&ColorfulTheme::default())
        .with_prompt(selection_two_prompt)
        .items(&context)
        .default(0)
        .interact_on_opt(&Term::stderr())?;

    match selection_two {
        Some(index) => env::set_var("CHOICE",context[index]),
        None => println!("Usuário não selecionou uma opção!")
    }

    Ok(())
}

pub fn introduction() {
    explication();
    choices();
}
