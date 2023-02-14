package time

import (
  "time"
  "fmt"
)

// This function shows the current time in an infinite loop
func Nunc(){
  for{
  fmt.Println(time.Now())
  } 
}
