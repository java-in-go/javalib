package java

import "javalib/java/util"

type List = util.List
type ArrayList = util.ArrayList

func NewArrayList() *ArrayList {
	return util.NewArrayList()
}
