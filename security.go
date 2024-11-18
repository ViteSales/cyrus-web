// Copyright 2017 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package web

import (
	"github.com/vitesales/cyrus-pool/h"
	"github.com/vitesales/cyrus/models/security"
)

func init() {
	h.Filter().Methods().AllowAllToGroup(security.GroupEveryone)
}
