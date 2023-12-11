const fs = require('fs')
var data = fs.readFileSync('day01.in', 'utf8').trim()

function rocket(total, num) {
    return total + Math.floor(num/3)-2 
}

sol1 = data.split("\n").map(Number).reduce(rocket, 0)
console.log("Solution 1:", sol1)

function rocket2(total, num){
    var count = 0
    while(Math.floor(num/3)-2>0){
        num = Math.floor(num/3)-2
        count += num
    }
    return total + count
}

sol2 = data.split("\n").map(Number).reduce(rocket2, 0)
console.log("Solution 2:", sol2)
