package contree

import "strings"

type NodeConf struct {
	Name   string
	Path   string
	Level  int
	Value  string
	childs map[string]*NodeConf
}

func (n *NodeConf) Get(path string) (string, error) {
	pathSlice := strings.Split(path, ".")
	if len(pathSlice) > 1 {
		return n.Get(strings.Join(pathSlice[1:], "."))
	}
	if len(pathSlice) == 1 {
		if n.Name == pathSlice[0] {
			return n.Value, nil
		}
	}
	return "", &ConfigurationValueNotFoundError{}
}

func (n *NodeConf) Set(path string, value string) {
	pathSlice := strings.Split(path, ".")
	if len(pathSlice) > 1 {
		n.Set(strings.Join(pathSlice[1:], "."), value)
	}
	if len(pathSlice) == 1 {
		n.Value = value
	}
}

func (n *NodeConf) Add(node *NodeConf) {
	node.Path = n.Path + "." + node.Name
	node.Level = n.Level + 1
	n.childs[node.Name] = node
}
