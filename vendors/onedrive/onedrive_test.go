package onedrive_test

import (
	"github.com/synctv-org/vendors/vendors/onedrive"
)

var (
	client *onedrive.Client = onedrive.NewClient("", "", "")
)

// func TestUserInfo(t *testing.T) {
// 	resp, err := client.GetUserInfo()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	b, err := json.MarshalIndent(resp, "", "  ")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Logf("%s\n", b)
// }

// func TestDriveList(t *testing.T) {
// 	resp, err := client.DriveList()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	b, err := json.MarshalIndent(resp, "", "  ")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Logf("%s\n", b)
// }

// func TestFsList(t *testing.T) {
// 	resp, err := client.FsList("")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	b, err := json.MarshalIndent(resp, "", "  ")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Logf("%s\n", b)
// }

// func TestFsGet(t *testing.T) {
// 	resp, err := client.FsGet("aaaa.txt")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	b, err := json.MarshalIndent(resp, "", "  ")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Logf("%s\n", b)
// }

// func TestFsDownload(t *testing.T) {
// 	resp, err := client.FsDownload("aaaa.txt")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Logf("%s\n", resp)
// }
