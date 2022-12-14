const fs = require("fs")

const lines = []
const file = fs.readFileSync(`${process.argv[2] ?? "sample"}.txt`, "utf-8")


let sea_level = 5
function isPointInLines(point) {
  // console.log(point, sea_level)
  if(point[1] === sea_level + 2) {
    return 1
  }
  for(let i of lines) {
    //horizontal line
    if(i[0][1] == i[1][1]) {
      if(point[1] == i[0][1] && point[0] >= Math.min(i[0][0], i[1][0]) && point[0] <= Math.max(i[1][0], i[0][0])) {
        return 1
      }
    } else {
      if(point[0] == i[0][0] && point[1] >= Math.min(i[0][1], i[1][1]) && point[1] <= Math.max(i[1][1], i[0][1])) {
        return 1
      }
    }
  }
  return 0
}

let n = 0
file.split("\n").forEach((f) => {
  const splitfn = f.split(" -> ")
  f.split(" -> ").forEach((fl, i) => {
    xyval = fl.split(",").map(c => parseInt(c))
    if(xyval[1] > sea_level) {
      sea_level = xyval[1]
    }
    if (i != 0) {
      lines[n - 1].push(xyval)
    }
    if (i != splitfn.length - 1) {
      lines.push([xyval])
      n++
    }
  })
})

let curr_pos = [500, 0]
let count = 0
while(true) {
  if(!isPointInLines([curr_pos[0], curr_pos[1]+1])) {
    curr_pos[1]++
  } else if(!isPointInLines([curr_pos[0]-1, curr_pos[1]+1]))  {
    curr_pos[0]--
    curr_pos[1]++
  } else if(!isPointInLines([curr_pos[0]+1, curr_pos[1]+1])) {
    curr_pos[0]++
    curr_pos[1]++
  } else {

    count++
    if(curr_pos[0] == 500 && curr_pos[1] == 0) {
      break
    }
    lines.push([curr_pos, curr_pos])
    // console.log(lines)
    // break
    curr_pos = [500, 0]
  }

  // if(curr_pos[1] == sea_level+2) {
  //   lines.push([curr_pos, curr_pos])
  //   count++
  //   curr_pos = [500, 0]
  // }
}
console.log( count)
// console.log(lines)


