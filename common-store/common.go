package common_store

type ThingToStore interface {}

type storeMyThing func() ThingToStore

var things map[string]ThingToStore = make(map[string]ThingToStore)

func Store(thingfunc storeMyThing, key string) {
   things[key] = thingfunc()
}

func Retrieve(key string) ThingToStore {
    return things[key]
}
