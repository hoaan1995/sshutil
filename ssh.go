package sshutil

import (
	"golang.org/x/crypto/ssh"
)

func Connect(ip, port, user, password string) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	connection, err := ssh.Dial("tcp", ip+":"+port, config)
	if err != nil {
		return nil, err
	}

	return connection, nil
}

func RunCommand(connection *ssh.Client, command string) error {
	session, err := connection.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	err = session.Run(command)
	if err != nil {
		return err
	}

	return nil
}

func CloseConnection(connection *ssh.Client) {
	connection.Close()
}
