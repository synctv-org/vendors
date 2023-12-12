package emby_test

import (
	"testing"

	"github.com/synctv-org/vendors/vendors/emby"
)

var (
	cli *emby.Client
)

// func TestLibrary(t *testing.T) {
// 	flr, err := cli.Library()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Logf("%+v", flr)
// }

// func TestGetItems(t *testing.T) {
// 	flr, err := cli.GetItems(527)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Logf("%+v", flr)
// }

func TestGetItem(t *testing.T) {
	flr, err := cli.GetItem("7944")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", flr)
}

func TestMe(t *testing.T) {
	flr, err := cli.Me()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", flr)
}
