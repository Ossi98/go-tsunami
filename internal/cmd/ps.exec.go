package cmd

import (
	"bytes"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Proc struct {
	cmdName    string
	classpath  string //classpath
	psPath     string
	jarName    string
	pluginPath string
	outputType string
	outputFile string
}

func NewProcessScan(v *viper.Viper) *Proc {
	return &Proc{
		cmdName:    "java", //usr/bin/java
		classpath:  "-cp",
		psPath:     fmt.Sprintf("%v", v.GetString("cmd.tsunami.path")),
		pluginPath: fmt.Sprintf("%v", v.GetString("cmd.tsunami.plugin")) + "/*",
		jarName:    fmt.Sprintf("%v", v.GetString("cmd.tsunami.jar")), // will be searched by regex in future
		outputType: fmt.Sprintf("%v", v.GetString("cmd.tsunami.output-type")),
		outputFile: fmt.Sprintf("%v", v.GetString("cmd.tsunami.output-file")),
	}
}

func (p *Proc) getCMD() string {
	return p.cmdName
}

func (p *Proc) getArgsBase() string {
	return fmt.Sprintf("%s \"%s:%s\" com.google.tsunami.main.cli.TsunamiCli",
		p.classpath,
		p.jarName,
		p.pluginPath,
	) //-Dtsunami-config.location=/home/c-tsunami/tsunami/tsunami.yaml
}

func (p *Proc) getOutputType() string {
	return p.outputType
}

func (p *Proc) getOutputFile() string {
	return p.outputFile
}

func (p *Proc) ScanType(typeScan string) (string, error) {
	switch typeScan {
	case "ipv4":
		return IpV4Target, nil
	case "ipv6":
		return IpV6Target, nil
	case "hostname":
		return HostnameTarget, nil
	case "uri":
		return UriTarget, nil
	default:
		return "", errors.New("no target scan match")
	}

}

func (p *Proc) scanTarget(typeScan, target string) string {
	return typeScan + "=" + target
}

// MakeCmdLineDynPt create dynamic part of the cmd tsunami
func (p *Proc) makeCmdLineDynPt(cmd []string) string {
	return strings.Join(cmd, " ")
}

func (p *Proc) RunScan(typeScan, target string) (string, error) {

	writer := NewScriptW()

	/*fName, err := writer.Create()
	if err != nil {
		return "", err
	}*/

	var cmd = make([]string, 2, 4)

	cmdName := p.cmdName
	classpathArgs := p.getArgsBase()
	targetArg := p.scanTarget(typeScan, target)

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	tmpFileName := make([]byte, length)

	for i := range tmpFileName {
		tmpFileName[i] = charset[seededRand.Intn(len(charset))]
	}

	output := ScanResOutputFormat + "=" + p.outputType + " " + ScanResOutputFilename + "=" + p.outputFile + string(tmpFileName) + ".json"

	cmd = append(cmd, cmdName, classpathArgs, targetArg, output)

	command := p.makeCmdLineDynPt(cmd)

	fName, err := writer.CreateAndWrite(command) //fName, err := writer.Create()
	if err != nil {
		return "", err
	}

	id := fName
	/*if err := writer.Write(command); err != nil {
		return "", err
	}*/

	ps := exec.Command("./proc/" + fName)

	var out, er bytes.Buffer

	ps.Stdout = &out
	ps.Stderr = &er

	err = ps.Run()
	//err := ps.Start()
	if err != nil {

		return "", errors.New(fmt.Sprintf("process error, Stderr=%v, err=%v", er.String(), err))
	}

	ps.Wait()

	writer.Delete()

	//rename result scan.json to file tle fname

	if err := os.Rename(p.outputFile+string(tmpFileName), p.outputFile+fName+".json"); err != nil {
		log.Infof("can not rename file, err=%v", err)
		id = string(tmpFileName)
	}

	log.Println(out.String())

	return id, nil

}
