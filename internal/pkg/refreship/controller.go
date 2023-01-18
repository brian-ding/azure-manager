package refreship

import (
	"context"
	"fmt"
	"net/http"

	"github.com/brian-ding/azure-manager/internal/pkg/mgr"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var record *Record

func RefreshHandler(c *gin.Context) {
	var req RefreshRequest

	if c.BindJSON(&req) != nil {
		return
	}

	uuid := uuid.New()
	record = &Record{ID: uuid, Status: notStarted}

	go refresh(req.SubsId, req.ItfGrNm, req.ItfResNm, req.IpGrNm, req.IpResNm)

	c.JSON(http.StatusAccepted, RefreshResponse{RecordId: uuid})
}

func CheckHandler(c *gin.Context) {
	var req CheckRequest

	if err := c.BindUri(&req); err != nil {
		fmt.Println(err.Error())
		return
	}

	if record.ID.String() != req.RecordId {
		c.JSON(http.StatusNotFound, CheckResponse{Message: "record does not exist"})
		return
	}

	c.JSON(http.StatusOK, CheckResponse{"Succeed", *record})
}

func refresh(subsId, itfGrNm, itfResNm, ipGrNm, ipResNm string) {
	ctx := context.Background()

	record.Status = getIpMgr
	ipMgr, err := mgr.NewIPAddressManager(subsId)
	if err != nil {
		record.Message = err.Error()
		record.Status = fail
		return
	}

	record.Status = getIp
	record.Message = "ip manager got, getting public ip"
	ip, err := ipMgr.GetIPAddress(ctx, ipGrNm, ipResNm)
	if err != nil {
		record.Message = err.Error()
		record.Status = fail
		return
	}

	oldAddr := *ip.Properties.IPAddress
	record.Message = fmt.Sprintf("old ip address: %s, getting interface manager", oldAddr)

	record.Status = getItfMgr
	itfMgr, err := mgr.NewInterfaceManager(subsId)
	if err != nil {
		record.Message = err.Error()
		record.Status = fail
		return
	}

	record.Status = getItf
	record.Message = "interface manager got, getting interface"
	itf, err := itfMgr.Get(ctx, itfGrNm, itfResNm)
	if err != nil {
		record.Message = err.Error()
		record.Status = fail
		return
	}

	record.Status = updateItf
	record.Message = "interface got, dissociating"
	itf.Properties.IPConfigurations[0].Properties.PublicIPAddress = nil
	flag, err := itfMgr.Update(ctx, itfGrNm, itfResNm, itf)
	if !flag || err != nil {
		record.Message = err.Error()
		record.Status = fail
		return
	}

	record.Status = updateItf
	record.Message = "dissociated, associating"
	itf.Properties.IPConfigurations[0].Properties.PublicIPAddress = ip
	flag, err = itfMgr.Update(ctx, itfGrNm, itfResNm, itf)
	if !flag || err != nil {
		record.Message = err.Error()
		record.Status = fail
		return
	}

	record.Status = getIp
	record.Message = "associated, getting public ip"
	ip, err = ipMgr.GetIPAddress(ctx, ipGrNm, ipResNm)
	if err != nil {
		record.Message = err.Error()
		record.Status = fail
		return
	}

	newAddr := *ip.Properties.IPAddress
	record.Status = succeed
	record.Message = fmt.Sprintf("old ip address: %s, new ip address: %s", oldAddr, newAddr)
}
