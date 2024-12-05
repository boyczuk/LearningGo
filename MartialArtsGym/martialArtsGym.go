package main

import "fmt"

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
	var athlete mmaAthlete

	athlete.about.Name = "John Doe"
	athlete.about.Age = 28
	athlete.about.Weight = 170.5
	athlete.about.Height = 5.9
	athlete.about.Reach = 73.0
	athlete.about.Style = "Muay Thai"

	fmt.Println(athlete)
}
