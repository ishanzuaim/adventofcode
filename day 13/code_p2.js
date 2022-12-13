const fs = require("fs")

const file = fs.readFileSync(`${process.argv[2] ?? "sample"}.txt`, "utf-8")

const compare = (left, right) => {
  if (Array.isArray(left) && Array.isArray(right)) {
    for (let i = 0; i < Math.min(left.length, right.length); i++) {
      const result = compare(left[i], right[i]) 
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
  return left > right ? -1 : (left === right ? 0: 1)
}
const ls = []
file.split("\n").filter(d=>d).map(f => JSON.parse(f)).map(f  => {
  let n = 0;
  for(let i of ls) {
    if(compare(i, f) !== -1) {
      n++
    } else {
      break;
    }
  }
  ls.splice(n, 0, f)
})
const res = [[[2]], [[6]]].map(item => {
  for(let i =0; i < ls.length; i++) {
    if(compare(ls[i], item) === -1) {
      ls.splice(i+1, 0, item)
      console.log(i)
      return i+1;
    }
  }
}).reduce((a, b) => a*b)
console.log(res)
