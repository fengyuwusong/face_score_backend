package cache

import (
	"fmt"
	log "github.com/howiechou/golog"
	"sync"
	"time"
)

//*****  CachePoolBase  *****
type CachePoolBase struct {
	sync.RWMutex
	iterms map[string]ICache
	stop   chan bool
	once   sync.Once
}

func (this *CachePoolBase) set(k string, v ICache) {
	this.Lock()
	this.iterms[k] = v
	this.Unlock()
}

func (this *CachePoolBase) get(k string) (ICache, bool) {
	this.RLock()
	v, b := this.iterms[k]
	this.RUnlock()

	if !b {
		return nil, false
	}

	if v.CanEliminate() {
		v.Close(true)
		this.Lock()
		delete(this.iterms, k)
		this.Unlock()
		return nil, false
	}

	return v, true
}

func (this *CachePoolBase) run() {
	this.stop = make(chan bool)
	ticker := time.NewTicker(60 * time.Second)
	for {
		select {
		case <-ticker.C:
			this.DeleteExpired()
		case <-this.stop:
			ticker.Stop()
			return
		}
	}
}

func (this *CachePoolBase) Get(k string) (ICache, bool) {
	return this.get(k)
}

func (this *CachePoolBase) Replace(k string, v ICache) error {
	//check
	_, b := this.get(k)
	if !b {
		err := fmt.Errorf("Item:%s dosen't exit", k)
		return err
	}

	this.set(k, v)
	return nil
}

func (this *CachePoolBase) Uninit() {
	this.Lock()
	defer this.Unlock()

	this.stop <- true
	this.iterms = nil
}

func (this *CachePoolBase) init() {
	this.iterms = make(map[string]ICache)
	go this.run()
	log.Debugf("cachepoolbase: init run")
}

func (this *CachePoolBase) Add(k string, v ICache) error {
	this.once.Do(this.init)
	//check
	_, b := this.get(k)
	if b {
		//	err := errors.New(fmt.Sprintf("Item:%s has already exit", k))
		log.Debugf("Item:%s has already exit", k)
		this.set(k, v) //存在就覆盖
		return nil
	}
	log.Debugf("cachepoolbase add %v", k)
	//add
	this.set(k, v)
	return nil
}

func (this *CachePoolBase) Count() int {
	this.Lock()
	defer this.Unlock()

	return len(this.iterms)
}

func (this *CachePoolBase) Flush() {
	this.Lock()
	defer this.Unlock()

	this.iterms = map[string]ICache{}
}

func (this *CachePoolBase) Delete(k string) {
	this.Lock()
	defer this.Unlock()

	delete(this.iterms, k)
}

func (this *CachePoolBase) DeleteExpired() {
	arr := this.eliminateKeys()

	if len(arr) <= 0 {
		return
	}

	this.Lock()
	defer this.Unlock()
	for _, k := range arr {
		//log.Debugf("delete key: %v", k)
		delete(this.iterms, k)
	}
}

func (this *CachePoolBase) eliminateKeys() []string {
	this.RLock()
	defer this.RUnlock()

	var keys []string
	for k, v := range this.iterms {
		if v.CanEliminate() {
			v.Close(true)
			keys = append(keys, k)
		}
	}

	return keys
}
