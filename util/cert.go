package util

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mundipagg/boleto-api/config"
)

//ListCert Lista os Certificados necessários e chama o método que faz a cópia
func ListCert() (string, error) {

	list := []string{
		config.Get().CertBoletoPathCrt,
		config.Get().CertBoletoPathKey,
		config.Get().CertBoletoPathCa,
		config.Get().CertICP_PathPkey,
		config.Get().CertICP_PathChainCertificates,
	}

	var err error
	var res string

	for _, v := range list {

		res, err = copyCert(v)

		if err != nil {
			return "", err
		}

	}

	return res, nil

}

func copyCert(c string) (string, error) {
	execPath, _ := os.Getwd()

	f := strings.Split(c, "/")

	fName := f[len(f)-1]

	srcFile, err := os.Open(execPath + "/boleto_orig/" + fName)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return "", err
	}
	defer srcFile.Close()

	destFile, err := os.Create(c)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return "", err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return "", err
	}

	err = destFile.Sync()
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return "", err
	}

	err = os.Chmod(c, 0777)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return "", err
	}

	res := fmt.Sprintf("Certificate Copy Sucessful: %s", c)

	return res, nil
}