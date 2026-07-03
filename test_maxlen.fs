fn maxLength(str) {

  str[0] = 1
  let occur = {}

  let left = 0
  let max = 0

  let length = len(str)

  for (let i = 0; i < length; i = i + 1) {
    let ch = str[i]
    let idx = occur[ch]

    if ( idx != null && idx >=0 && idx >= left) {
      left = idx + 1
    }

    occur[ch] = i

    let currentLen = i - left + 1
    if (currentLen > max) {
       max = currentLen
    }
  }

  puts(max)
  puts(occur)
}

maxLength("abca")
