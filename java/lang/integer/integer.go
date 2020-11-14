package integer

import "strconv"

type Integer struct {
	value int
}

func New(value int) Integer {
	return Integer{
		value: value,
	}
}
func FromStr(value string) Integer {
	return Integer{
		value: ParseInt(value),
	}
}
func (i Integer) IntValue() int {
	return i.value
}
func (i Integer) CompareTo(anotherInteger Integer) int {
	return Compare(i.value, anotherInteger.value)
}
func (i Integer) ToStr() string {
	return ToStr(i)
}

func ToStr(i Integer) string {
	str := strconv.Itoa(i.value)
	return str
}

func ParseInt(value string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic("For input string:" + value + " can not cast to Integer")
	}
	return intValue
}

func Compare(x int, y int) int {
	if x < y {
		return -1
	} else if x == y {
		return 0
	} else {
		return 1
	}
}
