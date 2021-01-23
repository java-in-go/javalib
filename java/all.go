package java

import (
	"javalib/java/time/format"
	"javalib/java/util"
)

type List = util.List
type ArrayList = util.ArrayList
type LinkedList = util.LinkedList
type Map = util.Map
type HashMap = util.HashMap
type LinkedHashMap = util.LinkedHashMap

func NewArrayList() *ArrayList {
	return util.NewArrayList()
}
func NewArrayListCap(initialCapacity int) *ArrayList {
	return util.NewArrayListCap(initialCapacity)
}
func NewLinkedList() *LinkedList {
	return util.NewLinkedList()
}
func NewHashMap() *HashMap {
	return util.NewHashMap()
}
func NewHashMapCap(initialCapacity int) *HashMap {
	return util.NewHashMapCap(initialCapacity)
}
func NewLinkedHashMap() *LinkedHashMap {
	return util.NewLinkedHashMap()
}

func If(condition bool, trueValue interface{}, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	}
	return falseValue
}

//---------------------date time formatter-------------------------

// java pattern: yyyy-MM-dd HH:mm:ss
var DateTimeFormat = format.DateTimeFormat

// java pattern: yyyy-MM-dd HH:mm:ss.SSS
var DateTimeMSFormat = format.DateTimeMSFormat

// java pattern: yyyy-MM-dd HH:mm
var DateTimeMinuteFormat = format.DateTimeMinuteFormat

// java pattern: yyyy-MM-dd
var DateFormat = format.DateFormat

// java pattern: HH:mm:ss
var TimeFormat = format.TimeFormat

// java pattern: yyyyMMddHHmmss
var DateTimePureFormat = format.DateTimePureFormat

// java pattern: yyyyMMddHHmmssSSS
var DateTimeMSPureFormat = format.DateTimeMSPureFormat

// java pattern: yyyyMMddHHmm
var DateTimeMinutePureFormat = format.DateTimeMinutePureFormat

// java pattern: yyyyMMdd
var DatePureFormat = format.DatePureFormat

// java pattern: HHmmss
var TimePureFormat = format.TimePureFormat
