package auth 

import (
	"testing"
	"os"
)


var jwtm = NewJWTManager(os.Getenv("JWT_SECRET"))

func TestGenerate(t *testing.T) {
	t.Log(jwtm)
}