package arquiteture

import (
	"runtime"
	"testing"
)

//Verificar se a arquitetura suporta amd64
func TesteDependente(t *testing.T) {
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitetura md64")
	}

	// ...
	t.Fail()
}