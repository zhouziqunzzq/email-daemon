package main

import "testing"

func TestHTTPServer(t *testing.T) {
	t.Logf("Running HTTP Server for testing, press Ctrl-C to stop it")
	HTTPServer()
}
