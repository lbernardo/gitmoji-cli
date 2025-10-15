package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"

	"github.com/lbernardo/gitmoji-cli/internal/commands"
	"github.com/lbernardo/gitmoji-cli/pkg"
)

//go:embed gitmojis.json
var gitmojis []byte

func main() {
	var gm pkg.Gitmojis
	err := json.Unmarshal(gitmojis, &gm)
	if err != nil {
		fmt.Printf("failed to unmarshal gitmojis: %v\n", err)
		os.Exit(1)
	}

	if err := commands.Execute(gm); err != nil {
		fmt.Printf("failed to execute command: %v\n", err)
		os.Exit(1)
	}
}
