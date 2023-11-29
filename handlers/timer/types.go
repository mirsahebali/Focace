package timer

import "net/http"

type Timer struct {
	UserId        string `json:"user_id"`
	FocusTime     int    `json:"focus_time"`
	BreakTime     int    `json:"break_time"`
	LongBreakTime int    `json:"long_break_time"`
	Sessions      int    `json:"sessions"`
}

func ChangeSecondsRoute(w http.ResponseWriter, r *http.Request) {
}
