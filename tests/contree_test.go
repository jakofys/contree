package tests

import (
	"testing"

	"github.com/jakofys/contree"
)

func TestContreeSetValue(t *testing.T) {
	conf := contree.NewContree("app")
	conf.Set("path.to.my.value", "valueToRegister")

	if str, err := conf.Get("path.to.my.value"); err != nil {
		t.Error("While contree get, found ", str, " instead of 'valueToRegister'")
	}
}

func TestContreeSprintf(t *testing.T) {
	conf := contree.NewContree("app")
	conf.Set("path.to.my.value", "registry of value")
	expected := "It just a simple data for registry of value"

	if str := conf.Sprintf("It just a simple data for %path.to.my.value%"); str != expected {
		t.Error("Error during contree get, found ", str, " instead of '"+expected+"'")
	}
}

// func TestContreeImportFrom(t *testing.T) {
// 	conf := contree.NewContree("app")
// 	conf.FromFile("../data/conf.env", contree.DOTENV)
// 	expected := "user"

// 	if str := conf.Get("test.longpath.user"); str != expected {
// 		t.Error("Error during contree get after import from .env file, found ", str, " instead of '"+expected+"'")
// 	}
// }
