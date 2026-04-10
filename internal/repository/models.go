// Package repository contains models for the portfolio
package repository

type Project struct {
	Name string `json:"name"`
	Repo string `json:"repo"`
	Info string `json:"info"`
}

var Projects = []Project{
	{
		Name: "tic tac toe",
		Repo: "https://github.com/adrr-dev/tic-tac-toe-app",
		Info: "A simple Tic Tac Toe WebApp.",
	},
}

type Skill struct {
	Name  string `json:"name"`  // the name of my skill (golang, etc)
	Level string `json:"level"` // the level of my skill (moderate golang)
}

var Skills = []Skill{
	{
		Name:  "golang",
		Level: "beginner.",
	},
}
