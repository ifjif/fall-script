#[macro(call)]
fn infx(a,b) {
  quote(
    unquote(a)
    +
    unquote(b)
  )
}

#[macro(attr)]
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
