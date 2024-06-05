package util

import (
  "fmt"
  "strings"
)

// in go 1.22 this will correctly output 1,2,3 but, in previous versions, it outputs 4,4,4
func ForLoopWithShared() {
    ints := []*int{}

    for i := 1; i <= 3; i++ {
      ints = append(ints, &i)
    }

    output := []string{}

    for _, i := range ints {
      output = append(output, fmt.Sprintf("%d", *i))
    }

    fmt.Println(strings.Join(output, ","))
}
