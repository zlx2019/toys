package prints

// 终端文字 ASCII 控制符

// 文字显示样式标识
// Default 	 0 重置(终端默认样式)
// High 	 1 高亮
// Underline 4 下划线
// Flash     5 闪烁(不可用)
// Reverse	 7 反白(将字体颜色作为背景颜色)
// Hide		 8 隐藏
const (
	Default   = 0
	High      = 1
	Underline = 4
	Flash     = 5
	Reverse   = 7
	Hide      = 8
)

// 终端前景颜色标识
// Black 	30 黑色
// Red   	31 红色
// Green 	32 绿色
// Yellow 	33 黄色
// Blue 	34 蓝色
// Purple 	35 紫色
// Cyan 	36 青色
// White 	37 白色
const (
	Black = iota + 30
	Red
	Green
	Yellow
	Blue
	Purple
	Cyan
	White
)

// 终端背景颜色标识
// BackgroundBlack 	40 黑色背景
// BackgroundRed   	41 红色背景
// BackgroundGreen 	42 绿色背景
// BackgroundYellow 43 黄色背景
// BackgroundBlue 	44 蓝色背景
// BackgroundPurple 45 紫色背景
// BackgroundCyan 	46 青色背景
// BackgroundWhite 	47 白色背景
const (
	BackgroundBlack = iota + 40
	BackgroundRed
	BackgroundGreen
	BackgroundYellow
	BackgroundBlue
	BackgroundPurple
	BackgroundCyan
	BackgroundWhite
)
