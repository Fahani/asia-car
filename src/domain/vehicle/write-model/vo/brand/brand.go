package brand

type Brand struct {
	value string
}

func FromValue(value string) (Brand, error) {
	return Brand{value: value}, nil
}

func (n Brand) GetValue() string {
	return n.value
}
