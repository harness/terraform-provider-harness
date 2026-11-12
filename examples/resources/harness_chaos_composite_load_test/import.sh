# Import id is org_id/project_id/identifier. identifier is the stage id.
# The pipeline id is {identifier}_pipeline and is not the Terraform id.
# Destroy deletes that pipeline. The load test and the probe stay in place.
terraform import harness_chaos_composite_load_test.checkout <org_id>/<project_id>/<identifier>
