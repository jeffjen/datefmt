package datefmt

import (
	"fmt"
	"testing"
	"time"
)

func TestStrftime(t *testing.T) {
	timestamp, err := Strftime(time.Now(), "%T")
	if err != nil {
		t.Errorf("failed to format current time: %s", err)
	} else {
		fmt.Println(timestamp)
	}
}

func TestStrptime(t *testing.T) {
	d, err := Strptime("16:34:55", "%T")
	if err != nil {
		t.Errorf("failed to parse time: %s", err)
	} else {
		fmt.Println(d)
	}
}
