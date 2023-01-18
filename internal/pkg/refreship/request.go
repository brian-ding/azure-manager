package refreship

type RefreshRequest struct {
	SubsId   string `json:"subsId"`
	ItfGrNm  string `json:"itfGrNm"`
	ItfResNm string `json:"itfResNm"`
	IpGrNm   string `json:"ipGrNm"`
	IpResNm  string `json:"ipResNm"`
}

type CheckRequest struct {
	RecordId string `uri:"id" binding:"required,uuid"`
}
