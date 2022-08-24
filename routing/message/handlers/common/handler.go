// Copyright (c) 2022 Contributors to the Eclipse Foundation
//
// See the NOTICE file(s) distributed with this work for additional
// information regarding copyright ownership.
//
// This program and the accompanying materials are made available under the
// terms of the Eclipse Public License 2.0 which is available at
// http://www.eclipse.org/legal/epl-2.0
//
// SPDX-License-Identifier: EPL-2.0

package command

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/eclipse-kanto/azure-connector/config"
)

// MessageHandler represents the internal interface for implementing a Watermill message handler.
type MessageHandler interface {
	Init(settings *config.AzureSettings, connSettings *config.AzureConnectionSettings) error
	HandleMessage(message *message.Message) ([]*message.Message, error)
	Name() string
	Topics() []string
}
