package biz

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	pwd := hashPassowrd("abc")
	spew.Dump(pwd)
}

func TestVerifyPassword(t *testing.T) {
	pwd := hashPassowrd("abc")
	res := verifyPassowrd(pwd, "abc")
	assert.True(t, res)
}