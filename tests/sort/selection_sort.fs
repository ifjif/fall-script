fn selection_sort(arr) {
  let n = len(arr)

  for(let i = 0; i < n - 1; i = i + 1) {
    let min = i

    for(let j = i + 1; j < n; j = j + 1) {
      if(arr[j] < arr[min]) {
        min = j
      }
    }

    let tem = arr[i]
    arr[i] = arr[min]
    arr[min] = tem
  }
}

let arr = [3,4,5,1,2,8,7,6]
puts(arr)
selection_sort(arr)
puts(arr)
