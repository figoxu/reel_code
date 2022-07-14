package tree

//https://leetcode.cn/problems/implement-trie-prefix-tree/solution/shi-xian-trie-qian-zhui-shu-by-leetcode-ti500/
//https://leetcode.cn/problems/implement-trie-prefix-tree/solution/by-lfool-k6hb/

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func NewTrie() Trie {
	return Trie{}
}

func (p *Trie) Insert(word string) {
	node := p
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (p *Trie) searchPrefix(prefix string) *Trie {
	node := p
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (p *Trie) Search(word string) bool {
	node := p.searchPrefix(word)
	return node != nil && node.isEnd
}

func (p *Trie) StartsWith(prefix string) bool {
	return p.searchPrefix(prefix) != nil
}
