fn shell_sort(arr) {
  let n = len(arr)

  for (let gap = n / 2; gap > 0 ; gap = gap / 2) {

    for (let i = gap; i < n; i = i + 1) {
      let tem = arr[i]
      let j = i
/*
n:8
gap:4
04 15 26 37
gap:2
02 13 24 35 46 57

0246 1357
gap:1
01 12 23 34 56 67
*/
      for(;j >= gap && arr[j - gap] > arr[j]; j = j - gap) {
        arr[j] = arr[j - gap]
      }

      arr[j] = tem
    }
  }
}

let arr = [3,4,5,1,2,8,7,6]
puts(arr)
shell_sort(arr)
puts(arr)
