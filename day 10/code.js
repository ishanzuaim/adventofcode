const fs = require("fs");

const file = fs.readFileSync("input.txt", "utf8");
const ls = [20, 60, 100, 140, 180, 220];

let cycle = 0;
let overall = 0;
let add = 1;
file.split("\n").forEach((f) => {
  let flag = 0;
  if (f === "noop") {
    cycle += 1;
  } else {
    const add_amt = parseInt(f.split(" ")[1]);
    cycle += 1;
    if (ls.includes(cycle)) {
      console.log(add * cycle, cycle, add);
      overall += add * cycle;
    }
    cycle += 1;
    if (ls.includes(cycle)) {
      console.log(add * cycle, cycle, add);
      overall += add * cycle;
      flag = 1;
    }
    add += add_amt;
  }

  if (!flag && ls.includes(cycle)) {
    console.log(add * cycle, cycle, add);

    overall += add * cycle;
  }
});

console.log(overall);
