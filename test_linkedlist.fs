struct LinkedList {
  head:
  tail:
}
struct Node {
  value:
  next:
}

// 槽位顶部提升
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

export LinkedList
