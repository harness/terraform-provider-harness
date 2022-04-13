package helpers

import (
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func HandleApiError(err error, d *schema.ResourceData) diag.Diagnostics {
	e := err.(nextgen.GenericSwaggerError)

	if e.Code() == nextgen.ErrorCodes.ResourceNotFound {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}
	return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
}
