package command

type Command struct {
	Name     string
	Lname    string
	Run      func(string) error
	Commands *[]Command
}
