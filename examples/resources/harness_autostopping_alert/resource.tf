resource "harness_autostopping_alert" "specific-rule-alert" {
    name = "demo-alert"
      recipients {
        email = ["user1@example.com", "user2@example.com"]
        slack = ["slack-web-hook-1", "slack-web-hook-2"]
    }
    events = ["autostopping_rule_created", "autostopping_rule_updated",
              "autostopping_rule_deleted", "autostopping_warmup_failed",
              "autostopping_cooldown_failed"]
    rule_id_list = [1234]
}

resource "harness_autostopping_alert" "all-rule-alert" {
    name = "demo-alert"
      recipients {
        email = ["user1@example.com", "user2@example.com"]
        slack = ["slack-web-hook-1", "slack-web-hook-2"]
    }
    events = ["autostopping_rule_created", "autostopping_rule_updated",
              "autostopping_rule_deleted", "autostopping_warmup_failed",
              "autostopping_cooldown_failed"]
    applicable_to_all_rules = true
}