package main

import (
	"fmt"
	"os"
	"testing"
)

func TestGenerator(t *testing.T) {
	directoriesToMake := []string{
		"cmd",
		"internal",
		"pkg",
		"vendor",
		"api",
		"web",
		"config",
		"init",
		"script",
		"build",
		"deployments",
		"test",
	}

	generate("Test")

	for _, dir := range directoriesToMake {
		if _, err := os.Stat("Test/" + dir); err != nil {
			fmt.Println("Directory already created")
		}
	}
	os.RemoveAll("Test")
}
