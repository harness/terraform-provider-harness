resource "harness_application" "myapp" {
  name        = "My Application"
  description = "This is my first Harness application"

  tags = [
    "mytag:myvalue",
    "team:development"
  ]
}
