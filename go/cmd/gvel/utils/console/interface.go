package console

import "github.com/manifoldco/promptui"

type Prompt interface {
	RequestPassphrase() string
	RequestString(label string, validate promptui.ValidateFunc) string
	RequestHiddenString(label string, validate promptui.ValidateFunc) string
	RequestConfirmation(label string) bool
	RequestChoice(label string, choices []string, opt RequestChoiceOptions) int
}
