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

let s1 = St1{
    a:1,
    b:2
}

let s2 = St2 {
  St1:s1,
  a:11,
  b:22,
  c:33
}

s2.getAge()
s2.getName()
