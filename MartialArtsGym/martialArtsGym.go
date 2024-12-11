package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type mmaAthlete struct {
	about struct {
		Name   string
		Age    int
		Weight float64
		Height float64
		Reach  float64
		Style  string
	}
	skills struct {
		Striking struct {
			Punches      int
			Kicks        int
			Elbows       int
			Knees        int
			Clinch       int
			Footwork     int
			Blocking     int
			HeadMovement int
		}
		Grappling struct {
			WrestlingTakedowns int
			JudoTakedowns      int
			TakedownDefence    int
			SubmissionOffence  int
			SubmissionDefence  int
			TopPressure        int
			GroundAndPound     int
		}
		Conditioning struct {
			Explosiveness int
			Strength      int
			Cardio        int
			Speed         int
			Chin          int
		}
	}
}

func main() {
	var athlete *mmaAthlete
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Welcome to your gym!")
		fmt.Println("What would you like to do?")
		fmt.Println("(create) Create fighter\n(view) View fighters\n(exit) Exit")

		rawInput, _ := reader.ReadString('\n')
		input := strings.TrimSpace(strings.ToLower(rawInput))

		if input == "create" {
			fmt.Println("Creating")
			athlete = &mmaAthlete{}

			fmt.Println("What's your fighter's name?")
			rawInput, _ := reader.ReadString('\n')
			input := strings.TrimSpace(rawInput)

			athlete.about.Name = input
			// Premade athlete need to fix
			athlete.about.Age = 28
			athlete.about.Weight = 170.5
			athlete.about.Height = 5.9
			athlete.about.Reach = 73.0
			athlete.about.Style = "Muay Thai"

		} else if input == "view" {
			fmt.Println("viewing")
			if athlete == nil {
				fmt.Println("No athlete exists. Please create one.")
			} else {
				fmt.Printf("Athlete Details:\n")
				fmt.Printf("Name: %s\n", athlete.about.Name)
				fmt.Printf("Age: %d\n", athlete.about.Age)
				fmt.Printf("Weight: %.1f lbs\n", athlete.about.Weight)
				fmt.Printf("Height: %.1f ft\n", athlete.about.Height)
				fmt.Printf("Reach: %.1f in\n", athlete.about.Reach)
				fmt.Printf("Style: %s\n", athlete.about.Style)
			}
		} else if input == "exit" {
			break
		} else {
			fmt.Println("Input not recognized")
		}
	}
}
