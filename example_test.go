package noeq

import (
	"testing"
)

import (
	// "github.com/bmizerany/noeq.go"
	"log"
)

func TestExample(t *testing.T) {
	fl, err := New("", "localhost:4444")
	if err != nil {
		t.Fatal(err)
	}

	for n := 0; n < 100; n++ {
		id, err := genOne(fl, 3)
		if err != nil {
			log.Fatal("FATAL: unable to connect to weiqid")
		}
		log.Printf("%d", id)
	}
}

// The client attempts to give the user as much insight as possible
// It will *not* automaticly attempt a reconnect
// on error. However, it will attempt a reconnect if you ask for
// another id after an error has occured. This gives the user more
// control over how/when to attempt a reconnect.
func genOne(fl *Client, trys int) (id uint64, err error) {
	for ; trys > 0; trys-- {
		id, err = fl.GenOne()
		if err != nil {
			log.Println("ERROR:", err)
			continue
		}

		return
	}

	return
}
