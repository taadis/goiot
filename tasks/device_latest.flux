option task = {
    name: "task_device_info_latest",
    cron: "*/5 * * * *",
}

from(bucket: "my-bucket")
    |> range(start: -1h)
    |> toString()
    |> filter(fn: (r) => r["_measurement"] == "machine_info")
    |> group(columns: ["machine_sn"], mode: "by")
    //|> last()
    |> map(fn: (r) => ({r with _measurement: "device-info-latest1"}))
    |> aggregateWindow(every: 5s, fn: last, createEmpty: false)
    |> yield(name: "device-info-latest")
    |> to(bucket: "my-bucket", org: "taadis")