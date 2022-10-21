package textgameengine

import (
  "fmt"
  "time"
  "math/rand"
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
// Events
//-----------------------------------------------------------------------------

// Add an event to events map
//func EventAdd(events *[]map[string]bool, event string, b bool) *[]map[string]bool {
    //(*events)[event] = b
    //return events
//}


//-----------------------------------------------------------------------------
// End Events
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

//-----------------------------------------------------------------------------
// End Random Functions
//-----------------------------------------------------------------------------
