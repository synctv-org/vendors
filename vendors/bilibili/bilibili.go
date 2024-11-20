package bilibili

type qrcodeResp struct {
	Data struct {
		URL       string `json:"url"`
		QrcodeKey string `json:"qrcode_key"`
	} `json:"data"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	TTL     int    `json:"ttl"`
}

type videoPageInfo struct {
	Message string `json:"message"`
	Data    struct {
		Premiere interface{} `json:"premiere"`
		Owner    struct {
			Name string `json:"name"`
			Face string `json:"face"`
			Mid  int    `json:"mid"`
		} `json:"owner"`
		VtDisplay string `json:"vt_display"`
		LikeIcon  string `json:"like_icon"`
		Tname     string `json:"tname"`
		Bvid      string `json:"bvid"`
		Pic       string `json:"pic"`
		Title     string `json:"title"`
		Dynamic   string `json:"dynamic"`
		UserGarb  struct {
			URLImageAniCut string `json:"url_image_ani_cut"`
		} `json:"user_garb"`
		Desc string `json:"desc"`
		Stat struct {
			Evaluation string `json:"evaluation"`
			ArgueMsg   string `json:"argue_msg"`
			Share      int    `json:"share"`
			Reply      int    `json:"reply"`
			Favorite   int    `json:"favorite"`
			Coin       int    `json:"coin"`
			Aid        int    `json:"aid"`
			NowRank    int    `json:"now_rank"`
			HisRank    int    `json:"his_rank"`
			Like       int    `json:"like"`
			Dislike    int    `json:"dislike"`
			Danmaku    int    `json:"danmaku"`
			View       int    `json:"view"`
			Vt         int    `json:"vt"`
		} `json:"stat"`
		Subtitle struct {
			List        []interface{} `json:"list"`
			AllowSubmit bool          `json:"allow_submit"`
		} `json:"subtitle"`
		DescV2 []struct {
			RawText string `json:"raw_text"`
			Type    int    `json:"type"`
			BizID   int    `json:"biz_id"`
		} `json:"desc_v2"`
		HonorReply struct {
			Honor []struct {
				Desc               string `json:"desc"`
				Aid                int    `json:"aid"`
				Type               int    `json:"type"`
				WeeklyRecommendNum int    `json:"weekly_recommend_num"`
			} `json:"honor"`
		} `json:"honor_reply"`
		Pages []struct {
			From       string `json:"from"`
			Part       string `json:"part"`
			Vid        string `json:"vid"`
			Weblink    string `json:"weblink"`
			FirstFrame string `json:"first_frame"`
			Dimension  struct {
				Width  int `json:"width"`
				Height int `json:"height"`
				Rotate int `json:"rotate"`
			} `json:"dimension"`
			Cid      uint64 `json:"cid"`
			Page     int    `json:"page"`
			Duration int    `json:"duration"`
		} `json:"pages"`
		UgcSeason struct {
			Title    string `json:"title"`
			Cover    string `json:"cover"`
			Intro    string `json:"intro"`
			Sections []struct {
				Title    string `json:"title"`
				Episodes []struct {
					Title string `json:"title"`
					Bvid  string `json:"bvid"`
					Page  struct {
						From      string `json:"from"`
						Part      string `json:"part"`
						Vid       string `json:"vid"`
						Weblink   string `json:"weblink"`
						Dimension struct {
							Width  int `json:"width"`
							Height int `json:"height"`
							Rotate int `json:"rotate"`
						} `json:"dimension"`
						Cid      int `json:"cid"`
						Page     int `json:"page"`
						Duration int `json:"duration"`
					} `json:"page"`
					Arc struct {
						DescV2 interface{} `json:"desc_v2"`
						Author struct {
							Name string `json:"name"`
							Face string `json:"face"`
							Mid  int    `json:"mid"`
						} `json:"author"`
						Dynamic   string `json:"dynamic"`
						TypeName  string `json:"type_name"`
						VtDisplay string `json:"vt_display"`
						Pic       string `json:"pic"`
						Title     string `json:"title"`
						Desc      string `json:"desc"`
						Stat      struct {
							Evaluation string `json:"evaluation"`
							ArgueMsg   string `json:"argue_msg"`
							Share      int    `json:"share"`
							Reply      int    `json:"reply"`
							Fav        int    `json:"fav"`
							Coin       int    `json:"coin"`
							Aid        int    `json:"aid"`
							NowRank    int    `json:"now_rank"`
							HisRank    int    `json:"his_rank"`
							Like       int    `json:"like"`
							Dislike    int    `json:"dislike"`
							Danmaku    int    `json:"danmaku"`
							View       int    `json:"view"`
							Vt         int    `json:"vt"`
							Vv         int    `json:"vv"`
						} `json:"stat"`
						Rights struct {
							Bp            int `json:"bp"`
							Elec          int `json:"elec"`
							Download      int `json:"download"`
							Movie         int `json:"movie"`
							Pay           int `json:"pay"`
							Hd5           int `json:"hd5"`
							NoReprint     int `json:"no_reprint"`
							Autoplay      int `json:"autoplay"`
							UgcPay        int `json:"ugc_pay"`
							IsCooperation int `json:"is_cooperation"`
							UgcPayPreview int `json:"ugc_pay_preview"`
							ArcPay        int `json:"arc_pay"`
							FreeWatch     int `json:"free_watch"`
						} `json:"rights"`
						Dimension struct {
							Width  int `json:"width"`
							Height int `json:"height"`
							Rotate int `json:"rotate"`
						} `json:"dimension"`
						Duration           int  `json:"duration"`
						State              int  `json:"state"`
						TypeID             int  `json:"type_id"`
						Ctime              int  `json:"ctime"`
						Aid                int  `json:"aid"`
						Pubdate            int  `json:"pubdate"`
						Videos             int  `json:"videos"`
						EnableVt           int  `json:"enable_vt"`
						Copyright          int  `json:"copyright"`
						IsChargeableSeason bool `json:"is_chargeable_season"`
						IsBlooper          bool `json:"is_blooper"`
					} `json:"arc"`
					SeasonID  int    `json:"season_id"`
					SectionID int    `json:"section_id"`
					ID        int    `json:"id"`
					Aid       int    `json:"aid"`
					Cid       uint64 `json:"cid"`
					Attribute int    `json:"attribute"`
				} `json:"episodes"`
				SeasonID int `json:"season_id"`
				ID       int `json:"id"`
				Type     int `json:"type"`
			} `json:"sections"`
			Stat struct {
				SeasonID int `json:"season_id"`
				View     int `json:"view"`
				Danmaku  int `json:"danmaku"`
				Reply    int `json:"reply"`
				Fav      int `json:"fav"`
				Coin     int `json:"coin"`
				Share    int `json:"share"`
				NowRank  int `json:"now_rank"`
				HisRank  int `json:"his_rank"`
				Like     int `json:"like"`
				Vt       int `json:"vt"`
				Vv       int `json:"vv"`
			} `json:"stat"`
			Attribute   int  `json:"attribute"`
			ID          int  `json:"id"`
			SignState   int  `json:"sign_state"`
			Mid         int  `json:"mid"`
			EpCount     int  `json:"ep_count"`
			SeasonType  int  `json:"season_type"`
			EnableVt    int  `json:"enable_vt"`
			IsPaySeason bool `json:"is_pay_season"`
		} `json:"ugc_season"`
		Rights struct {
			Bp            int `json:"bp"`
			Elec          int `json:"elec"`
			Download      int `json:"download"`
			Movie         int `json:"movie"`
			Pay           int `json:"pay"`
			Hd5           int `json:"hd5"`
			NoReprint     int `json:"no_reprint"`
			Autoplay      int `json:"autoplay"`
			UgcPay        int `json:"ugc_pay"`
			IsCooperation int `json:"is_cooperation"`
			UgcPayPreview int `json:"ugc_pay_preview"`
			NoBackground  int `json:"no_background"`
			CleanMode     int `json:"clean_mode"`
			IsSteinGate   int `json:"is_stein_gate"`
			Is360         int `json:"is_360"`
			NoShare       int `json:"no_share"`
			ArcPay        int `json:"arc_pay"`
			FreeWatch     int `json:"free_watch"`
		} `json:"rights"`
		Dimension struct {
			Width  int `json:"width"`
			Height int `json:"height"`
			Rotate int `json:"rotate"`
		} `json:"dimension"`
		Cid                int  `json:"cid"`
		Copyright          int  `json:"copyright"`
		Duration           int  `json:"duration"`
		TeenageMode        int  `json:"teenage_mode"`
		Aid                uint `json:"aid"`
		Videos             int  `json:"videos"`
		Tid                int  `json:"tid"`
		SeasonID           int  `json:"season_id"`
		EnableVt           int  `json:"enable_vt"`
		State              int  `json:"state"`
		Pubdate            int  `json:"pubdate"`
		Ctime              int  `json:"ctime"`
		NoCache            bool `json:"no_cache"`
		IsUpowerPlay       bool `json:"is_upower_play"`
		IsSeasonDisplay    bool `json:"is_season_display"`
		IsUpowerExclusive  bool `json:"is_upower_exclusive"`
		IsStory            bool `json:"is_story"`
		IsChargeableSeason bool `json:"is_chargeable_season"`
		NeedJumpBv         bool `json:"need_jump_bv"`
		DisableShowUpInfo  bool `json:"disable_show_up_info"`
	} `json:"data"`
	Code int `json:"code"`
	TTL  int `json:"ttl"`
}

type videoInfo struct {
	Message string `json:"message"`
	Data    struct {
		HighFormat        interface{} `json:"high_format"`
		From              string      `json:"from"`
		Result            string      `json:"result"`
		Message           string      `json:"message"`
		Format            string      `json:"format"`
		SeekType          string      `json:"seek_type"`
		AcceptFormat      string      `json:"accept_format"`
		SeekParam         string      `json:"seek_param"`
		AcceptQuality     []uint64    `json:"accept_quality"`
		AcceptDescription []string    `json:"accept_description"`
		Durl              []struct {
			BackupURL interface{} `json:"backup_url"`
			Ahead     string      `json:"ahead"`
			Vhead     string      `json:"vhead"`
			URL       string      `json:"url"`
			Order     int         `json:"order"`
			Length    int         `json:"length"`
			Size      int         `json:"size"`
		} `json:"durl"`
		SupportFormats []struct {
			Codecs         interface{} `json:"codecs"`
			Format         string      `json:"format"`
			NewDescription string      `json:"new_description"`
			DisplayDesc    string      `json:"display_desc"`
			Superscript    string      `json:"superscript"`
			Quality        int         `json:"quality"`
		} `json:"support_formats"`
		VideoCodecid int    `json:"video_codecid"`
		Timelength   int    `json:"timelength"`
		Quality      uint64 `json:"quality"`
		LastPlayTime int    `json:"last_play_time"`
		LastPlayCid  int    `json:"last_play_cid"`
	} `json:"data"`
	Code int `json:"code"`
	TTL  int `json:"ttl"`
}

type playerV2Info struct {
	Message string `json:"message"`
	Data    struct {
		BgmInfo      interface{} `json:"bgm_info"`
		OnlineSwitch struct {
			EnableGrayDashPlayback string `json:"enable_gray_dash_playback"`
			NewBroadcast           string `json:"new_broadcast"`
			RealtimeDm             string `json:"realtime_dm"`
			SubtitleSubmitSwitch   string `json:"subtitle_submit_switch"`
		} `json:"online_switch"`
		Name         string `json:"name"`
		Bvid         string `json:"bvid"`
		PreviewToast string `json:"preview_toast"`
		Role         string `json:"role"`
		Permission   string `json:"permission"`
		LoginMidHash string `json:"login_mid_hash"`
		IPInfo       struct {
			IP       string `json:"ip"`
			ZoneIP   string `json:"zone_ip"`
			Country  string `json:"country"`
			Province string `json:"province"`
			City     string `json:"city"`
			ZoneID   int    `json:"zone_id"`
		} `json:"ip_info"`
		PlayerIcon struct {
			URL1  string `json:"url1"`
			Hash1 string `json:"hash1"`
			URL2  string `json:"url2"`
			Hash2 string `json:"hash2"`
			Ctime int    `json:"ctime"`
		} `json:"player_icon"`
		ElecHighLevel struct {
			LevelStr      string `json:"level_str"`
			Title         string `json:"title"`
			Intro         string `json:"intro"`
			PrivilegeType int    `json:"privilege_type"`
		} `json:"elec_high_level"`
		GuideAttention []interface{} `json:"guide_attention"`
		OperationCard  []interface{} `json:"operation_card"`
		JumpCard       []interface{} `json:"jump_card"`
		ViewPoints     []interface{} `json:"view_points"`
		Subtitle       struct {
			Lan       string `json:"lan"`
			LanDoc    string `json:"lan_doc"`
			Subtitles []struct {
				Lan         string `json:"lan"`
				LanDoc      string `json:"lan_doc"`
				SubtitleURL string `json:"subtitle_url"`
				IDStr       string `json:"id_str"`
				ID          int64  `json:"id"`
				Type        int    `json:"type"`
				AiType      int    `json:"ai_type"`
				AiStatus    int    `json:"ai_status"`
				IsLock      bool   `json:"is_lock"`
			} `json:"subtitles"`
			AllowSubmit bool `json:"allow_submit"`
		} `json:"subtitle"`
		Vip struct {
			AvatarSubscriptURL string `json:"avatar_subscript_url"`
			NicknameColor      string `json:"nickname_color"`
			Label              struct {
				Path                  string `json:"path"`
				Text                  string `json:"text"`
				LabelTheme            string `json:"label_theme"`
				TextColor             string `json:"text_color"`
				BgColor               string `json:"bg_color"`
				BorderColor           string `json:"border_color"`
				ImgLabelURIHans       string `json:"img_label_uri_hans"`
				ImgLabelURIHant       string `json:"img_label_uri_hant"`
				ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"`
				ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"`
				BgStyle               int    `json:"bg_style"`
				UseImgLabel           bool   `json:"use_img_label"`
			} `json:"label"`
			AvatarSubscript int   `json:"avatar_subscript"`
			ThemeType       int   `json:"theme_type"`
			Type            int   `json:"type"`
			VipPayType      int   `json:"vip_pay_type"`
			DueDate         int64 `json:"due_date"`
			Role            int   `json:"role"`
			Status          int   `json:"status"`
			TvVipStatus     int   `json:"tv_vip_status"`
			TvVipPayType    int   `json:"tv_vip_pay_type"`
			TvDueDate       int   `json:"tv_due_date"`
		} `json:"vip"`
		LevelInfo struct {
			CurrentLevel int   `json:"current_level"`
			CurrentMin   int   `json:"current_min"`
			CurrentExp   int   `json:"current_exp"`
			NextExp      int   `json:"next_exp"`
			LevelUp      int64 `json:"level_up"`
		} `json:"level_info"`
		Fawkes struct {
			ConfigVersion int `json:"config_version"`
			FfVersion     int `json:"ff_version"`
		} `json:"fawkes"`
		LastPlayTime int `json:"last_play_time"`
		Aid          int `json:"aid"`
		NowTime      int `json:"now_time"`
		OnlineCount  int `json:"online_count"`
		LoginMid     int `json:"login_mid"`
		Cid          int `json:"cid"`
		PageNo       int `json:"page_no"`
		AnswerStatus int `json:"answer_status"`
		BlockTime    int `json:"block_time"`
		MaxLimit     int `json:"max_limit"`
		LastPlayCid  int `json:"last_play_cid"`
		Options      struct {
			Is360      bool `json:"is_360"`
			WithoutVip bool `json:"without_vip"`
		} `json:"options"`
		IsUgcPayPreview   bool `json:"is_ugc_pay_preview"`
		HasNext           bool `json:"has_next"`
		IsOwner           bool `json:"is_owner"`
		NeedLoginSubtitle bool `json:"need_login_subtitle"`
		ShowSwitch        struct {
			LongProgress bool `json:"long_progress"`
		} `json:"show_switch"`
		NoShare           bool `json:"no_share"`
		ToastBlock        bool `json:"toast_block"`
		IsUpowerExclusive bool `json:"is_upower_exclusive"`
		IsUpowerPlay      bool `json:"is_upower_play"`
		AllowBp           bool `json:"allow_bp"`
		DisableShowUpInfo bool `json:"disable_show_up_info"`
	} `json:"data"`
	Code int `json:"code"`
	TTL  int `json:"ttl"`
}

type seasonInfo struct {
	Message string `json:"message"`
	Result  struct {
		IconFont struct {
			Name string `json:"name"`
			Text string `json:"text"`
		} `json:"icon_font"`
		Evaluate      string `json:"evaluate"`
		Subtitle      string `json:"subtitle"`
		Title         string `json:"title"`
		Alias         string `json:"alias"`
		Record        string `json:"record"`
		SquareCover   string `json:"square_cover"`
		Staff         string `json:"staff"`
		JpTitle       string `json:"jp_title"`
		SeasonTitle   string `json:"season_title"`
		ShareSubTitle string `json:"share_sub_title"`
		ShareCopy     string `json:"share_copy"`
		Link          string `json:"link"`
		Actors        string `json:"actors"`
		ShareURL      string `json:"share_url"`
		Cover         string `json:"cover"`
		BkgCover      string `json:"bkg_cover"`
		Activity      struct {
			HeadBgURL string `json:"head_bg_url"`
			Title     string `json:"title"`
			ID        int    `json:"id"`
		} `json:"activity"`
		Episodes []struct {
			BadgeInfo struct {
				BgColor      string `json:"bg_color"`
				BgColorNight string `json:"bg_color_night"`
				Text         string `json:"text"`
			} `json:"badge_info"`
			ShareURL    string `json:"share_url"`
			ShareCopy   string `json:"share_copy"`
			Title       string `json:"title"`
			Bvid        string `json:"bvid"`
			Link        string `json:"link"`
			Cover       string `json:"cover"`
			ShortLink   string `json:"short_link"`
			Badge       string `json:"badge"`
			LongTitle   string `json:"long_title"`
			ReleaseDate string `json:"release_date"`
			From        string `json:"from"`
			Subtitle    string `json:"subtitle"`
			Vid         string `json:"vid"`
			Rights      struct {
				AllowDemand   int `json:"allow_demand"`
				AllowDm       int `json:"allow_dm"`
				AllowDownload int `json:"allow_download"`
				AreaLimit     int `json:"area_limit"`
			} `json:"rights"`
			Skip struct {
				Ed struct {
					End   int `json:"end"`
					Start int `json:"start"`
				} `json:"ed"`
				Op struct {
					End   int `json:"end"`
					Start int `json:"start"`
				} `json:"op"`
			} `json:"skip"`
			Dimension struct {
				Height int `json:"height"`
				Rotate int `json:"rotate"`
				Width  int `json:"width"`
			} `json:"dimension"`
			Cid                uint64 `json:"cid"`
			PubTime            int    `json:"pub_time"`
			Aid                int    `json:"aid"`
			EpID               uint64 `json:"ep_id"`
			ID                 int    `json:"id"`
			Duration           int    `json:"duration"`
			Pv                 int    `json:"pv"`
			BadgeType          int    `json:"badge_type"`
			Status             int    `json:"status"`
			EnableVt           bool   `json:"enable_vt"`
			ShowDrmLoginDialog bool   `json:"showDrmLoginDialog"`
			IsViewHide         bool   `json:"is_view_hide"`
		} `json:"episodes"`
		Positive struct {
			Title string `json:"title"`
			ID    int    `json:"id"`
		} `json:"positive"`
		Styles  []string `json:"styles"`
		Seasons []struct {
			BadgeInfo struct {
				BgColor      string `json:"bg_color"`
				BgColorNight string `json:"bg_color_night"`
				Text         string `json:"text"`
			} `json:"badge_info"`
			IconFont struct {
				Name string `json:"name"`
				Text string `json:"text"`
			} `json:"icon_font"`
			Badge               string `json:"badge"`
			Cover               string `json:"cover"`
			HorizontalCover1610 string `json:"horizontal_cover_1610"`
			HorizontalCover169  string `json:"horizontal_cover_169"`
			SeasonTitle         string `json:"season_title"`
			NewEp               struct {
				Cover     string `json:"cover"`
				IndexShow string `json:"index_show"`
				ID        int    `json:"id"`
			} `json:"new_ep"`
			Stat struct {
				Favorites    int `json:"favorites"`
				SeriesFollow int `json:"series_follow"`
				Views        int `json:"views"`
				Vt           int `json:"vt"`
			} `json:"stat"`
			MediaID    int  `json:"media_id"`
			BadgeType  int  `json:"badge_type"`
			SeasonID   int  `json:"season_id"`
			SeasonType int  `json:"season_type"`
			EnableVt   bool `json:"enable_vt"`
		} `json:"seasons"`
		Section []struct {
			Title      string        `json:"title"`
			EpisodeIDs []interface{} `json:"episode_ids"`
			Episodes   []struct {
				BadgeInfo struct {
					BgColor      string `json:"bg_color"`
					BgColorNight string `json:"bg_color_night"`
					Text         string `json:"text"`
				} `json:"badge_info"`
				IconFont struct {
					Name string `json:"name"`
					Text string `json:"text"`
				} `json:"icon_font"`
				Link         string `json:"link"`
				Badge        string `json:"badge"`
				Bvid         string `json:"bvid"`
				Title        string `json:"title"`
				Cover        string `json:"cover"`
				ReleaseDate  string `json:"release_date"`
				Subtitle     string `json:"subtitle"`
				LongTitle    string `json:"long_title"`
				Vid          string `json:"vid"`
				From         string `json:"from"`
				ShareURL     string `json:"share_url"`
				ShareCopy    string `json:"share_copy"`
				ShortLink    string `json:"short_link"`
				StatForUnity struct {
					Danmaku struct {
						Icon     string `json:"icon"`
						PureText string `json:"pure_text"`
						Text     string `json:"text"`
						Value    int    `json:"value"`
					} `json:"danmaku"`
					Vt struct {
						Icon     string `json:"icon"`
						PureText string `json:"pure_text"`
						Text     string `json:"text"`
						Value    int    `json:"value"`
					} `json:"vt"`
					Coin  int `json:"coin"`
					Likes int `json:"likes"`
					Reply int `json:"reply"`
				} `json:"stat_for_unity"`
				Stat struct {
					Coin     int `json:"coin"`
					Danmakus int `json:"danmakus"`
					Likes    int `json:"likes"`
					Play     int `json:"play"`
					Reply    int `json:"reply"`
					Vt       int `json:"vt"`
				} `json:"stat"`
				Rights struct {
					AllowDemand   int `json:"allow_demand"`
					AllowDm       int `json:"allow_dm"`
					AllowDownload int `json:"allow_download"`
					AreaLimit     int `json:"area_limit"`
				} `json:"rights"`
				Skip struct {
					Ed struct {
						End   int `json:"end"`
						Start int `json:"start"`
					} `json:"ed"`
					Op struct {
						End   int `json:"end"`
						Start int `json:"start"`
					} `json:"op"`
				} `json:"skip"`
				Dimension struct {
					Height int `json:"height"`
					Rotate int `json:"rotate"`
					Width  int `json:"width"`
				} `json:"dimension"`
				Duration           int  `json:"duration"`
				Pv                 int  `json:"pv"`
				PubTime            int  `json:"pub_time"`
				ID                 int  `json:"id"`
				EpID               int  `json:"ep_id"`
				Status             int  `json:"status"`
				Aid                int  `json:"aid"`
				Cid                int  `json:"cid"`
				BadgeType          int  `json:"badge_type"`
				IsViewHide         bool `json:"is_view_hide"`
				ShowDrmLoginDialog bool `json:"showDrmLoginDialog"`
				EnableVt           bool `json:"enable_vt"`
			} `json:"episodes"`
			Attr      int `json:"attr"`
			EpisodeID int `json:"episode_id"`
			ID        int `json:"id"`
			Type      int `json:"type"`
			Type2     int `json:"type2"`
		} `json:"section"`
		Areas []struct {
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"areas"`
		PlayStrategy struct {
			Strategies []string `json:"strategies"`
		} `json:"play_strategy"`
		NewEp struct {
			Desc  string `json:"desc"`
			Title string `json:"title"`
			ID    int    `json:"id"`
			IsNew int    `json:"is_new"`
		} `json:"new_ep"`
		Series struct {
			SeriesTitle string `json:"series_title"`
			DisplayType int    `json:"display_type"`
			SeriesID    int    `json:"series_id"`
		} `json:"series"`
		Freya struct {
			BubbleDesc    string `json:"bubble_desc"`
			BubbleShowCnt int    `json:"bubble_show_cnt"`
			IconShow      int    `json:"icon_show"`
		} `json:"freya"`
		Publish struct {
			PubTime       string `json:"pub_time"`
			PubTimeShow   string `json:"pub_time_show"`
			IsFinish      int    `json:"is_finish"`
			IsStarted     int    `json:"is_started"`
			UnknowPubDate int    `json:"unknow_pub_date"`
			Weekday       int    `json:"weekday"`
		} `json:"publish"`
		UpInfo struct {
			VipLabel struct {
				BgColor     string `json:"bg_color"`
				BorderColor string `json:"border_color"`
				Text        string `json:"text"`
				TextColor   string `json:"text_color"`
				BgStyle     int    `json:"bg_style"`
			} `json:"vip_label"`
			AvatarSubscriptURL string `json:"avatar_subscript_url"`
			Avatar             string `json:"avatar"`
			Uname              string `json:"uname"`
			NicknameColor      string `json:"nickname_color"`
			Pendant            struct {
				Image string `json:"image"`
				Name  string `json:"name"`
				Pid   int    `json:"pid"`
			} `json:"pendant"`
			IsFollow   int `json:"is_follow"`
			ThemeType  int `json:"theme_type"`
			Mid        int `json:"mid"`
			VerifyType int `json:"verify_type"`
			Follower   int `json:"follower"`
			VipStatus  int `json:"vip_status"`
			VipType    int `json:"vip_type"`
		} `json:"up_info"`
		Payment struct {
			Price             string `json:"price"`
			Promotion         string `json:"promotion"`
			Tip               string `json:"tip"`
			VipFirstPromotion string `json:"vip_first_promotion"`
			VipPrice          string `json:"vip_price"`
			VipPromotion      string `json:"vip_promotion"`
			PayType           struct {
				AllowDiscount    int `json:"allow_discount"`
				AllowPack        int `json:"allow_pack"`
				AllowTicket      int `json:"allow_ticket"`
				AllowTimeLimit   int `json:"allow_time_limit"`
				AllowVipDiscount int `json:"allow_vip_discount"`
				ForbidBb         int `json:"forbid_bb"`
			} `json:"pay_type"`
			Discount      int `json:"discount"`
			ViewStartTime int `json:"view_start_time"`
			VipDiscount   int `json:"vip_discount"`
		} `json:"payment"`
		Stat struct {
			FollowText string `json:"follow_text"`
			Coins      int    `json:"coins"`
			Danmakus   int    `json:"danmakus"`
			Favorite   int    `json:"favorite"`
			Favorites  int    `json:"favorites"`
			Likes      int    `json:"likes"`
			Reply      int    `json:"reply"`
			Share      int    `json:"share"`
			Views      int    `json:"views"`
			Vt         int    `json:"vt"`
		} `json:"stat"`
		Rights struct {
			Copyright       string `json:"copyright"`
			Resource        string `json:"resource"`
			CanWatch        int    `json:"can_watch"`
			AllowReview     int    `json:"allow_review"`
			AreaLimit       int    `json:"area_limit"`
			BanAreaShow     int    `json:"ban_area_show"`
			AllowBp         int    `json:"allow_bp"`
			AllowDownload   int    `json:"allow_download"`
			ForbidPre       int    `json:"forbid_pre"`
			FreyaWhite      int    `json:"freya_white"`
			IsCoverShow     int    `json:"is_cover_show"`
			IsPreview       int    `json:"is_preview"`
			OnlyVipDownload int    `json:"only_vip_download"`
			AllowBpRank     int    `json:"allow_bp_rank"`
			WatchPlatform   int    `json:"watch_platform"`
		} `json:"rights"`
		UserStatus struct {
			AreaLimit    int `json:"area_limit"`
			BanAreaShow  int `json:"ban_area_show"`
			Follow       int `json:"follow"`
			FollowStatus int `json:"follow_status"`
			Login        int `json:"login"`
			Pay          int `json:"pay"`
			PayPackPaid  int `json:"pay_pack_paid"`
			Sponsor      int `json:"sponsor"`
		} `json:"user_status"`
		Rating struct {
			Count int     `json:"count"`
			Score float64 `json:"score"`
		} `json:"rating"`
		Mode   int `json:"mode"`
		Status int `json:"status"`
		Show   struct {
			WideScreen int `json:"wide_screen"`
		} `json:"show"`
		MediaID        int  `json:"media_id"`
		Total          int  `json:"total"`
		Type           int  `json:"type"`
		SeasonID       int  `json:"season_id"`
		ShowSeasonType int  `json:"show_season_type"`
		EnableVt       bool `json:"enable_vt"`
	} `json:"result"`
	Code int `json:"code"`
}

type pgcURLInfo struct {
	Message string `json:"message"`
	Result  struct {
		RecordInfo struct {
			RecordIcon string `json:"record_icon"`
			Record     string `json:"record"`
		} `json:"record_info"`
		Result         string        `json:"result"`
		SeekParam      string        `json:"seek_param"`
		Message        string        `json:"message"`
		AcceptFormat   string        `json:"accept_format"`
		Type           string        `json:"type"`
		SeekType       string        `json:"seek_type"`
		From           string        `json:"from"`
		Format         string        `json:"format"`
		ClipInfoList   []interface{} `json:"clip_info_list"`
		AcceptQuality  []uint64      `json:"accept_quality"`
		SupportFormats []struct {
			DisplayDesc    string        `json:"display_desc"`
			Superscript    string        `json:"superscript"`
			Format         string        `json:"format"`
			Description    string        `json:"description"`
			NewDescription string        `json:"new_description"`
			Codecs         []interface{} `json:"codecs"`
			Quality        int           `json:"quality"`
			NeedLogin      bool          `json:"need_login,omitempty"`
		} `json:"support_formats"`
		Durl []struct {
			Ahead     string   `json:"ahead"`
			Vhead     string   `json:"vhead"`
			URL       string   `json:"url"`
			Md5       string   `json:"md5"`
			BackupURL []string `json:"backup_url"`
			Size      int      `json:"size"`
			Length    int      `json:"length"`
			Order     int      `json:"order"`
		} `json:"durl"`
		AcceptDescription []string `json:"accept_description"`
		Fnver             int      `json:"fnver"`
		NoRexcode         int      `json:"no_rexcode"`
		VideoCodecid      int      `json:"video_codecid"`
		Bp                int      `json:"bp"`
		Fnval             int      `json:"fnval"`
		Quality           uint64   `json:"quality"`
		Timelength        int      `json:"timelength"`
		IsPreview         int      `json:"is_preview"`
		Code              int      `json:"code"`
		Status            int      `json:"status"`
		IsDrm             bool     `json:"is_drm"`
		VideoProject      bool     `json:"video_project"`
		HasPaid           bool     `json:"has_paid"`
	} `json:"result"`
	Code int `json:"code"`
}

type Nav struct {
	Message string `json:"message"`
	Data    struct {
		VipLabel struct {
			Path                  string `json:"path"`
			Text                  string `json:"text"`
			LabelTheme            string `json:"label_theme"`
			TextColor             string `json:"text_color"`
			BgColor               string `json:"bg_color"`
			BorderColor           string `json:"border_color"`
			ImgLabelURIHans       string `json:"img_label_uri_hans"`
			ImgLabelURIHant       string `json:"img_label_uri_hant"`
			ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"`
			ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"`
			BgStyle               int    `json:"bg_style"`
			UseImgLabel           bool   `json:"use_img_label"`
		} `json:"vip_label"`
		WbiImg struct {
			ImgURL string `json:"img_url"`
			SubURL string `json:"sub_url"`
		} `json:"wbi_img"`
		OfficialVerify struct {
			Desc string `json:"desc"`
			Type int    `json:"type"`
		} `json:"officialVerify"`
		Uname            string `json:"uname"`
		Face             string `json:"face"`
		ShopURL          string `json:"shop_url"`
		VipNicknameColor string `json:"vip_nickname_color"`
		Pendant          struct {
			Name              string `json:"name"`
			Image             string `json:"image"`
			ImageEnhance      string `json:"image_enhance"`
			ImageEnhanceFrame string `json:"image_enhance_frame"`
			Pid               int    `json:"pid"`
			Expire            int    `json:"expire"`
			NPid              int    `json:"n_pid"`
		} `json:"pendant"`
		Official struct {
			Title string `json:"title"`
			Desc  string `json:"desc"`
			Role  int    `json:"role"`
			Type  int    `json:"type"`
		} `json:"official"`
		Vip struct {
			AvatarSubscriptURL string `json:"avatar_subscript_url"`
			NicknameColor      string `json:"nickname_color"`
			Label              struct {
				Path                  string `json:"path"`
				Text                  string `json:"text"`
				LabelTheme            string `json:"label_theme"`
				TextColor             string `json:"text_color"`
				BgColor               string `json:"bg_color"`
				BorderColor           string `json:"border_color"`
				ImgLabelURIHans       string `json:"img_label_uri_hans"`
				ImgLabelURIHant       string `json:"img_label_uri_hant"`
				ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"`
				ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"`
				BgStyle               int    `json:"bg_style"`
				UseImgLabel           bool   `json:"use_img_label"`
			} `json:"label"`
			AvatarSubscript int   `json:"avatar_subscript"`
			ThemeType       int   `json:"theme_type"`
			Type            int   `json:"type"`
			VipPayType      int   `json:"vip_pay_type"`
			DueDate         int64 `json:"due_date"`
			Role            int   `json:"role"`
			Status          int   `json:"status"`
			TvVipStatus     int   `json:"tv_vip_status"`
			TvVipPayType    int   `json:"tv_vip_pay_type"`
			TvDueDate       int   `json:"tv_due_date"`
		} `json:"vip"`
		Wallet struct {
			Mid           int     `json:"mid"`
			BcoinBalance  float64 `json:"bcoin_balance"`
			CouponBalance int     `json:"coupon_balance"`
			CouponDueTime int     `json:"coupon_due_time"`
		} `json:"wallet"`
		LevelInfo struct {
			CurrentLevel int `json:"current_level"`
			CurrentMin   int `json:"current_min"`
			CurrentExp   int `json:"current_exp"`
		} `json:"level_info"`
		MobileVerified     int   `json:"mobile_verified"`
		FaceNft            int   `json:"face_nft"`
		VipDueDate         int64 `json:"vipDueDate"`
		VipStatus          int   `json:"vipStatus"`
		VipType            int   `json:"vipType"`
		VipPayType         int   `json:"vip_pay_type"`
		VipThemeType       int   `json:"vip_theme_type"`
		Moral              int   `json:"moral"`
		VipAvatarSubscript int   `json:"vip_avatar_subscript"`
		EmailVerified      int   `json:"email_verified"`
		Mid                int   `json:"mid"`
		FaceNftType        int   `json:"face_nft_type"`
		IsSeniorMember     int   `json:"is_senior_member"`
		Scores             int   `json:"scores"`
		AllowanceCount     int   `json:"allowance_count"`
		AnswerStatus       int   `json:"answer_status"`
		HasShop            bool  `json:"has_shop"`
		IsLogin            bool  `json:"isLogin"`
		IsJury             bool  `json:"is_jury"`
	} `json:"data"`
	Code int `json:"code"`
	TTL  int `json:"ttl"` // Money          int `json:"money"` // 可能为float
}

type dashResp struct {
	Message string `json:"message"`
	Data    struct {
		HighFormat        interface{} `json:"high_format"`
		SeekParam         string      `json:"seek_param"`
		SeekType          string      `json:"seek_type"`
		Message           string      `json:"message"`
		From              string      `json:"from"`
		Result            string      `json:"result"`
		AcceptFormat      string      `json:"accept_format"`
		Format            string      `json:"format"`
		AcceptQuality     []int       `json:"accept_quality"`
		AcceptDescription []string    `json:"accept_description"`
		SupportFormats    []struct {
			Format         string   `json:"format"`
			NewDescription string   `json:"new_description"`
			DisplayDesc    string   `json:"display_desc"`
			Superscript    string   `json:"superscript"`
			Codecs         []string `json:"codecs"`
			Quality        int      `json:"quality"`
		} `json:"support_formats"`
		Dash struct {
			Dolby struct {
				Audio interface{} `json:"audio"`
				Type  int         `json:"type"`
			} `json:"dolby"`
			Flac  interface{} `json:"flac"`
			Video []struct {
				SegmentBase struct {
					Initialization string `json:"initialization"`
					IndexRange     string `json:"index_range"`
				} `json:"segment_base"`
				FrameRate    string   `json:"frame_rate"`
				BaseURL      string   `json:"base_url"`
				MimeType     string   `json:"mime_type"`
				Codecs       string   `json:"codecs"`
				Sar          string   `json:"sar"`
				BackupURL    []string `json:"backup_url"`
				Bandwidth    int64    `json:"bandwidth"`
				Height       int64    `json:"height"`
				Width        int64    `json:"width"`
				StartWithSap int64    `json:"start_with_sap"`
				ID           int      `json:"id"`
				Codecid      int      `json:"codecid"`
			} `json:"video"`
			Audio []struct {
				SegmentBase struct {
					Initialization string `json:"initialization"`
					IndexRange     string `json:"index_range"`
				} `json:"segment_base"`
				FrameRate    string   `json:"frame_rate"`
				BaseURL      string   `json:"base_url"`
				MimeType     string   `json:"mime_type"`
				Codecs       string   `json:"codecs"`
				Sar          string   `json:"sar"`
				BackupURL    []string `json:"backup_url"`
				Bandwidth    int64    `json:"bandwidth"`
				Height       int      `json:"height"`
				Width        int      `json:"width"`
				StartWithSap int64    `json:"start_with_sap"`
				ID           int      `json:"id"`
				Codecid      int      `json:"codecid"`
			} `json:"audio"`
			Duration      float64 `json:"duration"`
			MinBufferTime float64 `json:"min_buffer_time"`
		} `json:"dash"`
		VideoCodecid int `json:"video_codecid"`
		Quality      int `json:"quality"`
		Timelength   int `json:"timelength"`
		LastPlayTime int `json:"last_play_time"`
		LastPlayCid  int `json:"last_play_cid"`
	} `json:"data"`
	Code int `json:"code"`
	TTL  int `json:"ttl"`
}

type dashPGCResp struct {
	Message string `json:"message"`
	Result  struct {
		RecordInfo struct {
			RecordIcon string `json:"record_icon"`
			Record     string `json:"record"`
		} `json:"record_info"`
		Result       string `json:"result"`
		SeekParam    string `json:"seek_param"`
		AcceptFormat string `json:"accept_format"`
		Type         string `json:"type"`
		Message      string `json:"message"`
		Format       string `json:"format"`
		SeekType     string `json:"seek_type"`
		From         string `json:"from"`
		ClipInfoList []struct {
			ToastText  string `json:"toastText"`
			ClipType   string `json:"clipType"`
			MaterialNo int    `json:"materialNo"`
			Start      int    `json:"start"`
			End        int    `json:"end"`
		} `json:"clip_info_list"`
		AcceptQuality  []int `json:"accept_quality"`
		SupportFormats []struct {
			DisplayDesc    string   `json:"display_desc"`
			Superscript    string   `json:"superscript"`
			Format         string   `json:"format"`
			Description    string   `json:"description"`
			NewDescription string   `json:"new_description"`
			Codecs         []string `json:"codecs"`
			Quality        int      `json:"quality"`
			NeedLogin      bool     `json:"need_login,omitempty"`
			NeedVip        bool     `json:"need_vip,omitempty"`
		} `json:"support_formats"`
		AcceptDescription []string `json:"accept_description"`
		Dash              struct {
			Video []struct {
				SegmentBase struct {
					Initialization string `json:"initialization"`
					IndexRange     string `json:"index_range"`
				} `json:"segment_base"`
				MimeType     string   `json:"mime_type"`
				Md5          string   `json:"md5"`
				Sar          string   `json:"sar"`
				Codecs       string   `json:"codecs"`
				BaseURL      string   `json:"base_url"`
				FrameRate    string   `json:"frame_rate"`
				BackupURL    []string `json:"backup_url"`
				Codecid      int      `json:"codecid"`
				Size         int      `json:"size"`
				StartWithSap int64    `json:"start_with_sap"`
				Width        int64    `json:"width"`
				StartWithSAP int      `json:"startWithSAP"`
				ID           int      `json:"id"`
				Height       int64    `json:"height"`
				Bandwidth    int64    `json:"bandwidth"`
			} `json:"video"`
			Audio []struct {
				SegmentBase struct {
					Initialization string `json:"initialization"`
					IndexRange     string `json:"index_range"`
				} `json:"segment_base"`
				MimeType     string   `json:"mime_type"`
				Md5          string   `json:"md5"`
				Sar          string   `json:"sar"`
				Codecs       string   `json:"codecs"`
				BaseURL      string   `json:"base_url"`
				FrameRate    string   `json:"frame_rate"`
				BackupURL    []string `json:"backup_url"`
				Codecid      int      `json:"codecid"`
				Size         int      `json:"size"`
				StartWithSap int64    `json:"start_with_sap"`
				Width        int      `json:"width"`
				StartWithSAP int      `json:"startWithSAP"`
				ID           int      `json:"id"`
				Height       int      `json:"height"`
				Bandwidth    int64    `json:"bandwidth"`
			} `json:"audio"`
			Dolby struct {
				Audio []interface{} `json:"audio"`
				Type  int           `json:"type"`
			} `json:"dolby"`
			Duration      float64 `json:"duration"`
			MinBufferTime float64 `json:"min_buffer_time"`
		} `json:"dash"`
		Fnver        int  `json:"fnver"`
		NoRexcode    int  `json:"no_rexcode"`
		VideoCodecid int  `json:"video_codecid"`
		Bp           int  `json:"bp"`
		Quality      int  `json:"quality"`
		Timelength   int  `json:"timelength"`
		Fnval        int  `json:"fnval"`
		IsPreview    int  `json:"is_preview"`
		Code         int  `json:"code"`
		Status       int  `json:"status"`
		IsDrm        bool `json:"is_drm"`
		VideoProject bool `json:"video_project"`
		HasPaid      bool `json:"has_paid"`
	} `json:"result"`
	Code int `json:"code"`
}

type parseLivePageResp struct {
	Message string `json:"message"`
	Data    struct {
		Title      string `json:"title"`
		UserCover  string `json:"user_cover"`
		UID        uint64 `json:"uid"`
		RoomID     uint64 `json:"room_id"`
		SortID     uint64 `json:"sort_id"`
		LiveStatus uint64 `json:"live_status"`
	} `json:"data"`
	Code int `json:"code"`
}

type getLiveMasterInfoResp struct {
	Message string `json:"message"`
	Data    struct {
		Info struct {
			Username string `json:"uname"`
			Face     string `json:"face"`
			UID      uint64 `json:"uid"`
		} `json:"info"`
	} `json:"data"`
	Code int `json:"code"`
}

type getLiveStreamResp struct {
	Message string `json:"message"`
	Data    struct {
		AcceptQuality      []string `json:"accept_quality"`
		QualityDescription []struct {
			Desc    string `json:"desc"`
			Quality uint64 `json:"qn"`
		} `json:"quality_description"`
		DUrl []struct {
			URL   string `json:"url"`
			Order int    `json:"order"`
		} `json:"durl"`
		CurrentQuality uint64 `json:"current_quality"`
		CurrentQn      uint64 `json:"current_qn"`
	} `json:"data"`
	Code int `json:"code"`
	TTL  int `json:"ttl"`
}
