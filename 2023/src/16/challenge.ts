import assert from 'assert'
import { read } from '../shared/read'

const input_path = `${__dirname}/input`
const lines = read(input_path)

enum Direction {
  Left = 'left',
  Right = 'right',
  Up = 'up',
  Down = 'down'
}

interface Cell {
  type: '/' | '\\' | '|' | '-'
}

let width = lines[0].length
let height = lines.length

const getCellKey = (x: number, y: number): string => {
  return `${x};${y}`
}

const contraption: Map<string, Cell> = new Map()
lines.forEach((line, line_index) => {
  line.split('').forEach((character, character_index) => {
    if (character === '.') {
      return
    }

    const cell_key = getCellKey(character_index, line_index)
    contraption.set(cell_key, { type: character as Cell['type'] })
  })
})

const getNextX = (x: number, direction: Direction): number => {
  switch(direction) {
    case Direction.Left:
      return x - 1
    case Direction.Right:
      return x + 1
  }

  return x
}

const getNextY = (y: number, direction: Direction): number => {
  switch(direction) {
    case Direction.Up:
      return y - 1
    case Direction.Down:
      return y + 1
  }

  return y
}

const getNextCoords = (x: number, y: number, next_direction: Direction): {next_x: number, next_y: number} => {
  return {
    next_x: getNextX(x, next_direction),
    next_y: getNextY(y, next_direction)
  }
}

type Transformation = Record<Direction, Direction[]>

const slashMirrorTransformation: Transformation = {
  [Direction.Right]: [Direction.Up],
  [Direction.Left]: [Direction.Down],
  [Direction.Up]: [Direction.Right],
  [Direction.Down]: [Direction.Left]
}

const antiSlashMirrorTransformation: Transformation = {
  [Direction.Right]: [Direction.Down],
  [Direction.Left]: [Direction.Up],
  [Direction.Up]: [Direction.Left],
  [Direction.Down]: [Direction.Right]
}

const verticalSplitterTransformation: Transformation = {
  [Direction.Down]: [Direction.Down],
  [Direction.Up]: [Direction.Up],
  [Direction.Right]: [Direction.Up, Direction.Down],
  [Direction.Left]: [Direction.Up, Direction.Down]
}

const horizontalSplitterTransformation: Transformation = {
  [Direction.Right]: [Direction.Right],
  [Direction.Left]: [Direction.Left],
  [Direction.Up]: [Direction.Left, Direction.Right],
  [Direction.Down]: [Direction.Left, Direction.Right]
}

const getTransformation = (cell: Cell): Transformation => {
  switch(cell.type) {
    case '/':
      return slashMirrorTransformation
    case '\\':
      return antiSlashMirrorTransformation
    case '|':
      return verticalSplitterTransformation
    case '-':
      return horizontalSplitterTransformation
    default:
      throw new Error('unknown cell type')
  }
}

const energized_cells: Set<string> = new Set()

const printCells = () => {
  for(let y = 0 ; y < height ; y++) {
    let line = ''
    for(let x = 0 ; x < width ; x++) {
      const key = getCellKey(x, y)
      if (energized_cells.has(key)) {
        line += '#'
      } else {
        line += '.'
      }
    }
    console.log(line)
  }

  console.log()
}

const sentBeams: Set<string> = new Set()

const sendBeam = (x: number, y: number, direction: Direction): void => {
  if (x < 0 || x >= width) {
    return
  }
  if (y < 0 || y >= height) {
    return
  }

  const cell_key = getCellKey(x, y)
  if (sentBeams.has(`${cell_key};${direction}`)) {
    return
  }

  sentBeams.add(`${cell_key};${direction}`)

  energized_cells.add(cell_key)
  if (!contraption.has(cell_key)) {
    const next_x = getNextX(x, direction)
    const next_y = getNextY(y, direction)

    return sendBeam(next_x, next_y, direction)
  }

  const cell = contraption.get(cell_key)
  assert(cell)

  const transformation = getTransformation(cell)
  const next_directions = transformation[direction]
  next_directions.forEach(next_direction => {
    const { next_x, next_y } = getNextCoords(x, y, next_direction)
    sendBeam(next_x, next_y, next_direction)
  })
}

sendBeam(0, 0, Direction.Right)
printCells()
console.log(energized_cells.size)