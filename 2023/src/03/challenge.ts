import { read } from '../shared/read'

const input_path = `${__dirname}/input`
const lines = read(input_path)

interface Part {
  line: number
  start: number
  end: number
  value: number
}

interface Symbol {
  line: number
  position: number
  character: string
}

const parts: Part[] = []
const symbols: Symbol[] = []

lines.forEach((line, line_index) => {
  let part_value = ''
  let start = -1
  line.split('').forEach((character, character_index) => {
    if (character === '.') {
      if (part_value) {
        const part: Part = {
          line: line_index,
          start,
          end: character_index - 1,
          value: Number.parseInt(part_value, 10)
        }
        parts.push(part)
        part_value = ''
        start = -1
      }

      return
    }

    const is_number = character.match(/[0-9]/)
    if (is_number) {
      if (start === -1) {
        start = character_index
      }
      part_value = `${part_value}${character}`
      return
    }

    const symbol: Symbol = {
      line: line_index,
      position: character_index,
      character
    }
    symbols.push(symbol)
    if (part_value) {
      const part: Part = {
        line: line_index,
        start,
        end: character_index - 1,
        value: Number.parseInt(part_value, 10)
      }
      parts.push(part)
      part_value = ''
      start = -1
    }
  })
  
  if (part_value) {
    const part: Part = {
      line: line_index,
      start,
      end: line.length - 1,
      value: Number.parseInt(part_value, 10),
    }
    parts.push(part)
  }
})

// const good_parts: number[] = []
// symbols.forEach(symbol => {
//   [
//     {line: symbol.line - 1, position: symbol.position - 1},
//     {line: symbol.line - 1, position: symbol.position},
//     {line: symbol.line - 1, position: symbol.position + 1},

//     {line: symbol.line, position: symbol.position - 1},
//     {line: symbol.line, position: symbol.position + 1},

//     {line: symbol.line + 1, position: symbol.position - 1},
//     {line: symbol.line + 1, position: symbol.position},
//     {line: symbol.line + 1, position: symbol.position + 1},
//   ].forEach(linked => {
//     const linked_part_index = parts.findIndex(part => {
//       return part.line === linked.line && linked.position >= part.start && linked.position <= part.end 
//     })
//     if (linked_part_index > -1) {
//       good_parts.push(parts[linked_part_index].value)
//       parts.splice(linked_part_index, 1)
//     }
//   })
// })

// console.log(good_parts.reduce((acc, value)=> acc + value, 0))

const values: number[] = []
symbols.forEach(symbol => {
  const symbol_parts: { index: number, value: number }[] = [];
  [
    {line: symbol.line - 1, position: symbol.position - 1},
    {line: symbol.line - 1, position: symbol.position},
    {line: symbol.line - 1, position: symbol.position + 1},

    {line: symbol.line, position: symbol.position - 1},
    {line: symbol.line, position: symbol.position + 1},

    {line: symbol.line + 1, position: symbol.position - 1},
    {line: symbol.line + 1, position: symbol.position},
    {line: symbol.line + 1, position: symbol.position + 1},
  ].forEach(linked => {
    const linked_part_index = parts.findIndex(part => {
      return part.line === linked.line && linked.position >= part.start && linked.position <= part.end 
    })
    if (linked_part_index > -1 && !symbol_parts.some(symbol_part => symbol_part.index === linked_part_index)) {
      symbol_parts.push({
        index: linked_part_index,
        value: parts[linked_part_index].value
      })
    }
  })
  if (symbol.character === '*' && symbol_parts.length === 2) {
    values.push(symbol_parts[0].value * symbol_parts[1].value)
  }
})

console.log(values.reduce((acc, value) => acc+value, 0))