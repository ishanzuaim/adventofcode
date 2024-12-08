const fs = require("fs");
const readline = require("readline");

async function processLineByLine() {
  const fileStream = fs.createReadStream("inp.txt");

  const rl = readline.createInterface({
    input: fileStream,
  });
  let count = 0;
  for await (const line of rl) {
    const arr = line.split(" ").map(Number);
    for (let i = 0; i < arr.length; i++) {
      const newArr = arr.filter((_, idx) => idx !== i);
      const isValid = isIncreasingOrDecreasing(newArr) && meetsParams(newArr);
      if (isValid) {
        count += 1;
        break;
      }
    }
  }
  console.log(count);
}

function isIncreasingOrDecreasing(arr) {
  let prev = arr[0];
  if (arr[0] > arr[1]) {
    for (const elem of arr) {
      if (prev < elem) {
        return false;
      }
      prev = elem;
    }
  } else {
    for (const elem of arr) {
      if (prev > elem) {
        return false;
      }
      prev = elem;
    }
  }
  return true;
}

function meetsParams(arr) {
  let prev = arr[0];
  let flag = 0;
  for (const elem of arr) {
    // __AUTO_GENERATED_PRINT_VAR_START__
    const diff = Math.abs(elem - prev);
    if (flag && (diff < 1 || diff > 3)) {
      return false;
    }
    if (!flag) {
      flag = 1;
    } else {
      prev = elem;
    }
  }
  return true;
}

processLineByLine();
