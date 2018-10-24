package main

import (
  "fmt"
  "github.com/Tehsmash/go-sandbox/common-store"
)

type Thing1 struct {
  egg string
}

type Thing2 struct {
  chicken int
}

func thing1func() common_store.ThingToStore {
  return Thing1{"hello"}
}

func thing2func() common_store.ThingToStore {
  return Thing2{50}
}

func main() {
  common_store.Store(thing1func, "thing1")
  common_store.Store(thing2func, "thing2")

  thingback := common_store.Retrieve("thing1").(Thing1)
  thingback2 := common_store.Retrieve("thing2").(Thing2)

  fmt.Println(thingback)
  fmt.Println(thingback.egg)
  fmt.Println(thingback2.chicken)
}
