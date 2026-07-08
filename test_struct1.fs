struct st1 {
  a:
  b:
  c:
}

let a = st1 {
  a:1,
  b:2,
  c:3
}

puts(a)

puts(a.a,a.b,a.c)

/*
// 拿到st1的meta,得到index
fn(s1 st1) getName() {
  s1.a    getFieldByIndex 0
}

a.b.c

// 拿到犯法，判断是否是直接，不是 需要带offset
// 在方法内部的 字段获取都需要 + offset
// 在方法中访问字段，直接索引
*/
