package askpass

import (
	"fmt"
)

var (
	onesPassword               = "%s@%s's passowrd: "
	enterPassphraseForFilename = "Enter passphrase for \"%s\": "
)

func Scan(s string) (string, error) {
	var user, host string
	n, err := fmt.Sscanf(s, onesPassword, &user, &host)
	if err != nil {
		fmt.Printf("pass word pat: user = %s, host = %s, n = %d, err = %v\n", user, host, n, err)
	}
	var filename string
	n, err = fmt.Sscanf(s, enterPassphraseForFilename, &filename)
	if err != nil {
		fmt.Printf("passphrase pat: filename = %s, n = %d, err = %v\n", filename, n, err)
	}
	return filename, err
}
