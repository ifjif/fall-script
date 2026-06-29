import {plus as plus2, abc, nplus as np } from "./macro_test_mod2"

#[macro(call)]
fn minus(a,b) {
  quote(
    unquote(a)
    -
    unquote(b)
  )
}

fn nonmacro() {
  let caches = {}
  fn() {
    1 + 1
  }
}

let plusm = plus2(1,1)
let minusm = minus(2,1)

let result = plusm + minusm

puts(abc,result)

export nonmacro
export plus2
export np
export minus
