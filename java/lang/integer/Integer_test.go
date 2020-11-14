package integer

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	i := New(1)
	str := i.ToStr()
	if str != "1" {
		t.Error("New ToStr failed")
	}
	if i.IntValue() != 1 {
		t.Error("Integer IntValue failed")
	}

}

func TestFromStr(t *testing.T) {
	x := FromStr("2")
	value := x.IntValue()
	fmt.Println(value)
}
