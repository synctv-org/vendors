package bilibili_test

import (
	"fmt"
	"testing"

	"github.com/synctv-org/vendors/vendors/bilibili"
)

func TestMatch(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		url     string
		wantT   string
		wantID  string
		wantErr bool
	}{
		{
			name:    "bv",
			url:     "https://www.bilibili.com/video/BV1i5411y7fB",
			wantT:   "bv",
			wantID:  "BV1i5411y7fB",
			wantErr: false,
		},
		{
			name:    "av",
			url:     "https://www.bilibili.com/video/av1",
			wantT:   "av",
			wantID:  "1",
			wantErr: false,
		},
		{
			name:    "ss",
			url:     "https://www.bilibili.com/bangumi/play/ss1",
			wantT:   "ss",
			wantID:  "1",
			wantErr: false,
		},
		{
			name:    "ep",
			url:     "https://www.bilibili.com/bangumi/play/ep1",
			wantT:   "ep",
			wantID:  "1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotT, gotID, err := bilibili.Match(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Match() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotT != tt.wantT {
				t.Errorf("Match() gotT = %v, want %v", gotT, tt.wantT)
			}
			if gotID != tt.wantID {
				t.Errorf("Match() gotID = %v, want %v", gotID, tt.wantID)
			}
		})
	}
}

func TestGetDashVideoURL(t *testing.T) {
	c, err := bilibili.NewClient(nil)
	if err != nil {
		t.Fatal(err)
	}
	m, m2, err := c.GetDashVideoURL(0, "BV1y7411Q7Eq", 171776208)
	if err != nil {
		t.Fatal(err)
	}
	for _, as := range m.GetCurrentPeriod().AdaptationSets {
		for _, r := range as.Representations {
			t.Log(r.BaseURL)
		}
	}
	for _, as := range m2.GetCurrentPeriod().AdaptationSets {
		for _, r := range as.Representations {
			t.Log(r.BaseURL)
		}
	}
}

func TestGetDashVideoURLMPDFile(t *testing.T) {
	c, err := bilibili.NewClient(nil)
	if err != nil {
		t.Fatal(err)
	}
	m, _, err := c.GetDashVideoURL(0, "BV1y7411Q7Eq", 171776208)
	if err != nil {
		t.Fatal(err)
	}
	s, err := m.WriteToString()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
}

func TestEditAndGetDashVideoURLMPDFile(t *testing.T) {
	c, err := bilibili.NewClient(nil)
	if err != nil {
		t.Fatal(err)
	}
	m, _, err := c.GetDashVideoURL(0, "BV1y7411Q7Eq", 171776208)
	if err != nil {
		t.Fatal(err)
	}
	m.BaseURL = append(m.BaseURL, "/")
	id := 0
	for _, as := range m.GetCurrentPeriod().AdaptationSets {
		for _, r := range as.Representations {
			for i := range r.BaseURL {
				r.BaseURL[i] = fmt.Sprintf("/api/movie/proxy/roomid/movieid?id=%d", id)
				id++
			}
		}
	}
	s, err := m.WriteToString()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
}

func TestGetSubtitles(t *testing.T) {
	// need to login
	c, err := bilibili.NewClient(nil)
	if err != nil {
		t.Fatal(err)
	}
	subs, err := c.GetSubtitles(0, "BV1sw411Q7dr", 1248205613)
	if err != nil {
		t.Fatal(err)
	}
	for _, sub := range subs {
		t.Log(sub)
	}
}

func TestGetLiveStreams(t *testing.T) {
	c, err := bilibili.NewClient(nil)
	if err != nil {
		t.Fatal(err)
	}
	streams, err := c.GetLiveStreams(7777, true)
	if err != nil {
		t.Fatal(err)
	}
	for _, stream := range streams {
		t.Log(stream)
	}

	streams, err = c.GetLiveStreams(7777, false)
	if err != nil {
		t.Fatal(err)
	}
	for _, stream := range streams {
		t.Log(stream)
	}
}
