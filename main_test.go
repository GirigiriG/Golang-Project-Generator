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
		if _, err := os.Stat("Test/" + dir); err == nil {
			fmt.Println(err)
			t.Fatal("Directory already exist")
		}
	}

	os.RemoveAll("Test")
}
