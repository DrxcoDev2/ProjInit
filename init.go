package main

import (
	"fmt"
	"time"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	var projectName string
	var typeProject string
	var framework string
	var useTypescript bool

	promptName := &survey.Input{
		Message: "What is your project name?",
		Default: "mi-app",
	}
	survey.AskOne(promptName, &projectName)

	promptTypeProject := &survey.Select{
		Message: "Select your project type",
		Options: []string{"C", "C++", "GO", "FullStack", "Vanilla HTML/CSS/JS"},
		Default: "C",
	}
	survey.AskOne(promptTypeProject, &typeProject)

	switch typeProject {
	case "C":
		fmt.Println("Creando proyecto C...")
	case "C++":
		fmt.Println("Creando proyecto C++...")
	case "GO":
		fmt.Println("Creando proyecto GO...")
	case "FullStack":
		fmt.Println("Inicializando proyecto FullStack")

		promptFramework := &survey.Select{
			Message: "Which framework do you want to use?",
			Options: []string{"React", "Vue", "Svelte", "Next.js"},
			Default: "React",
		}
		survey.AskOne(promptFramework, &framework)

		promptTS := &survey.Confirm{
			Message: "Use TypeScript?",
			Default: true,
		}
		survey.AskOne(promptTS, &useTypescript)

		// Simular creación
		fmt.Println("\n🛠  Creating project...")
		time.Sleep(1 * time.Second)
		fmt.Printf("📁 Project: %s\n", projectName)
		fmt.Printf("⚛️  Framework: %s\n", framework)
		fmt.Printf("🧠 TypeScript: %v\n", useTypescript)

		time.Sleep(1 * time.Second)
		fmt.Println("\n✅ Done! Run:")
		fmt.Printf("   cd %s\n", projectName)
		fmt.Println("   npm start\n")
	}

}
