package onedrive

import "time"

type MicrosoftUserInfo struct {
	OdataContext      string `json:"@odata.context"`
	UserPrincipalName string `json:"userPrincipalName"`
	ID                string `json:"id"`
	DisplayName       string `json:"displayName"`
	Surname           string `json:"surname"`
	GivenName         string `json:"givenName"`
	PreferredLanguage string `json:"preferredLanguage"`
	Mail              string `json:"mail"`
}

type MicrosoftError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type microsoftUserInfoResp struct {
	MicrosoftUserInfo
	MicrosoftError `json:"error"`
}

type FsList struct {
	OdataContext string  `json:"@odata.context"`
	OdataCount   int     `json:"@odata.count"`
	Value        []Value `json:"value"`
}

type fsListResp struct {
	FsList
	MicrosoftError `json:"error"`
}

type Reactions struct {
	CommentCount int `json:"commentCount"`
}

type Application struct {
	DisplayName string `json:"displayName"`
	ID          string `json:"id"`
}

type Device struct {
	ID string `json:"id"`
}

type User struct {
	DisplayName string `json:"displayName"`
	ID          string `json:"id"`
}

type ParentReference struct {
	DriveID   string `json:"driveId"`
	DriveType string `json:"driveType"`
	ID        string `json:"id"`
	Path      string `json:"path"`
}

type FileSystemInfo struct {
	CreatedDateTime      time.Time `json:"createdDateTime"`
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
}

type View struct {
	ViewType  string `json:"viewType"`
	SortBy    string `json:"sortBy"`
	SortOrder string `json:"sortOrder"`
}
type Folder struct {
	ChildCount int  `json:"childCount"`
	View       View `json:"view"`
}

type SpecialFolder struct {
	Name string `json:"name"`
}

type OneDriveSync struct {
	OdataType string `json:"@odata.type"`
	ID        string `json:"id"`
}

type CreatedBy struct {
	Application  Application  `json:"application"`
	Device       Device       `json:"device"`
	User         User         `json:"user"`
	OneDriveSync OneDriveSync `json:"oneDriveSync"`
}

type LastModifiedBy struct {
	Application  Application  `json:"application"`
	Device       Device       `json:"device"`
	User         User         `json:"user"`
	OneDriveSync OneDriveSync `json:"oneDriveSync"`
}
type Hashes struct {
	QuickXorHash string `json:"quickXorHash"`
	Sha1Hash     string `json:"sha1Hash"`
	Sha256Hash   string `json:"sha256Hash"`
}

type File struct {
	MimeType string `json:"mimeType"`
	Hashes   Hashes `json:"hashes"`
}

type Value struct {
	CreatedDateTime           time.Time       `json:"createdDateTime"`
	CTag                      string          `json:"cTag"`
	ETag                      string          `json:"eTag"`
	ID                        string          `json:"id"`
	LastModifiedDateTime      time.Time       `json:"lastModifiedDateTime"`
	Name                      string          `json:"name"`
	Size                      int             `json:"size"`
	WebURL                    string          `json:"webUrl"`
	Reactions                 Reactions       `json:"reactions"`
	CreatedBy                 CreatedBy       `json:"createdBy,omitempty"`
	LastModifiedBy            LastModifiedBy  `json:"lastModifiedBy,omitempty"`
	ParentReference           ParentReference `json:"parentReference"`
	FileSystemInfo            FileSystemInfo  `json:"fileSystemInfo"`
	Folder                    Folder          `json:"folder,omitempty"`
	SpecialFolder             SpecialFolder   `json:"specialFolder,omitempty"`
	MicrosoftGraphDownloadURL string          `json:"@microsoft.graph.downloadUrl,omitempty"`
	File                      File            `json:"file,omitempty"`
}

type FsGet struct {
	CreatedBy struct {
		User struct {
			ID          string `json:"id"`
			DisplayName string `json:"displayName"`
		} `json:"user"`
	} `json:"createdBy"`
	CreatedDateTime time.Time `json:"createdDateTime"`
	CTag            string    `json:"cTag"`
	ETag            string    `json:"eTag"`
	Folder          struct {
		ChildCount int `json:"childCount"`
	} `json:"folder"`
	ID             string `json:"id"`
	LastModifiedBy struct {
		User struct {
			ID          string `json:"id"`
			DisplayName string `json:"displayName"`
		} `json:"user"`
	} `json:"lastModifiedBy"`
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
	Name                 string    `json:"name"`
	Root                 struct{}  `json:"root"`
	Size                 int       `json:"size"`
	WebURL               string    `json:"webUrl"`
}

type fsGetResp struct {
	FsGet
	MicrosoftError `json:"error"`
}
