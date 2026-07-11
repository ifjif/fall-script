let a = 1
puts(a)

fn name() {
  let a = 11
  puts(a)

  fn name2() {
    a = 22
    puts(a)
  }
  name2()
  puts(a)
}
name()
