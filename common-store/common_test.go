package common_store

import (
	"testing"
	"time"
)

type Thing1 struct {
	value string
}

func addfunc() *ThingToStore {
	var thing ThingToStore = Thing1{"hello"}
	return &thing
}

func removefunc(thing *ThingToStore) {
	return
}

func Test_CommonStore(t *testing.T) {
	store := NewCommonStore(addfunc, removefunc)
	store.AddThingToStore("shoe")
	store.AddThingToStore("eggs")
	store.AddThingToStore("cheese")

	time.Sleep(5 * time.Second)

	if len(store.RetrieveAll()) != 3 {
		t.Fatalf("Test failed not 3 items in map")
	}

	thing := *store.Retrieve("shoe")

	if thing == nil {
		t.Fatalf("Retreive failed")
	}

	storedThing := thing.(Thing1)

	if storedThing.value != "hello" {
		t.Fatalf("Stored value lost")
	}

	store.RemoveThingFromStore("shoe")
	time.Sleep(6 * time.Second)

	deletedthing := store.Retrieve("shoe")
	if deletedthing != nil {
		t.Fatalf("Thing not removed from store")
	}
}
