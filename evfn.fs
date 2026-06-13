let makefn = fn(n){
  if(n <= 1) {
    return 1
  }

  return makefn(n-1)
}

let ff = makefn(2)

puts(ff)


