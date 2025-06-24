use std::{fs::read_to_string, path::Path};

use regex::Regex;

pub struct NoteParser {
    title: String,
    all_content: String,
}

impl NoteParser {
    pub fn load_from_file(path: &Path) -> Result<NoteParser, Box<dyn std::error::Error>> {
        let all_content = read_to_string(path).unwrap();

        let sections: Vec<&str> = all_content.split("---").collect();
        if sections.len() != 3 {
            return Err(format!(
                "File '{}' doesn't have 3 sections separated by '---'",
                path.to_str().unwrap()
            )
            .as_str())?;
        }

        let title = path.file_name().unwrap().to_str().unwrap().to_string();

        return Ok(NoteParser { title, all_content });
    }

    pub fn new(all_content: String, title: String) -> NoteParser {
        let sections: Vec<&str> = all_content.split("---").collect();
        assert_eq!(sections.len(), 3);

        return NoteParser { title, all_content };
    }

    pub fn get_tags(&self) -> Vec<String> {
        let mut idx = 0;
        let mut tags: Vec<String> = Vec::new();
        let tag_section_chars: Vec<char> = self.all_content.chars().collect();

        while idx < tag_section_chars.len() {
            if tag_section_chars[idx] == '#' && !tag_section_chars[idx + 1].is_whitespace() {
                idx += 1; // add 1 to ignore the # in the output
                let start_tag_idx = idx;

                while idx < tag_section_chars.len() && !tag_section_chars[idx].is_whitespace() {
                    idx += 1;
                }

                let end_tag_idx = idx;

                tags.push(self.all_content[start_tag_idx..end_tag_idx].to_string());
            } else {
                idx += 1;
            }
        }

        return tags;
    }

    pub fn get_title(&self) -> String {
        return self.title.clone();
    }

    pub fn get_content(&self) -> String {
        return self.all_content.split("---").collect::<Vec<&str>>()[1].to_string();
    }

    pub fn get_linked_note_titles(&self) -> Vec<String> {
        let re = Regex::new(r"\[\[(.*)\]\]").unwrap();
        let mut titles: Vec<String> = Vec::new();

        for (_, [title]) in re.captures_iter(&self.all_content).map(|c| c.extract()) {
            titles.push(title.to_string());
        }

        return titles;
    }
}
