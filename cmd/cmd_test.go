package cmd

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGet2faCode(t *testing.T) {
	sharedSecret := "Oz4vd4C4yDTf11JdMtpbY63YI1Q="
	expectedCode := "57NKD"
	testTime := time.Unix(1519494428, 0)
	timestamp := uint64(testTime.Unix())

	actualCode, err := get2faCode(sharedSecret, timestamp)
	assert.NoError(t, err)
	assert.Equal(t, expectedCode, actualCode)
}

func TestGet2faCodeInvalidSecret(t *testing.T) {
	sharedSecret := ""
	testTime := time.Unix(1519494428, 0)
	timestamp := uint64(testTime.Unix())

	actualCode, err := get2faCode(sharedSecret, timestamp)
	assert.EqualError(t, err, "invalid shared secret")
	assert.Empty(t, actualCode)
}

func TestDecodeSecret(t *testing.T) {
	encodedSecret := "Oz4vd4C4yDTf11JdMtpbY63YI1Q="
	expectedSecret := []byte{59, 62, 47, 119, 128, 184, 200, 52, 223, 215, 82, 93, 50, 218, 91, 99, 173, 216, 35, 84}
	actualSecret, err := decodeSecret(encodedSecret)
	assert.NoError(t, err)
	assert.Equal(t, expectedSecret, actualSecret)
}

func TestDecodeSecretInvalidSecret(t *testing.T) {
	cases := []struct {
		name        string
		secretValue string
		err         string
	}{
		{
			name:        "empty shared secret",
			secretValue: "",
			err:         "empty secret",
		},
		{
			name:        "invalid shared secret",
			secretValue: "=",
			err:         "invalid secret",
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			actualSecret, err := decodeSecret(tCase.secretValue)
			assert.EqualError(t, err, tCase.err)
			assert.Nil(t, actualSecret)
		})
	}
}
