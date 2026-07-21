fn doMatch(v) {
   match v {
    0 => 0,
    true => true,
    false => false,
    "abc" => "abc",
    null => "是null",
    1 => {
        let a = 10
        a
    }
    2 | 3 | 4 => "2,3,4"
    20 if true => "20 if true",
    30 if false => "30 if false",
    _ => {}
  }
}

let result = doMatch(0) 
puts(result)

result = doMatch(true)
puts(result)

result = doMatch(false)
puts(result)

result = doMatch("abc")
puts(result)

result = doMatch(null)
puts(result)

result = doMatch(1)
puts(result)

result = doMatch(2)
puts(result)

result = doMatch(3)
puts(result)

result = doMatch(4)
puts(result)

result = doMatch(20)
puts(result)

result = doMatch(30)
puts(result)

result = doMatch(40)
puts(result)
