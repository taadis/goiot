package model

// 机床每次上报的数据
type Model struct {
	IDnum              int64  // ID
	MachineSN          string // 设备编号
	MachineType        string // 设备类型
	MachineIP          string // 设备IP
	MachineMode        string // 机床状态
	PowerOnTime        string // 运行时间
	RunningTime        string // 通电时间
	CuttingTime        string // 切削时间
	ProcessingPart     string // 程序注释
	CurrentProgramName string // 当前程序名
	ProcessingCount    int64  // 当班数量(机床计数)
	ProcessedCount     int64  // 总数量(机床计数)
	SpindleLoad        int64  // 主轴负载
	SpindleSpeed       int64  // 主轴转速
	FeedSpeed          int64  // 进给速度
	SpindleOverride    int64  // 主轴倍率
	FeedOverride       int64  // 进给轴倍率
	SpindleSpeedSet    int64  // 主轴设定转速S
	FeedSpeedSet       int64  // 进给指定速度F
	AlarmType          int64  // 报警类型
	AlarmNo            int64  // 报警号
	AlarmMessage       int64  // 机床报警信息
	Tool               string // 刀具使用信息
	Xload              int64  // X轴负载
	Yload              int64  // Y轴负载
	Zload              int64  // Z轴负载
	CreateTime         int64  // 创建时间
}
