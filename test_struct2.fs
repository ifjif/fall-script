struct St1 {
  a:
  b:
  c:
  d:
}

fn (self St1) getName() {
}

struct St2 {
  a:
  b:
  St1
  c:
}

let st1 = St1 {
  a:1,
  b:2,
  c:3,
  d:4
}

// [st1][getName][xx][xx] callMethod
// {
// } // 方法需要提前编译
//st1.getName()

let st2 = St2 {
  St1: st1,
  a:5,
  b:6,
  c:7
}

puts(st2.d)
puts(st1.a, st2.a, st2.St1.a)

st2.a = 55
st2.St1.a = 11
puts(st1.a, st2.a, st2.St1.a)
