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
		Name                  string    `json:"Name"`
		ServerID              string    `json:"ServerId"`
		ID                    string    `json:"Id"`
		GUID                  string    `json:"Guid"`
		Etag                  string    `json:"Etag"`
		DateCreated           time.Time `json:"DateCreated"`
		CanDelete             bool      `json:"CanDelete"`
		CanDownload           bool      `json:"CanDownload"`
		PresentationUniqueKey string    `json:"PresentationUniqueKey"`
		SortName              string    `json:"SortName"`
		ForcedSortName        string    `json:"ForcedSortName"`
		Path                  string    `json:"Path"`
		FileName              string    `json:"FileName"`
		ProviderIds           struct {
		} `json:"ProviderIds"`
		IsFolder                bool    `json:"IsFolder"`
		ParentID                string  `json:"ParentId"`
		Type                    string  `json:"Type"`
		DisplayPreferencesID    string  `json:"DisplayPreferencesId"`
		CollectionType          string  `json:"CollectionType"`
		LockData                bool    `json:"LockData"`
		PrimaryImageAspectRatio float64 `json:"PrimaryImageAspectRatio,omitempty"`
		ImageTags               struct {
			Primary string `json:"Primary"`
		} `json:"ImageTags,omitempty"`
	} `json:"Items"`
	TotalRecordCount int `json:"TotalRecordCount"`
}
type ExternalUrls struct {
	Name string `json:"Name"`
	URL  string `json:"Url"`
}
type MediaStreams struct {
	Codec                  string  `json:"Codec"`
	CodecTag               string  `json:"CodecTag"`
	Language               string  `json:"Language"`
	ColorTransfer          string  `json:"ColorTransfer"`
	ColorPrimaries         string  `json:"ColorPrimaries"`
	ColorSpace             string  `json:"ColorSpace"`
	Comment                string  `json:"Comment"`
	StreamStartTimeTicks   int     `json:"StreamStartTimeTicks"`
	TimeBase               string  `json:"TimeBase"`
	CodecTimeBase          string  `json:"CodecTimeBase"`
	Title                  string  `json:"Title"`
	Extradata              string  `json:"Extradata"`
	VideoRange             string  `json:"VideoRange"`
	DisplayTitle           string  `json:"DisplayTitle"`
	DisplayLanguage        string  `json:"DisplayLanguage"`
	NalLengthSize          string  `json:"NalLengthSize"`
	IsInterlaced           bool    `json:"IsInterlaced"`
	IsAVC                  bool    `json:"IsAVC"`
	ChannelLayout          string  `json:"ChannelLayout"`
	BitRate                int     `json:"BitRate"`
	BitDepth               int     `json:"BitDepth"`
	RefFrames              int     `json:"RefFrames"`
	Rotation               int     `json:"Rotation"`
	Channels               int     `json:"Channels"`
	SampleRate             int     `json:"SampleRate"`
	IsDefault              bool    `json:"IsDefault"`
	IsForced               bool    `json:"IsForced"`
	Height                 int     `json:"Height"`
	Width                  int     `json:"Width"`
	AverageFrameRate       float64 `json:"AverageFrameRate"`
	RealFrameRate          float64 `json:"RealFrameRate"`
	Profile                string  `json:"Profile"`
	Type                   string  `json:"Type"`
	AspectRatio            string  `json:"AspectRatio"`
	Index                  uint64  `json:"Index"`
	IsExternal             bool    `json:"IsExternal"`
	DeliveryMethod         string  `json:"DeliveryMethod"`
	DeliveryURL            string  `json:"DeliveryUrl"`
	IsExternalURL          bool    `json:"IsExternalUrl"`
	IsTextSubtitleStream   bool    `json:"IsTextSubtitleStream"`
	SupportsExternalStream bool    `json:"SupportsExternalStream"`
	Path                   string  `json:"Path"`
	Protocol               string  `json:"Protocol"`
	PixelFormat            string  `json:"PixelFormat"`
	Level                  int     `json:"Level"`
	IsAnamorphic           bool    `json:"IsAnamorphic"`
	ItemID                 string  `json:"ItemId"`
	ServerID               string  `json:"ServerId"`
	AttachmentSize         int     `json:"AttachmentSize"`
	MimeType               string  `json:"MimeType"`
	IsClosedCaptions       bool    `json:"IsClosedCaptions"`
	SubtitleLocationType   string  `json:"SubtitleLocationType"`
}

type RequiredHTTPHeaders struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
}

type MediaSources struct {
	Protocol                   string              `json:"Protocol"`
	ID                         string              `json:"Id"`
	Path                       string              `json:"Path"`
	EncoderPath                string              `json:"EncoderPath"`
	EncoderProtocol            string              `json:"EncoderProtocol"`
	Type                       string              `json:"Type"`
	Container                  string              `json:"Container"`
	Size                       int                 `json:"Size"`
	Name                       string              `json:"Name"`
	SortName                   string              `json:"SortName"`
	IsRemote                   bool                `json:"IsRemote"`
	RunTimeTicks               int                 `json:"RunTimeTicks"`
	ContainerStartTimeTicks    int                 `json:"ContainerStartTimeTicks"`
	SupportsTranscoding        bool                `json:"SupportsTranscoding"`
	SupportsDirectStream       bool                `json:"SupportsDirectStream"`
	SupportsDirectPlay         bool                `json:"SupportsDirectPlay"`
	IsInfiniteStream           bool                `json:"IsInfiniteStream"`
	RequiresOpening            bool                `json:"RequiresOpening"`
	OpenToken                  string              `json:"OpenToken"`
	RequiresClosing            bool                `json:"RequiresClosing"`
	LiveStreamID               string              `json:"LiveStreamId"`
	BufferMs                   int                 `json:"BufferMs"`
	RequiresLooping            bool                `json:"RequiresLooping"`
	SupportsProbing            bool                `json:"SupportsProbing"`
	Video3DFormat              string              `json:"Video3DFormat"`
	MediaStreams               []MediaStreams      `json:"MediaStreams"`
	Formats                    []string            `json:"Formats"`
	Bitrate                    int                 `json:"Bitrate"`
	Timestamp                  string              `json:"Timestamp"`
	RequiredHTTPHeaders        RequiredHTTPHeaders `json:"RequiredHttpHeaders"`
	DirectStreamURL            string              `json:"DirectStreamUrl"`
	TranscodingURL             string              `json:"TranscodingUrl"`
	TranscodingSubProtocol     string              `json:"TranscodingSubProtocol"`
	TranscodingContainer       string              `json:"TranscodingContainer"`
	AnalyzeDurationMs          int                 `json:"AnalyzeDurationMs"`
	ReadAtNativeFramerate      bool                `json:"ReadAtNativeFramerate"`
	DefaultAudioStreamIndex    int64               `json:"DefaultAudioStreamIndex"`
	DefaultSubtitleStreamIndex int64               `json:"DefaultSubtitleStreamIndex"`
	ItemID                     string              `json:"ItemId"`
	ServerID                   string              `json:"ServerId"`
}

type RemoteTrailers struct {
	URL  string `json:"Url"`
	Name string `json:"Name"`
}

type People struct {
	Name            string `json:"Name"`
	ID              string `json:"Id"`
	Role            string `json:"Role"`
	Type            string `json:"Type"`
	PrimaryImageTag string `json:"PrimaryImageTag"`
}

type Studios struct {
	Name string `json:"Name"`
	ID   int    `json:"Id"`
}

type GenreItems struct {
	Name string `json:"Name"`
	ID   int    `json:"Id"`
}

type TagItems struct {
	Name string `json:"Name"`
	ID   int    `json:"Id"`
}

type UserData struct {
	Rating                int       `json:"Rating"`
	PlayedPercentage      float32   `json:"PlayedPercentage"`
	UnplayedItemCount     int       `json:"UnplayedItemCount"`
	PlaybackPositionTicks int       `json:"PlaybackPositionTicks"`
	PlayCount             int       `json:"PlayCount"`
	IsFavorite            bool      `json:"IsFavorite"`
	LastPlayedDate        time.Time `json:"LastPlayedDate"`
	Played                bool      `json:"Played"`
	Key                   string    `json:"Key"`
	ItemID                string    `json:"ItemId"`
	ServerID              string    `json:"ServerId"`
}

type ArtistItems struct {
	Name string `json:"Name"`
	ID   string `json:"Id"`
}

type Composers struct {
	Name string `json:"Name"`
	ID   string `json:"Id"`
}

type AlbumArtists struct {
	Name string `json:"Name"`
	ID   string `json:"Id"`
}

type ImageTags struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
}

type Chapters struct {
	StartPositionTicks int    `json:"StartPositionTicks"`
	Name               string `json:"Name"`
	ImageTag           string `json:"ImageTag"`
	MarkerType         string `json:"MarkerType"`
	ChapterIndex       int    `json:"ChapterIndex"`
}

type Items struct {
	Name                 string           `json:"Name"`
	OriginalTitle        string           `json:"OriginalTitle"`
	ServerID             string           `json:"ServerId"`
	ID                   string           `json:"Id"`
	GUID                 string           `json:"Guid"`
	Etag                 string           `json:"Etag"`
	Prefix               string           `json:"Prefix"`
	PlaylistItemID       string           `json:"PlaylistItemId"`
	DateCreated          time.Time        `json:"DateCreated"`
	ExtraType            string           `json:"ExtraType"`
	SortIndexNumber      int              `json:"SortIndexNumber"`
	Container            string           `json:"Container"`
	SortName             string           `json:"SortName"`
	ForcedSortName       string           `json:"ForcedSortName"`
	Video3DFormat        string           `json:"Video3DFormat"`
	PremiereDate         time.Time        `json:"PremiereDate"`
	ExternalUrls         []ExternalUrls   `json:"ExternalUrls"`
	MediaSources         []MediaSources   `json:"MediaSources"`
	CriticRating         int              `json:"CriticRating"`
	GameSystemID         int              `json:"GameSystemId"`
	AsSeries             bool             `json:"AsSeries"`
	GameSystem           string           `json:"GameSystem"`
	ProductionLocations  []string         `json:"ProductionLocations"`
	Path                 string           `json:"Path"`
	OfficialRating       string           `json:"OfficialRating"`
	CustomRating         string           `json:"CustomRating"`
	ChannelID            string           `json:"ChannelId"`
	ChannelName          string           `json:"ChannelName"`
	Overview             string           `json:"Overview"`
	Taglines             []string         `json:"Taglines"`
	Genres               []string         `json:"Genres"`
	RunTimeTicks         int              `json:"RunTimeTicks"`
	Size                 int              `json:"Size"`
	FileName             string           `json:"FileName"`
	Bitrate              int              `json:"Bitrate"`
	PlayAccess           string           `json:"PlayAccess"`
	ProductionYear       int              `json:"ProductionYear"`
	Number               string           `json:"Number"`
	ChannelNumber        string           `json:"ChannelNumber"`
	IndexNumber          int              `json:"IndexNumber"`
	IndexNumberEnd       int              `json:"IndexNumberEnd"`
	ParentIndexNumber    int              `json:"ParentIndexNumber"`
	RemoteTrailers       []RemoteTrailers `json:"RemoteTrailers"`
	IsFolder             bool             `json:"IsFolder"`
	ParentID             string           `json:"ParentId"`
	Type                 string           `json:"Type"`
	People               []People         `json:"People"`
	Studios              []Studios        `json:"Studios"`
	GenreItems           []GenreItems     `json:"GenreItems"`
	TagItems             []TagItems       `json:"TagItems"`
	LocalTrailerCount    int              `json:"LocalTrailerCount"`
	UserData             UserData         `json:"UserData"`
	RecursiveItemCount   int              `json:"RecursiveItemCount"`
	ChildCount           int              `json:"ChildCount"`
	SeriesName           string           `json:"SeriesName"`
	SeriesID             string           `json:"SeriesId"`
	SeasonID             string           `json:"SeasonId"`
	SpecialFeatureCount  int              `json:"SpecialFeatureCount"`
	DisplayPreferencesID string           `json:"DisplayPreferencesId"`
	Status               string           `json:"Status"`
	AirDays              []string         `json:"AirDays"`
	Tags                 []string         `json:"Tags"`
	Artists              []string         `json:"Artists"`
	ArtistItems          []ArtistItems    `json:"ArtistItems"`
	Composers            []Composers      `json:"Composers"`
	Album                string           `json:"Album"`
	CollectionType       string           `json:"CollectionType"`
	DisplayOrder         string           `json:"DisplayOrder"`
	AlbumArtist          string           `json:"AlbumArtist"`
	AlbumArtists         []AlbumArtists   `json:"AlbumArtists"`
	SeasonName           string           `json:"SeasonName"`
	MediaStreams         []MediaStreams   `json:"MediaStreams"`
	PartCount            int              `json:"PartCount"`
	ImageTags            ImageTags        `json:"ImageTags"`
	BackdropImageTags    []string         `json:"BackdropImageTags"`
	ParentLogoImageTag   string           `json:"ParentLogoImageTag"`
	SeriesStudio         string           `json:"SeriesStudio"`
	ParentThumbItemID    string           `json:"ParentThumbItemId"`
	ParentThumbImageTag  string           `json:"ParentThumbImageTag"`
	Chapters             []Chapters       `json:"Chapters"`
	LocationType         string           `json:"LocationType"`
	MediaType            string           `json:"MediaType"`
	Width                int              `json:"Width"`
	Height               int              `json:"Height"`
	CameraMake           string           `json:"CameraMake"`
	CameraModel          string           `json:"CameraModel"`
	ExposureTime         int              `json:"ExposureTime"`
	ImageOrientation     string           `json:"ImageOrientation"`
	CompletionPercentage int              `json:"CompletionPercentage"`
	EpisodeTitle         string           `json:"EpisodeTitle"`
	IsMovie              bool             `json:"IsMovie"`
	Disabled             bool             `json:"Disabled"`
	MusicVideoCount      int              `json:"MusicVideoCount"`
	Subviews             []string         `json:"Subviews"`
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

type SystemInfoResp struct {
	SystemUpdateLevel                    string             `json:"SystemUpdateLevel"`
	OperatingSystemDisplayName           string             `json:"OperatingSystemDisplayName"`
	PackageName                          string             `json:"PackageName"`
	HasPendingRestart                    bool               `json:"HasPendingRestart"`
	IsShuttingDown                       bool               `json:"IsShuttingDown"`
	SupportsLibraryMonitor               bool               `json:"SupportsLibraryMonitor"`
	WebSocketPortNumber                  int32              `json:"WebSocketPortNumber"`
	CompletedInstallations               []InstallationInfo `json:"CompletedInstallations"`
	CanSelfRestart                       bool               `json:"CanSelfRestart"`
	CanSelfUpdate                        bool               `json:"CanSelfUpdate"`
	CanLaunchWebBrowser                  bool               `json:"CanLaunchWebBrowser"`
	ProgramDataPath                      string             `json:"ProgramDataPath"`
	ItemsByNamePath                      string             `json:"ItemsByNamePath"`
	CachePath                            string             `json:"CachePath"`
	LogPath                              string             `json:"LogPath"`
	InternalMetadataPath                 string             `json:"InternalMetadataPath"`
	TranscodingTempPath                  string             `json:"TranscodingTempPath"`
	HttpServerPortNumber                 int32              `json:"HttpServerPortNumber"`
	SupportsHttps                        bool               `json:"SupportsHttps"`
	HttpsPortNumber                      int32              `json:"HttpsPortNumber"`
	HasUpdateAvailable                   bool               `json:"HasUpdateAvailable"`
	SupportsAutoRunAtStartup             bool               `json:"SupportsAutoRunAtStartup"`
	HardwareAccelerationRequiresPremiere bool               `json:"HardwareAccelerationRequiresPremiere"`
	LocalAddress                         string             `json:"LocalAddress"`
	WanAddress                           string             `json:"WanAddress"`
	ServerName                           string             `json:"ServerName"`
	Version                              string             `json:"Version"`
	OperatingSystem                      string             `json:"OperatingSystem"`
	Id                                   string             `json:"Id"`
}

type InstallationInfo struct {
	Id              string  `json:"Id"`
	Name            string  `json:"Name"`
	AssemblyGuid    string  `json:"AssemblyGuid"`
	Version         string  `json:"Version"`
	UpdateClass     string  `json:"UpdateClass"`
	PercentComplete float64 `json:"PercentComplete,omitempty"`
}

type ItemsResp struct {
	Items            []Items `json:"Items"`
	TotalRecordCount uint64  `json:"TotalRecordCount"`
}

type PlayBackResp struct {
	MediaSources  []MediaSources `json:"MediaSources"`
	PlaySessionID string         `json:"PlaySessionId"`
}

type DirectPlayProfile struct {
	Container  string `json:"Container"`
	Type       string `json:"Type"`
	VideoCodec string `json:"VideoCodec,omitempty"`
	AudioCodec string `json:"AudioCodec,omitempty"`
}

type TranscodingProfile struct {
	Container           string `json:"Container"`
	Type                string `json:"Type"`
	AudioCodec          string `json:"AudioCodec"`
	Context             string `json:"Context"`
	Protocol            string `json:"Protocol,omitempty"`
	MaxAudioChannels    string `json:"MaxAudioChannels,omitempty"`
	MinSegments         string `json:"MinSegments,omitempty"`
	BreakOnNonKeyFrames bool   `json:"BreakOnNonKeyFrames,omitempty"`
	VideoCodec          string `json:"VideoCodec,omitempty"`
	CopyTimestamps      bool   `json:"CopyTimestamps,omitempty"`
	ManifestSubtitles   string `json:"ManifestSubtitles,omitempty"`
}

type Condition struct {
	Condition  string `json:"Condition"`
	Property   string `json:"Property"`
	Value      string `json:"Value"`
	IsRequired any    `json:"IsRequired"`
}

type CodecProfile struct {
	Type       string      `json:"Type"`
	Codec      string      `json:"Codec,omitempty"`
	Conditions []Condition `json:"Conditions"`
}

type SubtitleProfile struct {
	Format   string `json:"Format"`
	Method   string `json:"Method"`
	Protocol string `json:"Protocol,omitempty"`
}

type ResponseProfile struct {
	Type      string `json:"Type"`
	Container string `json:"Container"`
	MimeType  string `json:"MimeType"`
}

type DeviceProfile struct {
	MaxStaticBitrate                 int                  `json:"MaxStaticBitrate"`
	MaxStreamingBitrate              int                  `json:"MaxStreamingBitrate"`
	MusicStreamingTranscodingBitrate int                  `json:"MusicStreamingTranscodingBitrate"`
	DirectPlayProfiles               []DirectPlayProfile  `json:"DirectPlayProfiles"`
	TranscodingProfiles              []TranscodingProfile `json:"TranscodingProfiles"`
	CodecProfiles                    []CodecProfile       `json:"CodecProfiles"`
	SubtitleProfiles                 []SubtitleProfile    `json:"SubtitleProfiles"`
	ResponseProfiles                 []ResponseProfile    `json:"ResponseProfiles"`
}

type PlayBackReq struct {
	DeviceProfile                  DeviceProfile `json:"DeviceProfile"`
	ID                             string        `json:"Id,omitempty"`
	UserID                         string        `json:"UserId,omitempty"`
	MaxStreamingBitrate            int           `json:"MaxStreamingBitrate,omitempty"`
	StartTimeTicks                 int           `json:"StartTimeTicks,omitempty"`
	AudioStreamIndex               int           `json:"AudioStreamIndex,omitempty"`
	SubtitleStreamIndex            int           `json:"SubtitleStreamIndex,omitempty"`
	MaxAudioChannels               int           `json:"MaxAudioChannels,omitempty"`
	MediaSourceID                  string        `json:"MediaSourceId,omitempty"`
	LiveStreamID                   string        `json:"LiveStreamId,omitempty"`
	EnableDirectPlay               bool          `json:"EnableDirectPlay,omitempty"`
	EnableDirectStream             bool          `json:"EnableDirectStream,omitempty"`
	EnableTranscoding              bool          `json:"EnableTranscoding,omitempty"`
	AllowInterlacedVideoStreamCopy bool          `json:"AllowInterlacedVideoStreamCopy,omitempty"`
	AllowVideoStreamCopy           bool          `json:"AllowVideoStreamCopy,omitempty"`
	AllowAudioStreamCopy           bool          `json:"AllowAudioStreamCopy,omitempty"`
	IsPlayback                     bool          `json:"IsPlayback,omitempty"`
	AutoOpenLiveStream             bool          `json:"AutoOpenLiveStream,omitempty"`
	DirectPlayProtocols            []string      `json:"DirectPlayProtocols,omitempty"`
	CurrentPlaySessionID           string        `json:"CurrentPlaySessionId,omitempty"`
}
