package alist_test

import (
	"testing"

	"github.com/synctv-org/vendors/vendors/alist"
)

func TestFsGet(t *testing.T) {
	tests := []struct {
		path    string
		wantErr bool
	}{
		{
			path:    "/",
			wantErr: false,
		},
		{
			path:    "/test",
			wantErr: false,
		},
		{
			path:    "/test/test",
			wantErr: true,
		},
	}
	client, err := alist.NewClient("https://al.nn.ci/", "")
	if err != nil {
		t.Fatalf("NewClient error: %v", err)
	}
	for _, v := range tests {
		get, err := client.FsGet(&alist.FsGetReq{
			Path: v.path,
		})
		if (err != nil) != v.wantErr {
			t.Errorf("path: %s, wantErr: %v, gotErr: %v", v.path, v.wantErr, err)
			return
		}
		t.Logf("path: %s, get: %+v", v.path, get)
	}
}

func TestFsList(t *testing.T) {
	tests := []struct {
		path    string
		wantErr bool
	}{
		{
			path:    "/",
			wantErr: false,
		},
	}
	client, err := alist.NewClient("https://al.nn.ci/", "")
	if err != nil {
		t.Fatalf("NewClient error: %v", err)
	}
	for _, v := range tests {
		list, err := client.FsList(&alist.FsListReq{
			Path:    v.path,
			Page:    1,
			PerPage: 1,
		})
		if (err != nil) != v.wantErr {
			t.Errorf("path: %s, wantErr: %v, gotErr: %v", v.path, v.wantErr, err)
			return
		}
		t.Logf("path: %s, list: %+v", v.path, list)
	}
}

func TestMe(t *testing.T) {
	client, err := alist.NewClient("https://al.nn.ci/", "")
	if err != nil {
		t.Fatalf("NewClient error: %v", err)
	}
	me, err := client.Me()
	if err != nil {
		t.Fatalf("Me error: %v", err)
	}
	t.Logf("me: %+v", me)
}
