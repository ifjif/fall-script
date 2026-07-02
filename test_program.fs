fn delimiter() {
  puts("---------------------------------------")
}

/*
let a = 1 + 1
puts("加法: 1 + 1 = ")
puts(a)
delimiter()

let a2 = "hello " + "world"
puts("加法: 'hello ' + 'world' = ")
puts(a2)
delimiter()

puts("减法: 1 - 1 = ")
let a3 = 1 - 1
puts(a3)
delimiter()

let a4 = 2 * 2
puts("乘法: 2 * 2 = ")
puts(a4)
delimiter()

let a5 = 2 / 2
puts("除法: 2 / 2 = ")
puts(a5)
delimiter()

let a6 = -2
puts("取反：-2 = ")
puts(a6)
delimiter()

let a7 = !true
puts("非: !true = ")
puts(a7)
delimiter()

let a8 = !false
puts("非: !false = ")
puts(a8)
delimiter()

let a9 = true || false
puts("逻辑或: true || false = ")
puts(a9)
delimiter()

let b = true && false
puts("逻辑与: true && false = ")
puts(b)
delimiter()
*/

/*
puts("if表达式: 
let b1 = if(true) {
    b1 = 1
  }else {
    b1 = 2
  }
")
let b1 = if(true) {
  b1 = 1
}else {
  b1 = 2
}
puts(b1)
*/


puts("if表达式:
let b2 = if(false) {
  b2 = 1
}else {
  b2 = 2
}
")
let b2 = if(false) {
   1
}else {
   2
}
puts(b2)


puts("if表达式:
let b3 = if(false) {
  b3 = 1
}else if (true) {
  b3 = 3
}else {
  b3 = 2
}
")
let b3 = if(false) {
  b3 = 1
}else if (true) {
  b3 = 3
}else {
  b3 = 2
}
puts(b3)
delimiter()



/*
puts("for语句: 
let b4 = 0
for (let i = 0; i <= 10; i = i+1) {
  b4 = b4 + i
}
")
let b4 = 0
for(let i = 0; i <= 10; i = i+1) {
  b4 = b4 + i
}
puts(b4)
*/



/*
puts("for语句:
let b5 = 0;
for(;b5 <= 10;) {
  b5 = b5 + 1
}
")
let b5 = 0
for(;b5 <= 10;) {
  b5 = b5 + 1
}
puts(b5)
delimiter()
*/


/*
puts("while语句:
let b6 = 0;
while(b6 >= 10) {
  b6 = b6 + 1
}
")
let b6 = 0
while(b6 <= 10) {
  b6 = b6 + 1
}
puts(b6)
delimiter()
*/


/*
puts("do-while语句:
let b7 = 0;
do{
  b7 = b7 + 1
}while(b7 <= 10)
")
let b7 = 0
do{
  b7 = b7 + 1
}while(b7 <= 10)
puts(b7)
delimiter()
*/

let b8 = [1,true, false,"hello",1,2]
puts("数组：[1,true, false,'hello'] = ")
puts(b8)
delimiter()

