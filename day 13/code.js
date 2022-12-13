const fs = require("fs")

const file = fs.readFileSync(`${process.argv[2] ?? "sample"}.txt`, "utf-8")

const compare = (left, right) => {
  // console.log(left, right, left > right)
  if (Array.isArray(left) && Array.isArray(right)) {
    for (let i = 0; i < Math.min(left.length, right.length); i++) {
      const result = compare(left[i], right[i]) 
      // console.log({result})
      if (result) {
        return result
      }
    }
    return left.length > right.length ? -1 : (left.length === right.length ? 0: 1)
  } else if (Array.isArray(left)) {
    return compare(left, [right])
  } else if (Array.isArray(right)) {
    return compare([left], right)
  }
  // console.log({left, right})
  return left > right ? -1 : (left === right ? 0: 1)
}


const result = file.split("\n\n").map((f, i) => {
  [left, right] = f.split("\n").map(f => JSON.parse(f))
  // console.log(compare(left, right))
  return compare(left, right) == -1 ? 0 : i+1
}).reduce((a, b) => a+b)

console.log(result)

