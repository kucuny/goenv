package cli

type GoEnvClient struct {
	command    string
	subCommand []string
}

func NewGoEnvClient(args []string) *GoEnvClient {
	if len(args) < 1 {
		return nil
	}

	return &GoEnvClient{
		command:    args[1],
		subCommand: args[1:],
	}
}

func (cli *GoEnvClient) Run() {

}