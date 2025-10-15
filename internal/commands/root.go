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

func rootCmd() *cobra.Command {
	var gitCommitMessage string
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

			var options = []huh.Option[string]{}
			for _, g := range gitmojis.Gitmojis {
				options = append(options, huh.NewOption(fmt.Sprintf("%s %s", g.Emoji, g.Description), g.Emoji))
			}

			var selected string

			groups := []*huh.Group{}
			groups = append(groups, huh.NewGroup(huh.NewSelect[string]().
				Title("Select a gitmoji").Options(
				options...,
			).Value(&selected)))

			if gitCommitMessage == "" {
				groups = append(groups, huh.NewGroup(huh.NewInput().
					Title("Enter commit message").
					Placeholder("commit message").
					Value(&gitCommitMessage)))
			}

			form := huh.NewForm(
				groups...,
			)

			if err := form.Run(); err != nil {
				fmt.Printf("failed to run form: %v\n", err)
				os.Exit(1)
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
	return cm
}

func Execute(gm pkg.Gitmojis) error {
	gitmojis = gm
	return rootCmd().Execute()
}
