package main

type Config struct {
	InputFile  string            `json:input`
	OutputFile string            `json:output`
	Patterns   map[string]string `patterns	`
}
