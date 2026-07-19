struct A {
  a:
  b:
  c:
}
struct B {
  c:
  d:
  A
}

let bb = B {
  A: A{
    a:1,
    b:2,
    c:33
  },
  c:3,
  d:4
}

puts(bb.a, bb.b, bb.c, bb.d)

bb.a = 11
puts(bb.a)
bb.A.c = 111
puts(bb.A.c)
