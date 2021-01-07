package Client

import (
	"github.com/wuxia-server/game/Const"
	"math"
	"time"
)

// 更新体力值恢复
func (e *ClientModel) UpdateVigorRecover() {
	// 相差纳秒值
	diffNano := time.Now().Sub(e.User.VigorRecoverTime)
	// 相差秒值
	diffSecond := math.Floor(float64(diffNano) / float64(time.Second))

	// 可恢复的值
	recoverVigor := int(math.Floor(diffSecond / float64(Const.VigorRecoverTime)))

	// 超出上限
	if recoverVigor+e.User.Vigor >= Const.VigorLimit {
		e.User.Vigor = Const.VigorLimit
		e.User.VigorRecoverTime = time.Now()
	} else {
		e.User.Vigor += recoverVigor
		e.User.VigorRecoverTime = e.User.VigorRecoverTime.Add(time.Duration(recoverVigor*Const.VigorRecoverTime) * time.Second)
	}

	e.User.Save()
}

// 更新活力值恢复
func (e *ClientModel) UpdateVitalityRecover() {
	// 相差纳秒值
	diffNano := time.Now().Sub(e.User.VitalityRecoverTime)
	// 相差秒值
	diffSecond := math.Floor(float64(diffNano) / float64(time.Second))

	// 可恢复的值
	recoverVitality := int(math.Floor(diffSecond / float64(Const.VitalityRecoverTime)))

	// 超出上限
	if recoverVitality+e.User.Vitality >= Const.VitalityLimit {
		e.User.Vitality = Const.VitalityLimit
		e.User.VitalityRecoverTime = time.Now()
	} else {
		e.User.Vitality += recoverVitality
		e.User.VitalityRecoverTime = e.User.VitalityRecoverTime.Add(time.Duration(recoverVitality*Const.VitalityRecoverTime) * time.Second)
	}

	e.User.Save()
}
