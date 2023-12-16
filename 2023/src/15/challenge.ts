import { read } from '../shared/read'

const input_path = `${__dirname}/input`
const lines = read(input_path)
const line = lines[0]

interface Lens {
  key: string
  hash: number
  instruction: 'remove' | 'replace'
  focal_length?: number
}

interface Box {
  lenses: Lens[] 
}

const processHASH = (entry: string): number => {
  return entry.split('').reduce((acc, character) => {
    const increased = acc + character.charCodeAt(0)
    const multiplied = increased * 17
    return multiplied % 256
  }, 0)
}

// const part_1_total = line.split(',').reduce((acc, entry) => {
//   return acc + processHASH( entry)
// }, 0)

// console.log(part_1_total)

const lenses: Lens[] = line.split(',').map(entry => {
  const [key, focal_length] = entry.split('=')
  if (!focal_length) {
    const formatted_key = key.slice(0, key.length - 1)
    return {
      key: formatted_key,
      instruction: 'remove',
      hash: processHASH(formatted_key),
    }
  }

  return {
    key,
    instruction: 'replace',
    hash: processHASH(key),
    focal_length: Number.parseInt(focal_length),
  }
})

const boxes: Box[] = []
for (let i = 0 ; i < 256 ; i++) {
  boxes.push({ lenses: [] })
}

lenses.forEach(lens => {
  const box = boxes[lens.hash]
  const lens_index = box.lenses.findIndex(box_lens => box_lens.key === lens.key)
  switch(lens.instruction) {
    case 'remove':
      if (lens_index !== -1) {
        box.lenses.splice(lens_index, 1)
      }
      break
    case 'replace':
      if (lens_index === -1) {
        box.lenses.push(lens)
      } else {
        box.lenses[lens_index] = lens
      }
      break
  }
})

const print = () => {
  boxes.forEach((box, index) => {
    if (!box.lenses.length) {
      return
    }
    console.log(index, box)
  })
}

const processBoxValue = (box: Box, box_index: number): number => {
  return box.lenses.reduce((acc, lens, index) => {
    return acc + (box_index + 1) * (index + 1) * (lens.focal_length ?? 1)
  }, 0)
}

const total_box_value = boxes.reduce((acc, box, index) => acc + processBoxValue(box, index), 0)
console.log(total_box_value)