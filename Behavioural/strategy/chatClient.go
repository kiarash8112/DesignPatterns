package strategy

import "fmt"

/*
This class uses an encryption algorithm to encrypt a message before
sending it out.
What are the problems in this implementation? Refactor the code
using the strategy pattern. What are the benefits of the new
implementation?
*/

type chatClient struct {
	encryptionAlgorithm string
}

func newChatClient(alg string) chatClient {
	return chatClient{
		encryptionAlgorithm: alg,
	}
}

func (c chatClient) send(message string) {
	if c.encryptionAlgorithm == "DES" {
		fmt.Println("Encrypting message using DES")
	} else if c.encryptionAlgorithm == "AES" {
		fmt.Println("Encrypting message using AES")
	} else {
		fmt.Println("Unsupported encryption algorithm")
	}

	fmt.Println("Sending the encrypted message...")
}
