data "harness_delegate" "test" {
  name = "my-delegate"
}

resource "harness_delegate_approval" "test" {
  delegate_id = data.harness_delegate.test.id
  approve     = true
}
