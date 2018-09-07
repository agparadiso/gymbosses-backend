package mock

import (
	"encoding/json"

	"github.com/agparadiso/gymbosses/backend/pkg/clients"
	"github.com/davecgh/go-spew/spew"
)

type ClientsSrv struct{}

func NewClientsSrv() *ClientsSrv {
	return &ClientsSrv{}
}
func (c *ClientsSrv) SearchClientByID(id string) *clients.Client {
	data := []byte(`{"client_id": 5932, "name": "Karen", "profile_pic": "Forma%20Centro/profile_pic/20170816_082132.jpg", "state": 1, "last_name": "Martinez"}`)
	cli := clients.Client{}
	json.Unmarshal(data, &cli)
	return &cli
}

func (c *ClientsSrv) SearchClients(name string) *clients.SearchClientResponse {
	data := []byte(`{"clients": [{"client_id": 9955, "name": "Christian", "state": 1, "last_name": "Ledo"}, {"client_id": 8936, "name": "Juanfranco", "state": 0, "last_name": "Lima"}, {"client_id": 9738, "name": "Marcelo", "state": 1, "last_name": "Kurlander"}, {"client_id": 10296, "name": "Ezequiel", "state": 1, "last_name": "Regueira"}, {"client_id": 10279, "name": "Nicolas", "state": 0, "last_name": "Figueroa"}, {"client_id": 5932, "name": "Karen", "state": 1, "last_name": "Bertucci"}]}`)
	if name == "karen" {
		data = []byte(`{"clients": [{"client_id": 5932, "name": "Karen", "state": 1, "last_name": "Bertucci"}]}`)
	}
	scr := clients.SearchClientResponse{}
	json.Unmarshal(data, &scr)
	return &scr
}
func (c *ClientsSrv) CheckinHistory() *clients.CheckinHistoryResponse {
	dummyData := []byte(`{"checkin_history": [{"date": "2018-08-20 16:38", "client_id": 9955, "name": "Christian", "profile_pic": "profile_pic/no-img.jpg", "state": 1, "last_name": "Ledo"}, {"date": "2018-08-20 15:20", "client_id": 8936, "name": "Juanfranco", "profile_pic": "profile_pic/no-img.jpg", "state": 0, "last_name": "Lima"}, {"date": "2018-08-20 15:07", "client_id": 9738, "name": "Marcelo", "profile_pic": "profile_pic/no-img.jpg", "state": 1, "last_name": "Kurlander"}, {"date": "2018-08-20 14:57", "client_id": 10296, "name": "Ezequiel", "profile_pic": "profile_pic/no-img.jpg", "state": 1, "last_name": "Regueira"}, {"date": "2018-08-20 14:48", "client_id": 10279, "name": "Nicolas", "profile_pic": "profile_pic/no-img.jpg", "state": 0, "last_name": "Figueroa"}, {"date": "2018-08-20 14:29", "client_id": 5932, "name": "Karen", "profile_pic": "Forma%20Centro/profile_pic/20170816_082132.jpg", "state": 1, "last_name": "Martinez"}, {"date": "2018-08-20 14:00", "client_id": 10561, "name": "Liz", "profile_pic": "profile_pic/no-img.jpg", "state": 1, "last_name": "Udaquiola"}, {"date": "2018-08-20 13:54", "client_id": 5892, "name": "Albana", "profile_pic": "Forma%20Centro/profile_pic/Sin%20título%209.jpg", "state": 1, "last_name": "Rodriguez"}, {"date": "2018-08-20 13:51", "client_id": 7798, "name": "Federico", "profile_pic": "profile_pic/no-img.jpg", "state": 1, "last_name": "Barboza"}, {"date": "2018-08-20 13:11", "client_id": 10423, "name": "Santiago", "profile_pic": "profile_pic/no-img.jpg", "state": 1, "last_name": "Bertucci"}]}`)
	chr := clients.CheckinHistoryResponse{}
	json.Unmarshal(dummyData, &chr)
	return &chr
}

func (c *ClientsSrv) NewClient(cli *clients.Client) error {
	spew.Dump(cli)
	return nil
}
