import {St3} from "./test_struct_module3"
struct St2 {
  St3
  a:
  c:
  d:
}

fn (self St2) GetSt2A() {
  puts(self.a)
}

export St2
