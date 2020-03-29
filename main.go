package main

import (
	"fmt"
	"os"
)

func main() {

	var name string
	fmt.Print("Project Name: ")
	fmt.Scan(&name)
	
	generate(name)
}

func generate(name string) {
	directoriesToMake := []string{"cmd", "internal", "pkg", "vendor", "api", "web", "config", "init", "script", "build", "deployments", "test"}

	for _, dir := range directoriesToMake {
		os.Mkdir(name, 0777)

		if _, err := os.Stat(dir); os.IsNotExist(err) {

			if err := os.Mkdir(name+"/"+dir, 0777); err != nil {
				fmt.Println(err.Error())
			}

			if dir == "cmd" {
				f, _ := os.Create(name + "/" + "cmd/main.go")
				defer f.Close()
				f.WriteString("package main\n\n\nfunc main() {\n\n}")
			}
		}

	}
}
