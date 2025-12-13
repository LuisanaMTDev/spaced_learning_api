use chrono::Local;
use clap::Parser;
use owo_colors::OwoColorize;
use terminal_size::{Width, terminal_size};

use crate::library::commands::Args;

mod library;

fn main() {
    let args = Args::parse();

    if let Some((Width(w), _)) = terminal_size() {
        println!(
            "{:^width$}",
            format!("Today is {}", Local::now().date_naive().bold()),
            width = w as usize
        );
        println!(
            "{:^width$}",
            format!(
                "Topic \"{}\" has {} amount of cards, and your desired repetitions are {}",
                args.topic, args.amount_of_cards, args.repetitions_desired
            )
            .bold()
            .yellow(),
            width = w as usize
        );
        args.print_repetitions_dates();
    } else {
        println!(
            "{}",
            format!("Today is {}", Local::now().date_naive().bold()),
        );
        println!(
            "{}",
            format!(
                "Topic \"{}\" has {} amount of cards, and your desired repetitions are {}",
                args.topic, args.amount_of_cards, args.repetitions_desired
            )
            .bold()
            .yellow(),
        );
        args.print_repetitions_dates();
    }
}
