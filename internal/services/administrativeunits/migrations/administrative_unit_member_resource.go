// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package migrations

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/administrativeunits/parse"
)

func ResourceAdministrativeUnitMemberInstanceResourceV0() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Schema: map[string]*pluginsdk.Schema{
			"administrative_unit_object_id": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},

			"member_object_id": {
				Type:     pluginsdk.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceAdministrativeUnitMemberInstanceStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	log.Println("[DEBUG] Migrating ID from v0 to v1 format")
	oldId, err := parse.AdministrativeUnitMemberID(rawState["id"].(string))
	if err != nil {
		return rawState, fmt.Errorf("parsing ID for `azuread_administrative_unit_role_member`: %+v", err)
	}

	newId := stable.NewDirectoryAdministrativeUnitIdMemberID(oldId.AdministrativeUnitId, oldId.MemberId)
	rawState["id"] = newId.ID()
	return rawState, nil
}
