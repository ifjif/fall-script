// 需要 St1 所有字段 和 所有的方法名
import {St1} from "./test_struct_module2"

struct St2 {
  St1
  a:
  b:
}

fn (self St2) GetSelfA() {
  return self.a
}

let st1 = St1 {
  a:1,
  b:2
}

let st2 = St2 {
  St1:st1,
  a:11,
  b:22
}
