//go:build !android && !ios && !windows

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */

package device

import "github.com/tailscale/wireguard-go/conn"

const (
	QueueStagedSize            = conn.IdealBatchSize
	QueueOutboundSize          = 1024
	QueueInboundSize           = 1024
	QueueHandshakeSize         = 1024
	MaxSegmentSize             = 2500 // largest possible UDP datagram
	PreallocatedBuffersPerPool = 512  // Disable and allow for infinite memory growth
)
