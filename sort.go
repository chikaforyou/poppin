package main
import "fmt"

func main() {
var temp, i, imax int
  o := [6]int {2,4,3,7,6,9}
  
  pass := 0
  for pass < len(o)-1 {
    imax = pass
    i = pass+1
    for i < len(o) {
      if o[i]>o[imax] {
        imax = i
      }
      i++
     }
    temp = o[imax]
    o[imax] = o[pass]
    o[pass] = temp
    pass++
   }
  }
