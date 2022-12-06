import { read } from '../shared/read'

const input_path = `${__dirname}/input`
const lines = read(input_path)

const isMoveLine = (s: string): boolean => s.includes('move')
const isCrateLine = (s: string): boolean => s.includes('[')

interface Move {
  start: number
  end: number
  count: number
}

let crates: Array<Array<string>> = []
let moves: Array<Move> = []
lines.forEach(line => {
  if (isCrateLine(line)) {
    line.split('').forEach((char, index) => {
      if (char !== '[' && char !== ']' && char !== ' ') {
        const crate_number = index > 1 ? (index - 1) / 4 : index - 1
        if (!crates[crate_number]) {
          crates[crate_number] = []
        }

        crates[crate_number].unshift(char)
      }
    })
  }

  if (isMoveLine(line)) {
    const splitted = line.split(' ')
    const move: Move = {
      count: Number.parseInt(splitted[1]),
      start: Number.parseInt(splitted[3]) - 1,
      end: Number.parseInt(splitted[5]) - 1,
    }

    moves.push(move)
  }
})

moves.forEach(move => {
  const crates_to_move = []
  for (let i = 0; i < move.count; i++) {
    const crate = crates[move.start].pop()
    if (crate) {
      crates_to_move.unshift(crate)
    }
  }

  crates_to_move.forEach(crate => {
    crates[move.end].push(crate)
  })
})

const part1_result = crates.reduce((acc, crate) => {
  return acc + crate[crate.length - 1]
}, '')

// Part 1
console.log('Part 1 =', part1_result)
