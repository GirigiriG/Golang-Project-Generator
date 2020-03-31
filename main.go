package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	generate()
}

func generate() {
	directoriesToMake := []string{"cmd", "internal", "pkg", "vendor", "api", "web", "config", "init", "script", "build", "deployments", "test"}

	for _, dir := range directoriesToMake {
		if _, err := os.Stat(dir); os.IsNotExist(err) {

			if err := os.Mkdir("../"+dir, 0777); err != nil {
				fmt.Println(err.Error())
			}

			if dir == "cmd" {
				f, _ := os.Create("../" + "cmd/main.go")
				defer f.Close()
				f.WriteString("package main\n\n\nfunc main() {\n\n}")
			}
		}
	}

	cleanUp()
}

func cleanUp() {
	os.Chdir("../")
	os.RemoveAll("Golang-Project-Generator")
	cmd := exec.Command("cd", "..")
	cmd.start()
}
