import { read } from '../shared/read'

const input_path = `${__dirname}/input`
const lines = read(input_path)

const isFinalSequence = (sequence: number[]): boolean => {
    return sequence.every(value => value === 0)
}

const buildSequences = (line: string): number[][] => {
  const starting_sequence = line.split(' ').map(character => Number.parseInt(character, 10))
  const sequences = [starting_sequence]
  
  while (!isFinalSequence(sequences[sequences.length - 1])) {
    const last_sequence = sequences[sequences.length - 1]
    const new_sequence: number[] = []
    for (let i = 0 ; i < last_sequence.length - 1 ; i++) {
      new_sequence.push(last_sequence[i+1] - last_sequence[i])
    }
  
    sequences.push(new_sequence)
  }
  
  return sequences
}

// const part1_result = lines.reduce((acc, line) => {
//   const sequences = buildSequences(line)

//   for (let i = sequences.length - 2 ; i >= 0 ; i--) {
//     const current_sequence = sequences[i]
//     const last_sequence = sequences[i + 1]
//     current_sequence.push(last_sequence[last_sequence.length - 1] + current_sequence[current_sequence.length - 1])
//   }
  
//   const sequence_value = sequences[0][sequences[0].length - 1] 
//   console.log(sequence_value)
//   return acc + sequence_value
// }, 0)

// console.log(part1_result)

const part2_result = lines.reduce((acc, line) => {
  const sequences = buildSequences(line)

  for (let i = sequences.length - 2 ; i >= 0 ; i--) {
    const current_sequence = sequences[i]
    const last_sequence = sequences[i + 1]
    current_sequence.unshift(current_sequence[0] - last_sequence[0])
  }
  
  const sequence_value = sequences[0][0] 
  // console.log(sequence_value)
  return acc + sequence_value
}, 0)

console.log(part2_result)