package alist

import "time"

type AlistResp[T any] struct {
	Code    uint64 `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp = AlistResp[loginResp]

type loginResp struct {
	Token string `json:"token"`
}

type FsGetReq struct {
	Path     string `json:"path"`
	Password string `json:"password"`
}

type FsGetResp = AlistResp[FsGetRespData]

type FsGetRespData struct {
	Name     string          `json:"name"`
	Size     uint64          `json:"size"`
	IsDir    bool            `json:"is_dir"`
	Modified time.Time       `json:"modified"`
	Created  time.Time       `json:"created"`
	Sign     string          `json:"sign"`
	Thumb    string          `json:"thumb"`
	Type     uint64          `json:"type"`
	Hashinfo string          `json:"hashinfo"`
	HashInfo any             `json:"hash_info"`
	RawURL   string          `json:"raw_url"`
	Readme   string          `json:"readme"`
	Provider string          `json:"provider"`
	Related  []*FsGetRelated `json:"related"`
}

type FsGetRelated struct {
	Name     string    `json:"name"`
	Size     uint64    `json:"size"`
	IsDir    bool      `json:"is_dir"`
	Modified time.Time `json:"modified"`
	Created  time.Time `json:"created"`
	Sign     string    `json:"sign"`
	Thumb    string    `json:"thumb"`
	Type     uint64    `json:"type"`
	Hashinfo string    `json:"hashinfo"`
}

type FsListReq struct {
	Path     string `json:"path"`
	Password string `json:"password"`
	Page     uint64 `json:"page"`
	PerPage  uint64 `json:"per_page"`
	Refresh  bool   `json:"refresh"`
}

type FsListResp = AlistResp[fsListResp]

type fsListResp struct {
	Content []struct {
		Name     string    `json:"name"`
		Size     uint64    `json:"size"`
		IsDir    bool      `json:"is_dir"`
		Modified time.Time `json:"modified"`
		Sign     string    `json:"sign"`
		Thumb    string    `json:"thumb"`
		Type     uint64    `json:"type"`
	} `json:"content"`
	Total    uint64 `json:"total"`
	Readme   string `json:"readme"`
	Write    bool   `json:"write"`
	Provider string `json:"provider"`
}

const (
	FsOtherReqMethodVideoPreview = "video_preview"
)

type FsOtherReq struct {
	Path     string `json:"path"`
	Method   string `json:"method"`
	Password string `json:"password"`
}

type FsOtherResp = AlistResp[fsOtherResp]

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
			TemplateHeight uint64 `json:"template_height"`
			TemplateID     string `json:"template_id"`
			TemplateName   string `json:"template_name"`
			TemplateWidth  uint64 `json:"template_width"`
			URL            string `json:"url"`
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

type FsMkdirResp = AlistResp[any]

type FsRenameReq struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

type FsRenameResp = AlistResp[any]

type FsRemoveReq struct {
	Dir   bool     `json:"dir"`
	Names []string `json:"names"`
}

type FsRemoveResp = AlistResp[any]

type MeResp = AlistResp[meResp]

type meResp struct {
	ID         uint64 `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	BasePath   string `json:"base_path"`
	Role       uint64 `json:"role"`
	Disabled   bool   `json:"disabled"`
	Permission uint64 `json:"permission"`
	SsoID      string `json:"sso_id"`
	Otp        bool   `json:"otp"`
}

type FsSearchReq struct {
	Parent   string `json:"parent"`
	Keywords string `json:"keywords"`
	Scope    uint64 `json:"scope"`
	Page     uint64 `json:"page"`
	PerPage  uint64 `json:"per_page"`
	Password string `json:"password"`
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

type FsSearchResp = AlistResp[fsSearchResp]
