import { read } from '../shared/read'

const input_path = `${__dirname}/input`
const lines = read(input_path)

enum Direction {
  Right = 'R',
  Left = 'L',
  Down = 'D',
  Up = 'U'
}

interface DigInstruction {
  direction: Direction
  distance: number
}

interface Hole {
  x: number
  y: number
}

const parsePart1Instructions = (): DigInstruction[] => {
  return lines.map(line => {
    const splitted = line.split(' ')
    const direction = splitted[0] as Direction
    const distance = Number.parseInt(splitted[1], 10)
    return {
      direction,
      distance
    }
  })
}

const directionMapper: Record<string, Direction> = {
  '0': Direction.Right,
  '1': Direction.Down,
  '2': Direction.Left,
  '3': Direction.Up 
}

const parsePart2Instructions = (): DigInstruction[] => {
  return lines.map(line => {
    const instruction_string = line.split(' ')[2].replace('(', '').replace(')', '')
    const direction_character = instruction_string[instruction_string.length - 1]
    const direction = directionMapper[direction_character]
    const distance = Number.parseInt(instruction_string.slice(1, instruction_string.length - 1), 16)

    return {
      direction,
      distance
    }
  })
}

const buildHoles = (instructions: DigInstruction[]): { holes: Hole[], perimeter: number } => {
  const holes: Hole[] = []
  let x = 0
  let y = 0
  let perimeter = 0
  holes.push({ x, y })
  instructions.forEach(instruction => {
    switch(instruction.direction) {
      case Direction.Left:
        x -= instruction.distance
        break
      case Direction.Right:
        x += instruction.distance
        break
      case Direction.Down:
        y += instruction.distance
        break
      case Direction.Up:
        y -= instruction.distance
        break
    }

    perimeter += instruction.distance

    holes.push({ x: x, y: y })
  })
  
  return { holes, perimeter }
}

const getPolygonArea = (holes: Hole[]) => { 
  const all_x: number[] = []
  const all_y: number[] = []
  holes.forEach((hole) => {
    all_x.push(hole.x)
    all_y.push(hole.y)
  })

  const n = all_x.length
  let area = 0;

  let j = n - 1;
  for (let i = 0; i < n; i++) {
    area += (all_x[j] + all_x[i]) * (all_y[j] - all_y[i]);
    j = i;
  }

  return Math.abs(area / 2.0)
}

const getResult = (instructions: DigInstruction[]): number => {
  const { holes, perimeter } = buildHoles(instructions)
  return getPolygonArea(holes) + (perimeter + 4) / 2 - 1
}

const solvePart1 = () => {
  const instructions: DigInstruction[] = parsePart1Instructions()
  console.log(getResult(instructions))
}

const solvePart2 = () => {
  const instructions: DigInstruction[] = parsePart2Instructions()
  console.log(getResult(instructions))
}

solvePart2()