package models

type UpdatePostConnectDetails struct {
	UID        string `json:"uid"`
	CallStatus string `json:"status_code"`
	HangupBy   string `json:"hangup_by"`
}

type PushCallbackDetails struct {
	UID          string `json:"uid"`
	CallDuration int    `json:"call_duration"`
	DialDuration int    `json:"dial_duration"`
	CallStatus   string `json:"status_code"`
	HangupBy     string `json:"hangup_by"`
}
