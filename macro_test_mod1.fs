import {plus} from "./macro_test_mod2"

#[macro(call)]
fn minus(a,b) {
  quote(
    unquote(a)
    -
    unquote(b)
  )
}

let plusm = plus(1,1)
let minusm = minus(2,1)

let result = plusm + minusm

puts(result)

export minus
