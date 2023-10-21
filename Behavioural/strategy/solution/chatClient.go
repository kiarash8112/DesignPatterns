package solution

import "fmt"

type encryptionAlgorithm interface {
	encrypt()
}

var _ encryptionAlgorithm = DESEncryptionAlgorithm{}

type DESEncryptionAlgorithm struct {
}

func (DESEncryptionAlgorithm) encrypt() {
	fmt.Println("Encrypting message using DES")
}

type chatClient struct {
	encryptionAlgorithm string
}

func newChatClient(alg string) chatClient {
	return chatClient{
		encryptionAlgorithm: alg,
	}
}

func (c chatClient) send(message string, encAlg encryptionAlgorithm) {
	encAlg.encrypt()
	fmt.Println("Sending the encrypted message...")
}
