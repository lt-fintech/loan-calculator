package service

import (
	"testing"
)

func TestCreateContract(t *testing.T) {
	contract := generateContract(100)
	t.Log(contract)
}
