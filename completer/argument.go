package completer

import (
	"github.com/c-bata/go-prompt"
)

func (c *Completer) argumentsCompleter(repo string, args []string) []prompt.Suggest {
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(
			[]prompt.Suggest{
				{Text: "help", Description: "Help about any command"},
				{Text: "pr", Description: "Create, view, and checkout pull requests"},
				{Text: "issue", Description: "Create and view issues"},
				// Custom commands.
				{Text: "exit", Description: "Exit this program"},
			},
			args[0],
			true,
		)
	}

	switch args[0] {
	case "issue":
		if len(args) == 2 {
			return prompt.FilterHasPrefix(
				[]prompt.Suggest{
					{Text: "create", Description: "Create a new issue"},
					{Text: "list", Description: "List and filter issues in this repository"},
					{Text: "status", Description: "Show status of relevant issues"},
					{Text: "view", Description: "View an issue in the browser"},
				},
				args[1],
				true,
			)
		}
		if args[1] == "view" && len(args) == 3 {
			suggests := getIssueNumberSuggestions(c.client, c.repo)
			suggests = append(suggests, getIssueURLSuggestions(c.client, c.repo)...)
			return prompt.FilterHasPrefix(
				suggests,
				args[2],
				true,
			)
		}
	case "pr":
		if len(args) == 2 {
			return prompt.FilterHasPrefix(
				[]prompt.Suggest{
					{Text: "checkout", Description: "Check out a pull request in Git"},
					{Text: "create", Description: "Create a pull request"},
					{Text: "list", Description: "List and filter pull requests in this repository"},
					{Text: "status", Description: "Show status of relevant pull requests"},
					{Text: "view", Description: "View a pull request in the browser"},
				},
				args[1],
				true,
			)
		}
		if args[1] == "view" && len(args) == 3 {
			suggests := getPullRequestsNumberSuggestions(c.client, c.repo)
			suggests = append(suggests, getPullRequestsBranchSuggestions(c.client, c.repo)...)
			suggests = append(suggests, getPullRequestsURLSuggestions(c.client, c.repo)...)
			return prompt.FilterHasPrefix(
				suggests,
				args[2],
				true,
			)
		}
	}
	return []prompt.Suggest{}
}