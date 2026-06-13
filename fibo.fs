fn fib(n) {
  if(n <= 1) {
      return n
  }

  return fib(n - 1) + fib(n - 2)
}


fn makeFib() {
  let cache = {}

  fn fib(n) {
    if(cache[n]) {
      return cache[n]
    }

    if (n <= 1) {
        return n
    }

    let result = fib(n - 1) + fib(n - 2)
    cache[n] = result
    return result
  }

  return fib
}


let result = fib(10)
puts(result)


let fib2 = makeFib()
result = fib2(10)
puts(result)
