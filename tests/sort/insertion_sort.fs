fn insertion_sort(arr) {
  let length = len(arr)

  for (let i = 1; i < length; i = i + 1) {
    let key = arr[i]

    let j = i - 1
    for (; j >= 0 && arr[j] > key; j = j - 1) {
      arr[j + 1] = arr[j]
    }

    arr[j + 1] = key
  }
}


let arr = [3,4,5,1,2,8,7,6]
puts(arr)
insertion_sort(arr)
puts(arr)
