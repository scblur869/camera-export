package main

import (
	"encoding/csv"
	"flag"
	"fmt"
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
	}
	writer.Flush()
}
