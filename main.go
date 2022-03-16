package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func formatLine(str string, cfg Config) string {
	var fmtd string
	patterns := cfg.Patterns
	for k, v := range patterns {
		if strings.Contains(str, k) {
			fmtd = strings.Replace(str, k, v, -1)
			break
		}
	}
	return fmtd
}

func writeData(file string, datalines []string) error {
	str := strings.Join(datalines, "\n")
	return ioutil.WriteFile(file, []byte(str), 0644)
}

func main() {
	configPath := os.Args[1:]
	if len(configPath) != 1 {
		log.Fatal("Incorrect number of argumen.ts. Provide a config file path !!")
		return
	}
	config := Config{}
	configBytes, err := ioutil.ReadFile(configPath[0])
	if err != nil {
		log.Fatal("Unable to read the config file")
		return
	}
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		log.Fatal("Unable to unmarshall the json config data")
		return
	}
	var data []byte
	data, err = ioutil.ReadFile(config.InputFile)
	if err != nil {
		log.Fatal("Unable to read the file: ", config.InputFile)
		return
	}
	lines := strings.Split(string(data), "\n")
	oplines := make([]string, 0, len(lines))
	for _, line := range lines {
		fmtLine := formatLine(line, config)
		oplines = append(oplines, fmtLine)
	}
	if len(oplines) != len(lines) {
		log.Fatal("Error occurred while processing the data. Length mismatch!!")
	}

	err = writeData(config.OutputFile, oplines)
	if err != nil {
		log.Fatal("Error occurred while writing the data to the file")
	}
}
