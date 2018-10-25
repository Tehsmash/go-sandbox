package main

import (
	"fmt"
	"github.com/Tehsmash/go-sandbox/common-store"
	"time"
)

type Thing1 struct {
	egg string
}

func addfunc() common_store.ThingToStore {
	return Thing1{"hello"}
}

func removefunc(thing common_store.ThingToStore) {
	thingcorrect := thing.(Thing1)
	fmt.Println("Removed ", thingcorrect.egg)
}

func main() {
	store := common_store.NewCommonStore(addfunc, removefunc)
	store.AddThingToStore("shoe")
	store.AddThingToStore("eggs")
	store.AddThingToStore("cheese")

	time.Sleep(5 * time.Second)

	fmt.Println(store.Retrieve("shoe"))
	fmt.Println(store.RetrieveAll())

	store.RemoveThingFromStore("shoe")

	time.Sleep(6 * time.Second)

	fmt.Println(store.Retrieve("shoe"))
	fmt.Println(store.RetrieveAll())
}
