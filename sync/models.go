package sync

import "grout/romm"

type LocalSave struct {
	RomID       int
	RomName     string
	FSSlug      string
	FileName    string
	FilePath    string
	EmulatorDir string
	RomFileName string
}

type SyncAction int

const (
	ActionUpload SyncAction = iota
	ActionDownload
	ActionConflict
	ActionSkip
)

func (a SyncAction) String() string {
	switch a {
	case ActionUpload:
		return "upload"
	case ActionDownload:
		return "download"
	case ActionConflict:
		return "conflict"
	case ActionSkip:
		return "skip"
	default:
		return "unknown"
	}
}

type SyncItem struct {
	LocalSave      LocalSave
	RemoteSave     *romm.Save
	Action         SyncAction
	Success        bool
	ForceOverwrite bool
}

func (item *SyncItem) Resolve(action SyncAction) {
	item.Action = action
}

type SyncReport struct {
	Uploaded   int
	Downloaded int
	Conflicts  int
	Skipped    int
	Errors     int
	Items      []SyncItem
}
