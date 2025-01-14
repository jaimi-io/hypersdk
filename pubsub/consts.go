// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package pubsub

import (
	"time"

	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/jaimi-io/hypersdk/consts"
)

const (
	ReadBufferSize      = 64 * units.KiB
	WriteBufferSize     = 64 * units.KiB
	WriteWait           = 10 * time.Second
	PongWait            = 60 * time.Second
	PingPeriod          = (PongWait * 9) / 10
	MaxReadMessageSize  = consts.NetworkSizeLimit
	MaxWriteMessageSize = 16 * units.MiB
	MaxMessageWait      = 50 * time.Millisecond
	MaxPendingMessages  = 1024
)
