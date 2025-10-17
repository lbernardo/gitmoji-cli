package commands

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/go-git/go-git/v6"
	"github.com/lbernardo/gitmoji-cli/pkg"
	"github.com/spf13/cobra"
)

var gitmojis pkg.Gitmojis

var filterTypes = map[string]string{
	"feature":     "✨",
	"feat":        "✨",
	"bug":         "🐛",
	"fix":         "🐛",
	"improvement": "⚡️",
	"remove":      "🔥",
	"release":     "🔖",
	"deploy":      "🚀",
	"doc":         "📝",
	"docs":        "📝",
	"work":        "🚧",
	"in-progress": "🚧",
	"progress":    "🚧",
	"refactor":    "♻️",
	"hotfix":      "🚑️",
}

func rootCmd() *cobra.Command {
	var gitCommitMessage string
	var commitType string
	cm := &cobra.Command{
		Use:   "gitmoji",
		Short: "Git commit with emojis",
		RunE: func(cmd *cobra.Command, args []string) error {

			r, err := git.PlainOpen(".")
			if err != nil {
				return fmt.Errorf("failed to open git repository: %w", err)
			}
			w, err := r.Worktree()
			if err != nil {
				return fmt.Errorf("failed to get worktree: %w", err)
			}
			fields := []huh.Field{}
			var selected string

			selected, ok := filterTypes[commitType]

			if commitType == "" || !ok {
				var options = []huh.Option[string]{}
				for _, g := range gitmojis.Gitmojis {
					options = append(options, huh.NewOption(fmt.Sprintf("%s %s", g.Emoji, g.Description), g.Emoji))
				}

				fields = append(fields, huh.NewSelect[string]().
					Title("Select a gitmoji").Options(
					options...,
				).Value(&selected))
			}

			if gitCommitMessage == "" {
				fields = append(fields, huh.NewInput().
					Title("Enter commit message").
					Placeholder("commit message").
					Value(&gitCommitMessage))
			}

			if len(fields) > 0 {
				form := huh.NewForm(
					huh.NewGroup(
						fields...,
					),
				)

				if err := form.Run(); err != nil {
					fmt.Printf("failed to run form: %v\n", err)
					os.Exit(1)
				}
			}
			_, err = w.Commit(fmt.Sprintf("%s %s", selected, gitCommitMessage), &git.CommitOptions{})
			if err != nil {
				return fmt.Errorf("failed to commit: %w", err)
			}
			fmt.Printf("Committed: %s %s\n", selected, gitCommitMessage)

			return nil
		},
	}
	cm.Flags().StringVarP(&gitCommitMessage, "message", "m", "", "commit message")
	cm.Flags().StringVarP(&commitType, "type", "t", "", "commit type (feature, bug, improvement, release)")
	return cm
}

func Execute(gm pkg.Gitmojis) error {
	gitmojis = gm
	return rootCmd().Execute()
}
