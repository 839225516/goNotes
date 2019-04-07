package main

import (
	"errors"
	"fmt"
	"reflect"
)

// 状态接口
type State interface {
	// 获取状态名字
	Name() string

	// 该状态是否允许同状态转移
	EnableSameTransit() bool

	// 响应状态开始时
	OnBegin()

	// 响应状态结束时
	OnEnd()

	// 判断能否转移到某个状态
	CanTransitTo(name string) bool
}

// 从状态实例获取状态名
func StateName(s State) string {
	if s == nil {
		return "none"
	}

	// 使用反射获取状态的名称
	return reflect.TypeOf(s).Elem().Name()
}

// 状态信息
type StateInfo struct {
	// 状态名
	name string
}

// 状态名
func (s *StateInfo) Name() string {
	return s.name
}

// 提供给内部设置名字
func (s *StateInfo) setName(name string) {
	s.name = name
}

// 允许同状态转移
func (s *StateInfo) EnableSameTransit() bool {
	return false
}

// 默认将状态开启时实现
func (s *StateInfo) OnBegin() {
}

// 默认将状态结束时实现
func (s *StateInfo) OnEnd() {
}

// 默认可以转移到任何状态
func (s *StateInfo) CanTransitTo(name string) bool {
	return true
}

// 状态管理
type StateMmanger struct {
	// 已添加的状态
	stateByName map[string]State

	// 状态改变时的回调
	Onchange func(from, to State)

	// 当前状态
	curr State
}

// 添加一个状态到管理器中
func (sm *StateMmanger) Add(s State) {
	// 获取状态的名称
	name := StateName(s)

	// 将s转换为能设置名字的接口，然后调用该接口
	s.(interface {
		setName(name string)
	}).setName(name)

	// 根据状态名获取已经添加的状态，检查该状态是否存在
	if sm.Get(name) != nil {
		panic("duplicate state:" + name)
	}

	// 根据名字保存到map中
	sm.stateByName[name] = s
}

// 根据名字获取指定状态
func (sm *StateMmanger) Get(name string) State {
	if v, ok := sm.stateByName[name]; ok {
		return v
	}

	return nil
}

// 初始化状态管理器
func NewStateManager() *StateMmanger {
	return &StateMmanger{
		stateByName: make(map[string]State),
	}
}

// 在状态间转移

// 状态没有找到的错误
var ErrStateNotFound = errors.New("state not found")

// 禁止在同状态间转移
var ErrForbidSameStateTransit = errors.New("forbid same state transit")

// 不能转移到指定状态
var ErrCannotTransitToState = errors.New("cannot transit to state")

// 获取当前的状态
func (sm *StateMmanger) CurrState() State {
	return sm.curr
}

// 当前能否转移到目标状态
func (sm *StateMmanger) CanCurrTransitTo(name string) bool {
	if sm.curr == nil {
		return true
	}

	// 相同的状态不用转移
	if sm.curr.Name() == name && !sm.curr.EnableSameTransit() {
		return true
	}

	// 使用当前状态，检查能否转移到指定名字的状态
	return sm.curr.CanTransitTo(name)
}

// 转移到指定状态
func (sm *StateMmanger) Transit(name string) error {
	// 获取目标状态
	next := sm.Get(name)

	// 目标不存在
	if next == nil {
		return ErrStateNotFound
	}

	// 记录转移前的状态
	pre := sm.curr

	// 当前有状态
	if sm.curr != nil {
		// 相同状态不用转换
		if sm.curr.Name() == name && !sm.curr.EnableSameTransit() {
			return ErrForbidSameStateTransit
		}

		// 不同转移到目标状态
		if !sm.curr.CanTransitTo(name) {
			return ErrCannotTransitToState
		}

		// 结束当前状态
		sm.curr.OnEnd()
	}

	// 将当前状态切换到要转移的目标状态
	sm.curr = next

	// 调用新状态的开始
	sm.curr.OnBegin()

	// 通知回调
	if sm.Onchange != nil {
		sm.Onchange(pre, sm.curr)
	}

	return nil
}

// 自定义状态实现状态接口
// 闲置状态
type IdleState struct {
	StateInfo // 使用StateInfo 实现基础接口
}

// 重新实现状态开始
func (i *IdleState) OnBegin() {
	fmt.Println("IdleState begin")
}

// 重新实现状态结束
func (i *IdleState) OnEnd() {
	fmt.Println("IdleState end")
}

type MoveState struct {
	StateInfo
}

func (m *MoveState) OnBegin() {
	fmt.Println("MoveState begin")
}

// 允许移动状态互相转移
func (m *MoveState) EnableSameTransit() bool {
	return true
}

// 跳跃状态
type JumpState struct {
	StateInfo
}

func (j *JumpState) OnBegin() {
	fmt.Println("JumpState begin")
}

// 跳跃状态不能转移到移动状态
func (j *JumpState) CanTransitTo(name string) bool {
	return name != "MoveState"
}

// 使用状态机
func main() {
	// 实例化一个状态管理器
	sm := NewStateManager()

	// 响应状态转移的通知
	sm.Onchange = func(from, to State) {
		//打印状态转移的流向
		fmt.Printf("%s ---> %s\n\n", StateName(from), StateName(to))
	}

	// 添加3个状态
	sm.Add(new(IdleState))
	sm.Add(new(MoveState))
	sm.Add(new(JumpState))

	// 在不同状态间转移
	transitAndReport(sm, "IdleState")

	transitAndReport(sm, "MoveState")
	transitAndReport(sm, "MoveState")

	transitAndReport(sm, "JumpState")
	transitAndReport(sm, "JumpState")

	transitAndReport(sm, "IdleState")

}

// 封装转移状态和输出日志
func transitAndReport(sm *StateMmanger, target string) {
	if err := sm.Transit(target); err != nil {
		fmt.Printf("FAILED! %s --> %s, %s\n\n", sm.CurrState().Name(), target, err.Error())
	}
}
