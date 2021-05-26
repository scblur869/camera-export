package _http

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"local/camera-export/models"
	"net/http"
	"strconv"
	"time"
)

func OpenDeviceCheck(ip string, port string) (models.DeviceInfoACK, error) {

	var infoReq models.Req
	var device models.DeviceInfoACK
	infoReq.Name = "DeviceInfoREQ"
	url := "http://" + ip + ":" + port + "/request"
	jsonBody, err := json.Marshal(&infoReq)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonBody))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println("get device check:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)

	error := json.Unmarshal([]byte(body), &device)

	if error != nil {
		fmt.Println("JSON parse error: ", error)

	}

	return device, error

}

func GetPersonListFromDevice(payload models.DeviceAuth) (models.PersonListResponse, error) {
	url := "http://" + payload.DeviceIP + ":8011/Request"
	var personList models.PersonListResponse

	UserName := payload.User
	PassWord := payload.Password
	uuid := payload.UUID
	ts := int(time.Now().Unix())
	strData := []byte(uuid + ":" + UserName + ":" + PassWord + ":" + strconv.Itoa(ts))
	hasher := md5.New()
	hasher.Write([]byte(strData))
	signedData := hex.EncodeToString(hasher.Sum(nil))

	var jsonBody = []byte(`
	{
	"Name": "personListRequest",
	"UUID": "` + uuid + `",
	"Session": "` + uuid + `_` + strconv.Itoa(ts) + `",
	"TimeStamp": ` + strconv.Itoa(ts) + `,
	"Sign": "` + signedData + `",
	  	"Data": {
		      "Action": "getPersonList", 
					"PersonType": 2,
					"PageNo": 1,
				  "PageSize": 1000
		        }
	}
	 `)
	fmt.Println(string(jsonBody))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	// 	"Sign": "` + signedData + `",
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println("get person list:", resp.Status)
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		fmt.Println(error)
	}
	json.Unmarshal([]byte(body), &personList)

	return personList, error
}

func PersonDetailsRequest(payload models.DeviceAuth, personId string, personType int) (models.Response, error) {
	url := "http://" + payload.DeviceIP + ":8011/Request"
	var response models.Response

	UserName := payload.User
	PassWord := payload.Password
	uuid := payload.UUID
	ts := int(time.Now().Unix())
	strData := []byte(uuid + ":" + UserName + ":" + PassWord + ":" + strconv.Itoa(ts))
	hasher := md5.New()
	hasher.Write([]byte(strData))
	signedData := hex.EncodeToString(hasher.Sum(nil))
	pIdStr := strconv.Itoa(personType)
	var jsonBody = []byte(`
	{
	"Name": "personListRequest",
	"UUID": "` + uuid + `",
	"Session": "` + uuid + `_` + strconv.Itoa(ts) + `",
	"TimeStamp": ` + strconv.Itoa(ts) + `,
	"Sign": "` + signedData + `",
	  	"Data": {
		      "Action": "getPerson", 
			  "PersonType": ` + pIdStr + `,
			  "PersonId": "` + personId + `",
			  "GetPhoto": 1
		        }
	}
	 `)
	// fmt.Println(string(jsonBody))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println("get person details:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("body resp", string(body))
	error := json.Unmarshal([]byte(body), &response)

	if error != nil {
		fmt.Println("JSON parse error: ", error)

	}

	return response, error

}
