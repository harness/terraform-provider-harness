# Only Locust and K6 can be imported. JMeter, Gatling, and Custom are rejected.
terraform import harness_chaos_load_test.example <org_id>/<project_id>/<identifier>

# The first apply sends the configuration as the full desired state.
# Copy variable, resources, max_duration_sec, tags, tool_config, cleanup_policy,
# and managed_infra from state before applying. An omitted variable list,
# resources block, duration, or tags set clears the stored value. tool_config
# replaces the script. cleanup_policy defaults to delete and managed_infra to false.
# Destroy deletes the load test definition. Recorded runs stay stored, and a
# run that is already in progress is not stopped.
