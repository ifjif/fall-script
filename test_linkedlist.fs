struct LinkedList {
  head:
  tail:
}
struct Node {
  value:
  next:
}

// 顶部提升
fn NewNode(value) {
  return Node {
    value:value,
    next:null
  }
}

fn (self LinkedList) append(value) {
  let node = NewNode(value)
  if (!self.head && !self.tail) {
    self.head = self.tail = node
  }
  self.tail.next = node
  self.tail = node
}


let ll = LinkedList{head:null, tail:null}

ll.append(1)
ll.append(2)

for (let a = ll.head; a; a = a.next) {
  puts(a.value)
}
