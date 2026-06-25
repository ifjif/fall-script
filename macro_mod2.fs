import {a} from "./macro_mod1"
#[macro(call)]
fn infix(a, b) {
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
      puts(unquote(attrs))
    }
  )
}

let infix2 = 1+1

export infix
export infix2
export a
export log
