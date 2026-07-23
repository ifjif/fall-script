/*
 * init
 * prepend
 * append
 * insert_at_index
 */
struct SingleLinkedList {
  head:
  tail:
  size:
}

// 初始化
fn (self SingleLinkedList) init() {
  self.head = null
  self.tail = null
  self.size = 0
}

struct Node {
  value:
  next:
}

fn NewNode(value) {
  return Node {
    value:value,
    next:null
  }
}

/*
 * 改变大小
 */
fn changeSize(linkedList, inc) {
  let size = linkedList.size
  linkedList.size = size + inc
}

/*
 * 判断头尾是否存在
 */
fn exist(self, node) {
  if (!self.head || !self.tail) {
    self.head = node
    self.tail = node
    self.size = 1
    return false
  }
  return true
}

/*
 * 改变头
 */
fn changeHead(linkedList, node) {
  let head = linkedList.head
  linkedList.head = node
  node.next = head
  let size = linkedList.size
  changeSize(linkedList, 1)
}

/*
 * 改变尾
fn changeTail(linkedList, node) {
  let tail = linkedList.tail
  tail.next = node
  linkedList.tail = node
  changeSize(linkedList, 1)
}
*/

/*
 * 插入到头部
 */
fn (self SingleLinkedList) prepend(value) {
  let node = NewNode(value)
  if (exist(self, node)) {
    changeHead(self, node)
  }
}

/*
 * 插入到尾部
fn (self SingleLinkedList) append(value) {
  let node = NewNode(value)
  if (exist(self, node)) {
    changeTail(self,node)
  }
}
*/
/*

/*
 * 插入到指定索引处
 
fn (self SingleLinkedList) insert_at_index(value, index) {
  let node = NewNode(value)
  if (!exist(self, node)) {
    return
  }

  if (index < 0) {
    changeHead(node)
    return
  } 

  if (index > self.size) {
    changeTail(node)
    return
  }

  let prev = self.head
  let target = self.head
  for (let i = 0; i < self.size; i = i + 1) {
    if (index == i) {
      self.do_index_insert(prev, target, node)
      return
    }
    prev = target
    target = target.next
  }
}

fn (self SingleLinkedList) do_index_insert(prev, move, new) {
  prev.next = new
  new.next = move
}
*/


fn (self SingleLinkedList) print() {
  for(let n = self.head; n ; n = n.next) {
    puts(n.value)
  }
}
export SingleLinkedList
