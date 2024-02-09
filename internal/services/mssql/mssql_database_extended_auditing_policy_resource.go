// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssql

import (
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v5.0/sql" // nolint: staticcheck
	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/terraform-provider-azurerm/helpers/tf"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/mssql/parse"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
	"github.com/hashicorp/terraform-provider-azurerm/internal/timeouts"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
)

func resourceMsSqlDatabaseExtendedAuditingPolicy() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Create: resourceMsSqlDatabaseExtendedAuditingPolicyCreateUpdate,
		Read:   resourceMsSqlDatabaseExtendedAuditingPolicyRead,
		Update: resourceMsSqlDatabaseExtendedAuditingPolicyCreateUpdate,
		Delete: resourceMsSqlDatabaseExtendedAuditingPolicyDelete,

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.DatabaseExtendedAuditingPolicyID(id)
			return err
		}),

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(30 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(30 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"database_id": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: commonids.ValidateSqlDatabaseID,
			},

			"enabled": {
				Type:     pluginsdk.TypeBool,
				Optional: true,
				Default:  true,
			},

			"storage_endpoint": {
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsURLWithHTTPS,
			},

			"storage_account_access_key": {
				Type:         pluginsdk.TypeString,
				Optional:     true,
				Sensitive:    true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"storage_account_access_key_is_secondary": {
				Type:     pluginsdk.TypeBool,
				Optional: true,
				Default:  false,
			},

			"retention_in_days": {
				Type:         pluginsdk.TypeInt,
				Optional:     true,
				Default:      0,
				ValidateFunc: validation.IntBetween(0, 3285),
			},

			"log_monitoring_enabled": {
				Type:     pluginsdk.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func resourceMsSqlDatabaseExtendedAuditingPolicyCreateUpdate(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).MSSQL.DatabaseExtendedBlobAuditingPoliciesClient
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	log.Printf("[INFO] preparing arguments for MsSql Database Extended Auditing Policy creation.")

	dbId, err := commonids.ParseSqlDatabaseID(d.Get("database_id").(string))
	if err != nil {
		return err
	}

	if d.IsNewResource() {
		existing, err := client.Get(ctx, dbId.ResourceGroupName, dbId.ServerName, dbId.DatabaseName)
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("checking for the presence of existing %s: %+v", dbId, err)
			}
		}

		// if state is not disabled, we should import it.
		if existing.ID != nil && *existing.ID != "" && existing.ExtendedDatabaseBlobAuditingPolicyProperties != nil && existing.ExtendedDatabaseBlobAuditingPolicyProperties.State != sql.BlobAuditingPolicyStateDisabled {
			return tf.ImportAsExistsError("azurerm_mssql_database_extended_auditing_policy", *existing.ID)
		}
	}

	params := sql.ExtendedDatabaseBlobAuditingPolicy{
		ExtendedDatabaseBlobAuditingPolicyProperties: &sql.ExtendedDatabaseBlobAuditingPolicyProperties{
			StorageEndpoint:             utils.String(d.Get("storage_endpoint").(string)),
			IsStorageSecondaryKeyInUse:  utils.Bool(d.Get("storage_account_access_key_is_secondary").(bool)),
			RetentionDays:               utils.Int32(int32(d.Get("retention_in_days").(int))),
			IsAzureMonitorTargetEnabled: utils.Bool(d.Get("log_monitoring_enabled").(bool)),
		},
	}

	if d.Get("enabled").(bool) {
		params.ExtendedDatabaseBlobAuditingPolicyProperties.State = sql.BlobAuditingPolicyStateEnabled
	} else {
		params.ExtendedDatabaseBlobAuditingPolicyProperties.State = sql.BlobAuditingPolicyStateDisabled
	}

	if v, ok := d.GetOk("storage_account_access_key"); ok {
		params.ExtendedDatabaseBlobAuditingPolicyProperties.StorageAccountAccessKey = utils.String(v.(string))
	}

	if _, err = client.CreateOrUpdate(ctx, dbId.ResourceGroupName, dbId.ServerName, dbId.DatabaseName, params); err != nil {
		return fmt.Errorf("creating extended auditing policy for %s: %+v", dbId, err)
	}

	read, err := client.Get(ctx, dbId.ResourceGroupName, dbId.ServerName, dbId.DatabaseName)
	if err != nil {
		return fmt.Errorf("retrieving the extended auditing policy for %s: %+v", dbId, err)
	}

	if read.ID == nil || pointer.From(read.ID) == "" {
		return fmt.Errorf("the extended auditing policy ID for %s is 'nil' or 'empty'", dbId.String())
	}

	// TODO: update this to use the Database ID - requiring a State Migration
	readId, err := parse.DatabaseExtendedAuditingPolicyID(pointer.From(read.ID))
	if err != nil {
		return err
	}

	d.SetId(readId.ID())

	return resourceMsSqlDatabaseExtendedAuditingPolicyRead(d, meta)
}

func resourceMsSqlDatabaseExtendedAuditingPolicyRead(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).MSSQL.DatabaseExtendedBlobAuditingPoliciesClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.DatabaseExtendedAuditingPolicyID(d.Id())
	if err != nil {
		return err
	}

	resp, err := client.Get(ctx, id.ResourceGroup, id.ServerName, id.DatabaseName)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("reading MsSql Database %s Extended Auditing Policy (MsSql Server Name %q / Resource Group %q): %s", id.DatabaseName, id.ServerName, id.ResourceGroup, err)
	}

	databaseId := commonids.NewSqlDatabaseID(id.SubscriptionId, id.ResourceGroup, id.ServerName, id.DatabaseName)
	d.Set("database_id", databaseId.ID())

	if props := resp.ExtendedDatabaseBlobAuditingPolicyProperties; props != nil {
		d.Set("storage_endpoint", props.StorageEndpoint)
		d.Set("storage_account_access_key_is_secondary", props.IsStorageSecondaryKeyInUse)
		d.Set("retention_in_days", props.RetentionDays)
		d.Set("log_monitoring_enabled", props.IsAzureMonitorTargetEnabled)
		d.Set("enabled", props.State == sql.BlobAuditingPolicyStateEnabled)
	}

	return nil
}

func resourceMsSqlDatabaseExtendedAuditingPolicyDelete(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).MSSQL.DatabaseExtendedBlobAuditingPoliciesClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.DatabaseExtendedAuditingPolicyID(d.Id())
	if err != nil {
		return err
	}

	params := sql.ExtendedDatabaseBlobAuditingPolicy{
		ExtendedDatabaseBlobAuditingPolicyProperties: &sql.ExtendedDatabaseBlobAuditingPolicyProperties{
			State: sql.BlobAuditingPolicyStateDisabled,
		},
	}

	if _, err := client.CreateOrUpdate(ctx, id.ResourceGroup, id.ServerName, id.DatabaseName, params); err != nil {
		return fmt.Errorf("deleting MsSql Database %q Extended Auditing Policy( MsSql Server %q / Resource Group %q): %+v", id.DatabaseName, id.ServerName, id.ResourceGroup, err)
	}

	return nil
}
