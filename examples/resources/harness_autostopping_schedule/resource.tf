resource "harness_autostopping_schedule" "MondayWholeDayUp" {
    name = "MondayWholeDayUp"
    schedule_type = "uptime"
    time_zone = "UTC"    

    repeats {
        days = "MON"
    }

    rules = [ 123 ]
}

resource "harness_autostopping_schedule" "MondayUpTill4:30pm" {
    name = "MondayUpTill4:30pm"
    schedule_type = "uptime"
    time_zone = "UTC"    

    repeats {
        days = "MON"
        end_time = "16:30"
    }

    rules = [ 123 ]
}
