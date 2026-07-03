fn delimiter() {
  puts("----------------------------------")
}
let arr = [1,2,3,4,5,6]
puts(arr)
delimiter()

puts("忽略 start, 忽略 end , step 默认")
let arr2 = arr[:]
puts(arr2)
delimiter()

puts("start = 1, 忽略 end, step 默认")
let arr3 = arr[1:]
puts(arr3)
delimiter()

puts("修改 arr3[0] = 22, 则 arr[1], arr2[1] 都变化")
arr3[0] = 22
puts(arr, arr2, arr3)
delimiter()

puts("忽略 start, 忽略 end, step = -1")
let arr4 = arr[::-1]
puts(arr4)
delimiter()

puts("修改 arr4[0] = 66, 则 arr, arr2, arr3 不变")
arr4[0] = 66
puts(arr4, arr, arr2, arr3)
delimiter()

puts("忽略 start, 忽略 end, step = 2")
let arr5 = arr[::2]
puts(arr5)
delimiter()

puts("字符串 slice")
puts("忽略 start, 忽略 end, step = -1")
let arr6 = "abcde"
let arr7 = arr6[::-1]
puts(arr6, arr7)
delimiter()

puts("start = 1")
let arr8 = arr6[1:]
puts(arr8)
