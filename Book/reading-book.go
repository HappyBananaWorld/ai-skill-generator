package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func input(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func inputInt(prompt string, min, max int) int {
	for {
		value := input(prompt)
		i, err := strconv.Atoi(value)
		if err == nil && i >= min && i <= max {
			return i
		}
		fmt.Printf("❌ Please enter a valid number between %d and %d\n", min, max)
	}
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

func generatePrompt(
	title, description, levelStr, goalStr string,
	chapters int,
	practical, activeRecall bool,
	preferredContext string,
) string {

	projectSection := ""
	if practical {
		projectSection = "Provide real-world examples, code snippets, and anti-patterns."
	}

	questionSection := ""
	if activeRecall {
		questionSection = "Generate review questions: conceptual, why, and scenario-based."
	}

	levelMap := map[string]string{
		"1": "Beginner",
		"2": "Intermediate",
		"3": "Advanced",
	}

	goalMap := map[string]string{
		"1": "Deep conceptual understanding",
		"2": "Project-based learning",
		"3": "Fast comprehension",
	}

	prompt := fmt.Sprintf(`
You are my personal study assistant.
I am studying a technical / scientific book and I want to understand it deeply, quickly, and practically.
Your role is NOT to summarize blindly, but to help me build real understanding.

---

📘 Book Metadata
- Book Title: %s
- Book Topic / Description: %s
- Book Level: %s
- Number of Chapters: %d

---

🎯 My Goal for This Book
%s

---

🧠 How I Want You To Help Me

For each chapter, do the following in order:

1️⃣ Chapter Understanding
- Identify the core concepts of the chapter.
- Explain why these concepts exist (what problem they solve).
- Explain how they connect to previous chapters.

2️⃣ Concept Breakdown (Very Important)
For every key concept, provide:
- Simple definition (plain English)
- Why it matters
- Common misunderstandings or mistakes
- Related concepts in this book

Use this structure:

Concept:
Definition:
Why it exists:
Common mistakes:
Connected concepts:

3️⃣ Practical Layer (if applicable)
%s

4️⃣ Active Recall & Questions
%s

5️⃣ Mental Model & Map
- Create a text-based mental map showing how concepts relate:
Concept A → Concept B → Concept C
Concept B ≠ Concept D

---

📝 Notes & Study Output
- Maintain a running glossary of terms.
- Maintain chapter-by-chapter notes that can later be exported as a study sheet.
- Write clean, structured Markdown.

---

⚠️ Important Rules
- Do NOT oversimplify.
- Do NOT skip reasoning.
- If a concept depends on earlier chapters, remind me.
- If something is abstract, make it concrete.
- Teach me as if I will use this knowledge in real projects.

---

✅ Final Objective
By the end of this book, I should be able to:
- Explain the concepts in my own words
- Apply them in practice
- Recognize good vs bad usage
- Connect ideas across chapters

%s
`, title, description, levelMap[levelStr], chapters, goalMap[goalStr], projectSection, questionSection, func() string {
		if preferredContext != "" {
			return fmt.Sprintf("Preferred Context / Tech Stack: %s", preferredContext)
		}
		return ""
	}())

	return prompt
}

func main() {
	title := input("📘 Book Title: ")
	description := input("📝 Description: ")

	level := inputInt(
		"🎯 Level:\n  1) Beginner\n  2) Intermediate\n  3) Advanced\nSelect: ", 1, 3,
	)

	goal := inputInt(
		"🚀 Goal:\n  1) Deep understanding\n  2) Project-based learning\n  3) Fast review\nSelect: ", 1, 3,
	)

	chapters := inputInt("📚 Total Chapters: ", 1, 100)

	practical := inputYesNo("💡 Include practical examples and code snippets?")
	activeRecall := inputYesNo("🧠 Include questions & active recall?")

	preferredContext := input("⚙️ Preferred Context / Tech Stack (optional): ")

	fmt.Println("\n===== GENERATED PROMPT =====\n")
	prompt := generatePrompt(title, description, strconv.Itoa(level), strconv.Itoa(goal), chapters, practical, activeRecall, preferredContext)
	fmt.Println(prompt)
}
