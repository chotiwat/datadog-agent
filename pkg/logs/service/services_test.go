// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2018 Datadog, Inc.

package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddService(t *testing.T) {
	services := NewServices()
	service := NewService("foo", "1234", Before)

	services.AddService(service)

	added := services.GetAddedServices("foo")
	assert.Equal(t, 0, len(added))

	services.AddService(service)
	assert.Equal(t, 1, len(added))
}

func TestRemoveService(t *testing.T) {
	services := NewServices()
	service := NewService("foo", "1234", Before)

	services.RemoveService(service)

	added := services.GetAddedServices("foo")
	assert.Equal(t, 0, len(added))

	services.RemoveService(service)
	assert.Equal(t, 1, len(added))
}
