# 基于go实现的java基础类库

#### 设计原则

尽量使用最少的代码，完成类似的功能。

由于go不支持泛型和重载，部分方法进行了重命名

由于go调用其他包结构体需要带包名，例如

```go
import 	"javalib/java/util"

func main() {
    util.NewArrayList()
}
```

因此将所有结构体在java包all.go中进行重命名，调用时可通过`java.NewArrayList()`的方式快速调用

```go
import 	"javalib/java"

func main() {
    java.NewArrayList()
}
```



#### 预期实现的功能

+ java.lang.Object
+ java.util.List
  + [x] `Size()`
  + [x] `IsEmpty()`
  + [x] `Contains(interface{})`
  + [x] `Size() int`
  + [x] `IsEmpty() bool`
  + [x] `Contains(obj interface{}) bool`
  + [x] `Add(obj interface{}) bool`
  + [x] `Remove(index int) interface{}`
  + [x] `RemoveObj(obj interface{}) bool`
  + [ ] `ContainsAll() bool`
  + [ ] `AddAll(list List) bool`
  + [ ] `RemoveAll(list List) bool`
  + [x] `Clear()`
  + [x] `Get(index int) interface{}`
  + [x] `Set(index int, obj interface{}) interface{}`
  + [x] `IndexOf(obj interface{}) int`
+ java.util.ArrayList
+ java.util.LinkedList
+ java.time.LocalDateTime
+ java.time.LocalDate
+ java.time.LocalTime
+ java.time.format.DateTimeFormatter