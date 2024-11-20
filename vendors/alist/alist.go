package alist

import "time"

type Resp[T any] struct {
	Data    T      `json:"data"`
	Message string `json:"message"`
	Code    uint64 `json:"code"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp = Resp[loginResp]

type loginResp struct {
	Token string `json:"token"`
}

type FsGetReq struct {
	Path     string `json:"path"`
	Password string `json:"password"`
}

type FsGetResp = Resp[FsGetRespData]

type FsGetRespData struct {
	Modified time.Time       `json:"modified"`
	Created  time.Time       `json:"created"`
	HashInfo any             `json:"hash_info"`
	Sign     string          `json:"sign"`
	Thumb    string          `json:"thumb"`
	Hashinfo string          `json:"hashinfo"`
	Name     string          `json:"name"`
	RawURL   string          `json:"raw_url"`
	Readme   string          `json:"readme"`
	Provider string          `json:"provider"`
	Related  []*FsGetRelated `json:"related"`
	Size     uint64          `json:"size"`
	Type     uint64          `json:"type"`
	IsDir    bool            `json:"is_dir"`
}

type FsGetRelated struct {
	Modified time.Time `json:"modified"`
	Created  time.Time `json:"created"`
	Name     string    `json:"name"`
	Sign     string    `json:"sign"`
	Thumb    string    `json:"thumb"`
	Hashinfo string    `json:"hashinfo"`
	Size     uint64    `json:"size"`
	Type     uint64    `json:"type"`
	IsDir    bool      `json:"is_dir"`
}

type FsListReq struct {
	Path     string `json:"path"`
	Password string `json:"password"`
	Page     uint64 `json:"page"`
	PerPage  uint64 `json:"per_page"`
	Refresh  bool   `json:"refresh"`
}

type FsListResp = Resp[fsListResp]

type fsListResp struct {
	Readme   string `json:"readme"`
	Provider string `json:"provider"`
	Content  []struct {
		Modified time.Time `json:"modified"`
		Name     string    `json:"name"`
		Sign     string    `json:"sign"`
		Thumb    string    `json:"thumb"`
		Size     uint64    `json:"size"`
		Type     uint64    `json:"type"`
		IsDir    bool      `json:"is_dir"`
	} `json:"content"`
	Total uint64 `json:"total"`
	Write bool   `json:"write"`
}

const (
	FsOtherReqMethodVideoPreview = "video_preview"
)

type FsOtherReq struct {
	Path     string `json:"path"`
	Method   string `json:"method"`
	Password string `json:"password"`
}

type FsOtherResp = Resp[fsOtherResp]

type fsOtherResp struct {
	DriveID              string `json:"drive_id"`
	FileID               string `json:"file_id"`
	PlayCursor           any    `json:"play_cursor"`
	VideoPreviewPlayInfo struct {
		Category                        string `json:"category"`
		LiveTranscodingSubtitleTaskList []struct {
			Language string `json:"language"`
			Status   string `json:"status"`
			URL      string `json:"url"`
		} `json:"live_transcoding_subtitle_task_list"`
		LiveTranscodingTaskList []struct {
			Stage          string `json:"stage"`
			Status         string `json:"status"`
			TemplateID     string `json:"template_id"`
			TemplateName   string `json:"template_name"`
			URL            string `json:"url"`
			TemplateHeight uint64 `json:"template_height"`
			TemplateWidth  uint64 `json:"template_width"`
		} `json:"live_transcoding_task_list"`
		Meta struct {
			Duration float64 `json:"duration"`
			Height   uint64  `json:"height"`
			Width    uint64  `json:"width"`
		} `json:"meta"`
	} `json:"video_preview_play_info"`
}

type FsMkdirReq struct {
	Path string `json:"path"`
}

type FsMkdirResp = Resp[any]

type FsRenameReq struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

type FsRenameResp = Resp[any]

type FsRemoveReq struct {
	Names []string `json:"names"`
	Dir   bool     `json:"dir"`
}

type FsRemoveResp = Resp[any]

type MeResp = Resp[meResp]

type meResp struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	BasePath   string `json:"base_path"`
	SsoID      string `json:"sso_id"`
	ID         uint64 `json:"id"`
	Role       uint64 `json:"role"`
	Permission uint64 `json:"permission"`
	Disabled   bool   `json:"disabled"`
	Otp        bool   `json:"otp"`
}

type FsSearchReq struct {
	Parent   string `json:"parent"`
	Keywords string `json:"keywords"`
	Password string `json:"password"`
	Scope    uint64 `json:"scope"`
	Page     uint64 `json:"page"`
	PerPage  uint64 `json:"per_page"`
}

type fsSearchResp struct {
	Content []struct {
		Parent string `json:"parent"`
		Name   string `json:"name"`
		IsDir  bool   `json:"is_dir"`
		Size   uint64 `json:"size"`
		Type   uint64 `json:"type"`
	} `json:"content"`
	Total uint64 `json:"total"`
}

type FsSearchResp = Resp[fsSearchResp]
