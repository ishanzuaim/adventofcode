const fs = require("fs");
const readline = require("readline");

async function processLineByLine() {
  const fileStream = fs.createReadStream("inp.txt");

  const rl = readline.createInterface({
    input: fileStream,
  });
  const arr1 = [];
  const arr2 = [];
  for await (const line of rl) {
    const [a, b] = line.split("  ");
    arr1.push(+a);
    arr2.push(+b);
  }
  arr1.sort();
  arr2.sort();
  let total = 0;

  arr1.forEach((elem) => {
    total += elem * findHowMany(elem, arr2);
  });
  console.log(total);
}

function findHowMany(elem, arr) {
  let count = 0;
  for (const i of arr) {
    if(elem === i) {
      count++;
    }
  }
  return count;
}

processLineByLine();
