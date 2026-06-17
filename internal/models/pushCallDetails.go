package models

type UpdatePostConnectDetails struct {
	UID        string `json:"uid"`
	CallStatus string `json:"status_code"`
	HangupBy   string `json:"hangup_by"`
}

type PushCallbackDetails struct {
	UID            string `json:"uid"`
	CallDuration   string    `json:"call_duration"`
	DialedDuration string    `json:"dialed_duration"`
	CallStatus     string `json:"status_code"`
	HangupBy       string `json:"hangup_by"`
}
