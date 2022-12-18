const fs = require("fs")

const file = fs.readFileSync(`${process.argv[2] ?? "sample"}.txt`, "utf-8")

function calc(n) {
  return [n.slice(2).split(",")[0], n.slice(2).split(", ")[1].slice(2)].map(x => parseInt(x))
}
const coords = []
let count = 0;
file.split("\n").forEach(f => {
  const [_, senspos, beacpos] = f.split("at ")
  const sensor = calc(senspos)
  const beacon = calc(beacpos)

  const range = Math.abs(sensor[0] - beacon[0]) + Math.abs(sensor[1] - beacon[1])

  let n = 0;
  let ls = [];
  for (let i = sensor[0] - range; i < range + sensor[0]; i++) {
    let y = sensor[1] - n
    // coords.add(`${i} ${y}`)
    y == 2_000_000 && ls.push(i)
    y = sensor[1] + n
    // coords.add(`${i} ${y}`)
    y == 2_000_000 && ls.push(i)

    if (i < sensor[0]) {
      n++
    } else {
      n--
    }
  }
coords.push(ls)
})

// console.log(coords)
// let count = 0
// for(let i =-1 ; i <= 26; i++) {
//   count += coords.has(`${i} 10`)
// }
const points = new Set();
for(let i of coords.filter(l => l.length == 2)) {
  for(let n = Math.min(i[0], i[1]); n < Math.max(i[0], i[1]); n++) {
    points.add(n)
  }
}

console.log(points.size)
