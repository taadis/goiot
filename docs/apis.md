# APIs

## API001:整体统计接口

## API002:历史告警统计查询接口

```
// output
{
    "code": "告警号码",
    "total": "总数量",
    "current": "24小时内数量"
}
```

## API003:设备状态统计查询接口

```
// input
{
    "machineSN": "设备编号"
}

// output
{
    "machineSN": "设备编号",
    "machineStatus": "设备状态",
    "total": "该状态的总时间",
    "current": "24小时内的统计"
}
```

## API004:单设备最新采集的数据

Task:连续查询每台设备最新的数据至独立表，有多少台设备就在这个表记录多少条数据。

```
// input
{
    "machineSN": "设备编号"
}

// output
{
    "allRawKey": "allRawValue"
}
```

## API005:按天/单台设备及某个子项的统计