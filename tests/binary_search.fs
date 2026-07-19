let data = [1,2,3,4,5]

fn binarySearch(data, target) {
  let left = 0
  let right = len(data)

  while(left < right) {
    let md = (left + right) / 2

    if (data[md] > target) {
      right = md - 1
    }else if (data[md] < target) {
      left = md + 1
    }else {
      return md
    }
  }
}

let a = binarySearch(data, 10)
puts(a)
