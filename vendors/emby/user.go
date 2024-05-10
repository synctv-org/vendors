package emby

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	json "github.com/json-iterator/go"
)

type GetApiKeyResp struct {
	AccessToken string `json:"AccessToken"`
	UserID      string `json:"User"`
	ServerID    string `json:"ServerId"`
}

func (c *Client) GetApiKey(username, password string) (*GetApiKeyResp, error) {
	req, err := c.NewRequest(http.MethodPost, "/emby/Users/authenticatebyname", &LoginReq{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Emby Client=\"Emby Web\", Device=\"SyncTV\", DeviceId=\"%s\", Version=\"4.7.14.0\"", uuid.NewString()))
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(b))
	}
	var loginResp LoginResp
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		return nil, err
	}
	return &GetApiKeyResp{
		AccessToken: loginResp.AccessToken,
		UserID:      loginResp.User.ID,
		ServerID:    loginResp.ServerID,
	}, nil
}

func (c *Client) Login(username, password string) error {
	s, err := c.GetApiKey(username, password)
	if err != nil {
		return err
	}
	c.SetKey(s.AccessToken)
	c.SetUserID(s.UserID)
	return nil
}

func (c *Client) Logout() error {
	req, err := c.NewRequest(http.MethodPost, "/emby/Sessions/Logout", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Emby Client=\"Emby Web\", Device=\"SyncTV\", DeviceId=\"%s\", Version=\"4.7.14.0\"", uuid.NewString()))
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(b))
	}
	c.SetKey("")
	c.SetUserID("")
	return nil
}

func (c *Client) Me() (*MeResp, error) {
	if c.userID == "" {
		return nil, errors.New("user id not set")
	}
	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("/emby/Users/%s", c.userID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("status code %d: %s", resp.StatusCode, string(b))
	}
	var meResp MeResp
	if err := json.NewDecoder(resp.Body).Decode(&meResp); err != nil {
		return nil, err
	}
	return &meResp, nil
}

func (c *Client) UserViews(opt ...QueryFunc) (*ItemsResp, error) {
	if c.userID == "" {
		return nil, errors.New("user id not set")
	}
	o := map[string]string{}
	for _, f := range opt {
		f(o)
	}
	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("/emby/Users/%s/Views", c.userID), o)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("status code %d: %s", resp.StatusCode, string(b))
	}
	var viewsResp ItemsResp
	if err := json.NewDecoder(resp.Body).Decode(&viewsResp); err != nil {
		return nil, err
	}
	return &viewsResp, nil
}

func (c *Client) UserItemsByID(id string) (*Items, error) {
	if c.userID == "" {
		return nil, errors.New("user id not set")
	}
	if id == "" {
		return nil, errors.New("id not set")
	}
	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("/emby/Users/%s/Items/%s", c.userID, id), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("status code %d: %s", resp.StatusCode, string(b))
	}
	var itemsResp Items
	if err := json.NewDecoder(resp.Body).Decode(&itemsResp); err != nil {
		return nil, err
	}
	return &itemsResp, nil
}

func (c *Client) UserItems(opt ...QueryFunc) (*ItemsResp, error) {
	if c.userID == "" {
		return nil, errors.New("user id not set")
	}
	o := map[string]string{
		"Fields": "MediaSources,ParentId,Container",
	}
	for _, f := range opt {
		f(o)
	}
	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("/emby/Users/%s/Items", c.userID), nil, o)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("status code %d: %s", resp.StatusCode, string(b))
	}
	var itemsResp ItemsResp
	if err := json.NewDecoder(resp.Body).Decode(&itemsResp); err != nil {
		return nil, err
	}
	return &itemsResp, nil
}

func (c *Client) Seasons(id string, opt ...QueryFunc) (*ItemsResp, error) {
	if c.userID == "" {
		return nil, errors.New("user id not set")
	}
	o := map[string]string{}
	for _, f := range opt {
		f(o)
	}
	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("/emby/Shows/%s/Seasons", id), o)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("status code %d: %s", resp.StatusCode, string(b))
	}
	var itemsResp ItemsResp
	if err := json.NewDecoder(resp.Body).Decode(&itemsResp); err != nil {
		return nil, err
	}
	return &itemsResp, nil
}

func (c *Client) Episodes(seriesID, seasonId string, opt ...QueryFunc) (*ItemsResp, error) {
	if c.userID == "" {
		return nil, errors.New("user id not set")
	}
	o := map[string]string{
		"SeasonId": seasonId,
	}
	for _, f := range opt {
		f(o)
	}
	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("/emby/Shows/%s/Episodes", seriesID), nil, o)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("status code %d: %s", resp.StatusCode, string(b))
	}
	var itemsResp ItemsResp
	if err := json.NewDecoder(resp.Body).Decode(&itemsResp); err != nil {
		return nil, err
	}
	return &itemsResp, nil
}

type UserPlaybackInfoOption func(*PlayBackReq)

func WithMaxStreamingBitrate(bitrate int) UserPlaybackInfoOption {
	return func(p *PlayBackReq) {
		p.MaxStreamingBitrate = bitrate
	}
}

func WithSubtitleStreamIndex(index int) UserPlaybackInfoOption {
	return func(p *PlayBackReq) {
		p.SubtitleStreamIndex = index
	}
}

func WithAudioStreamIndex(index int) UserPlaybackInfoOption {
	return func(p *PlayBackReq) {
		p.AudioStreamIndex = index
	}
}

func WithMediaSourceID(id string) UserPlaybackInfoOption {
	return func(p *PlayBackReq) {
		p.MediaSourceID = id
	}
}

func (c *Client) UserPlaybackInfo(id string, opts ...UserPlaybackInfoOption) (*PlayBackResp, error) {
	if c.userID == "" {
		return nil, errors.New("user id not set")
	}
	cfg := &PlayBackReq{
		DeviceProfile: DeviceProfile{
			MaxStaticBitrate:                 140000000,
			MaxStreamingBitrate:              140000000,
			MusicStreamingTranscodingBitrate: 192000,
			DirectPlayProfiles: []DirectPlayProfile{
				{
					Container:  "mp4,m4v",
					Type:       "Video",
					VideoCodec: "h264",
					AudioCodec: "mp3,aac",
				},
				{
					Container:  "mkv",
					Type:       "Video",
					VideoCodec: "h264",
					AudioCodec: "mp3,aac",
				},
			},
			TranscodingProfiles: []TranscodingProfile{
				{
					Container:        "mkv",
					Type:             "Video",
					AudioCodec:       "mp3,aac",
					VideoCodec:       "h264",
					Context:          "Static",
					MaxAudioChannels: "2",
					CopyTimestamps:   true,
				},
				{
					Container:           "ts",
					Type:                "Video",
					AudioCodec:          "mp3,aac",
					VideoCodec:          "h264",
					Context:             "Streaming",
					Protocol:            "hls",
					MaxAudioChannels:    "2",
					MinSegments:         "1",
					BreakOnNonKeyFrames: true,
					ManifestSubtitles:   "vtt",
				},
				{
					Container:  "mp4",
					Type:       "Video",
					AudioCodec: "mp3,aac",
					VideoCodec: "h264",
					Context:    "Static",
					Protocol:   "http",
				},
			},
			CodecProfiles: []CodecProfile{
				{
					Type:  "VideoAudio",
					Codec: "aac",
					Conditions: []Condition{
						{
							Condition:  "Equals",
							Property:   "IsSecondaryAudio",
							Value:      "false",
							IsRequired: "false",
						},
					},
				},
				{
					Type: "VideoAudio",
					Conditions: []Condition{
						{
							Condition:  "Equals",
							Property:   "IsSecondaryAudio",
							Value:      "false",
							IsRequired: "false",
						},
					},
				},
				{
					Type:  "Video",
					Codec: "h264",
					Conditions: []Condition{
						{
							Condition:  "EqualsAny",
							Property:   "VideoProfile",
							Value:      "high|main|baseline|constrained baseline|high 10",
							IsRequired: false,
						},
						{
							Condition:  "LessThanEqual",
							Property:   "VideoLevel",
							Value:      "62",
							IsRequired: false,
						},
						{
							Condition:  "LessThanEqual",
							Property:   "Width",
							Value:      "3840",
							IsRequired: false,
						},
					},
				},
				{
					Type: "Video",
					Conditions: []Condition{
						{
							Condition:  "LessThanEqual",
							Property:   "Width",
							Value:      "3840",
							IsRequired: false,
						},
					},
				},
			},
			SubtitleProfiles: []SubtitleProfile{
				{
					Format: "vtt",
					Method: "Hls",
				},
				{
					Format: "vtt",
					Method: "External",
				},
				{
					Format: "ass",
					Method: "External",
				},
			},
			ResponseProfiles: []ResponseProfile{
				{
					Type:      "Video",
					Container: "m4v",
					MimeType:  "video/mp4",
				},
			},
		},
		UserID:              c.userID,
		StartTimeTicks:      0,
		IsPlayback:          true,
		AutoOpenLiveStream:  true,
		MaxStreamingBitrate: 8000001,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	req, err := c.NewRequest(http.MethodPost, fmt.Sprintf("/emby/Items/%s/PlaybackInfo", id), cfg, map[string]string{
		"reqformat": "json",
	})
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("status code %d: %s", resp.StatusCode, string(b))
	}
	var playBackResp PlayBackResp
	if err := json.NewDecoder(resp.Body).Decode(&playBackResp); err != nil {
		return nil, err
	}
	return &playBackResp, nil
}

func (c *Client) DeleteActiveEncodeings(playSessionID string) error {
	if playSessionID == "" {
		return errors.New("play session id is empty")
	}
	req, err := c.NewRequest(http.MethodPost, "/emby/Videos/ActiveEncodings/Delete", nil, map[string]string{
		"PlaySessionId": playSessionID,
	})
	if err != nil {
		return err
	}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("status code %d: %s", resp.StatusCode, string(b))
	}
	return nil
}
