package contree

import (
	"regexp"
	"strings"
)

type ConfigurationValueNotFoundError struct{}

func (c *ConfigurationValueNotFoundError) Error() string {
	return "Configuration path not found"
}

type Contree struct {
	name   string
	childs map[string]*NodeConf
}

func NewConf(name string) *Contree {
	return &Contree{
		name:   name,
		childs: make(map[string]*NodeConf),
	}
}

func (c *Contree) Load(contree *Contree) error {
	return nil
}

func (c *Contree) Sprintf(str string) string {
	regex := regexp.MustCompile(`%[\w.]+%`)
	for _, found := range regex.FindAllString(str, -1) {
		str = strings.ReplaceAll(str, found, c.Get(found[1:len(found)-1]))
	}
	return str
}

func (c *Contree) Set(path string, value string) {
	pathSlice := strings.Split(path, ".")
	if node, ok := c.childs[pathSlice[0]]; ok {
		node.Set(strings.Join(pathSlice[1:], "."), value)
	}
}

func (c *Contree) Get(path string) string {
	pathSlice := strings.Split(path, ".")
	if node, ok := c.childs[pathSlice[0]]; ok {
		if str, err := node.Get(strings.Join(pathSlice[1:], ".")); err == nil {
			return str
		}
	}
	return ""
}

func (c *Contree) Add(node *NodeConf) {
	node.Level = 0
	c.childs[node.Name] = node
}
