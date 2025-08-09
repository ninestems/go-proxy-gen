package log

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func captureOutput(f func()) (stdout, stderr string) {
	var outBuf, errBuf bytes.Buffer

	rOut, wOut, _ := os.Pipe()
	rErr, wErr, _ := os.Pipe()

	origStdout := os.Stdout
	origStderr := os.Stderr
	os.Stdout = wOut
	os.Stderr = wErr

	SetWriters(wOut, wErr)

	done := make(chan struct{})
	go func() {
		_, _ = outBuf.ReadFrom(rOut)
		_, _ = errBuf.ReadFrom(rErr)
		close(done)
	}()

	f()

	_ = wOut.Close()
	_ = wErr.Close()

	os.Stdout = origStdout
	os.Stderr = origStderr

	SetWriters(os.Stdout, os.Stderr)

	<-done
	return outBuf.String(), errBuf.String()
}

func TestInfoLogLevelInfo(t *testing.T) {
	SetLevel("info")
	stdout, _ := captureOutput(func() {
		Info("test info")
	})
	require.Contains(t, stdout, "test info")
}

func TestDebugLogLevelDebug(t *testing.T) {
	SetLevel("debug")
	stdout, _ := captureOutput(func() {
		Debug("test debug")
	})
	require.Contains(t, stdout, "test debug")
}

func TestDebugLogLevelInfo(t *testing.T) {
	SetLevel("info")
	stdout, _ := captureOutput(func() {
		Debug("should not appear")
	})
	require.NotContains(t, stdout, "should not appear")
}

func TestErrorLog(t *testing.T) {
	_, stderr := captureOutput(func() {
		Error("test error")
	})
	require.Contains(t, stderr, "test error")
}

func TestInfofLogLevelInfo(t *testing.T) {
	SetLevel("info")
	stdout, _ := captureOutput(func() {
		Infof("formatted %s", "info")
	})
	require.Contains(t, stdout, "formatted info")
}

func TestDebugfLogLevelDebug(t *testing.T) {
	SetLevel("debug")
	stdout, _ := captureOutput(func() {
		Debugf("formatted %s", "debug")
	})
	require.Contains(t, stdout, "formatted debug")
}

func TestDebugfLogLevelInfo(t *testing.T) {
	SetLevel("info")
	stdout, _ := captureOutput(func() {
		Debugf("should not %s", "appear")
	})
	require.NotContains(t, stdout, "should not appear")
}

func TestErrorf(t *testing.T) {
	_, stderr := captureOutput(func() {
		Errorf("formatted %s", "error")
	})
	require.Contains(t, stderr, "formatted error")
}

func TestSetLevelCaseInsensitive(t *testing.T) {
	SetLevel("DeBuG")
	stdout, _ := captureOutput(func() {
		Debug("mixed case debug")
	})
	require.Contains(t, stdout, "mixed case debug")
}
