const fs = require("fs");
const readline = require("readline");

async function processLineByLine() {
  const fileStream = fs.createReadStream("input.txt");

  const rl = readline.createInterface({
    input: fileStream,
  });
  const grid = [];
  let i = 0;
  for await (const line of rl) {
    grid[i] = [];
    let j = 0;
    for (const char of line) {
      grid[i][j] = char;
      j++;
    }
    i++;
  }

  let total = 0;
  grid.forEach((row, rowIdx) => {
    row.forEach((col, colIdx) => {
      // if (col == "X") {
      //   total += checkRight(colIdx, row);
      //   total += checkDown(grid, rowIdx, colIdx);
      //   total += checkLeft(colIdx, row)
      //   total += checkUp(grid, rowIdx, colIdx);
      //   total+= checkDiagonals(grid, rowIdx, colIdx)
      // }
      if (col == "A") {
        let match = 0
        for(const word of checkDiagonalsP2(grid, rowIdx, colIdx)) {
          if(word == "MAS") {
            match++;
          }
        }
        total += +(match >= 2)
      }
    });
  });
  console.log(total);
}

const VAL = "XMAS";
function checkRight(colIdx, line, word = VAL) {
  const lineCheck = line.slice(colIdx, colIdx + word.length).join("");
  return word === lineCheck;
}

function checkDown(grid, rowIdx, colIdx) {
  const line = extractVertical(grid, colIdx);
  return checkRight(rowIdx, line);
}

function checkLeft(colIdx, line) {
  const newLine = rotateLine(line)
  return checkRight(line.length - 1 - colIdx , newLine)
}

function checkUp(grid, rowIdx, colIdx) {
  const line = extractVertical(grid, colIdx);
  return checkLeft(rowIdx, line)
}

function checkDiagonalsP2(grid, rowIdx, colIdx) {
  const words = []
  if(rowIdx - 1 >= 0 && colIdx -1 >= 0 && rowIdx + 1 < grid.length && colIdx + 1 < grid[0].length) {
    words.push(`${grid[rowIdx-1][colIdx-1]}${grid[rowIdx][colIdx]}${grid[rowIdx+1][colIdx+1]}`)
    words.push(`${grid[rowIdx-1][colIdx+1]}${grid[rowIdx][colIdx]}${grid[rowIdx+1][colIdx-1]}`)
    words.push(rotateLine(words[0]).join(""))
    words.push(rotateLine(words[1]).join(""))
  }
  return words
}

function checkDiagonals(grid, rowIdx, colIdx) {
  const left = getLeftDiagonal(grid, rowIdx, colIdx)
  const right = getRightDiagonal(grid, rowIdx, colIdx)
  return (
        checkRight(0, left.before)
      + checkRight(0, left.after)
      + checkRight(0, right.before)
      + checkRight(0, right.after)
  )
}

function getLeftDiagonal(grid, rowIdx, colIdx) {
  const after = [];
  const before = [];

  const afterState = [rowIdx, colIdx]
  const beforeState = [rowIdx, colIdx]
  while(true) {
    let flag = 0;
    if(afterState[0] < grid.length && afterState[1] < grid[afterState[0]].length) {
      after.push(grid[afterState[0]][afterState[1]])
      afterState[0]++;
      afterState[1]++;
      flag = 1;
    }
    if(beforeState[0] >= 0 && beforeState[1] >= 0) {
      before.push(grid[beforeState[0]][beforeState[1]])
      beforeState[0]--;
      beforeState[1]--;
      flag =1;
    }
    if(!flag) {
      break;
    }
  }
  return {before, after}
}

//   x
//  x
// x

function getRightDiagonal(grid, rowIdx, colIdx) {
  const after = [];
  const before = [];

  const afterState = [rowIdx, colIdx]
  const beforeState = [rowIdx, colIdx]
  while(true) {
    let flag = 0;
    if(afterState[0] < grid.length && afterState[1] >= 0) {
      after.push(grid[afterState[0]][afterState[1]])
      afterState[0]++;
      afterState[1]--;
      flag = 1;
    }
    // console.log(beforeState)
    if(beforeState[0] >= 0 && beforeState[1] < grid[beforeState[0]].length) {
      before.push(grid[beforeState[0]][beforeState[1]])
      beforeState[0]--;
      beforeState[1]++;
      flag =1;
    }
    if(!flag) {
      break;
    }
  }
  return {before, after}
}

function rotateLine(line) {
  const newLine = []
  for(let i = line.length - 1; i >= 0; i--) {
    newLine.push(line[i])
  }
  return newLine
}


function extractVertical(grid, column) {
  const line = [];
  for (i = 0; i < grid.length; i++) {
    line[i] = grid[i][column];
  }
  return line;
}

processLineByLine();
