package tests

import (
	"testing"

	"github.com/jakofys/contree"
)

func TestNodeConfGetValue(t *testing.T) {

	node := &contree.NodeConf{
		Name:  "name1",
		Value: "value1",
	}

	if node.IsNamed("name") {
		if node.GetValue() != "value" {
			t.Error("Creation of node basicly failed", node)
		}
	}
}

func TestNodeConfInsertNode(t *testing.T) {

	node1 := &contree.NodeConf{
		Name:  "name1",
		Value: "value1",
	}

	node2 := contree.NodeConf{
		Name:  "name2",
		Value: "value2",
	}

	node1.Insert(&node2)

	if node1.Browse("name2") != "value2" {
		t.Error("Not browse to final node search, found: ", node1.Browse("name2"))
	}
}

func TestNodeConfDepthBrowse(t *testing.T) {
	node1 := &contree.NodeConf{
		Name:  "name1",
		Value: "value1",
	}

	node2 := contree.NodeConf{
		Name:  "name2",
		Value: "value2",
	}

	node3 := contree.NodeConf{
		Name:  "name3",
		Value: "value3",
	}

	node2.Insert(&node3)
	node1.Insert(&node2)

	if node1.Browse("name2.name3") != "value3" {
		t.Error("Not browse to final node search, found: ", node1.Browse("name2.name3"))
	}

	if node1.Browse("name2.name4") != "" {
		t.Error("Not browse to final node search, found: ", node1.Browse("name2.name4"))
	}
}

func TestNodeConfSetRecursivly(t *testing.T) {

	node := &contree.NodeConf{
		Name:  "name1",
		Value: "value1",
	}

	node.SetRecursivly("name1.name2.name3", "helloworld")
	node.SetRecursivly("path.to", "salut")
	node.SetRecursivly("path.to.yo", "12312")

	if node.Browse("name1.name2.name3") != "helloworld" {
		t.Error("Recursivly setting value failed, found: ", node.Browse("name1.name2.name3"))
	}

	if node.Browse("path.to.yo") == "helloworld" {
		t.Error("Recursivly setting value failed, found helloworld instead or: ", node.Browse("path.to.yo"))
	}
}

func TestMergeContree(t *testing.T) {

	node1 := &contree.NodeConf{
		Name:  "node1",
		Value: "value1",
	}

	node1.SetRecursivly("for.example.1", "helloworld")
	node1.SetRecursivly("for.example.1.1", "helloworld twice")
	node1.SetRecursivly("for.example.3", "basic info")

	node2 := &contree.NodeConf{
		Name:  "name1",
		Value: "value1",
	}
	node2.SetRecursivly("for.example.2", "bye bye")
	node2.SetRecursivly("for.example.2.1", "bye bye twice")
	node2.SetRecursivly("for.example.3", "rewrite info")

	node1.Merge(node2, true)

	if node1.Browse("for.example.2.1") != "bye bye twice" ||
		node1.Browse("for.example.1.1") != "helloworld twice" ||
		node1.Browse("for.example.3") != "rewrite info" {
		t.Error()
	}
}
