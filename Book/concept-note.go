package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// Python-like input
func input(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func inputYesNo(prompt string) bool {
	for {
		value := strings.ToLower(input(prompt + " (y/n): "))
		if value == "y" || value == "yes" {
			return true
		} else if value == "n" || value == "no" {
			return false
		}
		fmt.Println("❌ Please enter y or n")
	}
}

func generateConceptPrompt(
	bookTitle, conceptName, yourUnderstanding, examples, notes string,
	includeExamples bool,
) string {

	exampleSection := ""
	if includeExamples && examples != "" {
		exampleSection = fmt.Sprintf("Examples / Code (optional): %s", examples)
	}

	prompt := fmt.Sprintf(`You are my personal study assistant.

I have just learned a concept from the book "%s". 
Your task is to turn my understanding of this concept into a professional, structured study note that I can use as a part of my personal study guide.

Here are the details:

- Concept Name: %s
- My Understanding: %s
%s
- Additional Notes / Observations (optional): %s

Please create the study note using the following structure:

1️⃣ Concept Overview
- A clear, concise definition of the concept in plain English.
- Explain why it is important in the context of the book.
- Highlight any common misunderstandings.

2️⃣ Key Points / Bullet Summary
- Break down the main points of the concept in bullets.
- Include practical or theoretical notes if available.
- Add any relevant connections to other concepts in the book.

3️⃣ Examples / Applications
- If examples were provided, rewrite them in a clear, structured way.
- If code snippets are included, format them properly.
- Show how the concept can be applied in real scenarios.

4️⃣ Visual / Mental Map Suggestions (optional)
- Suggest ways to visualize the concept or its relationship with other concepts (can be textual if diagrams are not possible).

5️⃣ Notes for Revision
- Highlight key takeaways.
- Include tips for remembering or applying the concept.

Important Rules:
- Use clear, simple language suitable for personal study notes.
- Keep the note structured and easy to scan.
- Avoid unnecessary text; focus on clarity and usefulness.
- Ensure it can be exported as Markdown or text for later use.
`, bookTitle, conceptName, yourUnderstanding, exampleSection, notes)

	return prompt
}

func main() {
	fmt.Println("📚 NotebookLM Concept Prompt Generator")

	bookTitle := input("📘 Book Title: ")
	conceptName := input("💡 Concept Name: ")
	yourUnderstanding := input("✏️ Your Understanding (short description): ")
	includeExamples := inputYesNo("💻 Do you want to include examples / code?")
	examples := ""
	if includeExamples {
		examples = input("📌 Examples / Code: ")
	}
	notes := input("📝 Additional Notes / Observations (optional): ")

	fmt.Println("\n===== GENERATED PROMPT =====\n")
	prompt := generateConceptPrompt(bookTitle, conceptName, yourUnderstanding, examples, notes, includeExamples)
	fmt.Println(prompt)
}
