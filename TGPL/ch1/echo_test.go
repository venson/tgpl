package main

import (
	"os"
	"testing"
)

func BenchmarkEcho(b *testing.B) {
	setup()
	b.StartTimer()
	echo()
}
func BenchmarkEchoJoin(b *testing.B) {
	setup()
	b.StartTimer()
	echojoin()
}
func setup() {
	os.Args = append(os.Args, "--addr=http://b.com:566/something.avsc")
	os.Args = append(os.Args, "Get")
	os.Args = append(os.Args, `./some/resource/fred`)
}
