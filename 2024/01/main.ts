import fs from "node:fs";
function convertToArrays(data: string): number[][] {
  const lines = data.split("\n");
  let first = [];
  let second = [];
  lines.forEach((item: string) => {
    let parts = item.split("   ");
    if (parts.length < 2) return;
    first.push(Number(parts[0]));
    second.push(Number(parts[1]));
  });
  return [first, second];
}

function convertToMap(data: number[]):Map<number,number> {
  let res = new Map<number,number>()
  data.forEach((item) => {
    let mapItem = res.get(item) || 0
    res.set(item,mapItem+1)
  })
  return res
}

function solveSecond(list1: number[], list2: number[]):number {
  let similarity = 0
  let map = convertToMap(list2)

  list1.forEach(item => {
    similarity += (map.get(item) || 0) * item
  })
  return similarity
}

function solve(list1: number[], list2: number[]):number {
  list1.sort();
  list2.sort();
  let distance = 0;
  list1.forEach((item, index) => {
    distance += Math.abs(item - list2[index])
  });
  return distance
}

const data = fs.readFileSync("input.txt", "utf8");
const arrs = convertToArrays(data);
console.log(`distance: ${solve(arrs[0], arrs[1])}`)
console.log(`similarity: ${solveSecond(arrs[0], arrs[1])}`)
