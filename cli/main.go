package main

import (
	"fmt"
	"os"
	"os/exec"
)

func help() string {
	result := ""

	helpmap := map[string]string{
		"help": "Give helpful information",
		"new":  "Creates a new wywern project with a scaffold.",
	}

	result += "Help instructions for wywern CLI:\n"

	for k, v := range helpmap {
		result += fmt.Sprintf("    - %s: %s\n", k, v)
	}

	return result
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("No arguments found...")
		fmt.Print(help())
		os.Exit(0)
	}
	subcommand := args[0]

	switch subcommand {
	case "help":
		fmt.Print(help())
	case "new":
		if len(args) < 2 {
			fmt.Println("Invalid use of the `new` subcommand.")
			fmt.Println("Proper use: `wyvern new my_project`")
			os.Exit(0)
		}

		projectName := args[1]

		os.Mkdir(projectName, 0777)
		os.Mkdir(fmt.Sprintf("%s/controllers", projectName), 0777)
		os.Mkdir(fmt.Sprintf("%s/controllers/index", projectName), 0777)

		os.Create(fmt.Sprintf("%s/main.go", projectName))
		indexgo, err := os.Create(fmt.Sprintf("%s/controllers/index/index.go", projectName))
		if err != nil {
			fmt.Println("Failed to create index.go")
			os.Exit(1)
		}

		contents := `
package index

import (
	"fmt"
	"net/http"

	"github.com/busyLambda/wywern/route"
)

var indexController = route.NewController(
	func(w *http.ResponseWriter, r *http.Request, res *route.Resources) {
		controller := index_page()
		component.Render(context.Background(), *w)
	},
	nil,
)	

func init() {
	fmt.Println("My init!")
}
		`

		indexgo.Write([]byte(contents))

		indextemplate, err := os.Create(fmt.Sprintf("%s/controllers/index/index.templ", projectName))
		if err != nil {
			fmt.Println("Failed to create index.templ")
			os.Exit(1)
		}

		contents = `
package index

templ index_page() {
	<h1>
		Index
	</h1>
}
		`
		indextemplate.Write([]byte(contents))

		cmd := exec.Command("templ", "generate")
		cmd.Output()

		cmd = exec.Command("go", "mod", "init", "github.com/asd/asd")
		cmd.Dir = projectName
		cmd.Output()

		cmd = exec.Command("go", "get", "github.com/a-h/templ")
		cmd.Dir = projectName
		cmd.Output()
	case "run":
		cmd := exec.Command("templ", "generate")
		cmd.Output()

		cmd = exec.Command("go", "run", "main.go")
		cmd.Output()
	default:
		fmt.Printf("Wrong subcommand `%s`.\n", subcommand)
		os.Exit(1)
	}
}
