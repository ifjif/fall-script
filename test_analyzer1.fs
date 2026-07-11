let a = 1
let b = true
let c = false
let d = "abc"
let e = null
let a2 = [1,true, false, "ab", null, a]
let a3 = {1:1, true:false, "ab":"ab", a:null}
a
b
c
d
e
a2
a3

let g = 1 + 1
let f1 = a + b
let f2 = a - b
let f3 = a * b
let f4 = a / b
let f5 = a && b
let f6 = a || b
let f7 = a > b
let f8 = a >= b
let f9 = a < b
let f10 = a <= b
let f11 = a == b
let f12 = a != b
let f13 = (2+2)

let f14 = -a
let f15 = !b

fn name(a, b) {
  let c = 1
  let d = 2
  a
  b
  c
  d
  g
  name

  fn name2(aa) {
    a
    c
    g
    name
  }
}

let name2 = fn(a,b) {
  let c = 1
  g
  name2
}

if (true) {
  let a = 1
}else if (true){
  let b = 2
}else {
  let c = 3
}

a[1];
[1,2,3][1]
{1:1, true:false}[true]

a[:]
a[::]
a[:::]
a[1:]

a.name

a = 1
a[1] = 1
a.name = "ab"

a(1,2,b)

for (let a = 1; a < 2; a = a+1) {
  let a = 2
}
for(;;) {
  let a = 3
}

while(true) {
  let a = 1
}

do {
  let a = 1
}while(true)

return;
return 1
