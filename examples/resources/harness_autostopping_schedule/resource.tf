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

resource "harness_autostopping_schedule" "MondayThroughFridayUptimeFrom9amTo6pm" {
    name = "MondayThroughFridayUptimeFrom9amTo6pm"
    schedule_type = "uptime"
    time_zone = "UTC"    

    repeats {
        days = "MON, TUE, WED, THU, FRI"
        start_time = "09:00"
        end_time = "18:00"
    }

    rules = [ 123 ]
}

resource "harness_autostopping_schedule" "MondayThroughFridayUptimeFrom9amTo6pmStartingFromDate" {
    name = "MondayThroughFridayUptimeFrom9amTo6pmStartingFrom"
    schedule_type = "uptime"
    time_zone = "UTC"    

    starting_from = "2023-01-02 15:04:05"

    repeats {
        days = "MON, TUE, WED, THU, FRI"
        start_time = "09:00"
        end_time = "18:00"
    }

    rules = [ 123 ]
}
