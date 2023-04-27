package time

import (
  "time"
  "fmt"
)

// This function shows the current time in an infinite loop
func Nunc(){
  for{
    now := time.Now()
    fmt.Println(now.Format("Monday, Jan 02, 2006 3:04:05 PM"))
    time.Sleep(1 * time.Second)
  } 
}
