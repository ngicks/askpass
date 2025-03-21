package askpass

import (
	"testing"
)

func TestScan(t *testing.T) {
	for _, s := range []string{
		"foo@bar's password: ",
		"Enter passphrase for \"/home/ngicks/.ssh/id_edsa\": ",
	} {
		t.Run(s, func(t *testing.T) {
			Scan(s)
		})
	}
}
