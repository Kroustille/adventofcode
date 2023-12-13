import { read } from '../shared/read'

const input_path = `${__dirname}/input`
const lines = read(input_path)

enum Instruction {
  LEFT = 'L',
  RIGHT = 'R'
}

interface Node {
  code: string
  left_node: string
  right_node: string
}

const instructions: Instruction[] = lines[0].split('').map(character => {
  if (character === 'L') {
    return Instruction.LEFT
  } else {
    return Instruction.RIGHT
  }
})

const nodes: Node[] = lines.slice(2).map(line => {
  const [code, directions] = line.split(' = ')
  const [left, right] = directions.split(', ')
  return {
    code,
    left_node: left.replace('(', '').replace(')', ''),
    right_node: right.replace('(', '').replace(')', '')
  }
})

const nodes_indexed_by_code = nodes.reduce((acc, node) => {
  const { code, left_node, right_node } = node
  return {
    ...acc,
    [code]: {
      left_node,
      right_node
    }
  }
}, {} as Record<string, { left_node: string, right_node: string }>)

const isStartingNode = (node: Node): boolean => {
  return node.code[node.code.length - 1] === 'A'
}

const isEndingNode = (node_code: string): boolean => {
  return node_code[node_code.length - 1] === 'Z'
}

const areEndingNodes = (node_codes: string[]): boolean => {
  return !node_codes.some(code => !isEndingNode(code))
}

const starting_nodes = nodes.filter(node => isStartingNode(node))
console.log(starting_nodes)
const steps_per_node = starting_nodes.map(starting_node => {
  let required_steps = 0
  let current_node_code = starting_node.code
  while (!isEndingNode(current_node_code)) {
    for (let i = 0; i < instructions.length; i++) {
      if (isEndingNode(current_node_code)) {
        break
      }

      const instruction = instructions[i]
      const node = nodes_indexed_by_code[current_node_code]
      if (instruction === Instruction.LEFT) {
        current_node_code = node.left_node
      } else {
        current_node_code = node.right_node
      }

      required_steps++
    }
  }

  return required_steps
})


console.log(steps_per_node)

// let required_steps = 0
// let current_node_codes = starting_nodes.map(node => node.code)

// while (!areEndingNodes(current_node_codes)) {
//   for (let i = 0; i < instructions.length; i++) {
//     if (areEndingNodes(current_node_codes)) {
//       break
//     }

//     const instruction = instructions[i]
//     current_node_codes = current_node_codes.map(code => {
//       const node = nodes_indexed_by_code[code]
//       if (instruction === Instruction.LEFT) {
//         return node.left_node
//       } else {
//         return node.right_node
//       }
//     })

//     required_steps++
//   }
// }

// console.log(required_steps)


// const cache: Record<string, number> = {}
// const findStepsNumber = (code: string, instruction_index: number, all_steps: number): number => {
//   if (isEndingNode(code)) {
//     return all_steps
//   }

//   if (cache[code]) {
//     return cache[code]
//   }

//   const instruction = instructions[instruction_index]
//   const node = nodes_indexed_by_code[code]
//   const next_instruction_index = (instruction_index + 1) % instructions.length
//   let steps = 0
//   if (instruction === Instruction.LEFT) {
//     steps = findStepsNumber(node.left_node, next_instruction_index, all_steps + 1)
//   } else {
//     steps = findStepsNumber(node.right_node, next_instruction_index, all_steps + 1)
//   }


//   cache[code] = steps

//   return steps
// }

// console.log(findStepsNumber('11A', 0, 0))
// console.log(cache)




// let current_node = 'AAA'
// let required_steps = 0
// while (current_node !== 'ZZZ') {
//   for (let i = 0; i < instructions.length; i++) {
//     if (current_node === 'ZZZ') {
//       continue
//     }

//     const instruction = instructions[i]
//     const node = nodes_indexed_by_code[current_node]
//     if (instruction === Instruction.LEFT) {
//       current_node = node.left_node
//     } else {
//       current_node = node.right_node
//     }
//     required_steps++
//   }
// }

// console.log(required_steps)