# CompactArraysIndent

A Go utility that pretty-prints slices of structs as JSON, while compacting specified array fields like `[]int` and `[]string` into a single line.

## 💡 Why?

By default, `json.MarshalIndent` outputs every element of an array on its own line. This gets noisy fast—especially with data like:

```json
"intArray": [
  0,
  1,
  0,
  0,
  0,
  0
]

This library rework's Go's Marshall to remove extra new lines ("\n") in arrays within entries while preserving structure and indentation of your JSON.

```json
"intArray": [0, 1, 0, 0, 0, 0]

## Usage

```go
import "github.com/etrnl/compactarraysindent"

jsonBytes, err := compactarraysindent.CompactMarshalIndent(data, []string{
  "intArray", "category", "names", "dcause",
}, "", "  ")

## Example

```go
type Person struct {
  Name     string `json:"name"`
  Gender   []int  `json:"gender"`
  Tags     []string `json:"tags"`
}

people := []Person{
  {Name: "A.J. Cook", Gender: []int{0,1,0}, Tags: []string{"actor", "musician"}},
}

out, _ := compactarraysindent.CompactMarshalIndent(people, []string{"gender", "tags"}, "", "  ")
fmt.Println(string(out))