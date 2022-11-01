package cmd

import (
	log "github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"time"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length = 10

	procDir = "proc/"

	shell = "#!/bin/bash"
)

type scriptW struct {
	fileName string
	path     string
	shebang  string
	//command  string
}

func NewScriptW() *scriptW {
	//to avoid that an others process access rand
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)

	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return &scriptW{
		fileName: string(b),
		path:     procDir,
		shebang:  shell,
		//command:  cmd,
	}
}

func (s scriptW) Create() (string, error) {
	f := s.path + s.fileName
	// detect if file exists
	fileInfo, err := os.Stat(f)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(f)
		if err != nil {
			log.Infof("can not create file, err= %v", err)
			//os.Exit(0)
			return "", err
		}
		defer file.Close()
	}
	return fileInfo.Name(), err
}

func (s scriptW) Write(content string) error {
	f := s.path + s.fileName
	// open file using READ & WRITE & X permission
	file, err := os.OpenFile(f, os.O_RDWR, 0754) //0666
	if err != nil {
		log.Infof("can not open file, err= %v", err)

		return err
	}
	defer file.Close()

	// write some text to file
	_, err = file.WriteString(s.shebang + "\n" + content)
	if err != nil {
		log.Infof("can not write on file, err= %v", err)
		return err
	}

	// save changes
	err = file.Sync()
	if err != nil {
		log.Infof("can not save content on file, err= %v", err)
		return err
	}

	if err := s.SetPermission(file, 0756); err != nil {
		return err
	}

	return err
}

func (s scriptW) SetPermission(file *os.File, perm os.FileMode) error {
	err := file.Chmod(perm)
	if err != nil {
		log.Infof("can not set permission, err=%v", err)
	}
	return err
}

func (s scriptW) Delete() error {
	file := s.path + s.fileName
	// delete file
	err := os.Remove(file)
	if err != nil {
		log.Infof("can not remove file, err= %v", err)
	}
	return err
}
