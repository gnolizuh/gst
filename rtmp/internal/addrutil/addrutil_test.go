//
//
// Tencent is pleased to support the open source community by making tRPC available.
//
// Copyright (C) 2023 THL A29 Limited, a Tencent company.
// All rights reserved.
//
// If you have downloaded a copy of the tRPC source code from Tencent,
// please note that tRPC source code is licensed under the  Apache 2.0 License,
// A copy of the Apache 2.0 License is included in this file.
//
//

package addrutil_test

import (
	"github.com/gnolizuh/trpc-go-multimedia/rtmp/internal/addrutil"
	"github.com/stretchr/testify/require"
	"net"
	"testing"
)

func TestAddrToKey(t *testing.T) {
	laddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:10000")
	require.Nil(t, err)
	raddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:10001")
	require.Nil(t, err)
	key := addrutil.AddrToKey(laddr, raddr)
	require.Equal(t, key, laddr.Network()+"_"+laddr.String()+"_"+raddr.String())
}
