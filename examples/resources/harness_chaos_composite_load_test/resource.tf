resource "harness_chaos_composite_load_test" "checkout" {
  org_id     = "default"
  project_id = "demo"
  identifier = "checkout"
  name       = "Checkout"

  description = "Checkout path under load, with a chaos probe beside it."
  objective   = "Saturate checkout while the probe watches the payment service."

  load_test_ref = harness_chaos_load_test.checkout_k6.id

  tags = {
    team = "payments"
  }

  probe {
    identity        = "checkout_probe"
    infra_reference = harness_chaos_infrastructure_v2.payments.infra_id
    duration        = "5m"
  }
}
