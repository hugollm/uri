package main

import (
	"io/ioutil"
	"os/exec"
	"strings"
	"testing"
)

func TestSamples(t *testing.T) {
	inputs, outputs := readSamples("samples.txt")
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		output := runProgram(input)
		expected := outputs[i]
		if output != expected {
			msg := "\n\n-----------------------------------INPUT\n" + input
			msg += "-----------------------------------OUTPUT\n" + output
			msg += "-----------------------------------EXPECTED\n" + expected
			msg += "-----------------------------------\n"
			t.Log(msg)
			t.Fail()
		}
	}
}

func readSamples(path string) (inputs []string, outputs []string) {
	inputs = make([]string, 0)
	outputs = make([]string, 0)
	text := readFile(path)
	samples := strings.Split(text, "===")
	for _, sample := range samples {
		parts := strings.Split(sample, "---")
		input := strings.TrimSpace(parts[0]) + "\n"
		output := strings.TrimSpace(parts[1]) + "\n"
		inputs = append(inputs, input)
		outputs = append(outputs, output)
	}
	return inputs, outputs
}

func readFile(path string) string {
	bytes, _ := ioutil.ReadFile(path)
	return string(bytes)
}

func runProgram(input string) string {
	cmd := exec.Command("go", "run", "area_of_a_circle.go")
	cmd.Stdin = strings.NewReader(input)
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	return string(output)
}
