import {infix as ifn,  log} from "./macro_mod2"
import {infix2} from "./macro_mod2"

#[macro(attr)]
fn trace(attrs, ast) {
  quote(
    fn(){
      puts(unquote(attrs))
    }
  )
}

let a = 1

export log
export a
export ifn
export infix2
