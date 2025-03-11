// Copyright © 2022 Weald Technology Trading
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package conn_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wealdtech/ethereal/v2/conn"
)

// TestConnBad tests bad connection creation.
func TestConnBad(t *testing.T) {
	ctx := context.Background()
	_, err := conn.New(ctx, "Bad", false, false)
	require.NotNil(t, err)
}

// TestConn tests bad connection creation.
func TestConn(t *testing.T) {
	if os.Getenv("EXECUTION_URL") == "" {
		t.Skip("EXECUTION_URL not supplied; test not running")
	}

	ctx := context.Background()
	_, err := conn.New(ctx, os.Getenv("EXECUTION_URL"), false, false)
	require.Nil(t, err)
}
