import {St2} from "./test_struct_module2"
import {St3} from "./test_struct_module3"

struct St1 {
  St2
  a:
  b:
}

fn (self St1) GetSt1A() {
  puts(self.a)
}

let st3 = St3 {
  a:"st3a",
  e:"st3e",
  f:"st3f"
}

let st2 = St2 {
  St3:st3,
  a:"st2a",
  c:"st2c",
  d:"st2d"
}

let st1 = St1 {
  St2:st2,
  a:"st1a",
  b:"st1b"
}

puts("St1组合St2,St2组合St3:")
puts()
puts("字段访问:----------------------------------------------")
puts("输出st3直接访问数据")
puts(st3.a,st3.e, st3.f)
puts("--------------------")

puts("输出st2直接访问数据")
puts(st2.a,st2.c, st2.d, st2.e, st2.f)
puts("显示输出st2中st3的直接访问数据:st2.St3.xx")
puts(st2.St3.a, st2.St3.e, st2.St3.f)
puts("--------------------")

puts("输出st1直接访问数据")
puts(st1.a, st1.b, st1.c, st1.d)
puts("显式输出st1中st2的直接访问数据:st1.St2.xx")
puts(st1.St2.a,st1.St2.c, st1.St2.d, st1.St2.e, st1.St2.f)
puts("显示输出st1中st2中的st3直接访问数据:st1.St2.St3.xx")
puts(st1.St2.St3.a, st1.St2.St3.e, st1.St2.St3.f)
puts("--------------------")

puts("方法访问:----------------------------------------------")
puts("访问St1自己的方法GetSt1A: st1.GetSt1A")
st1.GetSt1A()
puts("直接访问St1组合的St2方法GetSt2A: st1.GetSt2A")
st1.GetSt2A()
puts("显示访问St1组合的St2方法GetST2A: st1.St2.GetSt2A")
st1.St2.GetSt2A()
puts("访问St1组合中St2组合的St3方法GetSt3A,st2直接访问: st1.St2.GetSt3A")
st1.St2.GetSt3A()
puts("访问St1组合中St2组合的St3方法GetSt3A,st2显示访问: st1.St2.St3.GetSt3A")
st1.St2.St3.GetSt3A()
