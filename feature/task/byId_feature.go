package feature

type ById interface {
	ById() (interface{}, error)
}

func (*task) ById() (interface{}, error) {
	var result = map[string]any{
		"success": true,
		"values":  "id is 9",
	}
	return result, nil
}
