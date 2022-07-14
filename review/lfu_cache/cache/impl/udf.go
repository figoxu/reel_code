package impl

import (
	"container/heap"
	"sync"
	"time"

	"lfu/cache"
)

type UdfCache struct {
	lock          sync.RWMutex
	cap           int
	data          map[string]*ScoreItem
	priorityItems ScoreItems
}

func NewUdfCache(c int) cache.UdfCache {
	cache := &UdfCache{
		cap:           c,
		data:          map[string]*ScoreItem{},
		priorityItems: ScoreItems{},
	}
	return cache
}

func (p *UdfCache) recycle(keepSize int) {
	for p.priorityItems.Len() > keepSize {
		item := p.priorityItems[0]
		item = heap.Pop(&p.priorityItems).(*ScoreItem)
		delete(p.data, item.key)
	}
}

func (p *UdfCache) Get(k string) string {
	p.lock.RLock()
	defer p.lock.RUnlock()
	v := p.data[k]
	if v == nil {
		return ""
	}
	v.count = v.count + 1
	v.tick = time.Now().UnixMicro()
	heap.Fix(&p.priorityItems, v.index)
	return v.val
}

func (p *UdfCache) Set(k, v string) {
	p.lock.Lock()
	defer p.lock.Unlock()
	old := p.data[k]
	if old == nil {
		len := p.priorityItems.Len()
		if len >= p.cap {
			p.recycle(p.cap - 1)
		}
		item := &ScoreItem{
			key:   k,
			val:   v,
			count: 0,
			tick:  time.Now().UnixMicro(),
			index: 0,
		}
		p.data[k] = item
		heap.Push(&p.priorityItems, item)
	} else {
		old.val = v
		old.count = old.count + 1
		old.tick = time.Now().UnixMicro()
		heap.Fix(&p.priorityItems, old.index)
	}
}

type ScoreItem struct {
	key   string
	val   string
	count int
	tick  int64
	index int
}

type ScoreItems []*ScoreItem

func (p ScoreItems) Len() int { return len(p) }

func (p ScoreItems) Less(i, j int) bool {
	if p[i].count == p[j].count {
		return p[i].tick > p[j].tick
	}
	if p[i].count < p[j].count {
		return true
	}
	return false
}

func (p ScoreItems) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = i
	p[j].index = j
}

func (p *ScoreItems) Push(x interface{}) {
	n := len(*p)
	item := x.(*ScoreItem)
	item.index = n
	*p = append(*p, item)
}

func (pq *ScoreItems) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
