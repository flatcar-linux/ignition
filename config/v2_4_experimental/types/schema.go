package types

// generated by "schematyper --package=types schema/ignition.json -o internal/config/types/schema.go --root-type=Config" -- DO NOT EDIT

type CaReference struct {
	HTTPHeaders  HTTPHeaders  `json:"httpHeaders,omitempty"`
	Source       string       `json:"source"`
	Verification Verification `json:"verification,omitempty"`
}

type Config struct {
	Ignition Ignition `json:"ignition"`
	Networkd Networkd `json:"networkd,omitempty"`
	Passwd   Passwd   `json:"passwd,omitempty"`
	Storage  Storage  `json:"storage,omitempty"`
	Systemd  Systemd  `json:"systemd,omitempty"`
}

type ConfigReference struct {
	HTTPHeaders  HTTPHeaders  `json:"httpHeaders,omitempty"`
	Source       string       `json:"source"`
	Verification Verification `json:"verification,omitempty"`
}

type Create struct {
	Force   bool           `json:"force,omitempty"`
	Options []CreateOption `json:"options,omitempty"`
}

type CreateOption string

type Device string

type Directory struct {
	Node
	DirectoryEmbedded1
}

type DirectoryEmbedded1 struct {
	Mode *int `json:"mode,omitempty"`
}

type Disk struct {
	Device     string      `json:"device"`
	Partitions []Partition `json:"partitions,omitempty"`
	WipeTable  bool        `json:"wipeTable,omitempty"`
}

type File struct {
	Node
	FileEmbedded1
}

type FileContents struct {
	Compression  string       `json:"compression,omitempty"`
	HTTPHeaders  HTTPHeaders  `json:"httpHeaders,omitempty"`
	Source       string       `json:"source,omitempty"`
	Verification Verification `json:"verification,omitempty"`
}

type FileEmbedded1 struct {
	Append   bool         `json:"append,omitempty"`
	Contents FileContents `json:"contents,omitempty"`
	Mode     *int         `json:"mode,omitempty"`
}

type Filesystem struct {
	Mount *Mount  `json:"mount,omitempty"`
	Name  string  `json:"name,omitempty"`
	Path  *string `json:"path,omitempty"`
}

type Group string

type HTTPHeader []HTTPHeaderItem

type HTTPHeaderItem string

type HTTPHeaders []HTTPHeader

type Ignition struct {
	Config   IgnitionConfig `json:"config,omitempty"`
	Proxy    Proxy          `json:"proxy,omitempty"`
	Security Security       `json:"security,omitempty"`
	Timeouts Timeouts       `json:"timeouts,omitempty"`
	Version  string         `json:"version,omitempty"`
}

type IgnitionConfig struct {
	Append  []ConfigReference `json:"append,omitempty"`
	Replace *ConfigReference  `json:"replace,omitempty"`
}

type Link struct {
	Node
	LinkEmbedded1
}

type LinkEmbedded1 struct {
	Hard   bool   `json:"hard,omitempty"`
	Target string `json:"target"`
}

type Mount struct {
	Create         *Create       `json:"create,omitempty"`
	Device         string        `json:"device"`
	Format         string        `json:"format"`
	Label          *string       `json:"label,omitempty"`
	Options        []MountOption `json:"options,omitempty"`
	UUID           *string       `json:"uuid,omitempty"`
	WipeFilesystem bool          `json:"wipeFilesystem,omitempty"`
}

type MountOption string

type Networkd struct {
	Units []Networkdunit `json:"units,omitempty"`
}

type NetworkdDropin struct {
	Contents string `json:"contents,omitempty"`
	Name     string `json:"name"`
}

type Networkdunit struct {
	Contents string           `json:"contents,omitempty"`
	Dropins  []NetworkdDropin `json:"dropins,omitempty"`
	Name     string           `json:"name"`
}

type NoProxyItem string

type Node struct {
	Filesystem string     `json:"filesystem"`
	Group      *NodeGroup `json:"group,omitempty"`
	Overwrite  *bool      `json:"overwrite,omitempty"`
	Path       string     `json:"path"`
	User       *NodeUser  `json:"user,omitempty"`
}

type NodeGroup struct {
	ID   *int   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type NodeUser struct {
	ID   *int   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Partition struct {
	GUID               string  `json:"guid,omitempty"`
	Label              *string `json:"label,omitempty"`
	Number             int     `json:"number,omitempty"`
	ShouldExist        *bool   `json:"shouldExist,omitempty"`
	Size               *int    `json:"size,omitempty"`
	SizeMiB            *int    `json:"sizeMiB,omitempty"`
	Start              *int    `json:"start,omitempty"`
	StartMiB           *int    `json:"startMiB,omitempty"`
	TypeGUID           string  `json:"typeGuid,omitempty"`
	WipePartitionEntry bool    `json:"wipePartitionEntry,omitempty"`
}

type Passwd struct {
	Groups []PasswdGroup `json:"groups,omitempty"`
	Users  []PasswdUser  `json:"users,omitempty"`
}

type PasswdGroup struct {
	Gid          *int   `json:"gid,omitempty"`
	Name         string `json:"name"`
	PasswordHash string `json:"passwordHash,omitempty"`
	System       bool   `json:"system,omitempty"`
}

type PasswdUser struct {
	Create            *Usercreate        `json:"create,omitempty"`
	Gecos             string             `json:"gecos,omitempty"`
	Groups            []Group            `json:"groups,omitempty"`
	HomeDir           string             `json:"homeDir,omitempty"`
	Name              string             `json:"name"`
	NoCreateHome      bool               `json:"noCreateHome,omitempty"`
	NoLogInit         bool               `json:"noLogInit,omitempty"`
	NoUserGroup       bool               `json:"noUserGroup,omitempty"`
	PasswordHash      *string            `json:"passwordHash,omitempty"`
	PrimaryGroup      string             `json:"primaryGroup,omitempty"`
	SSHAuthorizedKeys []SSHAuthorizedKey `json:"sshAuthorizedKeys,omitempty"`
	Shell             string             `json:"shell,omitempty"`
	System            bool               `json:"system,omitempty"`
	UID               *int               `json:"uid,omitempty"`
}

type Proxy struct {
	HTTPProxy  string        `json:"httpProxy,omitempty"`
	HTTPSProxy string        `json:"httpsProxy,omitempty"`
	NoProxy    []NoProxyItem `json:"noProxy,omitempty"`
}

type Raid struct {
	Devices []Device     `json:"devices"`
	Level   string       `json:"level"`
	Name    string       `json:"name"`
	Options []RaidOption `json:"options,omitempty"`
	Spares  int          `json:"spares,omitempty"`
}

type RaidOption string

type SSHAuthorizedKey string

type Security struct {
	TLS TLS `json:"tls,omitempty"`
}

type Storage struct {
	Directories []Directory  `json:"directories,omitempty"`
	Disks       []Disk       `json:"disks,omitempty"`
	Files       []File       `json:"files,omitempty"`
	Filesystems []Filesystem `json:"filesystems,omitempty"`
	Links       []Link       `json:"links,omitempty"`
	Raid        []Raid       `json:"raid,omitempty"`
}

type Systemd struct {
	Units []Unit `json:"units,omitempty"`
}

type SystemdDropin struct {
	Contents string `json:"contents,omitempty"`
	Name     string `json:"name"`
}

type TLS struct {
	CertificateAuthorities []CaReference `json:"certificateAuthorities,omitempty"`
}

type Timeouts struct {
	HTTPResponseHeaders *int `json:"httpResponseHeaders,omitempty"`
	HTTPTotal           *int `json:"httpTotal,omitempty"`
}

type Unit struct {
	Contents string          `json:"contents,omitempty"`
	Dropins  []SystemdDropin `json:"dropins,omitempty"`
	Enable   bool            `json:"enable,omitempty"`
	Enabled  *bool           `json:"enabled,omitempty"`
	Mask     bool            `json:"mask,omitempty"`
	Name     string          `json:"name"`
}

type Usercreate struct {
	Gecos        string            `json:"gecos,omitempty"`
	Groups       []UsercreateGroup `json:"groups,omitempty"`
	HomeDir      string            `json:"homeDir,omitempty"`
	NoCreateHome bool              `json:"noCreateHome,omitempty"`
	NoLogInit    bool              `json:"noLogInit,omitempty"`
	NoUserGroup  bool              `json:"noUserGroup,omitempty"`
	PrimaryGroup string            `json:"primaryGroup,omitempty"`
	Shell        string            `json:"shell,omitempty"`
	System       bool              `json:"system,omitempty"`
	UID          *int              `json:"uid,omitempty"`
}

type UsercreateGroup string

type Verification struct {
	Hash *string `json:"hash,omitempty"`
}
