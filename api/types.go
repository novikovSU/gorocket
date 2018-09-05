package api

type Channel struct {
	Id           string   `json:"_id"`
	Name         string   `json:"name"`
	MessageCount int      `json:"msgs"`
	UserNames    []string `json:"usernames"`

	User User `json:"u"`

	ReadOnly  bool   `json:"ro"`
	Timestamp string `json:"ts"`
	T         string `json:"t"`
	UpdatedAt string `json:"_updatedAt"`
	SysMes    bool   `json:"sysMes"`
}

//TODO: finish
type Attachment struct {
}

type Room struct {
	Id   string `json:"_id"`
	Type string `json:"t"`
	Name string `json:"name,omitempty"`
}

type Subscription struct {
	Type          string `json:"t"`
	TimeStamp     string `json:"ts"`
	Name          string `json:"name"`
	Fname         string `json:"fname,omitempty"`
	RoomId        string `json:"rid"`
	User          User   `json:"u"`
	Open          bool   `json:"open"`
	Alert         bool   `json:"alert"`
	Unread        int64  `json:"unread"`
	UserMentions  int64  `json:"userMentions"`
	GroupMentions int64  `json:"groupMentions"`
	UpdatedAt     string `json:"_updatedAt"`
	Id            string `json:"_id"`
}

type ReadReceipt struct {
	Id        string `json:"_id"`
	RoomId    string `json:"roomId"`
	UserId    string `json:"userId"`
	MessageId string `json:"messageId"`
	TimeStamp string `json:"ts"`
	User      User   `json:"user"`
}

type User struct {
	Id       string `json:"_id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
}

type UserCredentials struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"pass"`
}

type Message struct {
	Id        string `json:"_id"`
	ChannelId string `json:"rid"`
	Text      string `json:"msg"`
	Timestamp string `json:"ts"`
	User      User   `json:"u"`
	EditedAt  string `json:"editedAt,omitempty"`
	EditedBy  User   `json:"editedBy,omitempty"`
}

type Group struct {
	Id        string `json:"_id"`
	Name      string `json:"name"`
	T         string `json:"t"`
	Msgs      int64  `json:"msgs"`
	U         User   `json:"u"`
	Timestamp string `json:"ts"`
	Readonly  bool   `json:"ro"`
	Sysmes    bool   `json:"sysMes"`
	UpdatedAt string `json:"_updatedAt"`
}

type Info struct {
	Version string `json:"version"`

	Build struct {
		Date        string `json:"date"`
		NodeVersion string `json:"nodeVersion"`
		Arch        string `json:"arch"`
		Platform    string `json:"platform"`
		OsRelease   string `json:"osRelease"`
		TotalMemory int64  `json:"totalMemory"`
		FreeMemory  int64  `json:"freeMemory"`
		CpuCount    int    `json:"cpus"`
	} `json:"build"`

	Travis struct {
		BuildNumber string `json:"buildNumber"`
		Branch      string `json:"branch"`
		Tag         string `json:"tag"`
	} `json:"travis"`

	Commit struct {
		Hash    string `json:"hash"`
		Date    string `json:"date"`
		Author  string `json:"author"`
		Subject string `json:"subject"`
		Tag     string `json:"tag"`
		Branch  string `json:"branch"`
	} `json:"commit"`

	GraphicsMagick struct {
		Enabled bool `json:"enabled"`
	} `json:"GraphicsMagick"`

	ImageMagick struct {
		Enabled bool   `json:"enabled"`
		Version string `json:"version"`
	} `json:"ImageMagick"`
}
