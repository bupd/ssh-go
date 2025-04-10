package main

import (
	"log"
	"time"

	"golang.org/x/crypto/ssh"
)

func main() {
	config := &ssh.ClientConfig{
		User: "sftp-user1",
		Auth: []ssh.AuthMethod{
			ssh.Password("kumar"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", "44.202.50.230:22", config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}

	log.Println("we have a client")
	time.Sleep(5 * time.Second)

	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	log.Println("we have a session")
	defer session.Close()
}
