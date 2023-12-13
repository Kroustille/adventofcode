import assert from 'assert'
import { read } from '../shared/read'

const input_path = `${__dirname}/input`
const lines = read(input_path)

enum TerrainType {
  Ash = '.',
  Rock = '#'
}

interface Position {
  x: number
  y: number
}

interface Pattern {
  width: number
  height: number
  rocks: Position[]
}

interface Reflection {
  type: 'vertical' | 'horizontal'
  index: number
}

const patterns: Pattern[] = []
let positions: Position[] = []
let width = 0
let height = 0

lines.forEach(line => {
  if (line === '') {
    patterns.push({
      rocks: [...positions],
      width,
      height
    })
    positions = []
    width = 0
    height = 0
    return
  }

  width = line.length

  line.split('').forEach((character, character_index) => {
    if (character === TerrainType.Ash) {
      return
    }

    positions.push({
      x: character_index,
      y: height
    })
  })

  height++
})

const isValidVerticalReflection = (pattern: Pattern, reflection_index: number): boolean => {
  let interval = 1
  while (reflection_index - interval >= 0 && reflection_index + interval - 1 <= pattern.width - 1) {
    const left_col_rocks = pattern.rocks.filter(rock => rock.x === reflection_index - interval)
    const right_col_rocks = pattern.rocks.filter(rock => rock.x === reflection_index + interval - 1)
    const has_all_rocks_in_common = left_col_rocks.length === right_col_rocks.length && left_col_rocks.every(left_rock => right_col_rocks.some(right_rock => right_rock.y === left_rock.y))
    if (!has_all_rocks_in_common) {
      return false
    }

    interval++
  }

  return true
}

const findVerticalReflection = (pattern: Pattern, ignore_reflection?: Reflection): number => {
  for (let i = 1 ; i <= pattern.width - 1 ; i++) {
    if (ignore_reflection?.type === 'vertical' && ignore_reflection.index === i) {
      continue
    }
    if (isValidVerticalReflection(pattern, i)) {
      return i
    }
  }

  return -1
}

const isValidHorizontalReflection = (pattern: Pattern, reflection_index: number): boolean => {
  let interval = 1
  while (reflection_index - interval >= 0 && reflection_index + interval - 1 < pattern.height) {
    const up_row_rocks = pattern.rocks.filter(rock => rock.y === reflection_index - interval)
    const bottom_row_rocks = pattern.rocks.filter(rock => rock.y === reflection_index + interval - 1)
    const has_all_rocks_in_common = up_row_rocks.length === bottom_row_rocks.length && up_row_rocks.every(up_rock => bottom_row_rocks.some(bottom_rock => up_rock.x === bottom_rock.x))
    if (!has_all_rocks_in_common) {
      return false
    }

    interval++
  }

  return true
}

const findHorizontalReflection = (pattern: Pattern, ignore_reflection?: Reflection): number => {
  for (let i = 1 ; i <= pattern.height - 1 ; i++) {
    if (ignore_reflection?.type === 'horizontal' && ignore_reflection.index === i) {
      continue
    }
    if (isValidHorizontalReflection(pattern, i)) {
      return i
    }
  }

  return -1
}

const getReflection = (pattern: Pattern, ignore_reflection?: Reflection): Reflection | null => {
  const vertical_reflection_index = findVerticalReflection(pattern, ignore_reflection)
  if (vertical_reflection_index !== -1) {
    return {
      type: 'vertical',
      index: vertical_reflection_index
    }
  }

  const horizontal_reflection_index = findHorizontalReflection(pattern, ignore_reflection)
  if (horizontal_reflection_index !== -1) { 
    return {
      type: 'horizontal',
      index: horizontal_reflection_index
    }
  }

  return null
}

const getValue = (reflection: Reflection): number => {
  if (reflection.type === 'vertical') {
    return reflection.index
  }

  return reflection.index * 100
}

// const total = patterns.reduce((acc, pattern) => {
//   const reflection = getReflection(pattern)
//   return acc + getValue(reflection)
// }, 0)

// console.log(total)

const findOtherReflectionWithRocks = (pattern: Pattern, initial_reflection: Reflection): Reflection | null => {
  for(let x = 0 ; x < pattern.width ; x++) {
    for(let y = 0 ; y < pattern.height ; y++) {
      const has_rock = pattern.rocks.some(rock => rock.x === x && rock.y === y)
      if (has_rock) {
        continue
      }

      const reflection = getReflection({
        ...pattern,
        rocks: [...pattern.rocks, { x, y }]
      }, initial_reflection)
      
      if (reflection && (reflection.type !== initial_reflection.type || reflection.index !== initial_reflection.index)) {
        return reflection
      }
    }
  }

  return null
}

const findOtherReflectionWithAshes = (pattern: Pattern, initial_reflection: Reflection): Reflection | null => {
  for(let i = 0 ; i <= pattern.rocks.length ; i++) {
    const current_rock = pattern.rocks[i]

    const new_rocks = i === 0 ? pattern.rocks.slice(1) :  [...pattern.rocks.slice(i-1, i), ...pattern.rocks.slice(i+1)]
    const reflection = getReflection({
      ...pattern,
      rocks: new_rocks
    }, initial_reflection)
    
    if (reflection && (reflection.type !== initial_reflection.type || reflection.index !== initial_reflection.index)) {
      return reflection
    }
  }

  return null
}

const total = patterns.reduce((acc, pattern) => {
  const initial_reflection = getReflection(pattern)
  assert(initial_reflection)

  const other_reflection_with_rocks = findOtherReflectionWithRocks(pattern, initial_reflection)
  if (other_reflection_with_rocks) {
    return acc + getValue(other_reflection_with_rocks)
  }
  
  const other_reflection_with_ashes = findOtherReflectionWithAshes(pattern, initial_reflection)
  assert(other_reflection_with_ashes)
  
  return acc + getValue(other_reflection_with_ashes)
}, 0)

console.log(total)