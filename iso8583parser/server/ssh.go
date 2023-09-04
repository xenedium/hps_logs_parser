package server

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/melbahja/goph"
	"github.com/xenedium/hps_logs_parser/iso8583parser/parser"
	"golang.org/x/crypto/ssh"
	"os"
	"path"
)

type SSHEndpointIncomingData struct {
	Host        string `json:"host"`
	Port        uint   `json:"port"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	PrivateKey  string `json:"private_key"`
	Passphrase  string `json:"passphrase"`
	AbsoluteDir string `json:"absolute_dir"`
}

// SSHEndpoint is a handler for POST /api/v1/ssh endpoint
// Request body:
//
//	{
//		"host": "",
//		"port": "",
//		"username": "",
//		"password": "",
//	 	"private_key": "",
//		"passphrase": "",
//		"absolute_dir": ""
//	}
//
// (password) and (private_key, passphrase) are mutually exclusive
func SSHEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		var incomingData SSHEndpointIncomingData

		if err := c.ShouldBindJSON(&incomingData); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			return
		}

		if incomingData.Host == "" || incomingData.Username == "" || incomingData.AbsoluteDir == "" {
			c.AbortWithStatusJSON(400, gin.H{"error": "missing required fields"})
			return
		}

		if incomingData.Port == 0 {
			incomingData.Port = 22
		}

		sshClient, err := GetSSHClient(incomingData)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}

		sftpClient, err := sshClient.NewSftp()
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}
		tempDir, err := os.MkdirTemp(os.TempDir(), "*")
		if err != nil {
			fmt.Println(err)
			return
		}
		files, err := sftpClient.ReadDir(incomingData.AbsoluteDir)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}

		for _, file := range files {
			err := sshClient.Download(path.Join(incomingData.AbsoluteDir, file.Name()), path.Join(tempDir, file.Name()))
			if err != nil {
				fmt.Fprint(gin.DefaultErrorWriter, err.Error())
			}
		}

		logParser := parser.NewParser(tempDir)
		logParser.Parse(true)

		c.JSON(200, gin.H{"messages": logParser.Messages})
	}
}

func GetSSHClient(data SSHEndpointIncomingData) (*goph.Client, error) {

	if data.Password != "" {
		return goph.NewConn(&goph.Config{
			User:     data.Username,
			Addr:     data.Host,
			Port:     data.Port,
			Auth:     goph.Password(data.Password),
			Timeout:  goph.DefaultTimeout,
			Callback: ssh.InsecureIgnoreHostKey(),
		})
	} else if data.PrivateKey != "" {
		// untested code
		keyFile, err := os.CreateTemp(os.TempDir(), "keyfile.pem")
		if err != nil {
			return nil, err
		}
		defer os.Remove(keyFile.Name())
		_, err = keyFile.WriteString(data.PrivateKey)
		auth, err := goph.Key(path.Join(os.TempDir(), keyFile.Name()), data.Passphrase)
		if err != nil {
			return nil, err
		}

		return goph.NewConn(&goph.Config{
			User:     data.Username,
			Addr:     data.Host,
			Port:     data.Port,
			Auth:     auth,
			Timeout:  goph.DefaultTimeout,
			Callback: ssh.InsecureIgnoreHostKey(),
		})
	} else {
		return nil, errors.New("missing required fields")
	}
}
