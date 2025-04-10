package main

import (
	"io"
	"log"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {
	// SSH config
	config := &ssh.ClientConfig{
		User: "sftp-user1",
		Auth: []ssh.AuthMethod{
			ssh.Password("kumar"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to SSH server
	client, err := ssh.Dial("tcp", "44.202.50.230:22", config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	defer client.Close()
	log.Println("SSH connection established")

	// Create SFTP client
	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		log.Fatal("Failed to create SFTP client: ", err)
	}
	defer sftpClient.Close()
	log.Println("SFTP session established")

	// Open local file
	localFile, err := os.Open("README.md") // Change to your local file path
	if err != nil {
		log.Fatal("Failed to open local file: ", err)
	}
	defer localFile.Close()
	log.Println("opened", localFile.Name())

	// Create destination file on the remote server
	remoteFile, err := sftpClient.Create("/potta/suthu/uploadedfile.txt") // Adjust the remote path
	if err != nil {
		log.Fatal("Failed to create remote file: ", err)
	}
	defer remoteFile.Close()

	// Copy local file to remote file
	bytes, err := io.Copy(remoteFile, localFile)
	if err != nil {
		log.Fatal("Failed to copy file: ", err)
	}

	log.Printf("File transferred successfully (%d bytes)\n", bytes)
}
