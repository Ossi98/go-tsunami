package scanner

import (
	"Ossi98/go-tsunami/internal/cmd"
	"Ossi98/go-tsunami/internal/utils/validator"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Scanner struct {
	psExec *cmd.Proc
}

func NewScannerCtrl(ps *cmd.Proc) *Scanner {
	return &Scanner{
		psExec: ps,
	}
}

type ScanRequest struct {
	Target     string `json:"target" binding:"required"` // verify if address valid in case of ipv4,6 or domain in case of hostname or uri
	TypeTarget string `json:"type" binding:"required"`
}

func (s *Scanner) StartScan(c *gin.Context) {
	var sr = &ScanRequest{}

	if err := c.ShouldBindJSON(sr); err != nil {
		validator.HttpValidationError(c, err)
	}

	st, err := s.psExec.ScanType(sr.TypeTarget)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"error":   true,
			"status":  400,
			"message": fmt.Sprintf("can not exec scan, err= %v", err),
		})
		log.Errorf("can not exec scan, err= %v", err)
		return
	}

	//args := s.psExec.MakeCmdLineDyn(st, sr.Target)

	output, err := s.psExec.RunScan(st, sr.Target)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"error":   true,
			"status":  500,
			"message": fmt.Sprintf("%v", err),
		})

		log.Errorf("%v", err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"id": output,
	})
	return
}
