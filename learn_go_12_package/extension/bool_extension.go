package extension

type Bool struct {
	Value bool
}

func (obj Bool) Or(y bool) bool {
	return obj.Value || y
}

func (obj Bool) And(y bool) bool {
	return obj.Value && y
}
