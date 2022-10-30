package cmd

import (
	"bytes"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os/exec"
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
		cmdName:    "java", //fmt.Sprintf("%v", v.GetString("cmd.tsunami."))
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
	return fmt.Sprintf("%s %s:%s com.google.tsunami.main.cli.TsunamiCli", p.classpath, p.jarName, p.pluginPath)
}

func (p *Proc) getOutputType() string {
	return p.outputType
}

func (p *Proc) getOutputFile() string {
	return p.outputFile
}

func (p *Proc) ScanType(typeScan string) (string, error) {
	log.Infof("test=%s", typeScan)

	switch typeScan {
	case "ipv4":
		return IpV4Target, nil
	case "ipv6":
		return IpV6Target, nil
	case "hostname":
		return HostnameTarget, nil
	case "uri":
		return UriTarget, nil

		/*default:
		return "", errors.New("no target scan match")*/
	}
	return "", errors.New("no target scan match")
}

func (p *Proc) MakeCmdLine(typeScan, target string) []string {
	tgt := typeScan + "=" + target
	output := ScanResOutputFormat + "=" + p.outputType
	output += " " + ScanResOutputFilename + "=" + p.outputFile

	var args = make([]string, 2, 4)
	args = append(args, tgt, output)

	return args
}

func (p *Proc) RunScan(cArgs ...string) (string, error) {
	args := p.getArgsBase()

	for _, v := range cArgs {
		args += " " + v
	}

	ps := exec.Command(p.getCMD(), args)

	var out bytes.Buffer
	ps.Stdout = &out

	err := ps.Run()
	if err != nil {
		errors.New(fmt.Sprintf(" process error, err=%v", err))
		return "", err
	}

	//notify the interface with a websocket to preform get results of scan

	log.Println(out.String())

	return out.String(), nil

}
