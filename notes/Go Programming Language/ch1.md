# 程序结构

## 命名

* Go语言中的 *函数名*、*变量名*、*常量名*、*类型名*、*语句标号* 和*包名* 等所有的命名，都遵循一个简单的命名规则：**一个名字必须以一个字母（Unicode字母）或下划线开头，后面可以跟任意数量的字母、数字或下划线**。大写字母和小写字母是不同的。

* 25个**关键字**：
    ```
    break      default       func     interface   select
    case       defer         go       map         struct
    chan       else          goto     package     switch
    const      fallthrough   if       range       type
    continue   for           import   return      var
    ```
    关键字不能用于自定义名字，只能在特定语法结构中使用

* 30多个**预定义**的名字，主要对应内建的*常量*，*类型* 和*函数*
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

## 声明

* 声明语句定义了程序的各种**实体对象**以及**部分或全部的属性**。Go语言主要有四种类型的声明语句：*var*、*const*、*type*和*func*，分别对应*变量*、*常量*、*类型* 和*函数* 实体对象的声明。

* 一个Go程序对应一个或多个以.go为文件后缀名的源文件中。每个源文件以**包的声明语句**开始。包声明语句之后是 *import* 语句导入依赖的其它包，然后是包一级的类型、变量、常量、函数的声明语句，**包级的各种类型的声明语句的顺序无关紧要**（函数内部的名字则必须**先声明**之后才能使用）

* 包级声明语句声明的名字可在整个包对应的**每个源文件**中访问。相比之下，局部声明的名字就只能在**函数内部**很小的范围被访问

* 一个函数的声明由 *函数名字*、*参数列表*（由函数的调用者提供参数变量的具体值）、一个**可选**的*返回值列表*和包含函数定义的*函数体*组成。如果函数没有返回值，那么返回值列表是省略的。执行函数从函数的第一个语句开始，依次顺序执行直到遇到 *return* 返回语句，如果没有返回语句则是执行到函数末尾，然后返回到**函数调用者**