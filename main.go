package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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

			if dir != "cmd" {
				f, _ := os.Create("../" + dir + "/.gitignore")
				defer f.Close()
				f.WriteString("*\n!.gitignore")
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
	wd := getCurrentDirectoryName()
	wd = strings.ToLower(wd)
	cmd := exec.Command("npx", "create-react-app", wd+"_app")
	cmd.Start()

	os.RemoveAll("Golang-Project-Generator")
}

func getCurrentDirectoryName() string {
	pwd, _ := os.Getwd()
	splitDir := strings.Split(pwd, "/")
	result := splitDir[len(splitDir)-1]
	return result
}
