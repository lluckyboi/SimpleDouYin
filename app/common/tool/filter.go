package tool

import "fmt"

// TrieNode 敏感词前缀树节点
type TrieNode struct {
	childMap map[rune]*TrieNode // 本节点下的所有子节点
	Data     string             // 在最后一个节点保存完整的一个内容
	End      bool               // 标识是否最后一个节点
}

// SensitiveTrie 敏感词前缀树
type SensitiveTrie struct {
	replaceChar rune // 敏感词替换的字符
	root        *TrieNode
}

// NewSensitiveTrie 构造敏感词前缀树实例
func NewSensitiveTrie(replaceChar rune) *SensitiveTrie {
	return &SensitiveTrie{
		replaceChar: replaceChar,
		root:        &TrieNode{End: false},
	}
}

// AddChild 前缀树添加
func (tn *TrieNode) AddChild(c rune) *TrieNode {

	if tn.childMap == nil {
		tn.childMap = make(map[rune]*TrieNode)
	}

	if trieNode, ok := tn.childMap[c]; ok {
		// 存在不添加了
		return trieNode
	} else {
		// 不存在
		tn.childMap[c] = &TrieNode{
			childMap: nil,
			End:      false,
		}
		return tn.childMap[c]
	}
}

// AddWord 添加敏感词
func (st *SensitiveTrie) AddWord(sensitiveWord string) {
	// 将敏感词转换成rune类型(int32)
	tireNode := st.root
	sensitiveChars := []rune(sensitiveWord)
	for _, charInt := range sensitiveChars {
		// 添加敏感词到前缀树中
		tireNode = tireNode.AddChild(charInt)
	}
	tireNode.End = true
	tireNode.Data = sensitiveWord
}

// FindChild 前缀树寻找字节点
func (tn *TrieNode) FindChild(c rune) *TrieNode {
	if tn.childMap == nil {
		return nil
	}

	if trieNode, ok := tn.childMap[c]; ok {
		return trieNode
	}
	return nil
}

// replaceRune 字符替换
func (st *SensitiveTrie) replaceRune(chars []rune, begin int, end int) {
	for i := begin; i < end; i++ {
		chars[i] = st.replaceChar
	}
}

// Match 查找替换发现的敏感词
func (st *SensitiveTrie) Match(text string) (sensitiveWords []string, replaceText string) {
	if st.root == nil {
		return nil, text
	}

	textChars := []rune(text)
	textCharsCopy := make([]rune, len(textChars))
	sensitiveMap := make(map[string]*struct{}) // 利用map把相同的敏感词去重
	copy(textCharsCopy, textChars)
	for i, textLen := 0, len(textChars); i < textLen; i++ {
		trieNode := st.root.FindChild(textChars[i])
		if trieNode == nil {
			continue
		}

		// 匹配到了敏感词的前缀，从后一个位置继续
		j := i + 1
		for ; j < textLen && trieNode != nil; j++ {
			if trieNode.End {
				// 完整匹配到了敏感词，将匹配的文本的敏感词替换成 *
				st.replaceRune(textCharsCopy, i, j)
				if _, ok := sensitiveMap[trieNode.Data]; !ok {
					sensitiveWords = append(sensitiveWords, string(textChars[i:j]))
				}
				sensitiveMap[trieNode.Data] = nil
				st.replaceRune(textCharsCopy, i, j)
			}
			trieNode = trieNode.FindChild(textChars[j])
		}

		// 文本尾部命中敏感词情况
		if j == textLen && trieNode != nil && trieNode.End {
			if _, ok := sensitiveMap[trieNode.Data]; !ok {
				sensitiveWords = append(sensitiveWords, trieNode.Data)
			}
			sensitiveMap[trieNode.Data] = nil
			st.replaceRune(textCharsCopy, i, textLen)
		}
	}

	if len(sensitiveWords) > 0 {
		// 有敏感词
		replaceText = string(textCharsCopy)
		fmt.Println("有敏感词，已进行替换")
	} else {
		// 没有则返回原来的文本
		replaceText = text
		fmt.Println("无敏感词")
	}

	return sensitiveWords, replaceText
}

// Init 初始化
func (st *SensitiveTrie) Init(words []string) *SensitiveTrie {
	for _, word := range words {
		fmt.Println(word, []rune(word))
		st.AddWord(word)
	}
	return st
}
