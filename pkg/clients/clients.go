package clients

type ClientsSrv interface {
	NewClient(*Client) error
	CheckinHistory(id string) *CheckinHistoryResponse
	SearchClientByID(id string) *Client
	SearchClients(name string) *SearchClientResponse
}

type Client struct {
	ID               int    `json:"client_id"`
	Name             string `json:"name"`
	LastName         string `json:"last_name"`
	Gender           string `json:"gender"`
	Birthdate        string `json:"birthdate"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	EmergencyContact string `json:"emergency_contact"`
	Diseases         string `json:"diseases"`
	ProfilePic       string `json:"profile_pic"`
	MeetUs           string `json:"meet_us"`
	MedicalInsurance string `json:"medical_insurance"`
	MedicalEmergency string `json:"medical_emergency"`
	State            int    `json:"state"`
}

type CheckinHistoryEntry struct {
	Client
	Date string `json:"date"`
}

type CheckinHistoryResponse struct {
	Checkin []CheckinHistoryEntry `json:"checkin_history"`
}

type SearchClientEntry struct {
	ID       int    `json:"client_id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	State    int    `json:"state"`
}

type SearchClientResponse struct {
	Clients []SearchClientEntry `json:"clients"`
}
