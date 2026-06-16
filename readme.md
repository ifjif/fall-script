# Go脚本语言

## 1.操作符

- **二元运算符**
  1. `+`
  2. `-`
  3. `*`
  4. `/`

- **一元运算符**
  1. `-`
  2. `!`

- **比较符**
  1. `<`
  2. `<=`
  3. `>`
  4. `>=`
  5. `==`
  6. `!=`

- **逻辑运算符**
  1. `&&`
  2. `||`

- **位运算符(不支持)**
  1. `&`
  2. `|`

- **赋值符**
  1. `=`

- **属性符**
  1. `#`

- **分隔符**
  1. `(`
  2. `)`
  3. `;`
  4. `:`
  5. `{`
  6. `}`
  7. `[`
  8. `]`
  9. `,`
  10. `"`

- **注释符**
  1. `//`
  2. `/**/`

## 2.关键字

- `fn`
- `let`
- `for`
- `while`
- `do`
- `return`
- `if`
- `else`
- `false`
- `true`
- `null`

## 3.表达式

- `1+1`
- `1-1`
- `1*1`
- `1/1`
- `1*(1+1)`
- `-1`
- `false`
- `true`
- `!false`
- `1>1`
- `1>=1`
- `1<1`
- `1<=1`
- `1==1`
- `1!=1`
- `1>1 && 1<1`
- `1>1 || 1<1`
- `a`
- `a = 1`
- `a[0]`
- `a[index/key] = xxx (map中不存在，返回null)`
- `[1,2,3]`
- `{a:1, b:2,true:1, false:2, 1:1, 2:2, "abc":12}`
- `fn(...){...}`
- `fn name(...){...}`
- `if(..){...}else if (...){...}else {...}`

## 4. 运算符优先级

- `LOWEST`
- `ASSIGN_ (=)`
- `LOGIC (&&, ||)`
- `EQUALS (==, !=)`
- `LESSGREATER (>, >=, <, <=)`
- `SUM(+, -)`
- `PRODUCT(*, /)`
- `PREFIX(-, !)`
- `CALL(fn())`
- `INDEX(array[index])`

## 5.语句

- `let a = 表达式;`
- `for(xx;xx;xx){...}`
- `while(xx){...}`
- `do{...}while(xx)`
- `return xx;`

## **6.宏**

    macro用来定义宏
      - call 调用类型
      - attr 属性类型
    quote对AST进行包裹
    unquote对AST进行求值，然后生成新AST

    quote可以提取的AST:
    - block    √  提取块(如function的body)
    - params   ×  提取形参(如function的params)
    - args     ×  提取实参(如call的args)
    - ident    ×  提取标识符(如function的ident)

    #[macro(call)]  // 调用宏，可像函数一样进行调用
    fn infx(a,b) {
      quote(
        unquote(a)
        +
        unquote(b)
      )
    }

    #[macro(attr)]  // 属性宏，可作用在函数定义上
    fn log(attrs, ast) {
      quote(
        fn(){
          puts("log-start")
          let arr = unquote(attrs)
          puts(arr[0], arr[1])
          let result = fn(){unquote(ast["block"])}()
          puts("log-end")
          return result
        }
      )
    }

    #[macro(attr)]
    fn trace(attrs, ast) {
      quote(
        fn(){
          puts("trace start")
          let arr = unquote(attrs)
          puts(arr)
          let result = fn(){ unquote(ast["block"]) }()
          puts("trace end")
          return result
        }
      )
    }

    // 链式传递，前一个宏的输出作为下一个宏的输入
    #[log("log-Get", "log/login")]
    #[trace("trace-Post", "trace/leave")]
    fn router(a, b) {
      puts("router function")
      puts(a + b)
      return "router-result"
    }

    puts(infx(1+1, 2+2))
    let result = router(1,2)
    puts("result= ", result)

    结果:
      6
      trace start
      [trace-Post, trace/leave]
      log-start
      log-Get
      log/login
      router function
      3
      log-end
      trace end
      result=
      router-result

## 6.操作码

- `Nop(u8:0)`
- `Null_`
- `Const`
- `Pop`
- `Gt`
- `Ge`
- `Add`
- `Sub`
- `Mul`
- `Div`
- `Eq`
- `Neq`
- `Neg`
- `Not`
- `True`
- `False`
- `Array_`
- `Hash_`
- `Index`
- `Call`
- `Jump`
- `JumpIsFalse`
- `SetGlobal`
- `GetGlobal`
- `SetLocal`
- `GetLocal`
- `GetFree`
- `GetBuiltin`
- `Closure_`
- `CurClosure`
- `Dup`
- `Return`
- `XReturn`

## 7.二进制格式(大端序)

`header{
  SIGNATURE = "fallscript"
  MAJOR     = (u8)0
  MINOR     = (u8)1
  PATCH     = (u8)0
}
compiled_function {
  MaxStackDepth   (u8)
  LocalVarNum     (u8)
  constant-num    (u32)
  Constants{
    i64:                type-tag(u8):I64(1)               (i64)value
    string:             type-tag(u8):STR(2) length(u32)        value
    compiled_function:  type-tag(u8):CF(3)                     value
  }
  instruction-length (u32)
  Instructions
}`
