import { read } from '../shared/read'

const input_path = `${__dirname}/input`
const data = read(input_path)
const lines = data.split('\n')

const calories: number[] = []
let acc = 0
for (const calory of lines) {
  if (calory) {
    acc += Number.parseInt(calory, 10)
  } else {
    calories.push(acc)
    acc = 0
  }
}

// Part 1
console.log(Math.max(...calories))

// Part 2
const total_top_three_calories = calories.sort((a, b) => b - a)
  .slice(0, 3)
  .reduce((sum, current) => sum + current, 0)

console.log(total_top_three_calories)