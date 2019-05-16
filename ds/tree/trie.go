package tree

import (
	"fmt"
	"github.com/xanzy/go-gitlab"
)

//example

//根节点不包含字符，除根节点外每一个节点都只包含一个字符；
//从根节点到某一节点，路径上经过的字符连接起来，为该节点对应的字符串；
//每个节点的所有子节点包含的字符都不相同。

// dict tree

// define dict 26 a~z
const count = 26

// node
// Go rune  ==> C char
type node struct {
	children *[count]*node
	mark int
	char rune
}

func createNode(char rune, mark int) *node {
	n := &node{
		char:char,
		mark:mark,
		children:nil,
	}
	for i:=0;i<int(count); i++ {
		n.children[i] = (*node)(nil)
	}
	return n
}

func appendNode(n *node, char rune) bool {
	child := n.children[char - rune('a')]
	if child != nil {
		return false
	} else {
		n.children[char-rune('a')] =createNode(char,0)
		return true
	}
}

func addWord(root *node, char []rune) bool{
	flag := true
	for _, c := range char {
		if appendNode(root,c) {
			flag = false
		}
		root = root.children[c - rune('a')]
	}
	if root.mark == 0 {
		flag = false
		root.mark =1
	}
	return !flag
}

func travaersal(root *node, char []rune) {
	if root == nil {
		return
	}
	cc := make([]rune,0,len(char))
	copy(cc,char)

	if root.mark == 1 {
		fmt.Println(cc)
	}
	for i := 0; i < count; i ++ {
		travaersal(root.children[i],cc)
	}
	//cc = cc[:0]
}

func check(root *node, char []rune) bool{

}