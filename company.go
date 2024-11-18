// Copyright 2018 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package web

import (
	"github.com/vitesales/cyrus-pool/h"
	"github.com/vitesales/cyrus/models"
	"github.com/vitesales/cyrus/models/fields"
)

var fields_Company = map[string]models.FieldDefinition{
	"DashboardBackground": fields.Binary{},
}

func init() {
	h.Company().AddFields(fields_Company)
}
