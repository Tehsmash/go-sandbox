package common_store

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type ThingToStore interface{}

type storeMyThing func() *ThingToStore
type removeMyThing func(thing *ThingToStore)

type SafeQueue struct {
	queue *list.List
	lock  sync.Mutex
}

type SafeThings struct {
	things map[string]*ThingToStore
	m      sync.RWMutex
}

type Store struct {
	things      *SafeThings
	addQueue    *SafeQueue
	removeQueue *SafeQueue
	addFunc     storeMyThing
	removeFunc  removeMyThing
}

func NewCommonStore(addfunc storeMyThing, removefunc removeMyThing) Store {
	stopCh := make(chan struct{})
	store := Store{
		things:      &SafeThings{things: make(map[string]*ThingToStore)},
		addQueue:    &SafeQueue{queue: list.New()},
		removeQueue: &SafeQueue{queue: list.New()},
		addFunc:     addfunc,
		removeFunc:  removefunc,
	}
	go store.Run(stopCh)
	return store
}

func (s *Store) AddThingToStore(name string) {
	s.addQueue.lock.Lock()
	s.addQueue.queue.PushBack(name)
	s.addQueue.lock.Unlock()
}

func (s *Store) RemoveThingFromStore(name string) {
	s.removeQueue.lock.Lock()
	s.removeQueue.queue.PushBack(name)
	s.removeQueue.lock.Unlock()
}

func (s *Store) Run(stopCh <-chan struct{}) {
	for {
		fmt.Println("running")
		select {
		case <-stopCh:
			break
		default:
			s.runWorker()
			time.Sleep(5 * time.Second)
		}
	}
	fmt.Println("Ending")
}

func (s *Store) runWorker() {
	for {
		s.addQueue.lock.Lock()
		addwork := s.addQueue.queue.Front()
		if addwork != nil {
			s.store(s.addQueue.queue.Remove(addwork).(string))
		}
		s.addQueue.lock.Unlock()

		s.removeQueue.lock.Lock()
		removework := s.removeQueue.queue.Front()
		if removework != nil {
			s.remove(s.removeQueue.queue.Remove(removework).(string))
		}
		s.removeQueue.lock.Unlock()
		if addwork == nil && removework == nil {
			return
		}
	}
}

func (s *Store) store(key string) {
	s.things.m.Lock()
	s.things.things[key] = s.addFunc()
	s.things.m.Unlock()
}

func (s *Store) remove(key string) {
	fmt.Println("here")
	s.things.m.Lock()
	s.removeFunc(s.things.things[key])
	delete(s.things.things, key)
	s.things.m.Unlock()
}

func (s *Store) Retrieve(key string) *ThingToStore {
	s.things.m.RLock()
	defer s.things.m.RUnlock()
	return s.things.things[key]
}

func (s *Store) RetrieveAll() map[string]*ThingToStore {
	s.things.m.Lock()
	defer s.things.m.Unlock()
	return s.things.things
}
