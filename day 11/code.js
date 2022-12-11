const fs = require("fs");

const file = fs.readFileSync("test.txt", "utf-8");

const op_fn = (code) => {
  const code_val = code[2] !== "old" ? parseInt(code[2], 10): "old";
  if(code[1] == "*") {
   return (val) => BigInt(val)*((code_val === "old") ? BigInt(val) : BigInt(code_val));
  } else if(code[1] == "+") {
   return (val) => BigInt(val)+((code_val === "old") ? BigInt(val) : BigInt(code_val));
  } else if(code[1] == "/") {
   return (val) => BigInt(val)/((code_val === "old") ? BigInt(val) : BigInt(code_val));
  } else if(code[1] == "-") {
   return (val) => BigInt(val)-((code_val === "old") ? BigInt(val) : BigInt(code_val));
  }
}

const ls = []
let n = -1;
file.split("\n").forEach((f) => {
  if(f.startsWith("Monkey")) {
    ls.push({});
    n+=1;
  } else {
    const [inst, amt] = f.split(": ");
    if(inst.includes("Starting")) {
     ls[n]["items"] = [...amt.split(", ").map(i => parseInt(i, 10))];
    } else if(inst.includes("Operation")) {
      ls[n]["op"] = op_fn(amt.split(" = ")[1].split(" "));
    } else if(inst.includes("Test")) {
      ls[n]["divisible"] = parseInt(amt.split(" ")[2], 10);
    } else if(inst.includes("If true")) {
      ls[n]["true"]  = parseInt(amt.split(" ")[3], 10)
    } else if(inst.includes("If false")) {
      ls[n]["false"] = parseInt(amt.split(" ")[3], 10)
    }
  }
});

count = []
ls.forEach(l => {
  count.push(0)
})
for(let i = 0; i < 10000; i++) {

  for(let j = 0; j < ls.length; j++) {
    for(let k = 0; k < ls[j]["items"].length; k++) {
      const itemAtIndex = ls[j]["items"][k];
      // const value = Math.floor(ls[j]["op"](itemAtIndex)/3);
      const value= ls[j]["op"](itemAtIndex)
      // console.log(value)
      count[j]+=1;
      if(value  % BigInt(ls[j]["divisible"]) === 0) {
        ls[ls[j]["true"]]["items"].push(value)
      } else {
        ls[ls[j]["false"]]["items"].push(value)
      }
     }
  ls[j]["items"] = []
  }
  
}
console.log(count)
console.log(count.sort((a, b) => b -a).slice(0, 2).reduce((a, b) => a*b));