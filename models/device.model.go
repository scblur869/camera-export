package models

type DeviceInfoACK struct {
	Name       string    `json:"Name"`
	DeviceInfo DeviceAck `json:"DeviceInfo"`
}

type DeviceAck struct {
	DeviceId    string `json:"DeviceId"`
	DeviceUUID  string `json:"DeviceUUID"`
	DeviceMac   string `json:"DeviceMac"`
	LocalIp     string `json:"LocalIp"`
	WebVersion  string `json:"WebVersion"`
	CoreVersion string `json:"CoreVersion"`
	VersionDate string `json:"VerDate"`
}

type Req struct {
	Name      string `json:"Name"`
	Timestamp int    `json:"Timestamp"`
	Sign      string `json:"Sign"`
}

type ErrorResonse struct {
	Name    string `json:"Name"`
	Code    int    `json:"Code"`
	Message string `json:"Message"`
}

type DeviceCheckError struct {
	Result    string `json:"Result"`
	ErrorCode int    `json:"ErrorCode"`
}

type DeviceAuth struct {
	User     string `json:"user"`
	Password string `json:"password"`
	UUID     string `json:"uuid"`
	DeviceIP string `json:"deviceip"`
}
