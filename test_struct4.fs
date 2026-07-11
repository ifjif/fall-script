struct St1 {
  a:
  b:
}

fn (self St1) getName() {
  puts("getName", self.a)
}

struct St2 {
  a:
  b:
  c:
  St1
}

fn (self St2) getAge() {
  puts("getAge", self.a)
}
