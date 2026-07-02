/*
 * 1
 * false
 * false
 * 2
 * true
 * 1
 * false
 * true
 */
fn closure() {
  let a = 1
  let b = 2

  puts(a) // 1
  fn inner() {
    a = false
    puts(a) // false

    fn inner2() {
      puts(a) // false
      puts(b) // 2

      b = true
      puts(b) // true

      let b = 1
      puts(b) // 1
    }

    inner2()
  }

  inner()

  puts(a) // false
  puts(b) // true
}

closure()
