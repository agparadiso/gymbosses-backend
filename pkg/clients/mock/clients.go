package mock

import (
	"encoding/json"

	"github.com/agparadiso/gymbosses/backend/pkg/clients"
)

type ClientsSrv struct{}

func NewClientsSrv() *ClientsSrv {
	return &ClientsSrv{}
}

func (c *ClientsSrv) CheckinHistory() *clients.CheckinHistoryResponse {
	dummyData := []byte(`{"checkin_history": [{"date": "2018-08-20 16:38", "client_id": 9955, "client_name": "Christian", "client_profile_pic": "profile_pic/no-img.jpg", "state": 1, "client_last_name": "Ledo"}, {"date": "2018-08-20 15:20", "client_id": 8936, "client_name": "Juanfranco", "client_profile_pic": "profile_pic/no-img.jpg", "state": 1, "client_last_name": "Lima"}, {"date": "2018-08-20 15:07", "client_id": 9738, "client_name": "Marcelo", "client_profile_pic": "profile_pic/no-img.jpg", "state": 1, "client_last_name": "Kurlander"}, {"date": "2018-08-20 14:57", "client_id": 10296, "client_name": "Ezequiel", "client_profile_pic": "profile_pic/no-img.jpg", "state": 1, "client_last_name": "Regueira"}, {"date": "2018-08-20 14:48", "client_id": 10279, "client_name": "Nicolas", "client_profile_pic": "profile_pic/no-img.jpg", "state": 1, "client_last_name": "Figueroa"}, {"date": "2018-08-20 14:29", "client_id": 5932, "client_name": "Karen", "client_profile_pic": "profile_pic/no-img.jpg", "state": 1, "client_last_name": "Martinez"}, {"date": "2018-08-20 14:00", "client_id": 10561, "client_name": "Liz", "client_profile_pic": "profile_pic/no-img.jpg", "state": 1, "client_last_name": "Udaquiola"}, {"date": "2018-08-20 13:54", "client_id": 5892, "client_name": "Albana", "client_profile_pic": "profile_pic/no-img.jpg", "state": 1, "client_last_name": "Rodriguez"}, {"date": "2018-08-20 13:51", "client_id": 7798, "client_name": "Federico", "client_profile_pic": "profile_pic/no-img.jpg", "state": 1, "client_last_name": "Barboza"}, {"date": "2018-08-20 13:11", "client_id": 10423, "client_name": "Santiago", "client_profile_pic": "profile_pic/no-img.jpg", "state": 1, "client_last_name": "Bertucci"}]}`)
	chr := clients.CheckinHistoryResponse{}
	json.Unmarshal(dummyData, &chr)
	return &chr
}
