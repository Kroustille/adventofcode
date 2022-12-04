import { read } from '../shared/read'

const input_path = `${__dirname}/input`
const pairs = read(input_path)

interface Section {
  start: number
  end: number
}

interface Pair {
  left: Section
  right: Section
}

const parseSection = (s: string): Section => {
  const [start, end] = s.split('-')
  return {
    start: Number.parseInt(start, 10),
    end: Number.parseInt(end, 10),
  }
}

const parsePair = (s: string): Pair => {
  const [left, right] = s.split(',')
  return {
    left: parseSection(left),
    right: parseSection(right)
  }
}

const doesLeftCompletelyOverlapRight = ({ left, right }: Pair): boolean => {
  return left.start <= right.start && left.end >= right.end
}

const doesRightCompletelyOverlapLeft = ({ left, right }: Pair): boolean => {
  return right.start <= left.start && right.end >= left.end
}

const doesCompletelyOverlap = (p: Pair): boolean => {
  return doesLeftCompletelyOverlapRight(p) || doesRightCompletelyOverlapLeft(p)
}

const doesOverlap = (p: Pair): boolean => {
  return p.left.start <= p.right.start && p.left.end >= p.right.start ||
    p.left.start <= p.right.end && p.left.end >= p.right.end ||
    p.right.start <= p.left.start && p.right.end >= p.left.start ||
    p.right.start <= p.left.end && p.right.end >= p.left.end
}

const completely_overlapping_sections = pairs.reduce((acc, pair_string) => {
  if (!pair_string) {
    return acc
  }

  const pair = parsePair(pair_string)
  if (doesCompletelyOverlap(pair)) {
    return acc + 1
  }

  return acc
}, 0)

const overlapping_sections = pairs.reduce((acc, pair_string) => {
  if (!pair_string) {
    return acc
  }

  const pair = parsePair(pair_string)
  if (doesOverlap(pair)) {
    return acc + 1
  }

  return acc
}, 0)

// Part 1
console.log('Part 1 =', completely_overlapping_sections)

// Part 2
console.log('Part 2 =', overlapping_sections)
