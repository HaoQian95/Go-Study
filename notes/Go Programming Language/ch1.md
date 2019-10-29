# 程序结构

## 命名

* Go语言中的 *函数名*、*变量名*、*常量名*、*类型名*、*语句标号*和*包名*等所有的命名，都遵循一个简单的命名规则：**一个名字必须以一个字母（Unicode字母）或下划线开头，后面可以跟任意数量的字母、数字或下划线**。大写字母和小写字母是不同的。

* 25个**关键字**：
    ```
    break      default       func     interface   select
    case       defer         go       map         struct
    chan       else          goto     package     switch
    const      fallthrough   if       range       type
    continue   for           import   return      var
    ```
    关键字不能用于自定义名字，只能在特定语法结构中使用

* 30多个**预定义**的名字，主要对应内建的*常量*，*类型*和*函数*
    ```
    内建常量: true false iota nil

    内建类型: int int8 int16 int32 int64
             uint uint8 uint16 uint32 uint64 uintptr
             float32 float64 complex128 complex64
             bool byte rune string error

    内建函数: make len cap new append copy close delete
             complex real imag
             panic recover
    ```
    这些内部预先定义的名字并不是关键字，可以在定义中重新使用它们

* 函数**内部**定义的名字**只在函数内部有效**。函数外部定义的名字在**当前包**的所有文件中都可以访问。名字开头字母的**大小写**决定了名字在包外的可见性。如果一个名字是大写字母开头的（**必须是在函数外部定义的包级名字；包级函数名本身也是包级名字**），那么它将是导出的，也就是说可以被外部的包访问。**包名字**一般总是用小写字母。

* 尽量使用**短小**的名字，尤其对于**局部变量**。如果一个名字的作用域比较大，生命周期也比较长，那么用长的名字将会更有意义。

* 驼峰式命名（不用下划线）。缩略词避免使用大小写混合。