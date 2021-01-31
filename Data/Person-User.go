package Data

import (
	"errors"
	"fmt"
	"github.com/team-zf/framework/Network"
	"github.com/wuxia-server/game/Const"
	"github.com/wuxia-server/game/Rule"
	"github.com/wuxia-server/game/StaticTable"
	"math"
	"time"
)

// 用户转为JsonMap输出格式
func (e *Person) __UserToJsonMap() map[string]interface{} {
	return e.User.ToJsonMap()
}

// 更新体力值恢复
func (e *Person) UpdateVigorRecover() {
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
func (e *Person) UpdateVitalityRecover() {
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

// 消耗体力
func (e *Person) SubVigor(num int) (*Network.WebSocketDDM, error) {
	if num <= 0 {
		return nil, errors.New(fmt.Sprintf("传入的数值有误(%d).", num))
	}

	// 相差纳秒值
	diffNano := time.Now().Sub(e.User.VigorRecoverTime)
	// 相差秒值
	diffSecond := math.Floor(float64(diffNano) / float64(time.Second))

	// 可恢复的值
	recoverVigor := int(math.Floor(diffSecond / float64(Const.VigorRecoverTime)))

	if recoverVigor+e.User.Vigor < num {
		return nil, errors.New(fmt.Sprintf("体力不足(%d), 含恢复值的当前体力为(%d)", num, recoverVigor+e.User.Vigor))
	}

	// 超出上限
	if recoverVigor+e.User.Vigor >= Const.VigorLimit {
		e.User.Vigor = Const.VigorLimit
		e.User.VigorRecoverTime = time.Now()
	} else {
		e.User.Vigor += recoverVigor
		e.User.VigorRecoverTime = e.User.VigorRecoverTime.Add(time.Duration(recoverVigor*Const.VigorRecoverTime) * time.Second)
	}

	e.User.Vigor -= num
	e.User.Save()

	e.UpdateVitalityRecover()

	ddm := new(Network.WebSocketDDM)
	ddm.Mod(Rule.RULE_USER, e.User.ToJsonMap())
	return ddm, nil
}

// 消耗体力(异常直接抛panic)
func (e *Person) SubVigorV(num int) *Network.WebSocketDDM {
	ddm, err := e.SubVigor(num)
	if err != nil {
		panic(err)
	}
	return ddm
}

// 消耗活力
func (e *Person) SubVitality(num int) (*Network.WebSocketDDM, error) {
	if num <= 0 {
		return nil, errors.New(fmt.Sprintf("传入的数值有误(%d).", num))
	}

	// 相差纳秒值
	diffNano := time.Now().Sub(e.User.VitalityRecoverTime)
	// 相差秒值
	diffSecond := math.Floor(float64(diffNano) / float64(time.Second))

	// 可恢复的值
	recoverVitality := int(math.Floor(diffSecond / float64(Const.VitalityRecoverTime)))

	if recoverVitality+e.User.Vitality < num {
		return nil, errors.New(fmt.Sprintf("活力不足(%d), 含恢复值的当前活力为(%d)", num, recoverVitality+e.User.Vitality))
	}

	// 超出上限
	if recoverVitality+e.User.Vitality >= Const.VitalityLimit {
		e.User.Vitality = Const.VitalityLimit
		e.User.VitalityRecoverTime = time.Now()
	} else {
		e.User.Vitality += recoverVitality
		e.User.VitalityRecoverTime = e.User.VitalityRecoverTime.Add(time.Duration(recoverVitality*Const.VitalityRecoverTime) * time.Second)
	}

	e.User.Vitality -= num
	e.User.Save()

	e.UpdateVigorRecover()

	ddm := new(Network.WebSocketDDM)
	ddm.Mod(Rule.RULE_USER, e.User.ToJsonMap())
	return ddm, nil
}

// 消耗活力(异常直接抛panic)
func (e *Person) SubVitalityV(num int) *Network.WebSocketDDM {
	ddm, err := e.SubVitality(num)
	if err != nil {
		panic(err)
	}
	return ddm
}

// 增加经验
func (e *Person) AddExp(v int) (*Network.WebSocketDDM, error) {
	if v <= 0 {
		return nil, errors.New(fmt.Sprintf("传入的数值有误(%d).", v))
	}

	upgrade := false
	level := e.Level()
	limit := StaticTable.GetRoleLevelLimit()
	if limit >= level {
		exp := e.User.Exp + v
		for {
			rlv := StaticTable.GetRoleLv(level + 1)
			// 已满级或经验不足以升级
			if rlv == nil || rlv.NeedExp > exp {
				break
			}
			upgrade = true
			level = rlv.Level
		}
		e.User.Exp = exp
		e.User.Level = level
	} else {
		e.User.Exp += v
	}
	e.User.Save()

	ddm := new(Network.WebSocketDDM)
	ddm.Mod(Rule.RULE_USER, e.User.ToJsonMap())

	// 如果升级了, 做条件校验
	if upgrade || true {
		ddm.Join(e.CondVerify())
	}

	return ddm, nil
}

// 增加经验(异常直接抛panic)
func (e *Person) AddExpV(v int) *Network.WebSocketDDM {
	ddm, err := e.AddExp(v)
	if err != nil {
		panic(err)
	}
	return ddm
}
