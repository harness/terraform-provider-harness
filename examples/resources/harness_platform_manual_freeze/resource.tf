resource "harness_platform_manual_freeze" "example" {
  identifier = "identifier"
  org_id     = "orgIdentifier"
  project_id = "projectIdentifier"
  account_id = "accountIdentifier"
  yaml       = <<-EOT
      freeze:
        name: %[2]s
        identifier: %[1]s
        entityConfigs:
          - name: r1
            entities:
              - filterType: All
                type: Org
              - filterType: All
                type: Project
              - filterType: All
                type: Service
              - filterType: All
                type: EnvType
        status: Disabled
        description: hi
        windows:
        - timeZone: Asia/Calcutta
          startTime: 2023-05-03 04:16 PM
          duration: 30m
          recurrence:
            type: Daily
        notificationRules: []
        tags: {}
      EOT
}
