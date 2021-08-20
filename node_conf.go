package contree

import (
	"strings"
)

type NodeConf struct {
	Name   string
	Value  string
	Level  int
	childs map[string]*NodeConf
	isLeaf bool
}

func (n *NodeConf) IsNamed(str string) bool {
	return str == n.Name
}

func (n *NodeConf) GetValue() string {
	return n.Value
}

func (n *NodeConf) Insert(node *NodeConf) {
	if n.childs == nil {
		n.childs = make(map[string]*NodeConf)
	}
	node.Level = n.Level + 1
	node.isLeaf = true
	n.childs[node.Name] = node
}

func (n *NodeConf) Browse(path string) string {
	part := strings.Split(path, ".")
	if node, ok := n.childs[part[0]]; ok {
		if len(part) > 1 {
			return node.Browse(strings.Join(part[1:], "."))
		}
		if len(part) == 1 {
			return node.GetValue()
		}
	}
	return ""
}

func (n *NodeConf) SetRecursivly(path string, value string) {
	if len(path) <= 0 {
		n.Value = value
		return
	}
	part := strings.Split(path, ".")
	if len(part) >= 1 {
		n.isLeaf = false
		if node, ok := n.childs[part[0]]; ok {
			node.SetRecursivly(strings.Join(part[1:], "."), value)
		} else {
			node := &NodeConf{
				Name: part[0],
			}
			n.Insert(node)
			if n.Level > 10 {
				return
			}
			node.SetRecursivly(strings.Join(part[1:], "."), value)
		}
	}
}
