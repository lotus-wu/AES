package xiao

import (
	"github.com/OpenWhiteBox/AES/primitives/matrix"
)

// GenerateEncryptionKeys creates a white-boxed version of the AES key `key` for encryption, with any non-determinism
// generated by `seed`.
func GenerateEncryptionKeys(key, seed []byte) (out Construction, inputMask, outputMask matrix.Matrix) {
	panic("Encryption key generation isn't currently implemented!")
}

// GenerateDecryptionKeys creates a white-boxed version of the AES key `key` for decryption, with any non-determinism
// generated by `seed`.
func GenerateDecryptionKeys(key, seed []byte) (out Construction, inputMask, outputMask matrix.Matrix) {
	panic("Decryption key generation isn't currently implemented!")
}