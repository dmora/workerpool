package workerpool_test

import (
	"os"
	"testing"

	log "github.com/spf13/jwalterweatherman"
)

// TestMain setups main testing package setup and teardown
func TestMain(m *testing.M) {
	log.SetStdoutThreshold(log.LevelDebug)
	exitCode := m.Run()
	// Tear Down
	log.DEBUG.Print("Testing ending")

	// Exit
	os.Exit(exitCode)
}
