import {LinkedList} from "./test_linkedlist"

let ll = LinkedList{head:null, tail:null}

ll.append(1)
ll.append(2)

for (let a = ll.head; a; a = a.next) {
  puts(a.value)
}
