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

- `import`
- `export`
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
- `a[x:x:x:x]`
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

- `import {xx, xxx as aa} from "xx"`
- `import * as xx from "xx"`
- `export xxx`
- `let a = 表达式;`
- `for(xx;xx;xx){...}`
- `while(xx){...}`
- `do{...}while(xx)`
- `return xx;`

## 6.slice

     [start : end : step : max]

     step 不能为0, 默认为1
     当step 不是默认值，而是指定值，如果max存在，panic
     当step 是默认值，而slice是string, 如果max存在,panic

     step > 0 不为默认值 时
     0 <= start < end <= max <= rawMax (end, max, rawMax 为最大不可到达的索引位置)
     start 省略 默认 0
     end 省略 默认 len(arr)

     step < 0 时（会 copy）
     len(arr) > start > end >= -1
     start, end 可为 负数,需要保证
     len(arr) > len(arr) + start > len(arr) + end
     start 省略 默认为 len(arr) - 1
     end 省略 默认为 -1

     let arr = [1,2,3,4]

     arr[1]
     arr[1:]
     arr[1:2]
     arr[1:2:3]
     arr[1:2:3:4]

     arr[:1]
     arr[:1:2]
     arr[:1:2:3]

     arr[::1]
     arr[::1:2]

     arr[:::1]

     arr[:]
     arr[::]
     arr[:::]

## 7.宏

**不支持宏生宏**

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

## 8.内置函数

- `len`
- `byte`
- `puts`

## 9.操作码

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
- `Slice`
- `Call`
- `Jump`
- `JumpIsFalse`
- `SetGlobal`
- `GetGlobal`
- `SetLocal`
- `GetLocal`
- `NewBoxLocal`
- `SetBoxLocal`
- `GetBoxLocal`
- `SetFree`
- `GetFree`
- `GetFreeRaw`
- `GetBuiltin`
- `Closure_`
- `CurClosure`
- `Dup`
- `Return`
- `XReturn`

## 10.export元信息格式(大端序)

    count               u16
    export_meta{
      name_length       u16
      name              string
      source_length     u16
      source            string
      imported_length   u16
      imported          string
      export_idx        u8
      origin            u8
      ast {
        kind            u8
        内部string结构:{
          tag           u8(short_str:0, long_str:1)
          length        (0:u8 / 1:u32)
          content       string
        }
      }
    }

## 11.二进制格式(大端序)

    header{
      SIGNATURE = "fallscript"
      MAJOR     = (u8)0
      MINOR     = (u8)1
      PATCH     = (u8)0
    }
    module_meta{
      name_length     (u32)
      name string
      global_num      (u16)
      imports_num     (u16)
      imports {
        from          (u16)
        imported      (u16)
        local         (u16)
      }
      exports_num     (u16)
      exports {
        name          (u16)
        global_idx    (u16)
      }
    }
    compiled_function {
      MaxStackDepth   (u8)
      LocalVarNum     (u8)
      constant_num    (u16)
      Constants{
        i64:                type-tag(u8):I64(1)               (i64)value
        string:             type-tag(u8):STR(2) length(u32)        value
        compiled_function:  type-tag(u8):CF(3)                     value
      }
      instruction_length (u32)
      Instructions
    }
