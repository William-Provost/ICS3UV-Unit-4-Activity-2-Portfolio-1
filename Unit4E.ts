/**
 * @author William
 * @version 1.0.0
 * @date 2025-11-27
 * @fileoverview This program calculates base^exponent using a while loop.
 */

// variables
let base: number = 0;
let exponent: number = 0;
let count: number = 0;
let result: number = 1;

// get base from user
base = parseInt(prompt("Enter the base: ") || "0");

// get exponent from user
exponent = parseInt(prompt("Enter the exponent: ") || "0");

// while loop multiplies base exponent times
count = 0;
while (count < exponent) {
  result = result * base;
  count++;
}

// display result
console.log(base + " raised to the power of " + exponent + " is: " + result);

console.log("\nDone.");
