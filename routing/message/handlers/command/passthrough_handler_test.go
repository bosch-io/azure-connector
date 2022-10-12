// Copyright (c) 2022 Contributors to the Eclipse Foundation
//
// See the NOTICE file(s) distributed with this work for additional
// information regarding copyright ownership.
//
// This program and the accompanying materials are made available under the
// terms of the Eclipse Public License 2.0 which is available at
// https://www.eclipse.org/legal/epl-2.0, or the Apache License, Version 2.0
// which is available at https://www.apache.org/licenses/LICENSE-2.0.
//
// SPDX-License-Identifier: EPL-2.0 OR Apache-2.0

package command

import (
	"io"
	"log"
	"testing"

	"github.com/eclipse-kanto/azure-connector/config"
	"github.com/eclipse-kanto/suite-connector/connector"
	"github.com/eclipse-kanto/suite-connector/logger"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPassthroughMessageHandler(t *testing.T) {
	settings := &config.AzureSettings{
		ConnectionString:        "HostName=dummy-hub.azure-devices.net;DeviceId=dummy-device;SharedAccessKey=dGVzdGF6dXJlc2hhcmVkYWNjZXNza2V5",
		PassthroughCommandTopic: "testCommand",
	}
	logger := logger.NewLogger(log.New(io.Discard, "", log.Ldate), logger.INFO)
	connSettings, err := config.PrepareAzureConnectionSettings(settings, nil, logger)
	require.NoError(t, err)
	messageHandler := &commandPassthroughMessageHandler{}

	require.NoError(t, messageHandler.Init(settings, connSettings))
	assert.Equal(t, messageHandler.passthroughCommandTopic, settings.PassthroughCommandTopic)
	assert.Equal(t, commandPassthroughHandlerName, messageHandler.Name())
	assert.Equal(t, []string{settings.PassthroughCommandTopic}, messageHandler.Topics())
}

func TestHandleCommand(t *testing.T) {
	settings := &config.AzureSettings{
		ConnectionString:        "HostName=dummy-hub.azure-devices.net;DeviceId=dummy-device;SharedAccessKey=dGVzdGF6dXJlc2hhcmVkYWNjZXNza2V5",
		PassthroughCommandTopic: "testCommand",
	}
	logger := logger.NewLogger(log.New(io.Discard, "", log.Ldate), logger.INFO)
	connSettings, err := config.PrepareAzureConnectionSettings(settings, nil, logger)
	require.NoError(t, err)
	messageHandler := &commandPassthroughMessageHandler{}
	require.NoError(t, messageHandler.Init(settings, connSettings))

	azureMessages, err := messageHandler.HandleMessage(&message.Message{Payload: []byte("dummy_payload")})
	require.NoError(t, err)

	azureMsg := azureMessages[0]
	azureMsgTopic, _ := connector.TopicFromCtx(azureMsg.Context())

	assert.Equal(t, "testCommand", azureMsgTopic)
	assert.Equal(t, "dummy_payload", string(azureMsg.Payload))
}
