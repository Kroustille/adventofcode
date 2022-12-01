import { readFileSync } from "fs"

export const read = (path: string): string => {
  const data = readFileSync(path)
  return data.toString()
}