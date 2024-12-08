const fs = require("fs");

const file = fs.readFileSync("input.txt", "utf8");

function screendraw(curr_pos, draw, cycle) {
  if (cycle > 239) {
    return;
  }
  const index = Math.floor((cycle - 1) / 40);
  const mypos = cycle % 40;
  if ([curr_pos - 1, curr_pos, curr_pos + 1].includes(mypos - 1)) {
    draw[index].push("#");
  } else {
    draw[index].push(".");
  }
}

let cycle = 0;
let add = 1;
const draw = [[], [], [], [], [], []];

file.split("\n").forEach((f) => {
  if (f === "noop") {
    cycle += 1;
    screendraw(add, draw, cycle);
  } else {
    const add_amt = parseInt(f.split(" ")[1]);
    cycle += 1;
    screendraw(add, draw, cycle);
    cycle += 1;
    screendraw(add, draw, cycle);
    add += add_amt;
  }
});

draw.forEach((d) => {
  console.log(d.join(""));
});
