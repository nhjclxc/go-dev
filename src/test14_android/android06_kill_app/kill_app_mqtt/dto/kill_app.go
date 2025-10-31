package dto

type KillAppCmdReq struct {
	PackageList []string `json:"packageList"`
}

type KillAppCmdResp struct {
	Reason        string           `json:"reason"`
	CmdStatus     bool             `json:"cmdStatus"`
	KillAppStatus []*KillAppStatus `json:"killAppStatus"`
}

type KillAppStatus struct {
	Package       string `json:"package"`
	CmdStatus     bool   `json:"cmdStatus"`
	CloseStatus   bool   `json:"closeStatus"`
	FailureReason string `json:"failureReason"`
}
