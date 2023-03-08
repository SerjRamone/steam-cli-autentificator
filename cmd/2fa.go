package cmd

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type maFile struct {
	SharedSecret string `json:"shared_secret"`
}

var IsInfinite bool

// 2faCmd represents the 2fa command
var _2faCmd = &cobra.Command{
	Use:   "2fa /path/to/*.maFile",
	Short: "Generate 2fa code (Steam Guard) by maFile",
	Run:   commandHandler,
}

// commandHandler is 2fa command handler
func commandHandler(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatal("you didn't specify the additional arguments\nUse the --help flag for comprehensive help on how to use this tool")
	}
	maFilePath := args[0]

	if _, err := os.Stat(maFilePath); os.IsNotExist(err) {
		fmt.Printf("maFile does not exist by path '%s'\n", maFilePath)
		return
	}

	maFileData, err := os.ReadFile(maFilePath)
	if err != nil {
		fmt.Printf("error reading maFile by path '%s': %v\n", maFilePath, err)
		return
	}

	var mf maFile
	err = json.Unmarshal(maFileData, &mf)
	if err != nil {
		fmt.Printf("error parsing maFile by path '%s': %v\n", maFilePath, err)
		return
	}

	sharedSecret := mf.SharedSecret
	if sharedSecret == "" {
		fmt.Printf("maFile by path '%s' does not contain a shared_secret\n", maFilePath)
		return
	}

	var s *spinner.Spinner
	if IsInfinite {
		s = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Start()
	}

	decorator := color.New(color.FgCyan, color.Bold)
	for ok := true; ok; ok = IsInfinite {
		t := uint64(time.Now().Unix())

		code, err := get2faCode(sharedSecret, t)
		if err != nil {
			fmt.Printf("2fa code generating error: %v\n", err)
			return
		}

		if IsInfinite {
			s.Suffix = fmt.Sprintf(" Steam Guard code is: %s\n", decorator.Sprintf("%s", code))
			time.Sleep(1 * time.Second)
		} else {
			fmt.Printf("Steam Guard code is: %s\n", decorator.Sprintf("%s", code))
		}
	}
}

// get2faCode generate Steam Guard authorization
func get2faCode(sharedSecret string, t uint64) (string, error) {
	key, err := decodeSecret(sharedSecret)
	if err != nil {
		return "", errors.New("invalid shared secret")
	}

	var (
		// Range of possible chars for auth code.
		codeChars = []byte{
			// 2, 3, 4, 5, 6, 7, 8, 9, B, C, D, F, G
			50, 51, 52, 53, 54, 55, 56, 57, 66, 67, 68, 70, 71,
			// H, J, K, M, N, P, Q, R, T, V, W, X, Y
			72, 74, 75, 77, 78, 80, 81, 82, 84, 86, 87, 88, 89,
		}

		codeCharsLen = len(codeChars)
	)

	t /= 30                           // converting time for any reason
	tb := make([]byte, 8)             // 00 00 00 00 00 00 00 00
	binary.BigEndian.PutUint64(tb, t) // 00 00 00 00 xx xx xx xx

	// evaluate hash code for `tb` by key
	mac := hmac.New(sha1.New, key)
	mac.Write(tb)
	hashcode := mac.Sum(nil)

	// last 4 bits provide initial position
	// len(hashcode) = 20 bytes
	start := hashcode[19] & 0xf

	// extract 4 bytes at `start` and drop first bit
	fc32 := binary.BigEndian.Uint32(hashcode[start : start+4])
	fc32 &= 1<<31 - 1
	fullcode := int(fc32)

	// generate auth code
	code := make([]byte, 5)
	for i := range code {
		code[i] = codeChars[fullcode%codeCharsLen]
		fullcode /= codeCharsLen
	}

	return string(code[:]), nil
}

// decodeSecret decode shared secret string from base64 to byte slice
func decodeSecret(encodedSecret string) ([]byte, error) {
	if encodedSecret == "" {
		return nil, errors.New("empty secret")
	}

	bytes, err := base64.StdEncoding.DecodeString(encodedSecret)
	if err != nil {
		return nil, errors.New("invalid secret")
	}

	return bytes, nil
}

func init() {
	_2faCmd.Flags().BoolVarP(&IsInfinite, "infinite", "i", false, "infinite generating")
	rootCmd.AddCommand(_2faCmd)
}
