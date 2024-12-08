const fs = require("fs");
const readline = require("readline");

async function processLineByLine() {
  const fileStream = fs.createReadStream("inp.txt");

  const rl = readline.createInterface({
    input: fileStream,
  });
  let lines = "";
  for await (const line of rl) {
    lines += line + "\n";
  }
  const output = lines.match(/mul\(\d+,\d+\)|do\(\)|don't\(\)/g);
  let total = 0;
  let flag = 1;
  for (const op of output) {
    if (op === "do()") {
      flag = 1;
      continue;
    } else if (op === "don't()") {
      flag = 0;
      continue;
    }
    if (flag) {
      const vals = op.match(/\d+/g).map(Number);
      total += vals[0] * vals[1];
    }
  }
  console.log(total);
}

processLineByLine();
