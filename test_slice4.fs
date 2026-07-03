#[macro(attr)]
fn sliceM(attrs, ast) {
  quote(
    fn() {
      let a = unquote(attrs)
      let b = a[1:]
      puts(a,b)
    }
  )
}

export sliceM
