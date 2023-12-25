package keywords

import (
	"sync"
	"time"
)

type Keywords struct {
	rwMutex sync.RWMutex

	keywords []string
	trie     Trie
}

func NewKeywords(keywords []string, opts ...TrieOption) *Keywords {
	trie := NewTrie(keywords, opts...)

	return &Keywords{
		keywords: keywords,
		trie:     trie,
	}
}

func (inst *Keywords) Filter(text string) (afterFilter string, isChange bool) {
	inst.rwMutex.RLock()
	afterFilter, _, isChange = inst.trie.Filter(text)
	inst.rwMutex.RUnlock()

	return
}

func (inst *Keywords) Find(text string) (keywords []string) {
	inst.rwMutex.RLock()
	keywords = inst.trie.FindKeywords(text)
	inst.rwMutex.RUnlock()

	return
}

func (inst *Keywords) Contain(text string) (is bool) {
	inst.rwMutex.RLock()
	keywords := inst.trie.FindKeywords(text)
	inst.rwMutex.RUnlock()

	return len(keywords) > 0
}

func (inst *Keywords) RefreshKeywords(keywords []string) {
	inst.rwMutex.Lock()
	inst.trie = NewTrie(keywords)
	inst.rwMutex.Unlock()

	return
}

type RefreshFn func() (keywords []string, err error)

func (inst *Keywords) AutoRefreshKeywords(fn RefreshFn, duration time.Duration) {
	go func() {
		for {
			keywords, err := fn()
			if err != nil {
				time.Sleep(duration)
				continue
			}

			inst.RefreshKeywords(keywords)
			time.Sleep(duration)
		}
	}()
}
