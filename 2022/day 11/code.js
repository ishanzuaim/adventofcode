// const fs = require("fs");

// const file = fs.readFileSync("sample.txt", "utf-8");
test_ls = [
  {
    "items": [79, 98],
    "op": (val) => val * 19,
    "divisible": 23,
    "true": 2,
    "false": 3,
    "insp": 0
  }, {
    "items": [54, 65, 75, 74],
    "op": (val) => val + 6,
    "divisible": 19,
    "true": 2,
    "false": 0,
    "insp": 0
  }, {
    "items": [79, 60, 97],
    "op": (val) => val * val,
    "divisible": 13,
    "true": 1,
    "false": 3,
    "insp": 0
  }, {
    "items": [74],
    "op": (val) => val + 3,
    "divisible": 17,
    "true": 0,
    "false": 1,
    "insp": 0
  }]
ls = [
  {
    "items": [98, 70, 75, 80, 84, 89, 55, 98],
    "op": (val) => val * 2,
    "divisible": 11,
    "true": 1,
    "false": 4,
    "insp": 0
  },
  {
    "items": [59],
    "op": (val) => val * val,
    "divisible": 19,
    "true": 7,
    "false": 3,
    "insp": 0
  },
  {
    "items": [77, 95, 54, 65, 89],
    "op": (val) => val + 6,
    "divisible": 7,
    "true": 0,
    "false": 5,
    "insp": 0
  },
  {
    "items": [71, 64, 75],
    "op": (val) => val + 2,
    "divisible": 17,
    "true": 6,
    "false": 2,
    "insp": 0
  },
  {
    "items": [74, 55, 87, 98],
    "op": (val) => val * 11,
    "divisible": 3,
    "true": 1,
    "false": 7,
    "insp": 0
  },
  {
    "items": [90, 98, 85, 52, 91, 60],
    "op": (val) => val + 7,
    "divisible": 5,
    "true": 0,
    "false": 4,
    "insp": 0
  },
  {
    "items": [99, 51],
    "op": (val) => val + 1,
    "divisible": 13,
    "true": 5,
    "false": 2,
    "insp": 0
  },
  {
    "items": [98, 94, 59, 76, 51, 65, 75],
    "op": (val) => val + 5,
    "divisible": 2,
    "true": 3,
    "false": 6,
    "insp": 0
  },
]

const prod = ls.map(x => x.divisible).reduce((a, b) => a * b, 1);
for (let i = 0; i < 10_000; i++) {

  for (const l of ls) {
    while (l.items.length >= 1) {
      const item = l.items.shift();
      // const value = Math.floor(ls[j]["op"](item)/3);
      let value = l.op(item) % prod
      l.insp += 1;
      if (value % l.divisible === 0) {
        ls[l.true].items.push(value)
      } else {
        ls[l.false].items.push(value)
      }
    }
  }
}
console.log(ls.map(x => x.insp).sort((a, b) => b - a).slice(0, 2).reduce((a, b) => a * b));