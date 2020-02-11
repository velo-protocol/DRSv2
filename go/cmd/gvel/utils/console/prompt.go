package console

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"strings"
)

type prompt struct{}

func NewPrompt() Prompt {
	return &prompt{}
}

func (prompt *prompt) RequestPassphrase() string {
	passphrase := prompt.RequestHiddenString("🔑 Please enter passphrase", nil)

	_, err := (&promptui.Prompt{
		Label: "🔑 Please repeat a passphrase to confirm ",
		Mask:  '*',
		Validate: func(s string) error {
			if s != passphrase {
				return errors.New("passphrase does not match")
			}
			return nil
		},
	}).Run()
	if err != nil {
		ExitWithError(ExitBadArgs, err)
	}

	return passphrase
}

func (prompt *prompt) RequestHiddenString(label string, validate promptui.ValidateFunc) string {
	userInput, err := (&promptui.Prompt{
		Label:    fmt.Sprintf("%s ", label),
		Mask:     '*',
		Validate: validate,
	}).Run()

	if err != nil {
		ExitWithError(ExitBadArgs, err)
	}
	return userInput
}

func (prompt *prompt) RequestString(label string, validate promptui.ValidateFunc) string {
	userInput, err := (&promptui.Prompt{
		Label:    fmt.Sprintf("%s ", label),
		Validate: validate,
	}).Run()

	if err != nil {
		ExitWithError(ExitBadArgs, err)
	}

	return userInput
}

func (prompt *prompt) RequestConfirmation(label string) bool {
	userInput, err := (&promptui.Prompt{
		Label: fmt.Sprintf("%s (Yes/no)", label),
		Validate: func(s string) error {
			if strings.ToLower(s) == "yes" || strings.ToLower(s) == "no" {
				return nil
			}
			return errors.New(`please input "yes" or "no"`)
		},
	}).Run()
	if err != nil {
		ExitWithError(ExitBadArgs, err)
	}

	return strings.ToLower(userInput) == "yes"
}

func (prompt *prompt) RequestChoice(label string, choices []string, currentChoice string) int {
	labelledChoices := make([]string, len(choices))
	for i, choice := range choices {
		if strings.ToLower(currentChoice) == strings.ToLower(choice) {
			choice = fmt.Sprintf("%s (Current)", choice)
		}

		choice = fmt.Sprintf("%d. %s", i+1, choice)
		labelledChoices[i] = choice
	}

	index, _, err := (&promptui.Select{
		Label: fmt.Sprintf("%s ", label),
		Items: labelledChoices,
	}).Run()

	if err != nil {
		ExitWithError(ExitBadArgs, err)
	}

	return index
}
