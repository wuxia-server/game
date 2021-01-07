package Const

const (
	InitialVigor        int = 60    // 初始体力值
	InitialVitality     int = 15    // 初始气力值
	VigorLimit          int = 120   // 体力上限值
	VigorRecoverTime    int = 600   // 体力恢复时间(秒)
	VitalityLimit       int = 30    // 气力上限值
	VitalityRecoverTime int = 2400  // 气力恢复时间(秒)
	InitialChapterId    int = 101   // 初始章节ID, 在初始化用户数据的时候根据此ID开通第一章副本数据
	CaveProduceEGM      int = 10000 // 洞府产量值扩大倍数, 避免小数的丢失
)
