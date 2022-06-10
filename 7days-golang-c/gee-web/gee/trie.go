package gee

import "strings"

type node struct {
	pattern  string  // 叶子节点才有值，即完整路径
	part     string  // 当前部分路径
	children []*node // 所有子节点
	isWild   bool    // 是否模糊匹配，part含有 : 或 * 时为true
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查询
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		// len(parts) == height && n.pattern == "" 的例子
		// 注册了 /1/2/3/4/5，搜索 /1/2/3/4，没有该路由，因为node "4" 中 pattern == ""
		// n.part[0] == '*' && n.pattern == "" 表明带有*的part只能存在一个并且放到最后一段
		// 注册了 /1/2/*，搜索 /1/2/3/4，有该路由，因为node "*" 中 pattern != ""
		// 注册了 /1/2/*/4，搜索 /1/2/3/4，没有该路由，因为node "*" 中 pattern == ""
		// 事实上， 依据addRoute逻辑， /1/2/*/4 是只能注册为 /1/2/* 的

		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil

}
