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

func errorCheck(body []byte) models.DeviceCheckError {
	var errResponse models.DeviceCheckError
	error := json.Unmarshal([]byte(body), &errResponse)
	if error != nil {
		fmt.Println("JSON parse error check : ", error)

	}

	return errResponse

}

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

	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("body resp", string(body))
	error := json.Unmarshal([]byte(body), &device)

	if error != nil {
		fmt.Println("JSON parse error: ", error)

	}
	if device.Name == "" {
		errMessage := errorCheck(body)

		device.Name = errMessage.Result
		device.DeviceInfo.DeviceId = "****"
		device.DeviceInfo.DeviceMac = "****"
		device.DeviceInfo.DeviceUUID = "****"
		device.DeviceInfo.LocalIp = ip
		device.DeviceInfo.CoreVersion = "****"
		device.DeviceInfo.VersionDate = "****"
		device.DeviceInfo.WebVersion = "****"

	}
	return device, error

}

func GetPersonListFromDevice(payload models.DeviceAuth) models.PersonListResponse {
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

	fmt.Println("response Status:", resp.Status)
	fmt.Printf("%+v", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	json.Unmarshal([]byte(body), &personList)

	return personList
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

	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("body resp", string(body))
	error := json.Unmarshal([]byte(body), &response)

	if error != nil {
		fmt.Println("JSON parse error: ", error)

	}

	return response, error

}
