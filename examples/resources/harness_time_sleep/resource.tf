# Pause 30s after delegate is created before any dependent resource proceeds.
# Useful when the delegate needs time to register before it can accept work.
resource "harness_time_sleep" "wait_for_delegate" {
  create_duration = "30s"
  depends_on      = [harness_platform_delegate_token.mydelegate]
}

# Pause 10s after delegate is destroyed before tearing down dependents.
# Gives the delegate time to finish in-flight tasks and fully deregister.
resource "harness_time_sleep" "drain_delegate" {
  destroy_duration = "10s"
  depends_on       = [harness_platform_delegate_token.mydelegate]
}

# Pause on both create and destroy.
# triggers causes the sleep to re-fire whenever delegate_id changes —
# for example when the delegate is recreated with a new ID.
resource "harness_time_sleep" "wait" {
  create_duration  = "30s"
  destroy_duration = "10s"

  triggers = {
    delegate_id = harness_platform_delegate_token.mydelegate.id
  }
}
