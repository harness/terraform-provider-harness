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

resource "harness_autostopping_schedule" "DownTimeForISTLuchOnEveryFridayInBetweenDate" {
    name = "DownTimeForISTLuchOnEveryFridayInBetweenDate"
    schedule_type = "downtime"
    time_zone = "Asia/Calcutta"    

    starting_from = "2023-01-02 15:04:05"
    ending_on = "2023-02-02 15:04:05"

    repeats {
        days = "FRI"
        start_time = "12:30"
        end_time = "14:30"
    }

    rules = [ 123 ]
}

resource "harness_autostopping_schedule" "DowntimeEveryFridayAfterNoonTillEOD" {
    name = "DowntimeEveryFridayAfterNoonTillEOD"
    schedule_type = "downtime"
    time_zone = "Asia/Calcutta"    

    repeats {
        days = "FRI"
        start_time = "17:30"
    }

    rules = [ 123 ]
}

resource "harness_autostopping_schedule" "CompleteDownTimeOnWeekEnd" {
    name = "CompleteDownTimeOnWeekEnd"
    schedule_type = "downtime"
    time_zone = "UTC"    

    repeats {
        days = "SUN, SAT"
    }

    rules = [ 123 ]
}
