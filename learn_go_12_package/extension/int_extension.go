package extension

type Int struct {
	Value int
}

func New(i int) *Int {
	return &Int{Value: i}
}

func (obj Int) IsGreaterThan(y int) bool {
	return obj.Value > y
}

func (obj Int) IsLessThan(y int) bool {
	return obj.Value < y
}

func (obj Int) IsEqualTo(y int) bool {
	return obj.Value == y
}

func (obj Int) IsGreaterThanOrEqualTo(y int) bool {
	return Bool{obj.IsGreaterThan(y)}.Or(obj.IsEqualTo(y))
}
