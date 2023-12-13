import { read } from '../shared/read'

const input_path = `${__dirname}/input`
const lines = read(input_path)

const parseStringNumber = (value: string): number => {
  switch (value) {
    case '1':
    case 'one':
      return 1

    case '2':
    case 'two':
      return 2

    case '3':
    case 'three':
      return 3

    case '4':
    case 'four':
      return 4

    case '5':
    case 'five':
      return 5

    case '6':
    case 'six':
      return 6

    case '7':
    case 'seven':
      return 7

    case '8':
    case 'eight':
      return 8

    case '9':
    case 'nine':
      return 9

    default:
      throw new Error('string cannot be parsed')
  }
}

const findCalibrationValues = (line: string, regexp: RegExp): { first: number, last: number } => {
  const all_match = [...line.matchAll(regexp)]
  
  let max_index = -1
  let last_value = ''

  let min_index = Infinity
  let first_value = ''

  for (const match of all_match) {
    if (!match.index && match.index !== 0) {
      continue
    }

    if (match.index > max_index) {
      max_index = match.index
      last_value = match[1]
    } 
    if (match.index < min_index) {
      min_index = match.index
      first_value = match[1]
    }
  }

  return {
    first: parseStringNumber(first_value),
    last: parseStringNumber(last_value)
  }
}

const resolve = (regexp: RegExp): number => {
  return lines
    .map(line => {
      const { first, last} = findCalibrationValues(line, regexp)
      return Number.parseInt(`${first}${last}`)
    })
    .reduce((acc, value) => acc + value, 0)
}

const resolveFirstChallenge = () => {
  const regexp = new RegExp(/(?=(0|1|2|3|4|5|6|7|8|9))/, 'gi')
  return resolve(regexp)
}

const resolveSecondChallenge = () => {
  const regexp = new RegExp(/(?=(1|2|3|4|5|6|7|8|9|one|two|three|four|five|six|seven|eight|nine))/, 'gi')
  return resolve(regexp)
}

// Part 1
console.log(resolveFirstChallenge())

// Part 2
console.log(resolveSecondChallenge())