package notiUtils

import (
	"errors"
	"os"

	userEntity "github.com/Olapat/go-architecture/model/entity/user"
	apiUtils "github.com/Olapat/go-architecture/utils/api"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	Users []userEntity.User
}

type UsersToPush struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

func (n Notification) PushMessage(messageKey string, bodyParams interface{}, data interface{}) {
	users := make([]UsersToPush, 0)
	for _, v := range n.Users {
		users = append(users, UsersToPush{
			ID:    v.ID.Hex(),
			Token: *v.FcmToken,
		})
	}
	host := os.Getenv("HOST_NOTIFICATION")
	Api := apiUtils.CallAPI{URL: host + "/fcm/push"}
	values := map[string]interface{}{"users": users, "message_key": messageKey, "body_params": bodyParams, "data": data}
	Api.CallAPIPostJson(values)
}

func (n Notification) List() (interface{}, error) {
	var ress interface{}
	if len(n.Users) == 0 {

		return ress, errors.New("User empty")
	}

	host := os.Getenv("HOST_NOTIFICATION")
	Api := apiUtils.CallAPI{URL: host + "/notification/" + n.Users[0].ID.Hex()}
	// ress := res.([]notificationEntity.Response)

	ress = Api.CallAPIGetJson()
	return ress, nil
}

func (n Notification) Read(id string) (interface{}, error) {
	var ress interface{}
	if len(n.Users) == 0 {

		return ress, errors.New("User empty")
	}

	host := os.Getenv("HOST_NOTIFICATION")
	Api := apiUtils.CallAPI{URL: host + "/notification/read"}
	// ress := res.([]notificationEntity.Response)

	values := map[string]interface{}{"id": id, "user_id": n.Users[0].ID.Hex()}
	ress = Api.CallAPIPatchJson(values)
	return ress, nil
}

func (n Notification) DeleteOne(id string) (interface{}, error) {
	var ress interface{}
	if len(n.Users) == 0 {
		return ress, errors.New("User empty")
	}

	host := os.Getenv("HOST_NOTIFICATION")
	Api := apiUtils.CallAPI{URL: host + "/notification/" + id}

	values := map[string]interface{}{"user_id": n.Users[0].ID.Hex()}
	ress = Api.CallAPIDeleteJson(values)
	return ress, nil
}

func (n Notification) DeleteAll() (interface{}, error) {
	var ress interface{}
	if len(n.Users) == 0 {
		return ress, errors.New("User empty")
	}

	host := os.Getenv("HOST_NOTIFICATION")
	Api := apiUtils.CallAPI{URL: host + "/notification/all"}

	values := map[string]interface{}{"user_id": n.Users[0].ID.Hex()}
	ress = Api.CallAPIDeleteJson(values)
	return ress, nil
}

type SMTP struct {
	To      []string
	Subject string
	Params  primitive.M
}

func (s SMTP) Send(keyPath string) {
	host := os.Getenv("HOST_NOTIFICATION")
	Api := apiUtils.CallAPI{URL: host + "/smtp/send/" + keyPath}
	values := map[string]interface{}{"to": s.To, "subject": s.Subject, "params": s.Params}
	Api.CallAPIPostJson(values)
}
