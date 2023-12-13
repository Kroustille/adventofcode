import assert from 'assert'
import { read } from '../shared/read'

const input_path = `${__dirname}/input`
const lines = read(input_path)

interface Pipe {
  key: string
  connected_pipes: string[]
}

const getKey = (x: number, y: number): string => {
  return `${x};${y}`
}

const pipes: Record<string, string[]> = {}
let animal_key: string = ''

lines.forEach((line, line_index) => {
  line.split('').forEach((character, character_index) => {
    const key =  getKey(character_index, line_index) 
    let connected_pipes = []

    switch (character) {
      case '|':
        connected_pipes.push(
          getKey(character_index, line_index+1),
          getKey(character_index, line_index-1)
        )
        break

      case '-':
        connected_pipes.push(
          getKey(character_index-1, line_index),
          getKey(character_index+1, line_index)
        )
        break

      case 'L':
        connected_pipes.push(
          getKey(character_index, line_index-1),
          getKey(character_index+1, line_index)
        )
      break

      case 'J':
        connected_pipes.push(
          getKey(character_index, line_index-1),
          getKey(character_index-1, line_index)
        )
      break

      case '7':
        connected_pipes.push(
          getKey(character_index, line_index+1),
          getKey(character_index-1, line_index)
        )
      break

      case 'F':
        connected_pipes.push(
          getKey(character_index, line_index+1),
          getKey(character_index+1, line_index)
        )
      break

      case 'S':
        animal_key = key
        break;
    }

    if (connected_pipes.length) {
      pipes[key] = connected_pipes
    }
  })
})

const animal_pipe_keys = Object.keys(pipes).filter(key => {
  return pipes[key].find(connected_pipe => connected_pipe === animal_key)
})

pipes[animal_key] = animal_pipe_keys

const seen_pipes: Record<string, number> = {}
const pipes_to_process: string[] = []

let distance = 0
seen_pipes[animal_key] = distance

pipes_to_process.push(animal_key)

while(pipes_to_process.length) {
  const current_pipe_key = pipes_to_process.shift()
  assert(current_pipe_key)

  const connected_pipe_keys = pipes[current_pipe_key]

  connected_pipe_keys
    .filter(connected_pipe_key => seen_pipes[connected_pipe_key] === undefined)
    .forEach(connected_pipe_key => {
    if (seen_pipes[connected_pipe_key] === undefined) {
      pipes_to_process.push(connected_pipe_key)
      seen_pipes[connected_pipe_key] = seen_pipes[current_pipe_key]+1
    }
  })
}

console.log(Math.max(...Object.values(seen_pipes)))