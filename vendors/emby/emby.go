package emby

import "time"

type LoginReq struct {
	Username string `json:"Username"`
	Password string `json:"Pw"`
}

type LoginResp struct {
	User struct {
		Name                      string    `json:"Name"`
		ServerID                  string    `json:"ServerId"`
		Prefix                    string    `json:"Prefix"`
		DateCreated               time.Time `json:"DateCreated"`
		ID                        string    `json:"Id"`
		HasPassword               bool      `json:"HasPassword"`
		HasConfiguredPassword     bool      `json:"HasConfiguredPassword"`
		HasConfiguredEasyPassword bool      `json:"HasConfiguredEasyPassword"`
		LastLoginDate             time.Time `json:"LastLoginDate"`
		LastActivityDate          time.Time `json:"LastActivityDate"`
		Configuration             struct {
			AudioLanguagePreference    string `json:"AudioLanguagePreference"`
			PlayDefaultAudioTrack      bool   `json:"PlayDefaultAudioTrack"`
			DisplayMissingEpisodes     bool   `json:"DisplayMissingEpisodes"`
			SubtitleMode               string `json:"SubtitleMode"`
			EnableLocalPassword        bool   `json:"EnableLocalPassword"`
			HidePlayedInLatest         bool   `json:"HidePlayedInLatest"`
			RememberAudioSelections    bool   `json:"RememberAudioSelections"`
			RememberSubtitleSelections bool   `json:"RememberSubtitleSelections"`
			EnableNextEpisodeAutoPlay  bool   `json:"EnableNextEpisodeAutoPlay"`
			ResumeRewindSeconds        int    `json:"ResumeRewindSeconds"`
			IntroSkipMode              string `json:"IntroSkipMode"`
		} `json:"Configuration"`
		Policy struct {
			IsAdministrator                 bool   `json:"IsAdministrator"`
			IsHidden                        bool   `json:"IsHidden"`
			IsHiddenRemotely                bool   `json:"IsHiddenRemotely"`
			IsHiddenFromUnusedDevices       bool   `json:"IsHiddenFromUnusedDevices"`
			IsDisabled                      bool   `json:"IsDisabled"`
			IsTagBlockingModeInclusive      bool   `json:"IsTagBlockingModeInclusive"`
			EnableUserPreferenceAccess      bool   `json:"EnableUserPreferenceAccess"`
			EnableRemoteControlOfOtherUsers bool   `json:"EnableRemoteControlOfOtherUsers"`
			EnableSharedDeviceControl       bool   `json:"EnableSharedDeviceControl"`
			EnableRemoteAccess              bool   `json:"EnableRemoteAccess"`
			EnableLiveTvManagement          bool   `json:"EnableLiveTvManagement"`
			EnableLiveTvAccess              bool   `json:"EnableLiveTvAccess"`
			EnableMediaPlayback             bool   `json:"EnableMediaPlayback"`
			EnableAudioPlaybackTranscoding  bool   `json:"EnableAudioPlaybackTranscoding"`
			EnableVideoPlaybackTranscoding  bool   `json:"EnableVideoPlaybackTranscoding"`
			EnablePlaybackRemuxing          bool   `json:"EnablePlaybackRemuxing"`
			EnableContentDeletion           bool   `json:"EnableContentDeletion"`
			EnableContentDownloading        bool   `json:"EnableContentDownloading"`
			EnableSubtitleDownloading       bool   `json:"EnableSubtitleDownloading"`
			EnableSubtitleManagement        bool   `json:"EnableSubtitleManagement"`
			EnableSyncTranscoding           bool   `json:"EnableSyncTranscoding"`
			EnableMediaConversion           bool   `json:"EnableMediaConversion"`
			EnableAllChannels               bool   `json:"EnableAllChannels"`
			EnableAllFolders                bool   `json:"EnableAllFolders"`
			InvalidLoginAttemptCount        int    `json:"InvalidLoginAttemptCount"`
			EnablePublicSharing             bool   `json:"EnablePublicSharing"`
			RemoteClientBitrateLimit        int    `json:"RemoteClientBitrateLimit"`
			AuthenticationProviderID        string `json:"AuthenticationProviderId"`
			SimultaneousStreamLimit         int    `json:"SimultaneousStreamLimit"`
			EnableAllDevices                bool   `json:"EnableAllDevices"`
		} `json:"Policy"`
	} `json:"User"`
	SessionInfo struct {
		PlayState struct {
			CanSeek        bool   `json:"CanSeek"`
			IsPaused       bool   `json:"IsPaused"`
			IsMuted        bool   `json:"IsMuted"`
			RepeatMode     string `json:"RepeatMode"`
			SubtitleOffset int    `json:"SubtitleOffset"`
			PlaybackRate   int    `json:"PlaybackRate"`
		} `json:"PlayState"`
		RemoteEndPoint        string    `json:"RemoteEndPoint"`
		Protocol              string    `json:"Protocol"`
		PlaylistIndex         int       `json:"PlaylistIndex"`
		PlaylistLength        int       `json:"PlaylistLength"`
		ID                    string    `json:"Id"`
		ServerID              string    `json:"ServerId"`
		UserID                string    `json:"UserId"`
		UserName              string    `json:"UserName"`
		Client                string    `json:"Client"`
		LastActivityDate      time.Time `json:"LastActivityDate"`
		DeviceName            string    `json:"DeviceName"`
		InternalDeviceID      int       `json:"InternalDeviceId"`
		DeviceID              string    `json:"DeviceId"`
		ApplicationVersion    string    `json:"ApplicationVersion"`
		SupportsRemoteControl bool      `json:"SupportsRemoteControl"`
	} `json:"SessionInfo"`
	AccessToken string `json:"AccessToken"`
	ServerID    string `json:"ServerId"`
}

type LibraryResp struct {
	Items []struct {
		Name                  string        `json:"Name"`
		ServerID              string        `json:"ServerId"`
		ID                    string        `json:"Id"`
		GUID                  string        `json:"Guid"`
		Etag                  string        `json:"Etag"`
		DateCreated           time.Time     `json:"DateCreated"`
		CanDelete             bool          `json:"CanDelete"`
		CanDownload           bool          `json:"CanDownload"`
		PresentationUniqueKey string        `json:"PresentationUniqueKey"`
		SortName              string        `json:"SortName"`
		ForcedSortName        string        `json:"ForcedSortName"`
		ExternalUrls          []interface{} `json:"ExternalUrls"`
		Path                  string        `json:"Path"`
		Taglines              []interface{} `json:"Taglines"`
		FileName              string        `json:"FileName"`
		RemoteTrailers        []interface{} `json:"RemoteTrailers"`
		ProviderIds           struct {
		} `json:"ProviderIds"`
		IsFolder                bool          `json:"IsFolder"`
		ParentID                string        `json:"ParentId"`
		Type                    string        `json:"Type"`
		DisplayPreferencesID    string        `json:"DisplayPreferencesId"`
		CollectionType          string        `json:"CollectionType"`
		BackdropImageTags       []interface{} `json:"BackdropImageTags"`
		LockedFields            []interface{} `json:"LockedFields"`
		LockData                bool          `json:"LockData"`
		PrimaryImageAspectRatio float64       `json:"PrimaryImageAspectRatio,omitempty"`
		ImageTags               struct {
			Primary string `json:"Primary"`
		} `json:"ImageTags,omitempty"`
	} `json:"Items"`
	TotalRecordCount int `json:"TotalRecordCount"`
}

type ExternalUrl struct {
	Name string `json:"Name"`
	Url  string `json:"Url"`
}

type MediaSourceInfo struct {
	Protocol                   string `json:"Protocol"`
	Id                         string `json:"Id"`
	Path                       string `json:"Path"`
	EncoderPath                string `json:"EncoderPath"`
	EncoderProtocol            string `json:"EncoderProtocol"`
	Type                       string `json:"Type"`
	Container                  string `json:"Container"`
	Size                       int    `json:"Size,omitempty"`
	Name                       string `json:"Name"`
	IsRemote                   bool   `json:"IsRemote"`
	RunTimeTicks               int    `json:"RunTimeTicks,omitempty"`
	SupportsTranscoding        bool   `json:"SupportsTranscoding"`
	SupportsDirectStream       bool   `json:"SupportsDirectStream"`
	SupportsDirectPlay         bool   `json:"SupportsDirectPlay"`
	IsInfiniteStream           bool   `json:"IsInfiniteStream"`
	RequiresOpening            bool   `json:"RequiresOpening"`
	OpenToken                  string `json:"OpenToken"`
	RequiresClosing            bool   `json:"RequiresClosing"`
	LiveStreamId               string `json:"LiveStreamId"`
	BufferMs                   int    `json:"BufferMs,omitempty"`
	RequiresLooping            bool   `json:"RequiresLooping"`
	SupportsProbing            bool   `json:"SupportsProbing"`
	Video3DFormat              string `json:"Video3DFormat,omitempty"`
	MediaStreams               []MediaStream
	Formats                    []string
	Bitrate                    int               `json:"Bitrate,omitempty"`
	Timestamp                  string            `json:"Timestamp"`
	RequiredHttpHeaders        map[string]string `json:"RequiredHttpHeaders"`
	TranscodingUrl             string            `json:"TranscodingUrl"`
	TranscodingSubProtocol     string            `json:"TranscodingSubProtocol"`
	TranscodingContainer       string            `json:"TranscodingContainer"`
	AnalyzeDurationMs          int               `json:"AnalyzeDurationMs,omitempty"`
	ReadAtNativeFramerate      bool              `json:"ReadAtNativeFramerate"`
	DefaultAudioStreamIndex    int               `json:"DefaultAudioStreamIndex,omitempty"`
	DefaultSubtitleStreamIndex int               `json:"DefaultSubtitleStreamIndex,omitempty"`
}

type BaseItemPerson struct {
}

type NameLongIdPair struct {
}

type MediaUrl struct {
	Url  string `json:"Url"`
	Name string `json:"Name"`
}

type MediaStream struct {
	Codec                  string  `json:"codec"`
	CodecTag               string  `json:"codecTag"`
	Language               string  `json:"language"`
	ColorTransfer          string  `json:"colorTransfer"`
	ColorPrimaries         string  `json:"colorPrimaries"`
	ColorSpace             string  `json:"colorSpace"`
	Comment                string  `json:"comment"`
	TimeBase               string  `json:"timeBase"`
	CodecTimeBase          string  `json:"codecTimeBase"`
	Title                  string  `json:"title"`
	Extradata              string  `json:"extradata"`
	VideoRange             string  `json:"videoRange"`
	DisplayTitle           string  `json:"displayTitle"`
	DisplayLanguage        string  `json:"displayLanguage"`
	NalLengthSize          string  `json:"nalLengthSize"`
	IsInterlaced           bool    `json:"isInterlaced"`
	IsAVC                  bool    `json:"isAVC,omitempty"`
	ChannelLayout          string  `json:"channelLayout"`
	BitRate                int32   `json:"bitRate,omitempty"`
	BitDepth               int32   `json:"bitDepth,omitempty"`
	RefFrames              int32   `json:"refFrames,omitempty"`
	PacketLength           int32   `json:"packetLength,omitempty"`
	Channels               int32   `json:"channels,omitempty"`
	SampleRate             int32   `json:"sampleRate,omitempty"`
	IsDefault              bool    `json:"isDefault"`
	IsForced               bool    `json:"isForced"`
	Height                 int32   `json:"height,omitempty"`
	Width                  int32   `json:"width,omitempty"`
	AverageFrameRate       float32 `json:"averageFrameRate,omitempty"`
	RealFrameRate          float32 `json:"realFrameRate,omitempty"`
	Profile                string  `json:"profile"`
	Type                   string  `json:"type"`
	AspectRatio            string  `json:"aspectRatio"`
	Index                  int32   `json:"index"`
	Score                  int32   `json:"score,omitempty"`
	IsExternal             bool    `json:"isExternal"`
	DeliveryMethod         string  `json:"deliveryMethod"`
	DeliveryUrl            string  `json:"deliveryUrl"`
	IsExternalUrl          bool    `json:"isExternalUrl,omitempty"`
	IsTextSubtitleStream   bool    `json:"isTextSubtitleStream"`
	SupportsExternalStream bool    `json:"supportsExternalStream"`
	Path                   string  `json:"path"`
	PixelFormat            string  `json:"pixelFormat"`
	Level                  float64 `json:"level,omitempty"`
	IsAnamorphic           bool    `json:"isAnamorphic,omitempty"`
}

type UserItemDataDto struct {
	Rating                float64 `json:"Rating,omitempty"`
	PlayedPercentage      float64 `json:"PlayedPercentage,omitempty"`
	UnplayedItemCount     int     `json:"UnplayedItemCount,omitempty"`
	PlaybackPositionTicks int     `json:"PlaybackPositionTicks"`
	PlayCount             int     `json:"PlayCount"`
	IsFavorite            bool    `json:"IsFavorite"`
	Likes                 bool    `json:"Likes,omitempty"`
	LastPlayedDate        string  `json:"LastPlayedDate,omitempty"`
	Played                bool    `json:"Played"`
	Key                   string  `json:"Key"`
	ItemId                string  `json:"ItemId"`
}

type NameIdPair struct {
}

type ChapterInfo struct {
	StartPositionTicks int    `json:"StartPositionTicks"`
	Name               string `json:"Name"`
	ImageTag           string `json:"ImageTag"`
}

type MediaItem struct {
	Name                         string `json:"Name"`
	OriginalTitle                string `json:"OriginalTitle"`
	ServerId                     string `json:"ServerId"`
	Id                           string `json:"Id"`
	Etag                         string `json:"Etag"`
	PlaylistItemId               string `json:"PlaylistItemId"`
	DateCreated                  string `json:"DateCreated,omitempty"`
	ExtraType                    string `json:"ExtraType"`
	AirsBeforeSeasonNumber       int    `json:"AirsBeforeSeasonNumber,omitempty"`
	AirsAfterSeasonNumber        int    `json:"AirsAfterSeasonNumber,omitempty"`
	AirsBeforeEpisodeNumber      int    `json:"AirsBeforeEpisodeNumber,omitempty"`
	DisplaySpecialsWithSeasons   bool   `json:"DisplaySpecialsWithSeasons,omitempty"`
	CanDelete                    bool   `json:"CanDelete,omitempty"`
	CanDownload                  bool   `json:"CanDownload,omitempty"`
	HasSubtitles                 bool   `json:"HasSubtitles,omitempty"`
	SupportsResume               bool   `json:"SupportsResume,omitempty"`
	PreferredMetadataLanguage    string `json:"PreferredMetadataLanguage"`
	PreferredMetadataCountryCode string `json:"PreferredMetadataCountryCode"`
	SupportsSync                 bool   `json:"SupportsSync,omitempty"`
	Container                    string `json:"Container"`
	SortName                     string `json:"SortName"`
	ForcedSortName               string `json:"ForcedSortName"`
	Video3DFormat                string `json:"Video3DFormat,omitempty"`
	PremiereDate                 string `json:"PremiereDate,omitempty"`
	ExternalUrls                 []ExternalUrl
	MediaSources                 []MediaSourceInfo
	CriticRating                 float64 `json:"CriticRating,omitempty"`
	GameSystemId                 int     `json:"GameSystemId,omitempty"`
	GameSystem                   string  `json:"GameSystem"`
	ProductionLocations          []string
	Path                         string `json:"Path"`
	OfficialRating               string `json:"OfficialRating"`
	CustomRating                 string `json:"CustomRating"`
	ChannelId                    string `json:"ChannelId"`
	ChannelName                  string `json:"ChannelName"`
	Overview                     string `json:"Overview"`
	Taglines                     []string
	Genres                       []string
	CommunityRating              float64 `json:"CommunityRating,omitempty"`
	RunTimeTicks                 int     `json:"RunTimeTicks,omitempty"`
	PlayAccess                   string  `json:"PlayAccess"`
	AspectRatio                  string  `json:"AspectRatio"`
	ProductionYear               int     `json:"ProductionYear,omitempty"`
	Number                       string  `json:"Number"`
	ChannelNumber                string  `json:"ChannelNumber"`
	IndexNumber                  int     `json:"IndexNumber,omitempty"`
	IndexNumberEnd               int     `json:"IndexNumberEnd,omitempty"`
	ParentIndexNumber            int     `json:"ParentIndexNumber,omitempty"`
	RemoteTrailers               []MediaUrl
	ProviderIds                  map[string]string `json:"ProviderIds"`
	IsFolder                     bool              `json:"IsFolder,omitempty"`
	ParentId                     string            `json:"ParentId"`
	Type                         string            `json:"Type"`
	People                       []BaseItemPerson
	Studios                      []NameLongIdPair
	GenreItems                   []NameLongIdPair
	ParentLogoItemId             string `json:"ParentLogoItemId"`
	ParentBackdropItemId         string `json:"ParentBackdropItemId"`
	ParentBackdropImageTags      []string
	LocalTrailerCount            int `json:"LocalTrailerCount,omitempty"`
	UserData                     UserItemDataDto
	RecursiveItemCount           int    `json:"RecursiveItemCount,omitempty"`
	ChildCount                   int    `json:"ChildCount,omitempty"`
	SeriesName                   string `json:"SeriesName"`
	SeriesId                     string `json:"SeriesId"`
	SeasonId                     string `json:"SeasonId"`
	SpecialFeatureCount          int    `json:"SpecialFeatureCount,omitempty"`
	DisplayPreferencesId         string `json:"DisplayPreferencesId"`
	Status                       string `json:"Status"`
	AirTime                      string `json:"AirTime"`
	AirDays                      []string
	Tags                         []string
	PrimaryImageAspectRatio      float64 `json:"PrimaryImageAspectRatio,omitempty"`
	Artists                      []string
	ArtistItems                  []NameIdPair
	Album                        string `json:"Album"`
	CollectionType               string `json:"CollectionType"`
	DisplayOrder                 string `json:"DisplayOrder"`
	AlbumId                      string `json:"AlbumId"`
	AlbumPrimaryImageTag         string `json:"AlbumPrimaryImageTag"`
	SeriesPrimaryImageTag        string `json:"SeriesPrimaryImageTag"`
	AlbumArtist                  string `json:"AlbumArtist"`
	AlbumArtists                 []NameIdPair
	SeasonName                   string `json:"SeasonName"`
	MediaStreams                 []MediaStream
	PartCount                    int               `json:"PartCount,omitempty"`
	ImageTags                    map[string]string `json:"ImageTags"`
	BackdropImageTags            []string
	ParentLogoImageTag           string `json:"ParentLogoImageTag"`
	ParentArtItemId              string `json:"ParentArtItemId"`
	ParentArtImageTag            string `json:"ParentArtImageTag"`
	SeriesThumbImageTag          string `json:"SeriesThumbImageTag"`
	SeriesStudio                 string `json:"SeriesStudio"`
	ParentThumbItemId            string `json:"ParentThumbItemId"`
	ParentThumbImageTag          string `json:"ParentThumbImageTag"`
	ParentPrimaryImageItemId     string `json:"ParentPrimaryImageItemId"`
	ParentPrimaryImageTag        string `json:"ParentPrimaryImageTag"`
	Chapters                     []ChapterInfo
	LocationType                 string `json:"LocationType"`
	MediaType                    string `json:"MediaType"`
	EndDate                      string `json:"EndDate,omitempty"`
	LockedFields                 []string
	LockData                     bool    `json:"LockData,omitempty"`
	Width                        int     `json:"Width,omitempty"`
	Height                       int     `json:"Height,omitempty"`
	CameraMake                   string  `json:"CameraMake"`
	CameraModel                  string  `json:"CameraModel"`
	Software                     string  `json:"Software"`
	ExposureTime                 float64 `json:"ExposureTime,omitempty"`
	FocalLength                  float64 `json:"FocalLength,omitempty"`
	ImageOrientation             string  `json:"ImageOrientation"`
	Aperture                     float64 `json:"Aperture,omitempty"`
	ShutterSpeed                 float64 `json:"ShutterSpeed,omitempty"`
	Latitude                     float64 `json:"Latitude,omitempty"`
	Longitude                    float64 `json:"Longitude,omitempty"`
	Altitude                     float64 `json:"Altitude,omitempty"`
	IsoSpeedRating               int     `json:"IsoSpeedRating,omitempty"`
	SeriesTimerId                string  `json:"SeriesTimerId"`
	ChannelPrimaryImageTag       string  `json:"ChannelPrimaryImageTag"`
	StartDate                    string  `json:"StartDate,omitempty"`
	CompletionPercentage         float64 `json:"CompletionPercentage,omitempty"`
	IsRepeat                     bool    `json:"IsRepeat,omitempty"`
	IsNew                        bool    `json:"IsNew,omitempty"`
	EpisodeTitle                 string  `json:"EpisodeTitle"`
	IsMovie                      bool    `json:"IsMovie,omitempty"`
	IsSports                     bool    `json:"IsSports,omitempty"`
	IsSeries                     bool    `json:"IsSeries,omitempty"`
	IsLive                       bool    `json:"IsLive,omitempty"`
	IsNews                       bool    `json:"IsNews,omitempty"`
	IsKids                       bool    `json:"IsKids,omitempty"`
	IsPremiere                   bool    `json:"IsPremiere,omitempty"`
	TimerId                      string  `json:"TimerId"`
	CurrentProgram               map[string]interface{}
	MovieCount                   int `json:"MovieCount,omitempty"`
	SeriesCount                  int `json:"SeriesCount,omitempty"`
	AlbumCount                   int `json:"AlbumCount,omitempty"`
	SongCount                    int `json:"SongCount,omitempty"`
	MusicVideoCount              int `json:"MusicVideoCount,omitempty"`
}

type GetItemResp struct {
	Items            []MediaItem `json:"Items"`
	TotalRecordCount uint64      `json:"TotalRecordCount"`
}

type MeResp struct {
	Name                      string    `json:"Name"`
	ServerID                  string    `json:"ServerId"`
	Prefix                    string    `json:"Prefix"`
	DateCreated               time.Time `json:"DateCreated"`
	ID                        string    `json:"Id"`
	HasPassword               bool      `json:"HasPassword"`
	HasConfiguredPassword     bool      `json:"HasConfiguredPassword"`
	HasConfiguredEasyPassword bool      `json:"HasConfiguredEasyPassword"`
	LastLoginDate             time.Time `json:"LastLoginDate"`
	LastActivityDate          time.Time `json:"LastActivityDate"`
	Policy                    struct {
		IsAdministrator                 bool   `json:"IsAdministrator"`
		IsHidden                        bool   `json:"IsHidden"`
		IsHiddenRemotely                bool   `json:"IsHiddenRemotely"`
		IsHiddenFromUnusedDevices       bool   `json:"IsHiddenFromUnusedDevices"`
		IsDisabled                      bool   `json:"IsDisabled"`
		IsTagBlockingModeInclusive      bool   `json:"IsTagBlockingModeInclusive"`
		EnableUserPreferenceAccess      bool   `json:"EnableUserPreferenceAccess"`
		EnableRemoteControlOfOtherUsers bool   `json:"EnableRemoteControlOfOtherUsers"`
		EnableSharedDeviceControl       bool   `json:"EnableSharedDeviceControl"`
		EnableRemoteAccess              bool   `json:"EnableRemoteAccess"`
		EnableLiveTvManagement          bool   `json:"EnableLiveTvManagement"`
		EnableLiveTvAccess              bool   `json:"EnableLiveTvAccess"`
		EnableMediaPlayback             bool   `json:"EnableMediaPlayback"`
		EnableAudioPlaybackTranscoding  bool   `json:"EnableAudioPlaybackTranscoding"`
		EnableVideoPlaybackTranscoding  bool   `json:"EnableVideoPlaybackTranscoding"`
		EnablePlaybackRemuxing          bool   `json:"EnablePlaybackRemuxing"`
		EnableContentDeletion           bool   `json:"EnableContentDeletion"`
		EnableContentDownloading        bool   `json:"EnableContentDownloading"`
		EnableSubtitleDownloading       bool   `json:"EnableSubtitleDownloading"`
		EnableSubtitleManagement        bool   `json:"EnableSubtitleManagement"`
		EnableSyncTranscoding           bool   `json:"EnableSyncTranscoding"`
		EnableMediaConversion           bool   `json:"EnableMediaConversion"`
		EnableAllChannels               bool   `json:"EnableAllChannels"`
		EnableAllFolders                bool   `json:"EnableAllFolders"`
		InvalidLoginAttemptCount        int    `json:"InvalidLoginAttemptCount"`
		EnablePublicSharing             bool   `json:"EnablePublicSharing"`
		RemoteClientBitrateLimit        int    `json:"RemoteClientBitrateLimit"`
		AuthenticationProviderID        string `json:"AuthenticationProviderId"`
		SimultaneousStreamLimit         int    `json:"SimultaneousStreamLimit"`
		EnableAllDevices                bool   `json:"EnableAllDevices"`
	} `json:"Policy"`
	Configuration struct {
		PlayDefaultAudioTrack      bool     `json:"PlayDefaultAudioTrack"`
		DisplayMissingEpisodes     bool     `json:"DisplayMissingEpisodes"`
		SubtitleMode               string   `json:"SubtitleMode"`
		EnableLocalPassword        bool     `json:"EnableLocalPassword"`
		OrderedViews               []string `json:"OrderedViews"`
		LatestItemsExcludes        []string `json:"LatestItemsExcludes"`
		MyMediaExcludes            []string `json:"MyMediaExcludes"`
		HidePlayedInLatest         bool     `json:"HidePlayedInLatest"`
		RememberAudioSelections    bool     `json:"RememberAudioSelections"`
		RememberSubtitleSelections bool     `json:"RememberSubtitleSelections"`
		EnableNextEpisodeAutoPlay  bool     `json:"EnableNextEpisodeAutoPlay"`
		ResumeRewindSeconds        int      `json:"ResumeRewindSeconds"`
		IntroSkipMode              string   `json:"IntroSkipMode"`
	} `json:"Configuration,omitempty"`
}
