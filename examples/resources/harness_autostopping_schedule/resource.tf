resource "harness_autostopping_schedule" "MondayWholeDayUp" {
    name = "MondayWholeDayUp"
    schedule_type = "uptime"
    time_zone = "UTC"    

    repeats {
        days = "MON"
    }

    rules = [ 123 ]
}

