package server

import (
	"context"
	"errors"
	"github.com/melbahja/goph"
	"github.com/xenedium/hps_logs_parser/parser"
	protocolBuffer "github.com/xenedium/hps_logs_parser/server/gRPC"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"path"
)

func (s *gRPCServer) SSHParse(ctx context.Context, incomingData *protocolBuffer.SSHRequest) (*protocolBuffer.Response, error) {

	log.Println("Received SSHParse request")

	if incomingData.Host == "" || incomingData.User == "" || incomingData.AbsoluteDir == "" {
		log.Printf("Missing required fields: %v\n", incomingData)
		return nil, errors.New("missing required fields")
	}

	if incomingData.Port == 0 {
		log.Printf("Port not specified, using default port 22\n")
		incomingData.Port = 22
	}

	log.Printf("Incoming data: %v\n", incomingData)
	sshClient, err := GetSSHClientFromGRPC(incomingData)
	if err != nil {
		log.Printf("Failed to get SSH client: %v\n", err)
		return nil, err
	}

	log.Printf("Connecting to %v\n", incomingData.Host)
	sftpClient, err := sshClient.NewSftp()
	if err != nil {
		log.Printf("Failed to get SFTP client: %v\n", err)
		return nil, err
	}

	log.Printf("Creating temp dir\n")
	tempDir, err := os.MkdirTemp(os.TempDir(), "*")
	if err != nil {
		log.Printf("Failed to create temp dir: %v\n", err)
		return nil, err
	}

	log.Printf("Reading remote dir: %v\n", incomingData.AbsoluteDir)
	files, err := sftpClient.ReadDir(incomingData.AbsoluteDir)
	if err != nil {
		log.Printf("Failed to read dir: %v\n", err)
		return nil, err
	}

	for _, file := range files {
		log.Printf("Downloading file: %v\n", file.Name())
		err := sshClient.Download(path.Join(incomingData.AbsoluteDir, file.Name()), path.Join(tempDir, file.Name()))
		if err != nil {
			log.Printf("Failed to download file: %v", err)
		}
	}

	log.Printf("Parsing files from %v\n", incomingData.AbsoluteDir)
	logParser := parser.NewParser(tempDir)
	logParser.Parse(true)

	log.Printf("Returning response")
	return &protocolBuffer.Response{
		Messages: logParser.Messages,
	}, nil
}

func GetSSHClientFromGRPC(data *protocolBuffer.SSHRequest) (*goph.Client, error) {

	if *data.Password != "" {
		return goph.NewConn(&goph.Config{
			User:     data.User,
			Addr:     data.Host,
			Port:     uint(data.Port),
			Auth:     goph.Password(*data.Password),
			Timeout:  goph.DefaultTimeout,
			Callback: ssh.InsecureIgnoreHostKey(),
		})
	} else if *data.PrivateKey != "" {
		// untested code
		keyFile, err := os.CreateTemp(os.TempDir(), "keyfile.pem")
		if err != nil {
			return nil, err
		}
		defer os.Remove(keyFile.Name())
		_, err = keyFile.WriteString(*data.PrivateKey)
		auth, err := goph.Key(path.Join(os.TempDir(), keyFile.Name()), *data.Passphrase)
		if err != nil {
			return nil, err
		}

		return goph.NewConn(&goph.Config{
			User:     data.User,
			Addr:     data.Host,
			Port:     uint(data.Port),
			Auth:     auth,
			Timeout:  goph.DefaultTimeout,
			Callback: ssh.InsecureIgnoreHostKey(),
		})
	} else {
		return nil, errors.New("missing required fields")
	}
}
