const fs = require("fs")

const file = fs.readFileSync("actual.txt", "utf-8");

function Graph() {
  let neighbors = this.neighbors = {};

  this.addEdge = function(u, v) {
    if (neighbors[u] === undefined) {
      neighbors[u] = [];
    }
    neighbors[u].push(v);
  };

  return this;
}
const ls = []
function bfs(graph, source, end) {
  const queue = [{ vertex: source, count: 0 }]
  const visited = { [source]: true }
  let tail = 0;
  while (tail < queue.length) {
    const { count, vertex } = queue[tail];
    tail++
    if (vertex === end) {
      return count
    }
    graph.neighbors[vertex]?.forEach(v => {
      if (!visited[v]) {
        visited[v] = true;
        queue.push({ vertex: v, count: count + 1 });
      }
    });
  }
}

const graph = new Graph();
const f_arr = file.split("\n").map(f => f.split(""));
let start;
let end;
for (let i = 0; i < f_arr.length; i++) {
  for (let j = 0; j < f_arr[0].length; j++) {

    if (f_arr[i][j] == "S") {
      start = `${i}-${j}`;
      f_arr[i][j] = 'a';
    } else if (f_arr[i][j] === "E") {
      end = `${i}-${j}`;
      f_arr[i][j] = 'z';
    }
    if (f_arr[i][j] == "a") {
      ls.push(`${i}-${j}`)
    }
    for (let x = -1; x < 2; x++) {
      for (let y = -1; y < 2; y++) {
        if ((x != 0 && y != 0) || x == y) {
          continue;
        }
        if (x + i < 0 || y + j < 0 || x + i >= f_arr.length || y + j >= f_arr[0].length) {
          continue;
        }
        if (f_arr[i + x][j + y].charCodeAt() <= f_arr[i][j].charCodeAt() + 1) {
          graph.addEdge(`${i}-${j}`, `${i + x}-${j + y}`);
        }
      }
    }
  }
}

// console.log(bfs(graph, start, end))
const paths = []
ls.forEach(l => {
  paths.push(bfs(graph, l, end));
})
console.log(Math.min(...(paths.filter(d => d))))
