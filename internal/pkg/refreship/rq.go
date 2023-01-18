package refreship

type Request struct {
	SubsId   string `json:"subsId"`
	ItfGrNm  string `json:"itfGrNm"`
	ItfResNm string `json:"itfResNm"`
	IpGrNm   string `json:"ipGrNm"`
	IpResNm  string `json:"ipResNm"`
}
