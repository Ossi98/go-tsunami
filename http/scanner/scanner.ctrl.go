package scanner

import (
	"Ossi98/go-tsunami/internal/cmd"
	"Ossi98/go-tsunami/internal/utils/validator"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"os"
)

type Scanner struct {
	psExec *cmd.Proc
	viper  *viper.Viper
}

func NewScannerCtrl(ps *cmd.Proc, vp *viper.Viper) *Scanner {
	return &Scanner{
		psExec: ps,
		viper:  vp,
	}
}

type scanRequest struct {
	Target     string `json:"target" binding:"required"` // verify if address valid in case of ipv4,6 or domain in case of hostname or uri
	TypeTarget string `json:"type" binding:"required"`
}

func (s *Scanner) StartScan(c *gin.Context) {
	var sr = &scanRequest{}

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

type scanReadURI struct {
	Id string `uri:"id" binding:"required"`
}

func (s *Scanner) ReadScanFile(c *gin.Context) {
	var uri = &scanReadURI{}

	if err := c.ShouldBindUri(uri); err != nil {
		log.Errorf("uri error %s", err)
		return
	}

	// Open our jsonFile
	jsonFile, err := os.Open(viper.GetString("cmd.tsunami.path") + uri.Id + ".json")

	// if we os.Open returns an error then handle it
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"error":   true,
			"status":  500,
			"message": fmt.Sprintf("%v", err),
		})
		log.Errorf("can not open file, err=%v", err)
		return
	}

	log.Info("Successfully Opened users.json")

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal(byteValue, &result)

	fmt.Println(result)

	c.Data(http.StatusOK, "application/json", byteValue)
	/*c.JSON(http.StatusOK, map[string]string{
		"id": uri.Id,
	})*/
	return
}
