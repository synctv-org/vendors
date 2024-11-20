package emby_test

import (
	"testing"

	"github.com/synctv-org/vendors/vendors/emby"
)

var cli *emby.Client

func TestLibrary(t *testing.T) {
	flr, err := cli.Library()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", flr)
}

func TestGetItem(t *testing.T) {
	flr, err := cli.GetItem("1")
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

func TestSystemInfo(t *testing.T) {
	flr, err := cli.SystemInfo()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", flr)
}

func TestViews(t *testing.T) {
	flr, err := cli.UserViews()
	if err != nil {
		t.Fatal(err)
	}
	// t.Logf("%+v", flr)
	item, err := cli.UserItemsByID(flr.Items[1].ID)
	if err != nil {
		t.Fatal(err)
	}
	// t.Logf("%+v", item)
	items, err := cli.UserItems(
		emby.WithParentID(item.ID),
		emby.WithIncludeItemTypes("Series"),
		emby.WithSortOrderAsc(),
		emby.WithSortBy("SortName"),
		emby.WithRecursive(),
	)
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range items.Items {
		t.Logf("%s", item.Name)
	}
}
