import assert from 'assert'
import { read } from '../shared/read'

const input_path = `${__dirname}/test_input`
const lines = read(input_path)

enum RuleType {
  Accepted = 'A',
  Rejected = 'R'
}

interface Rule {
  condition?: {
    property: 'x' | 'm' | 'a' | 's'
    symbol: '>' | '<'
    value: number
  }
  destination: RuleType | string
}

interface Workflow {
  code: string
  rules: Rule[]
}

interface Part {
  x: number
  m: number
  a: number
  s: number
}

const parseWorkflow = (line: string): Workflow => {
  const [code, rest] = line.split('{')
  const instructions = rest.replace('}', '').split(',')
  const rules: Rule[] = instructions.map(instruction => {
    const is_condition = instruction.includes(':')
    if (is_condition) {
      const property = instruction[0]
      assert(property === 'x' || property === 'm' || property === 'a' || property === 's')

      const symbol = instruction[1]
      assert(symbol === '<' || symbol === '>')

      const [string_value, destination] = instruction.slice(2).split(':')
      return {
        condition: {
          property,
          symbol,
          value: Number.parseInt(string_value, 10)
        },
        destination
      }
    }

    return {
      destination: instruction
    }
  })

  return {
    code,
    rules
  }
}

const parseValue = (part_value: string): number => {
  return Number.parseInt(part_value.split('=')[1], 10)
}

const parsePart = (line: string): Part => {
  const [x, m, a, s] = line.replace('{', '').replace('}', '').split(',')
  return {
    x: parseValue(x),
    m: parseValue(m),
    a: parseValue(a),
    s: parseValue(s),
  }
}

const parse = (): { parts: Part[], workflows: Workflow[] } => {
  const parts: Part[] = []
  const workflows: Workflow[] = []
  let is_parsing_workflows = true
  lines.forEach((line) => {
    if (line === '') {
      is_parsing_workflows = false
      return
    }

    if (is_parsing_workflows) {
      const workflow = parseWorkflow(line)
      workflows.push(workflow)
    } else {
      const part = parsePart(line)
      parts.push(part)
    }
  })

  return {
    parts,
    workflows
  }
}

const { parts, workflows } = parse()

const getDestination = (part: Part, workflow: Workflow): string => {
  const { rules } = workflow

  for(let i = 0 ; i < rules.length ; i++) {
    const rule = rules[i]
    if (!rule.condition) {
      return rule.destination
    }

    switch(rule.condition.symbol) {
      case '>':
        if(part[rule.condition.property] > rule.condition.value) {
          return rule.destination
        }
        break
      case '<':
        if(part[rule.condition.property] < rule.condition.value) {
          return rule.destination
        }
        break
    }
  }

  throw new Error('not found')
}

// const isPartAccepted = (part: Part): boolean => {
//   let destination = 'in'

//   while(destination !== 'A' && destination !== 'R') {
//     const current_workflow = workflows.find(workflow => workflow.code === destination)
//     assert(current_workflow)

//     destination = getDestination(part, current_workflow)
//   }

//   return destination === 'A'
// }

// const accepted_parts = parts.filter(isPartAccepted)
// const final_value = accepted_parts.reduce((acc, part) => acc + part.x + part.m + part.a + part.s, 0)
// console.log(final_value)

interface Range {
  start: number
  end: number
}

export interface Ranges {
  x: Range[]
  m: Range[]
  a: Range[]
  s: Range[]
}

const findRanges = (destination: string, ranges: Ranges) => {
  const workflow = workflows.find(({code}) => code === destination)
  assert(workflow)

  workflow.rules.forEach(rule => {
    if (rule.condition) {
      if (rule.destination === RuleType.Rejected) {
        const range = ranges[rule.condition.property]
        if (rule.condition.symbol === '<') {
          const old_range_index = range.findIndex(r => r.start > rule.condition!.value)
          if (old_range_index !== -1) {

          }
        }
      } else if (rule.destination === RuleType.Accepted) {
        return ranges
      } else {
        return findRanges(rule.destination, ranges)
      }
    } else {
      if (rule.destination === RuleType.Rejected) {

      } else if (rule.destination === RuleType.Accepted) {
        return ranges
      } else {
        return findRanges(rule.destination, ranges)
      }
    }
  })
}

const initial_ranges: Ranges = {
  x: [{ start: 1, end: 4000 }],
  m: [{ start: 1, end: 4000 }],
  a: [{ start: 1, end: 4000 }],
  s: [{ start: 1, end: 4000 }]
}

const final_ranges = findRanges('in', initial_ranges)
