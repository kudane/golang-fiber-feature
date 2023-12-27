package core

type CreateByIdCommand struct {
	Id int
}

type CreateByIdCommandResponse struct {
	Success bool
	Item    int
}

func (*service) ById(command *CreateByIdCommand) (*CreateByIdCommandResponse, error) {
	result := &CreateByIdCommandResponse{
		Success: true,
		Item:    command.Id,
	}
	return result, nil
}
