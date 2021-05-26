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

	"github.com/nfnt/resize"
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
	personTypeArr := [3]int{1, 2, 3}

	for _, ptype := range personTypeArr {
		listResponse, err := _http.GetPersonListFromDevice(authInfo, ptype)
		if err != nil {
			fmt.Println(err)
		}
		for _, s := range listResponse.Data.PersonList {

			res, err := _http.PersonDetailsRequest(authInfo, s.PersonId, ptype)
			if err != nil {
				fmt.Println(err)
			}
			userList = append(userList, res)
		}
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
		if list.Data.PersonInfo.PersonPhoto != "" {
			CreatePhotoFile(list.Data.PersonInfo.PersonPhoto, list.Data.PersonInfo.PersonName)
		}
	}
	writer.Flush()

}

func CreateDirectory() {
	_, err := os.Stat("./exported-photos")
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("./exported-photos", 0755)
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

	newImage := resize.Resize(960, 960, im, resize.Lanczos3)
	f, err := os.OpenFile("exported-photos/"+name+".jpg", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic("Cannot open file")
	}
	jpeg.Encode(f, newImage, nil)
}
