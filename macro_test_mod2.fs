import {minus, np,plus2} from "./macro_test_mod1"
#[macro(call)]
fn plus(a, b) {
  quote(
    unquote(a)
    +
    unquote(b)
  )
}

fn nplus() {
   puts("nplus")
}

let abc = minus(11,1)
let d = plus2(1,1)

puts(d)

export abc
export plus
export nplus
