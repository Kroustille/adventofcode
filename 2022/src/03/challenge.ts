import { read } from '../shared/read'

const input_path = `${__dirname}/input`
const rucksacks = read(input_path)

const isUpperCase = (s: string): boolean => {
  return s === s.toUpperCase()
}

const getPriority = (s: string): number => {
  const char_code = s.charCodeAt(0)
  const char_code_difference = isUpperCase(s) ? 38 : 96
  const priority = char_code - char_code_difference
  return priority
}

const total_priority_part1 = rucksacks.reduce((acc, rucksack) => {
  const [first_compartment, second_compartment] = [rucksack.substring(0, rucksack.length / 2), rucksack.substring(rucksack.length / 2)]
  const duplicated_item = first_compartment.split('').find(first_compartment_item => {
    return second_compartment.includes(first_compartment_item)
  })

  if (!duplicated_item) {
    return acc
  }

  const char_code = duplicated_item.charCodeAt(0)
  const char_code_difference = isUpperCase(duplicated_item) ? 38 : 96
  const priority = char_code - char_code_difference

  return acc + priority
}, 0)

// Part 1
console.log('Part 1 =', total_priority_part1)

const grouped_rucksacks = rucksacks.reduce((acc, rucksack) => {
  if (!rucksack) {
    return acc
  }

  if (!acc[acc.length - 1] || acc[acc.length - 1].length == 3) {
    acc.push([])
  }

  acc[acc.length - 1].push(rucksack)

  return acc
}, new Array<Array<string>>())

const badges = grouped_rucksacks.map((grouped_rucksack) => {
  return grouped_rucksack[0].split('').find(first_item => {
    return grouped_rucksack[1].split('').includes(first_item) && grouped_rucksack[2].split('').includes(first_item)
  })
})

const total_priority_part2 = badges.filter(Boolean).reduce((sum, badge) => sum + getPriority(badge!), 0)

// Part 2
console.log('Part 2 =', total_priority_part2)