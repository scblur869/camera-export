package models

type User struct {
	Name    string  `json:"name"`
	UUID    string  `json:"uuid"`
	Session string  `json:"session"`
	Sign    string  `json:"sign"`
	Data    Request `json:"data"`
}

type Request struct {
	Action     string `json:"action"`
	PersonType int    `json:"persontype"`
	PersonId   string `json:"personid"`
	GetPhoto   int    `json:"getphoto"`
}

type Response struct {
	Name    string       `json:"name"`
	Session string       `json:"session"`
	Data    DataResponse `json:"data"`
	Code    int          `json:"code"`
	Message string       `json:"message"`
}

type DataResponse struct {
	Action     string     `json:"action"`
	PersonType int        `json:"persontype"`
	PersonInfo PERSONINFO `json:"personinfo"`
}

type PERSONINFO struct {
	PersonId        string    `json:"PersonId"`
	PersonName      string    `json:"PersonName"`
	Sex             int       `json:"Sex"`
	IDCard          string    `json:"IDCard"`
	Nation          string    `json:"Nation"`
	Birthday        string    `json:"Birthday"`
	Phone           string    `json:"Phone"`
	Address         string    `json:"Address"`
	SaveTime        string    `json:"SaveTime"`
	LimitTime       int       `json:"LimitTime"`
	EndTime         string    `json:"EndTime"`
	Label           string    `json:"Label"`
	PersonExtension PERSONEXT `json:"PersonExtension"`
	PersonPhoto     string    `json:"PersonPhoto"`
}

type PERSONEXT struct {
	PersonCode1       string `json:"PersonCode1"`
	PersonCode2       string `json:"PersonCode2"`
	PersonCode3       string `json:"PersonCode3"`
	PersonReserveName string `json:"PersonReserveName"`
	PersonParam1      int    `json:"PersonParam1"`
	PersonParam2      int    `json:"PersonParam2"`
	PersonParam3      int    `json:"PersonParam3"`
	PersonParam4      int    `json:"PersonParam4"`
	PersonParam5      int    `json:"PersonParam5"`
	PersonData1       string `json:"PersonData1"`
	PersonData2       string `json:"PersonData2"`
	PersonData3       string `json:"PersonData3"`
	PersonData4       string `json:"PersonData4"`
	PersonData5       string `json:"PersonData5"`
}

type PersonListResponse struct {
	Name    string         `json:"name"`
	Data    PersonListData `json:"data"`
	Code    int            `json:"code"`
	Message string         `json:"message"`
}

type PersonListData struct {
	PersonType    int          `json:"persontype"`
	Action        string       `json:"action"`
	PageNo        int          `json:"pageno"`
	PageSize      int          `json:"pagesize"`
	PageCount     int          `json:"pagecount"`
	PersonCount   int          `json:"personcount"`
	PersonListNum int          `json:"personlistnum"`
	PersonList    []PersonList `json:"personlist"`
}

type PersonList struct {
	PersonId   string `json:"personid"`
	PersonName string `json:"personname"`
}
