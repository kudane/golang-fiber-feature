package feature

type List interface {
	List() (interface{}, error)
}

func (*priority) List() (interface{}, error) {
	var result = map[string]any{
		"success": true,
		"values":  []int{1, 2, 3},
	}
	return result, nil
}
