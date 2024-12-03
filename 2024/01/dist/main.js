"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const node_fs_1 = __importDefault(require("node:fs"));
function convertToArrays(data) {
    const lines = data.split("\n");
    let first = [];
    let second = [];
    lines.forEach((item) => {
        let parts = item.split("   ");
        if (parts.length < 2)
            return;
        first.push(Number(parts[0]));
        second.push(Number(parts[1]));
    });
    return [first, second];
}
function convertToMap(data) {
    let res = new Map();
    data.forEach((item) => {
        let mapItem = res.get(item) || 0;
        res.set(item, mapItem + 1);
    });
    return res;
}
function solveSecond(list1, list2) {
    let similarity = 0;
    let map = convertToMap(list2);
    list1.forEach(item => {
        similarity += (map.get(item) || 0) * item;
    });
    return similarity;
}
function solve(list1, list2) {
    list1.sort();
    list2.sort();
    let distance = 0;
    list1.forEach((item, index) => {
        distance += Math.abs(item - list2[index]);
    });
    return distance;
}
const data = node_fs_1.default.readFileSync("input.txt", "utf8");
const arrs = convertToArrays(data);
console.log(`distance: ${solve(arrs[0], arrs[1])}`);
console.log(`similarity: ${solveSecond(arrs[0], arrs[1])}`);
//# sourceMappingURL=main.js.map