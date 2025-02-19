package emby

import "time"

type LoginReq struct {
	Username string `json:"Username"`
	Password string `json:"Pw"`
}

type LoginResp struct {
	AccessToken string `json:"AccessToken"`
	ServerID    string `json:"ServerId"`
	SessionInfo struct {
		LastActivityDate   time.Time `json:"LastActivityDate"`
		DeviceID           string    `json:"DeviceId"`
		RemoteEndPoint     string    `json:"RemoteEndPoint"`
		Protocol           string    `json:"Protocol"`
		ApplicationVersion string    `json:"ApplicationVersion"`
		ID                 string    `json:"Id"`
		ServerID           string    `json:"ServerId"`
		UserID             string    `json:"UserId"`
		UserName           string    `json:"UserName"`
		Client             string    `json:"Client"`
		DeviceName         string    `json:"DeviceName"`
		PlayState          struct {
			RepeatMode     string `json:"RepeatMode"`
			SubtitleOffset int    `json:"SubtitleOffset"`
			PlaybackRate   int    `json:"PlaybackRate"`
			CanSeek        bool   `json:"CanSeek"`
			IsPaused       bool   `json:"IsPaused"`
			IsMuted        bool   `json:"IsMuted"`
		} `json:"PlayState"`
		PlaylistLength        int  `json:"PlaylistLength"`
		InternalDeviceID      int  `json:"InternalDeviceId"`
		PlaylistIndex         int  `json:"PlaylistIndex"`
		SupportsRemoteControl bool `json:"SupportsRemoteControl"`
	} `json:"SessionInfo"`
	User struct {
		DateCreated      time.Time `json:"DateCreated"`
		LastLoginDate    time.Time `json:"LastLoginDate"`
		LastActivityDate time.Time `json:"LastActivityDate"`
		Configuration    struct {
			AudioLanguagePreference    string `json:"AudioLanguagePreference"`
			SubtitleMode               string `json:"SubtitleMode"`
			IntroSkipMode              string `json:"IntroSkipMode"`
			ResumeRewindSeconds        int    `json:"ResumeRewindSeconds"`
			PlayDefaultAudioTrack      bool   `json:"PlayDefaultAudioTrack"`
			DisplayMissingEpisodes     bool   `json:"DisplayMissingEpisodes"`
			EnableLocalPassword        bool   `json:"EnableLocalPassword"`
			HidePlayedInLatest         bool   `json:"HidePlayedInLatest"`
			RememberAudioSelections    bool   `json:"RememberAudioSelections"`
			RememberSubtitleSelections bool   `json:"RememberSubtitleSelections"`
			EnableNextEpisodeAutoPlay  bool   `json:"EnableNextEpisodeAutoPlay"`
		} `json:"Configuration"`
		Name     string `json:"Name"`
		ServerID string `json:"ServerId"`
		Prefix   string `json:"Prefix"`
		ID       string `json:"Id"`
		Policy   struct {
			AuthenticationProviderID        string `json:"AuthenticationProviderId"`
			InvalidLoginAttemptCount        int    `json:"InvalidLoginAttemptCount"`
			SimultaneousStreamLimit         int    `json:"SimultaneousStreamLimit"`
			RemoteClientBitrateLimit        int    `json:"RemoteClientBitrateLimit"`
			EnableVideoPlaybackTranscoding  bool   `json:"EnableVideoPlaybackTranscoding"`
			EnableContentDownloading        bool   `json:"EnableContentDownloading"`
			EnableUserPreferenceAccess      bool   `json:"EnableUserPreferenceAccess"`
			EnableRemoteControlOfOtherUsers bool   `json:"EnableRemoteControlOfOtherUsers"`
			EnableSharedDeviceControl       bool   `json:"EnableSharedDeviceControl"`
			EnableRemoteAccess              bool   `json:"EnableRemoteAccess"`
			EnableLiveTvManagement          bool   `json:"EnableLiveTvManagement"`
			EnableLiveTvAccess              bool   `json:"EnableLiveTvAccess"`
			EnableMediaPlayback             bool   `json:"EnableMediaPlayback"`
			EnableAudioPlaybackTranscoding  bool   `json:"EnableAudioPlaybackTranscoding"`
			IsAdministrator                 bool   `json:"IsAdministrator"`
			EnablePlaybackRemuxing          bool   `json:"EnablePlaybackRemuxing"`
			EnableContentDeletion           bool   `json:"EnableContentDeletion"`
			IsTagBlockingModeInclusive      bool   `json:"IsTagBlockingModeInclusive"`
			EnableSubtitleDownloading       bool   `json:"EnableSubtitleDownloading"`
			EnableSubtitleManagement        bool   `json:"EnableSubtitleManagement"`
			EnableSyncTranscoding           bool   `json:"EnableSyncTranscoding"`
			EnableMediaConversion           bool   `json:"EnableMediaConversion"`
			EnableAllChannels               bool   `json:"EnableAllChannels"`
			EnableAllFolders                bool   `json:"EnableAllFolders"`
			IsDisabled                      bool   `json:"IsDisabled"`
			EnablePublicSharing             bool   `json:"EnablePublicSharing"`
			IsHiddenFromUnusedDevices       bool   `json:"IsHiddenFromUnusedDevices"`
			IsHiddenRemotely                bool   `json:"IsHiddenRemotely"`
			IsHidden                        bool   `json:"IsHidden"`
			EnableAllDevices                bool   `json:"EnableAllDevices"`
		} `json:"Policy"`
		HasPassword               bool `json:"HasPassword"`
		HasConfiguredPassword     bool `json:"HasConfiguredPassword"`
		HasConfiguredEasyPassword bool `json:"HasConfiguredEasyPassword"`
	} `json:"User"`
}

type LibraryResp struct {
	Items []struct {
		ProviderIDs    struct{}  `json:"ProviderIds"`
		DateCreated    time.Time `json:"DateCreated"`
		ForcedSortName string    `json:"ForcedSortName"`
		ParentID       string    `json:"ParentId"`
		Etag           string    `json:"Etag"`
		ID             string    `json:"Id"`
		ImageTags      struct {
			Primary string `json:"Primary"`
		} `json:"ImageTags,omitempty"`
		CollectionType          string  `json:"CollectionType"`
		PresentationUniqueKey   string  `json:"PresentationUniqueKey"`
		SortName                string  `json:"SortName"`
		Name                    string  `json:"Name"`
		Path                    string  `json:"Path"`
		FileName                string  `json:"FileName"`
		ServerID                string  `json:"ServerId"`
		DisplayPreferencesID    string  `json:"DisplayPreferencesId"`
		GUID                    string  `json:"Guid"`
		Type                    string  `json:"Type"`
		PrimaryImageAspectRatio float64 `json:"PrimaryImageAspectRatio,omitempty"`
		IsFolder                bool    `json:"IsFolder"`
		CanDownload             bool    `json:"CanDownload"`
		LockData                bool    `json:"LockData"`
		CanDelete               bool    `json:"CanDelete"`
	} `json:"Items"`
	TotalRecordCount int `json:"TotalRecordCount"`
}
type ExternalUrls struct {
	Name string `json:"Name"`
	URL  string `json:"Url"`
}
type MediaStreams struct {
	DisplayLanguage        string  `json:"DisplayLanguage"`
	Codec                  string  `json:"Codec"`
	Language               string  `json:"Language"`
	ColorTransfer          string  `json:"ColorTransfer"`
	ColorPrimaries         string  `json:"ColorPrimaries"`
	ColorSpace             string  `json:"ColorSpace"`
	Comment                string  `json:"Comment"`
	Path                   string  `json:"Path"`
	TimeBase               string  `json:"TimeBase"`
	CodecTimeBase          string  `json:"CodecTimeBase"`
	Title                  string  `json:"Title"`
	Extradata              string  `json:"Extradata"`
	Type                   string  `json:"Type"`
	CodecTag               string  `json:"CodecTag"`
	ChannelLayout          string  `json:"ChannelLayout"`
	NalLengthSize          string  `json:"NalLengthSize"`
	SubtitleLocationType   string  `json:"SubtitleLocationType"`
	MimeType               string  `json:"MimeType"`
	Profile                string  `json:"Profile"`
	PixelFormat            string  `json:"PixelFormat"`
	Protocol               string  `json:"Protocol"`
	ItemID                 string  `json:"ItemId"`
	DeliveryURL            string  `json:"DeliveryUrl"`
	DeliveryMethod         string  `json:"DeliveryMethod"`
	ServerID               string  `json:"ServerId"`
	DisplayTitle           string  `json:"DisplayTitle"`
	AspectRatio            string  `json:"AspectRatio"`
	VideoRange             string  `json:"VideoRange"`
	BitDepth               int     `json:"BitDepth"`
	AverageFrameRate       float64 `json:"AverageFrameRate"`
	RealFrameRate          float64 `json:"RealFrameRate"`
	Width                  int     `json:"Width"`
	Height                 int     `json:"Height"`
	AttachmentSize         int     `json:"AttachmentSize"`
	Index                  uint64  `json:"Index"`
	SampleRate             int     `json:"SampleRate"`
	Channels               int     `json:"Channels"`
	Rotation               int     `json:"Rotation"`
	RefFrames              int     `json:"RefFrames"`
	Level                  int     `json:"Level"`
	BitRate                int     `json:"BitRate"`
	StreamStartTimeTicks   int     `json:"StreamStartTimeTicks"`
	IsDefault              bool    `json:"IsDefault"`
	SupportsExternalStream bool    `json:"SupportsExternalStream"`
	IsTextSubtitleStream   bool    `json:"IsTextSubtitleStream"`
	IsAnamorphic           bool    `json:"IsAnamorphic"`
	IsExternalURL          bool    `json:"IsExternalUrl"`
	IsExternal             bool    `json:"IsExternal"`
	IsForced               bool    `json:"IsForced"`
	IsAVC                  bool    `json:"IsAVC"`
	IsClosedCaptions       bool    `json:"IsClosedCaptions"`
	IsInterlaced           bool    `json:"IsInterlaced"`
}

type RequiredHTTPHeaders struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
}

type MediaSources struct {
	RequiredHTTPHeaders        RequiredHTTPHeaders `json:"RequiredHttpHeaders"`
	DirectStreamURL            string              `json:"DirectStreamUrl"`
	TranscodingContainer       string              `json:"TranscodingContainer"`
	EncoderPath                string              `json:"EncoderPath"`
	EncoderProtocol            string              `json:"EncoderProtocol"`
	Type                       string              `json:"Type"`
	Container                  string              `json:"Container"`
	ServerID                   string              `json:"ServerId"`
	Name                       string              `json:"Name"`
	SortName                   string              `json:"SortName"`
	LiveStreamID               string              `json:"LiveStreamId"`
	Path                       string              `json:"Path"`
	Video3DFormat              string              `json:"Video3DFormat"`
	TranscodingSubProtocol     string              `json:"TranscodingSubProtocol"`
	TranscodingURL             string              `json:"TranscodingUrl"`
	Protocol                   string              `json:"Protocol"`
	ID                         string              `json:"Id"`
	Timestamp                  string              `json:"Timestamp"`
	OpenToken                  string              `json:"OpenToken"`
	ItemID                     string              `json:"ItemId"`
	Formats                    []string            `json:"Formats"`
	MediaStreams               []MediaStreams      `json:"MediaStreams"`
	RunTimeTicks               int                 `json:"RunTimeTicks"`
	BufferMs                   int                 `json:"BufferMs"`
	Size                       int                 `json:"Size"`
	DefaultSubtitleStreamIndex int64               `json:"DefaultSubtitleStreamIndex"`
	AnalyzeDurationMs          int                 `json:"AnalyzeDurationMs"`
	Bitrate                    int                 `json:"Bitrate"`
	DefaultAudioStreamIndex    int64               `json:"DefaultAudioStreamIndex"`
	ContainerStartTimeTicks    int                 `json:"ContainerStartTimeTicks"`
	SupportsDirectPlay         bool                `json:"SupportsDirectPlay"`
	SupportsProbing            bool                `json:"SupportsProbing"`
	SupportsTranscoding        bool                `json:"SupportsTranscoding"`
	IsInfiniteStream           bool                `json:"IsInfiniteStream"`
	IsRemote                   bool                `json:"IsRemote"`
	ReadAtNativeFramerate      bool                `json:"ReadAtNativeFramerate"`
	RequiresOpening            bool                `json:"RequiresOpening"`
	SupportsDirectStream       bool                `json:"SupportsDirectStream"`
	RequiresClosing            bool                `json:"RequiresClosing"`
	RequiresLooping            bool                `json:"RequiresLooping"`
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

type TagItems struct {
	Name string `json:"Name"`
	ID   int    `json:"Id"`
}

type UserData struct {
	LastPlayedDate        time.Time `json:"LastPlayedDate"`
	Key                   string    `json:"Key"`
	ItemID                string    `json:"ItemId"`
	ServerID              string    `json:"ServerId"`
	Rating                int       `json:"Rating"`
	UnplayedItemCount     int       `json:"UnplayedItemCount"`
	PlaybackPositionTicks int       `json:"PlaybackPositionTicks"`
	PlayCount             int       `json:"PlayCount"`
	PlayedPercentage      float32   `json:"PlayedPercentage"`
	IsFavorite            bool      `json:"IsFavorite"`
	Played                bool      `json:"Played"`
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
	Name               string `json:"Name"`
	ImageTag           string `json:"ImageTag"`
	MarkerType         string `json:"MarkerType"`
	StartPositionTicks int    `json:"StartPositionTicks"`
	ChapterIndex       int    `json:"ChapterIndex"`
}

type Items struct {
	DateCreated          time.Time        `json:"DateCreated"`
	PremiereDate         time.Time        `json:"PremiereDate"`
	ImageTags            ImageTags        `json:"ImageTags"`
	SeasonName           string           `json:"SeasonName"`
	CustomRating         string           `json:"CustomRating"`
	Etag                 string           `json:"Etag"`
	Prefix               string           `json:"Prefix"`
	PlaylistItemID       string           `json:"PlaylistItemId"`
	ID                   string           `json:"Id"`
	ExtraType            string           `json:"ExtraType"`
	EpisodeTitle         string           `json:"EpisodeTitle"`
	Container            string           `json:"Container"`
	SortName             string           `json:"SortName"`
	ForcedSortName       string           `json:"ForcedSortName"`
	Video3DFormat        string           `json:"Video3DFormat"`
	ServerID             string           `json:"ServerId"`
	ImageOrientation     string           `json:"ImageOrientation"`
	CameraModel          string           `json:"CameraModel"`
	CameraMake           string           `json:"CameraMake"`
	MediaType            string           `json:"MediaType"`
	LocationType         string           `json:"LocationType"`
	GameSystem           string           `json:"GameSystem"`
	ParentThumbImageTag  string           `json:"ParentThumbImageTag"`
	Path                 string           `json:"Path"`
	OfficialRating       string           `json:"OfficialRating"`
	ParentThumbItemID    string           `json:"ParentThumbItemId"`
	ChannelID            string           `json:"ChannelId"`
	ChannelName          string           `json:"ChannelName"`
	Overview             string           `json:"Overview"`
	SeriesStudio         string           `json:"SeriesStudio"`
	ParentLogoImageTag   string           `json:"ParentLogoImageTag"`
	OriginalTitle        string           `json:"OriginalTitle"`
	Name                 string           `json:"Name"`
	FileName             string           `json:"FileName"`
	AlbumArtist          string           `json:"AlbumArtist"`
	PlayAccess           string           `json:"PlayAccess"`
	CollectionType       string           `json:"CollectionType"`
	Number               string           `json:"Number"`
	ChannelNumber        string           `json:"ChannelNumber"`
	Album                string           `json:"Album"`
	Status               string           `json:"Status"`
	DisplayPreferencesID string           `json:"DisplayPreferencesId"`
	SeasonID             string           `json:"SeasonId"`
	SeriesID             string           `json:"SeriesId"`
	ParentID             string           `json:"ParentId"`
	Type                 string           `json:"Type"`
	SeriesName           string           `json:"SeriesName"`
	DisplayOrder         string           `json:"DisplayOrder"`
	GUID                 string           `json:"Guid"`
	ExternalUrls         []ExternalUrls   `json:"ExternalUrls"`
	TagItems             []TagItems       `json:"TagItems"`
	Tags                 []string         `json:"Tags"`
	AirDays              []string         `json:"AirDays"`
	People               []People         `json:"People"`
	MediaSources         []MediaSources   `json:"MediaSources"`
	RemoteTrailers       []RemoteTrailers `json:"RemoteTrailers"`
	Chapters             []Chapters       `json:"Chapters"`
	ProductionLocations  []string         `json:"ProductionLocations"`
	AlbumArtists         []AlbumArtists   `json:"AlbumArtists"`
	Subviews             []string         `json:"Subviews"`
	MediaStreams         []MediaStreams   `json:"MediaStreams"`
	Artists              []string         `json:"Artists"`
	ArtistItems          []ArtistItems    `json:"ArtistItems"`
	Taglines             []string         `json:"Taglines"`
	Genres               []string         `json:"Genres"`
	BackdropImageTags    []string         `json:"BackdropImageTags"`
	Composers            []Composers      `json:"Composers"`
	UserData             UserData         `json:"UserData"`
	MusicVideoCount      int              `json:"MusicVideoCount"`
	GameSystemID         int              `json:"GameSystemId"`
	Bitrate              int              `json:"Bitrate"`
	RunTimeTicks         int              `json:"RunTimeTicks"`
	ProductionYear       int              `json:"ProductionYear"`
	IndexNumber          int              `json:"IndexNumber"`
	Size                 int              `json:"Size"`
	IndexNumberEnd       int              `json:"IndexNumberEnd"`
	ParentIndexNumber    int              `json:"ParentIndexNumber"`
	SpecialFeatureCount  int              `json:"SpecialFeatureCount"`
	LocalTrailerCount    int              `json:"LocalTrailerCount"`
	PartCount            int              `json:"PartCount"`
	Width                int              `json:"Width"`
	Height               int              `json:"Height"`
	CriticRating         int              `json:"CriticRating"`
	RecursiveItemCount   int              `json:"RecursiveItemCount"`
	ExposureTime         int              `json:"ExposureTime"`
	ChildCount           int              `json:"ChildCount"`
	CompletionPercentage int              `json:"CompletionPercentage"`
	SortIndexNumber      int              `json:"SortIndexNumber"`
	IsMovie              bool             `json:"IsMovie"`
	Disabled             bool             `json:"Disabled"`
	IsFolder             bool             `json:"IsFolder"`
	AsSeries             bool             `json:"AsSeries"`
}

type MeResp struct {
	DateCreated      time.Time `json:"DateCreated"`
	LastLoginDate    time.Time `json:"LastLoginDate"`
	LastActivityDate time.Time `json:"LastActivityDate"`
	Configuration    struct {
		IntroSkipMode              string   `json:"IntroSkipMode"`
		SubtitleMode               string   `json:"SubtitleMode"`
		MyMediaExcludes            []string `json:"MyMediaExcludes"`
		OrderedViews               []string `json:"OrderedViews"`
		LatestItemsExcludes        []string `json:"LatestItemsExcludes"`
		ResumeRewindSeconds        int      `json:"ResumeRewindSeconds"`
		EnableLocalPassword        bool     `json:"EnableLocalPassword"`
		HidePlayedInLatest         bool     `json:"HidePlayedInLatest"`
		RememberAudioSelections    bool     `json:"RememberAudioSelections"`
		RememberSubtitleSelections bool     `json:"RememberSubtitleSelections"`
		EnableNextEpisodeAutoPlay  bool     `json:"EnableNextEpisodeAutoPlay"`
		PlayDefaultAudioTrack      bool     `json:"PlayDefaultAudioTrack"`
		DisplayMissingEpisodes     bool     `json:"DisplayMissingEpisodes"`
	} `json:"Configuration,omitempty"`
	Name     string `json:"Name"`
	ServerID string `json:"ServerId"`
	Prefix   string `json:"Prefix"`
	ID       string `json:"Id"`
	Policy   struct {
		AuthenticationProviderID        string `json:"AuthenticationProviderId"`
		InvalidLoginAttemptCount        int    `json:"InvalidLoginAttemptCount"`
		SimultaneousStreamLimit         int    `json:"SimultaneousStreamLimit"`
		RemoteClientBitrateLimit        int    `json:"RemoteClientBitrateLimit"`
		EnableVideoPlaybackTranscoding  bool   `json:"EnableVideoPlaybackTranscoding"`
		EnableContentDownloading        bool   `json:"EnableContentDownloading"`
		EnableUserPreferenceAccess      bool   `json:"EnableUserPreferenceAccess"`
		EnableRemoteControlOfOtherUsers bool   `json:"EnableRemoteControlOfOtherUsers"`
		EnableSharedDeviceControl       bool   `json:"EnableSharedDeviceControl"`
		EnableRemoteAccess              bool   `json:"EnableRemoteAccess"`
		EnableLiveTvManagement          bool   `json:"EnableLiveTvManagement"`
		EnableLiveTvAccess              bool   `json:"EnableLiveTvAccess"`
		EnableMediaPlayback             bool   `json:"EnableMediaPlayback"`
		EnableAudioPlaybackTranscoding  bool   `json:"EnableAudioPlaybackTranscoding"`
		IsAdministrator                 bool   `json:"IsAdministrator"`
		EnablePlaybackRemuxing          bool   `json:"EnablePlaybackRemuxing"`
		EnableContentDeletion           bool   `json:"EnableContentDeletion"`
		IsTagBlockingModeInclusive      bool   `json:"IsTagBlockingModeInclusive"`
		EnableSubtitleDownloading       bool   `json:"EnableSubtitleDownloading"`
		EnableSubtitleManagement        bool   `json:"EnableSubtitleManagement"`
		EnableSyncTranscoding           bool   `json:"EnableSyncTranscoding"`
		EnableMediaConversion           bool   `json:"EnableMediaConversion"`
		EnableAllChannels               bool   `json:"EnableAllChannels"`
		EnableAllFolders                bool   `json:"EnableAllFolders"`
		IsDisabled                      bool   `json:"IsDisabled"`
		EnablePublicSharing             bool   `json:"EnablePublicSharing"`
		IsHiddenFromUnusedDevices       bool   `json:"IsHiddenFromUnusedDevices"`
		IsHiddenRemotely                bool   `json:"IsHiddenRemotely"`
		IsHidden                        bool   `json:"IsHidden"`
		EnableAllDevices                bool   `json:"EnableAllDevices"`
	} `json:"Policy"`
	HasPassword               bool `json:"HasPassword"`
	HasConfiguredPassword     bool `json:"HasConfiguredPassword"`
	HasConfiguredEasyPassword bool `json:"HasConfiguredEasyPassword"`
}

type SystemInfoResp struct {
	ServerName                           string             `json:"ServerName"`
	WanAddress                           string             `json:"WanAddress"`
	PackageName                          string             `json:"PackageName"`
	ID                                   string             `json:"Id"`
	OperatingSystem                      string             `json:"OperatingSystem"`
	Version                              string             `json:"Version"`
	OperatingSystemDisplayName           string             `json:"OperatingSystemDisplayName"`
	SystemUpdateLevel                    string             `json:"SystemUpdateLevel"`
	InternalMetadataPath                 string             `json:"InternalMetadataPath"`
	LocalAddress                         string             `json:"LocalAddress"`
	TranscodingTempPath                  string             `json:"TranscodingTempPath"`
	ProgramDataPath                      string             `json:"ProgramDataPath"`
	ItemsByNamePath                      string             `json:"ItemsByNamePath"`
	CachePath                            string             `json:"CachePath"`
	LogPath                              string             `json:"LogPath"`
	CompletedInstallations               []InstallationInfo `json:"CompletedInstallations"`
	WebSocketPortNumber                  int32              `json:"WebSocketPortNumber"`
	HTTPServerPortNumber                 int32              `json:"HttpServerPortNumber"`
	HTTPSPortNumber                      int32              `json:"HttpsPortNumber"`
	SupportsHTTPS                        bool               `json:"SupportsHttps"`
	HasUpdateAvailable                   bool               `json:"HasUpdateAvailable"`
	SupportsAutoRunAtStartup             bool               `json:"SupportsAutoRunAtStartup"`
	HardwareAccelerationRequiresPremiere bool               `json:"HardwareAccelerationRequiresPremiere"`
	CanSelfUpdate                        bool               `json:"CanSelfUpdate"`
	CanSelfRestart                       bool               `json:"CanSelfRestart"`
	CanLaunchWebBrowser                  bool               `json:"CanLaunchWebBrowser"`
	SupportsLibraryMonitor               bool               `json:"SupportsLibraryMonitor"`
	IsShuttingDown                       bool               `json:"IsShuttingDown"`
	HasPendingRestart                    bool               `json:"HasPendingRestart"`
}

type InstallationInfo struct {
	ID              string  `json:"Id"`
	Name            string  `json:"Name"`
	AssemblyGUID    string  `json:"AssemblyGuid"`
	Version         string  `json:"Version"`
	UpdateClass     string  `json:"UpdateClass"`
	PercentComplete float64 `json:"PercentComplete,omitempty"`
}

type ItemsResp struct {
	Items            []*Items `json:"Items"`
	TotalRecordCount uint64   `json:"TotalRecordCount"`
}

type PlayBackResp struct {
	PlaySessionID string         `json:"PlaySessionId"`
	MediaSources  []MediaSources `json:"MediaSources"`
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
	VideoCodec          string `json:"VideoCodec,omitempty"`
	ManifestSubtitles   string `json:"ManifestSubtitles,omitempty"`
	BreakOnNonKeyFrames bool   `json:"BreakOnNonKeyFrames,omitempty"`
	CopyTimestamps      bool   `json:"CopyTimestamps,omitempty"`
}

type Condition struct {
	IsRequired any    `json:"IsRequired"`
	Condition  string `json:"Condition"`
	Property   string `json:"Property"`
	Value      string `json:"Value"`
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
	DirectPlayProfiles               []DirectPlayProfile  `json:"DirectPlayProfiles"`
	TranscodingProfiles              []TranscodingProfile `json:"TranscodingProfiles"`
	CodecProfiles                    []CodecProfile       `json:"CodecProfiles"`
	SubtitleProfiles                 []SubtitleProfile    `json:"SubtitleProfiles"`
	ResponseProfiles                 []ResponseProfile    `json:"ResponseProfiles"`
	MaxStaticBitrate                 int                  `json:"MaxStaticBitrate"`
	MaxStreamingBitrate              int                  `json:"MaxStreamingBitrate"`
	MusicStreamingTranscodingBitrate int                  `json:"MusicStreamingTranscodingBitrate"`
}

type PlayBackReq struct {
	ID                             string        `json:"Id,omitempty"`
	UserID                         string        `json:"UserId,omitempty"`
	MediaSourceID                  string        `json:"MediaSourceId,omitempty"`
	LiveStreamID                   string        `json:"LiveStreamId,omitempty"`
	CurrentPlaySessionID           string        `json:"CurrentPlaySessionId,omitempty"`
	DirectPlayProtocols            []string      `json:"DirectPlayProtocols,omitempty"`
	DeviceProfile                  DeviceProfile `json:"DeviceProfile"`
	MaxStreamingBitrate            int           `json:"MaxStreamingBitrate,omitempty"`
	StartTimeTicks                 int           `json:"StartTimeTicks,omitempty"`
	AudioStreamIndex               int           `json:"AudioStreamIndex,omitempty"`
	SubtitleStreamIndex            int           `json:"SubtitleStreamIndex,omitempty"`
	MaxAudioChannels               int           `json:"MaxAudioChannels,omitempty"`
	EnableDirectPlay               bool          `json:"EnableDirectPlay,omitempty"`
	AllowInterlacedVideoStreamCopy bool          `json:"AllowInterlacedVideoStreamCopy,omitempty"`
	AllowVideoStreamCopy           bool          `json:"AllowVideoStreamCopy,omitempty"`
	AllowAudioStreamCopy           bool          `json:"AllowAudioStreamCopy,omitempty"`
	IsPlayback                     bool          `json:"IsPlayback,omitempty"`
	AutoOpenLiveStream             bool          `json:"AutoOpenLiveStream,omitempty"`
	EnableTranscoding              bool          `json:"EnableTranscoding,omitempty"`
	EnableDirectStream             bool          `json:"EnableDirectStream,omitempty"`
}
