import { read } from '../shared/read'

const input_path = `${__dirname}/input`
const lines = read(input_path)
const datastream = lines[0]

let last_characters: string[] = []
const marker_size = 14

const index = datastream.split('').findIndex(character => {
  if (last_characters.length === marker_size) {
    last_characters = last_characters.slice(1)
  }
  if (last_characters.length < marker_size) {
    last_characters.push(character)
  }

  return new Set(last_characters).size === marker_size
})

// Part 1
console.log('Part 1 =', index + 1, datastream.slice(index))
