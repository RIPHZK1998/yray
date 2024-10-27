package command

type Command struct {
	name     string
	Lname    string
	Run      func(string) error
	Commands *[]Command
}
