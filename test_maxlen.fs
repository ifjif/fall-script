fn maxLength(str) {
  let occurred = {}

  let left = 0
  let max = 0
  let startIdx = 0

  let length = len(str)

  for (let i = 0; i < length; i = i + 1) {
    let ch = str[i]
    let idx = occurred[ch]

    if ( idx != null && idx >= left) {
      left = idx + 1
    }

    occurred[ch] = i

    let currentLen = i - left + 1
    if (currentLen > max) {
       max = currentLen
       startIdx = left
    }
  }

  puts(max)
  puts(occurred)
  puts(str[startIdx:startIdx+max])
}

maxLength("abca")
