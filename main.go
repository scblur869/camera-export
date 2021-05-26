package main

import (
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"flag"
	"fmt"
	"image/jpeg"
	"local/camera-export/_http"
	"local/camera-export/models"
	"os"
	"strconv"
)

func main() {
	var authInfo models.DeviceAuth
	var userList []models.Response
	deviceIp := flag.String("deviceip", "192.168.1.20", "ip address of camera")
	userName := flag.String("user", "admin", "admin user account")
	password := flag.String("pass", "admin", "admin user account password")
	port := flag.String("port", "8011", "device default port")
	flag.Parse()
	device, err := _http.OpenDeviceCheck(*deviceIp, *port)
	if err != nil {
		fmt.Println(err)
	}
	authInfo.DeviceIP = *deviceIp
	authInfo.UUID = device.DeviceInfo.DeviceUUID
	authInfo.User = *userName
	authInfo.Password = *password

	listResponse := _http.GetPersonListFromDevice(authInfo)

	for _, s := range listResponse.Data.PersonList {

		res, err := _http.PersonDetailsRequest(authInfo, s.PersonId, listResponse.Data.PersonType)
		if err != nil {
			fmt.Println(err)
		}
		userList = append(userList, res)
	}
	csvFile, err := os.Create("device-" + authInfo.UUID + ".csv")
	if err != nil {
		fmt.Println(err)
	}
	CreateDirectory()
	writer := csv.NewWriter(csvFile)
	var header []string
	header = append(header, "Name")
	header = append(header, "User ID")
	header = append(header, "Gender")
	header = append(header, "ID Number")
	header = append(header, "Telephone")
	header = append(header, "IC Card")
	header = append(header, "Picture Name")
	header = append(header, "Address")
	header = append(header, "Remarks")
	writer.Write(header)

	for _, list := range userList {
		var row []string
		row = append(row, list.Data.PersonInfo.PersonName)
		row = append(row, list.Data.PersonInfo.PersonId)
		row = append(row, strconv.Itoa(list.Data.PersonInfo.Sex))
		row = append(row, list.Data.PersonInfo.IDCard)
		row = append(row, list.Data.PersonInfo.Phone)
		row = append(row, strconv.Itoa(list.Data.PersonInfo.PersonExtension.PersonParam5))
		row = append(row, list.Data.PersonInfo.PersonName+".jpg")
		row = append(row, list.Data.PersonInfo.Address)
		row = append(row, list.Data.PersonInfo.PersonExtension.PersonData4)
		writer.Write(row)
		CreatePhotoFile(list.Data.PersonInfo.PersonPhoto, list.Data.PersonInfo.PersonName)
	}
	writer.Flush()

}

func CreateDirectory() {
	_, err := os.Stat("./photos")
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("./photos", 0755)
		if errDir != nil {
			fmt.Println(err)
		}
	}
}

func CreatePhotoFile(b64String string, name string) {
	unbased, err := base64.StdEncoding.DecodeString(b64String)
	if err != nil {
		panic("Cannot decode b64")
	}
	r := bytes.NewReader(unbased)
	im, err := jpeg.Decode(r)
	if err != nil {
		panic("Bad jpeg")
	}

	f, err := os.OpenFile("photos/"+name+".jpg", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic("Cannot open file")
	}

	jpeg.Encode(f, im, nil)
}
