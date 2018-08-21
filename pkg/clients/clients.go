package clients

type ClientsSrv interface {
	CheckinHistory() *CheckinHistoryResponse
}

type CheckinHistoryResponse struct {
	Checkin []CheckinHistoryEntry `json:"checkin_history"`
}

type CheckinHistoryEntry struct {
	Date       string `json:"date"`
	ID         int    `json:"client_id"`
	Name       string `json:"client_name"`
	LastName   string `json:"client_last_name"`
	ProfilePic string `json:"client_profile_pic"`
	State      int    `json:"state"`
}
