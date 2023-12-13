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
// 	flr, err := cli.GetItems("11")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Logf("%+v", flr)
// }

func TestGetItem(t *testing.T) {
	flr, err := cli.GetItem("1")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", flr)
}

// func TestMe(t *testing.T) {
// 	flr, err := cli.Me()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Logf("%+v", flr)
// }

// func TestSystemInfo(t *testing.T) {
// 	flr, err := cli.SystemInfo()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Logf("%+v", flr)
// }
