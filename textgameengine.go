package github.com/unclassedpenguin/textgameengine

import (
  "fmt"
  "time"
  "math/rand"
  "strings"
)

//-----------------------------------------------------------------------------
// Player 
//-----------------------------------------------------------------------------

type Player struct {
  Name      string
  Score     int
  Inventory []string
}

//-----------------------------------------------------------------------------
// End Player 
//-----------------------------------------------------------------------------

type Game struct {
  Events    map[string]bool
  TermWidth int
}

//-----------------------------------------------------------------------------
// Inventory Functions
//-----------------------------------------------------------------------------

// Pretty print Inventory
func PrintInventory(inventory *[]string) {
  if len(*inventory) > 0 {
    fmt.Println("Inventory:")
    fmt.Println("--------------------------")
    for _, v := range *inventory {
      fmt.Println("-", v)
    }
    fmt.Println("--------------------------")
  } else {
    fmt.Println("Your inventory appears to be empty...")
  }
}

// Adds an item to the inventory, or just returns the inventory
func InventoryAdd(inventory *[]string, item string) *[]string {
  *inventory = append(*inventory, item)
  return inventory
}

// Removes an item from the inventory
func InventoryRemove(inventory *[]string, item string) *[]string {
  indexOfItem := InventoryIndexOf(inventory, item)
  if indexOfItem != -1 {
    *inventory = append((*inventory)[:indexOfItem], (*inventory)[indexOfItem+1:]...)
  }
  return inventory
}

// Check if item is contained in inventory 
func InventoryContains(inventory *[]string, item string) bool {
  for _, v := range *inventory {
    if v == item {
      return true
    }
  }
  return false
}

// Get index of item in slice
func InventoryIndexOf(inventory *[]string, item string) (int) {
  for k, v := range *inventory {
    if item == v {
      return k
    }
  }
  return -1    //not found.
}

//-----------------------------------------------------------------------------
// End Inventory Functions
//-----------------------------------------------------------------------------


//-----------------------------------------------------------------------------
// Random Functions
//-----------------------------------------------------------------------------

// Get a single random number
func RandNumber(max int) int {
  rand.Seed(time.Now().UnixNano())
  rn := rand.Intn(max)
  return rn
}

// Word wrap used by PrintSlow
func wordWrap(str string, lineWidth int) (wrapped string) {
    //words := strings.Fields(strings.TrimSpace(text))
    words := strings.Fields(str)
    if len(words) == 0 {
        return str
    }
    wrapped = words[0]
    spaceLeft := lineWidth - len(wrapped)
    for _, word := range words[1:] {
        if len(word)+1 > spaceLeft {
            wrapped += "\n" + word
            spaceLeft = lineWidth - len(word)
        } else {
            wrapped += " " + word
            spaceLeft -= 1 + len(word)
        }
    }
    return
}

// Prints the text character by character.
func PrintSlow(str string, termWidth int, slowMode bool) {
  str = wordWrap(str, termWidth)
  if slowMode {
    stringSplit := strings.Split(str, "")
    for _, l := range stringSplit {
      if l != " " {
        fmt.Print(l)
        time.Sleep(25 * time.Millisecond)
      } else {
        fmt.Print(l)
        time.Sleep(55 * time.Millisecond)
      }
    }
    fmt.Print("\n")
  } else {
    fmt.Println(str)
  }
}

//-----------------------------------------------------------------------------
// End Random Functions
//-----------------------------------------------------------------------------
