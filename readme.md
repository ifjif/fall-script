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
- `macro(不支持)`

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

## 6.操作码

- `Nop`
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

## 7.二进制格式

**sign: fallscript0.0.1**
