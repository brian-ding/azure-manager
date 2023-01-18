package refreship

import (
	"context"
	"net/http"

	"github.com/brian-ding/azure-manager/internal/pkg/mgr"
	"github.com/gin-gonic/gin"
)

func RefreshIP(c *gin.Context) {
	var req Request

	if c.BindJSON(&req) != nil {
		return
	}

	ctx := context.Background()

	ipMgr, err := mgr.NewIPAddressManager(req.SubsId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	ip, err := ipMgr.GetIPAddress(ctx, req.IpGrNm, req.IpResNm)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)

		return
	}

	oldAddr := *ip.Properties.IPAddress

	itfMgr, err := mgr.NewInterfaceManager(req.SubsId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	itf, err := itfMgr.Get(ctx, req.ItfGrNm, req.ItfResNm)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)

		return
	}

	itf.Properties.IPConfigurations[0].Properties.PublicIPAddress = nil
	flag, err := itfMgr.Update(ctx, req.ItfGrNm, req.ItfResNm, itf)
	if !flag || err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	itf.Properties.IPConfigurations[0].Properties.PublicIPAddress = ip
	flag, err = itfMgr.Update(ctx, req.ItfGrNm, req.ItfResNm, itf)
	if !flag || err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	ip, err = ipMgr.GetIPAddress(ctx, req.IpGrNm, req.IpResNm)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)

		return
	}

	newAddr := *ip.Properties.IPAddress

	c.String(200, "IP Address before dissociation: %s\nIP Address after association: %s\n", oldAddr, newAddr)
}
