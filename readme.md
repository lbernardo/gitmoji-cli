# gitmoji-cli

This project is an alternative implementation to the original repo: https://github.com/carloscuesta/gitmoji-cli

`gitmoji-cli` is a small command-line utility to create Git commits prefixed with a Gitmoji (emoji) selection. It provides an interactive UI to pick a gitmoji and enter a commit message, and it can also accept a message via the -m flag. The tool uses an embedded gitmojis list and performs commits directly using go-git.

## Features

- Interactive selection of a gitmoji from the official list.
- Option to provide the commit message via the `-m` / `--message` flag.
- Commits using the local Git repository (no external git binary required).

## Installation

Build from source with Go 1.25 or newer:

```bash
# from project root
go build ./...

# or install globally (module path)
go install github.com/lbernardo/gitmoji-cli/cmd/gitmoji@latest
```

After `go install`, the `gitmoji` binary will be available in your `$GOPATH/bin` or `$(go env GOPATH)/bin`.

## Usage

Run the tool inside any Git repository:

```bash
gitmoji
```

The command opens an interactive form:

- First you select a gitmoji (emoji + description).
- Then you enter the commit message (unless you provided `-m`).

Example interactive flow:

- Select: ✨ Introduce new features
- Message: Add user profile page

The tool will run the commit and print a confirmation like:

```
Committed: ✨ Add user profile page
```

### Non-interactive message

You can provide the commit message with `-m` to skip the message input but keep the emoji selection interactive:

```bash
gitmoji -m "Update README with usage examples"
```

If you want a fully scripted commit (emoji + message) you can combine the `-m` flag with a non-interactive approach by using tools that send input to the program; however, this tool is primarily designed for interactive use.

## Flags

- `-m`, `--message` string: Provide the commit message. If omitted, the tool will prompt for it interactively.

## Behavior and implementation notes

- The program uses an embedded `gitmojis.json` shipped in `cmd/gitmoji/`.
- Commits are made via the go-git library directly on the current repository (`.`).
- If the working directory is not a git repository the command will return an error.
- The selected emoji and message are combined as `"<emoji> <message>"` and used as the commit message.

## Contributing

Contributions, fixes and suggestions are welcome. To work on the project:

1. Fork the repository.
2. Make changes and run `go build`.
3. Open a pull request with a clear description of your changes.

## License

This repository does not include an explicit license file in the workspace snapshot. If you publish this project, add a LICENSE file and choose a license.

## Acknowledgements

This project embeds the Gitmoji list from the official source and uses these libraries:

- github.com/charmbracelet/huh for interactive prompts
- github.com/go-git/go-git for git operations
- github.com/spf13/cobra for the CLI command

---

## About gitmoji.dev

[gitmoji.dev](https://gitmoji.dev/) is the official site and project that maintains the Gitmoji standard: a curated list of emojis to prefix git commit messages to make their intent clearer. The project provides an API and a JSON schema with the canonical list of gitmojis. This repository embeds a copy of that list (`cmd/gitmoji/gitmojis.json`) so the CLI uses the same emoji choices and descriptions as the official source.

If you want a different or updated list you can replace the embedded `gitmojis.json` with a newer one obtained from the gitmoji.dev API.

If you want, I can also:

- add examples showing how to script the tool non-interactively,
- run a quick build to validate the project builds cleanly on your machine, or
- generate a trimmed `README` with badges and a short usage table.

