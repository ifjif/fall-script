fn bubble_sort(arr) {
  let n = len(arr)

  for (let i = 0; i < n-1; i = i + 1) {
    let swapped = false

    for (let j = 0;j < n-i-1; j = j + 1) {
      if (arr[j] > arr[j+1]) {
        let tem = arr[j]
        arr[j] = arr[j+1]
        arr[j+1] = tem
        swapped = true
      }
    }

    if (!swapped) {
      return
    }
  }
}

let arr = [3,4,5,1,2,8,7,6]
puts(arr)
bubble_sort(arr)
puts(arr)
