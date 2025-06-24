pub mod parser;
pub mod noteservice {
    tonic::include_proto!("noteservice");
}

use crate::{noteservice::CreateNoteRequest, parser::NoteParser};
use clap::{Parser, ValueEnum};
use noteservice::note_service_client::NoteServiceClient;
use std::{fs::read_dir, path::Path};

#[derive(Debug, ValueEnum, Clone)]
enum NoteParseModes {
    /// Read note or notes and create them in the remote system
    Create,
    /// Parse notes and only update links in the remote system
    Link,
}

/// Program that parses content from a zettel (note) and creates data or updates data
/// via a gRPC connection to a remote server.
#[derive(Parser, Debug)]
#[command(version)]
struct Args {
    /// URL to connect to gRPC client
    #[arg(short, long, default_value_t = String::from("http://127.0.0.1"))]
    url: String,

    /// Port to connect to gRPC client
    #[arg(short, long, default_value_t = String::from("50051"))]
    port: String,

    #[clap(value_enum, short, long, default_value_t = NoteParseModes::Create)]
    mode: NoteParseModes,

    /// Path to the file or directory of files to parse
    #[arg(long)]
    path: String,
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let args = Args::parse();
    let mut client =
        NoteServiceClient::connect(format!("{}:{}", args.url, args.port).to_string()).await?;

    let mut note_parsers: Vec<NoteParser> = Vec::new();
    let path = Path::new(&args.path);

    if path.is_dir() {
        for path in read_dir(path).unwrap() {
            let parser = NoteParser::load_from_file(path.unwrap().path().as_path());
            match parser {
                Ok(parser) => note_parsers.push(parser),
                Err(_) => {}
            }
        }
    } else {
        let parser = NoteParser::load_from_file(path);
        match parser {
            Ok(parser) => note_parsers.push(parser),
            Err(_) => {}
        }
    }

    match args.mode {
        NoteParseModes::Create => {
            let mut failures = 0;

            for note_parser in &note_parsers {
                let res = client
                    .create_note(CreateNoteRequest {
                        content: note_parser.get_content(),
                        linked_note_titles: vec![],
                        tags: note_parser.get_tags(),
                        title: note_parser.get_title(),
                    })
                    .await;

                match res {
                    Err(e) => {
                        failures += 1;
                        if !e.code().eq(&tonic::Code::AlreadyExists) {
                            println!("{:?}", e);
                        }
                    }
                    Ok(_) => {}
                }
            }

            println!(
                "Completed {}/{} notes. {} unknown errors.",
                note_parsers.len() - failures,
                note_parsers.len(),
                failures,
            );
        }
        NoteParseModes::Link => {
            // TODO: Linking not yet supported
            todo!("Linking not yet supported")
        }
    }

    return Ok(());
}
