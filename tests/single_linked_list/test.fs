import {SingleLinkedList} from "./single_linked_list"

let link = SingleLinkedList{}
let a = link.init()
puts(link.head, link.tail,link.size)

link.prepend(1)
link.prepend(2)

link.print()
