package ioutils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nebulaclouds/nebula/nebulastdlib/storage"
)

func TestSimpleInputFilePath_GetInputPath(t *testing.T) {
	s := SimpleInputFilePath{
		pathPrefix: "s3://nebulaclouds-modelbuilder/metadata/propeller/staging/nebulaexamples-development-jf193q0cqo/odd-nums-task/data",
		store:      storage.URLPathConstructor{},
	}

	assert.Equal(t, "s3://nebulaclouds-modelbuilder/metadata/propeller/staging/nebulaexamples-development-jf193q0cqo/odd-nums-task/data/inputs.pb", s.GetInputPath().String())
}
