package contree

import (
	"errors"
	"regexp"
	"strings"
)

type Contree struct {
	root *NodeConf
}

func NewContree(name string) *Contree {
	return &Contree{
		root: &NodeConf{
			Name: name,
		},
	}
}

func (c *Contree) Set(path string, value string) {
	c.root.SetRecursivly(path, value)
}

func (c *Contree) Get(path string) (string, error) {
	value := c.root.Browse(path)
	if value == "" {
		return "", errors.New("Path value not found")
	}
	return value, nil
}

func (c *Contree) Sprintf(str string) string {
	regex := regexp.MustCompile(`%[\w.]+%`)
	for _, found := range regex.FindAllString(str, -1) {
		if value := c.root.Browse(found[1 : len(found)-1]); value != "" {
			str = strings.ReplaceAll(str, found, value)
		}
	}
	return str
}

// func (c *Contree) From(reader io.Reader, codec Codec) error {
// 	var buff []byte
// 	_, err := reader.Read(buff)
// 	if err != nil {
// 		return err
// 	}

// 	node, err := codec.Decode(buff)
// 	if err != nil {
// 		return err
// 	}

// 	c.Load(node)
// 	return nil
// }

// func (c *Contree) FromFile(filename string, codec Codec) error {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()
// 	return c.From(file, codec)
// }
