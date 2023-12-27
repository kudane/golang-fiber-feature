package feature

type List interface {
	List() (interface{}, error)
}

func (*task) List() (interface{}, error) {
	var result = map[string]any{
		"success": true,
		"values":  []int{4, 5, 6},
	}
	return result, nil
}
