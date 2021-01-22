# 基于go实现的Java util库

#### 设计原则

尽量使用最少的代码，完成类似的功能。

####  实现的功能
由于go不支持泛型和重载，部分方法进行了重命名
+ java.util.List
  + java.util.ArrayList
  + java.util.LinkedList
  + [x] `Size() int`
  + [x] `IsEmpty() bool`
  + [x] `Contains(obj interface{}) bool`
  + [x] `Add(obj interface{}) bool`
  + [x] `Remove(index int) interface{}`
  + [x] `RemoveObj(obj interface{}) bool`
  + [x] `Clear()`
  + [x] `Get(index int) interface{}`
  + [x] `Set(index int, obj interface{}) interface{}`
  + [x] `IndexOf(obj interface{}) int`
  + [x] `ForEach(f func(value interface{}))`
  + [x] `ForEachIndex(f func(index int, value interface{}))`
  + [x] `Map(f func(value interface{}) interface{}) List`
  + [x] `Reduce(f func(v1 interface{}, v2 interface{}) interface{}) interface{}`

+ java.util.Map
  + java.util.HashMap
  + java.util.LinkedHashMap
  + [x] `Size() int`
  + [x] `IsEmpty() bool`
  + [x] `ContainsKey(key interface{}) bool`
  + [x] `ContainsValue(value interface{}) bool`
  + [x] `Get(key interface{}) interface{}`
  + [x] `Put(key interface{}, value interface{}) interface{}`
  + [x] `Remove(key interface{}) interface{}`
  + [x] `PutAll(m Map)`
  + [x] `Clear()`
  + [x] `Keys() List`
  + [x] `Values() List`
  + [x] `ForEach(f func(key interface{},value interface{}))`
        
+ java.time.format.DateTimeFormatter
    + [x] `OfPattern(javaPattern string) *DateTimeFormatter`
    + [x] `Format(t time.Time) string`

#### 使用方法

go调用其他包结构体默认需要带包名，例如

```go
package main
import 	"javalib/java/util"

func main() {
    util.NewArrayList()
}
```

因此将所有结构体在java包all.go中进行重命名，调用时可通过`java.NewArrayList()`的方式快速调用

```go
package main
import 	"javalib/java"

func main() {
    java.NewArrayList()
}
```
或者直接使用
```go
package main
import . "javalib/java"

func main() {
    NewArrayList()
}
```

