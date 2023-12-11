const fs = require('fs')
var inp = fs.readFileSync('day02.in', 'utf8').trim().split(',').map(Number)

function solve(data, val1, val2){
    data[1] = val1
    data[2] = val2
    for(var i = 0; i < data.length; i+=4){
        if(data[i] == 99) break
        var k = data[i+1]
        var l = data[i+2]
        var m = data[i+3]
        switch(data[i]) {
            case 1:
                var num = data[k] + data[l]
                break
            case 2:
                var num = data[k] * data[l]
                break
            }
        data[m] = num
    }
    return data[0]
}

function findnounverb(){
    for(var i = 0; i < 99; i++){
        for(var j = 0; j < 99; j++){
            if(solve(Array.from(inp), i, j) == 19690720) return 100*i + j
        }
    }
}

console.log("Solution 1:", solve(Array.from(inp), 12,2))
console.log("Solution 2:", findnounverb())
