#[macro(call)]
fn plus(a, b) {
  quote(
    unquote(a)
    +
    unquote(b)
  )
}

export plus
