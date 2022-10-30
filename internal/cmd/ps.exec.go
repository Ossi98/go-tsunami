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
		cmdName:    "/usr/bin/java", //fmt.Sprintf("%v", v.GetString("cmd.tsunami."))
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
	return fmt.Sprintf("%s \"%s:%s\" com.google.tsunami.main.cli.TsunamiCli -Dtsunami-config.location=/home/c-tsunami/tsunami/tsunami.yaml", p.classpath, p.jarName, p.pluginPath)
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
	log.Infof("%s", fmt.Sprintf(args))
	//ps := exec.Command(p.getCMD(), args)
	/*ps := exec.Command("java",
		"-cp",
		"\"/home/c-tsunami/tsunami/tsunami-main-0.0.15-SNAPSHOT-cli.jar:/home/c-tsunami/tsunami/plugins/*\"",
		"com.google.tsunami.main.cli.TsunamiCli", //<-there is the prob
		"--ip-v4-target=127.0.0.1",
		"--scan-results-local-output-format=JSON",
		"--scan-results-local-output-filename=/home/c-tsunami/tsunami/tsunami-output.json",
	)*/
	ps := exec.Command("./script-tsunami.sh")

	var out, er bytes.Buffer

	ps.Stdout = &out
	ps.Stderr = &er

	err := ps.Run()
	if err != nil {

		return "", errors.New(fmt.Sprintf(" process error, Stderr=%v, err=%v", er.String(), err))
	}

	//notify the interface with a websocket to preform get results of scan

	log.Println(out.String())

	return out.String(), nil

}
