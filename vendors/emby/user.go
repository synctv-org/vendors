package emby

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	json "github.com/json-iterator/go"
)

type GetAPIKeyResp struct {
	AccessToken string `json:"AccessToken"`
	UserID      string `json:"User"`
	ServerID    string `json:"ServerId"`
}

func (c *Client) GetAPIKey(username, password string) (*GetAPIKeyResp, error) {
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
	return &GetAPIKeyResp{
		AccessToken: loginResp.AccessToken,
		UserID:      loginResp.User.ID,
		ServerID:    loginResp.ServerID,
	}, nil
}

func (c *Client) Login(username, password string) error {
	s, err := c.GetAPIKey(username, password)
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
	req, err := c.NewRequest(http.MethodGet, "/emby/Users/"+c.userID, nil)
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

func (c *Client) Episodes(seriesID, seasonID string, opt ...QueryFunc) (*ItemsResp, error) {
	if c.userID == "" {
		return nil, errors.New("user id not set")
	}
	o := map[string]string{
		"SeasonId": seasonID,
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
			MaxStaticBitrate:                 100000000,
			MaxStreamingBitrate:              120000000,
			MusicStreamingTranscodingBitrate: 384000,
			DirectPlayProfiles: []DirectPlayProfile{
				{
					Container:  "webm",
					Type:       "Video",
					VideoCodec: "vp8,vp9,av1",
					AudioCodec: "vorbis,opus",
				},
				{
					Container:  "mp4,m4v",
					Type:       "Video",
					VideoCodec: "h264,hevc,vp9,av1",
					AudioCodec: "aac,mp3,mp2,opus,flac,vorbis",
				},
				{
					Container:  "mov",
					Type:       "Video",
					VideoCodec: "h264",
					AudioCodec: "aac,mp3,mp2,opus,flac,vorbis",
				},
				{
					Container: "opus",
					Type:      "Audio",
				},
				{
					Container:  "webm",
					Type:       "Audio",
					AudioCodec: "opus",
				},
				{
					Container:  "ts",
					Type:       "Audio",
					AudioCodec: "mp3",
				},
				{
					Container: "mp3",
					Type:      "Audio",
				},
				{
					Container: "aac",
					Type:      "Audio",
				},
				{
					Container:  "m4a",
					Type:       "Audio",
					AudioCodec: "aac",
				},
				{
					Container:  "m4b",
					Type:       "Audio",
					AudioCodec: "aac",
				},
				{
					Container: "flac",
					Type:      "Audio",
				},
				{
					Container: "webma",
					Type:      "Audio",
				},
				{
					Container:  "webm",
					Type:       "Audio",
					AudioCodec: "webma",
				},
				{
					Container: "wav",
					Type:      "Audio",
				},
				{
					Container: "ogg",
					Type:      "Audio",
				},
				{
					Container:  "hls",
					Type:       "Video",
					VideoCodec: "av1,hevc,h264,vp9",
					AudioCodec: "aac,mp2,opus,flac",
				},
				{
					Container:  "hls",
					Type:       "Video",
					VideoCodec: "h264",
					AudioCodec: "aac,mp3,mp2",
				},
			},
			TranscodingProfiles: []TranscodingProfile{
				{
					Container:              "mp4",
					Type:                   "Audio",
					AudioCodec:             "aac",
					Context:                "Streaming",
					Protocol:               "hls",
					MaxAudioChannels:       "2",
					MinSegments:            "2",
					BreakOnNonKeyFrames:    true,
					EnableAudioVbrEncoding: true,
				},
				{
					Container:        "aac",
					Type:             "Audio",
					AudioCodec:       "aac",
					Context:          "Streaming",
					Protocol:         "http",
					MaxAudioChannels: "2",
				},
				{
					Container:        "mp3",
					Type:             "Audio",
					AudioCodec:       "mp3",
					Context:          "Streaming",
					Protocol:         "http",
					MaxAudioChannels: "2",
				},
				{
					Container:        "opus",
					Type:             "Audio",
					AudioCodec:       "opus",
					Context:          "Streaming",
					Protocol:         "http",
					MaxAudioChannels: "2",
				},
				{
					Container:        "wav",
					Type:             "Audio",
					AudioCodec:       "wav",
					Context:          "Streaming",
					Protocol:         "http",
					MaxAudioChannels: "2",
				},
				{
					Container:        "opus",
					Type:             "Audio",
					AudioCodec:       "opus",
					Context:          "Static",
					Protocol:         "http",
					MaxAudioChannels: "2",
				},
				{
					Container:        "mp3",
					Type:             "Audio",
					AudioCodec:       "mp3",
					Context:          "Static",
					Protocol:         "http",
					MaxAudioChannels: "2",
				},
				{
					Container:        "aac",
					Type:             "Audio",
					AudioCodec:       "aac",
					Context:          "Static",
					Protocol:         "http",
					MaxAudioChannels: "2",
				},
				{
					Container:        "wav",
					Type:             "Audio",
					AudioCodec:       "wav",
					Context:          "Static",
					Protocol:         "http",
					MaxAudioChannels: "2",
				},
				{
					Container:           "mp4",
					Type:                "Video",
					AudioCodec:          "aac,mp2,opus,flac",
					VideoCodec:          "av1,hevc,h264,vp9",
					Context:             "Streaming",
					Protocol:            "hls",
					MaxAudioChannels:    "2",
					MinSegments:         "2",
					BreakOnNonKeyFrames: true,
				},
				{
					Container:           "ts",
					Type:                "Video",
					AudioCodec:          "aac,mp3,mp2",
					VideoCodec:          "h264",
					Context:             "Streaming",
					Protocol:            "hls",
					MaxAudioChannels:    "2",
					MinSegments:         "2",
					BreakOnNonKeyFrames: true,
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
							IsRequired: false,
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
							IsRequired: false,
						},
					},
				},
				{
					Type:  "Video",
					Codec: "h264",
					Conditions: []Condition{
						{
							Condition:  "NotEquals",
							Property:   "IsAnamorphic",
							Value:      "true",
							IsRequired: false,
						},
						{
							Condition:  "EqualsAny",
							Property:   "VideoProfile",
							Value:      "high|main|baseline|constrained baseline|high 10",
							IsRequired: false,
						},
						{
							Condition:  "EqualsAny",
							Property:   "VideoRangeType",
							Value:      "SDR",
							IsRequired: false,
						},
						{
							Condition:  "LessThanEqual",
							Property:   "VideoLevel",
							Value:      "52",
							IsRequired: false,
						},
						{
							Condition:  "NotEquals",
							Property:   "IsInterlaced",
							Value:      "true",
							IsRequired: false,
						},
					},
				},
				{
					Type:  "Video",
					Codec: "hevc",
					Conditions: []Condition{
						{
							Condition:  "NotEquals",
							Property:   "IsAnamorphic",
							Value:      "true",
							IsRequired: false,
						},
						{
							Condition:  "EqualsAny",
							Property:   "VideoProfile",
							Value:      "main|main 10",
							IsRequired: false,
						},
						{
							Condition:  "EqualsAny",
							Property:   "VideoRangeType",
							Value:      "SDR|HDR10|HLG",
							IsRequired: false,
						},
						{
							Condition:  "LessThanEqual",
							Property:   "VideoLevel",
							Value:      "183",
							IsRequired: false,
						},
						{
							Condition:  "NotEquals",
							Property:   "IsInterlaced",
							Value:      "true",
							IsRequired: false,
						},
					},
				},
				{
					Type:  "Video",
					Codec: "vp9",
					Conditions: []Condition{
						{
							Condition:  "EqualsAny",
							Property:   "VideoRangeType",
							Value:      "SDR|HDR10|HLG",
							IsRequired: false,
						},
					},
				},
				{
					Type:  "Video",
					Codec: "av1",
					Conditions: []Condition{
						{
							Condition:  "NotEquals",
							Property:   "IsAnamorphic",
							Value:      "true",
							IsRequired: false,
						},
						{
							Condition:  "EqualsAny",
							Property:   "VideoProfile",
							Value:      "main",
							IsRequired: false,
						},
						{
							Condition:  "EqualsAny",
							Property:   "VideoRangeType",
							Value:      "SDR|HDR10|HLG",
							IsRequired: false,
						},
						{
							Condition:  "LessThanEqual",
							Property:   "VideoLevel",
							Value:      "19",
							IsRequired: false,
						},
					},
				},
			},
			SubtitleProfiles: []SubtitleProfile{
				{
					Format: "vtt",
					Method: "External",
				},
				{
					Format: "ass",
					Method: "External",
				},
				{
					Format: "ssa",
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
