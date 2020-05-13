/*
starting point:

package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {

}
*/

package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())

  isHeistOn := true
  eludedGuards := rand.Intn(100)

  if eludedGuards >= 50 {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  } else {
    fmt.Println("Plan a better disguise next time?")
    isHeistOn = false
  }

  openedVault := rand.Intn(100)

  if openedVault>= 70 {
    fmt.Println("Grab and GO!")
  } else {
    fmt.Println("The vault can’t be opened.")
    isHeistOn = false
  }

  leftSafely := rand.Intn(5)

  if isHeistOn == false {
    switch leftSafely {
    case 0:
      isHeistOn = false
      fmt.Println("You faild, now you must die!")
    case 1: 
      isHeistOn = false 
      fmt.Println("Turns out vault doors don't open from the inside...") 
    case 2: 
      isHeistOn = false 
      fmt.Println("book him dano")
    default:
      fmt.Println("Start the getaway car!")
    }

  }
  
  if isHeistOn {
    amtStolen := 10000 + rand.Intn(1000000)
    fmt.Println(amtStolen)
  }
}

