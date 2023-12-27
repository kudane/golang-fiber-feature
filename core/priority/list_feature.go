package core

type CreateListCommand struct {
	Id int
}

type CreateListCommandResponse struct {
	Success bool
	Items   []int
}

func (*service) List(command *CreateListCommand) (*CreateListCommandResponse, error) {
	result := &CreateListCommandResponse{
		Success: true,
		Items:   []int{4, 5, 6},
	}
	return result, nil
}
