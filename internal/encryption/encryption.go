package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/wizedkyle/cvecli/internal/logging"
	"io"
)

func DecryptData(cipherText string) string {
	cipherTextByte, err := hex.DecodeString(cipherText)
	if err != nil {
		logging.ConsoleLogger().Error().Err(err).Msg("failed to convert ciphertext to byte array")
	}
	block, err := aes.NewCipher(hashMachineId(getMachineId()))
	if err != nil {
		logging.ConsoleLogger().Error().Err(err).Msg("failed to generate cipher block")
	}
	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		logging.ConsoleLogger().Error().Err(err).Msg("failed to wrap block")
	}
	nonceSize := aesGcm.NonceSize()
	nonce, encryptedText := cipherTextByte[:nonceSize], cipherTextByte[nonceSize:]
	plaintext, err := aesGcm.Open(nil, nonce, encryptedText, nil)
	if err != nil {
		logging.ConsoleLogger().Error().Err(err).Msg("failed to decrypt ciphertext")
	}
	return string(plaintext)
}

func EncryptData(secret string) string {
	plaintext := []byte(secret)
	block, err := aes.NewCipher(hashMachineId(getMachineId()))
	if err != nil {
		logging.ConsoleLogger().Error().Err(err).Msg("failed to generate cipher block")
	}
	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		logging.ConsoleLogger().Error().Err(err).Msg("failed to wrap block")
	}
	nonce := make([]byte, aesGcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		logging.ConsoleLogger().Error().Err(err).Msg("failed to generate nonce")
	}
	cipherText := aesGcm.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", cipherText)
}

func hashMachineId(machineId string) []byte {
	hash := sha256.New()
	hash.Write([]byte(machineId))
	computedHash := hash.Sum(nil)
	return computedHash
}
