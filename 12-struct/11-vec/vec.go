package main

import (
	"fmt"
	"math"
)

type Vec2 struct {
	X, Y float32
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		v.X + other.X,
		v.Y + other.Y,
	}
}

// 使用矢量减去另外一个矢量，生产新的矢量
func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{
		v.X - other.X,
		v.Y - other.Y,
	}
}

func (v Vec2) Scale(s float32) Vec2 {
	return Vec2{
		v.X * s,
		v.Y * s,
	}
}

// 计算两个矢量的距离
func (v Vec2) DistanceTo(other Vec2) float32 {
	dx := v.X - other.X
	dy := v.Y - other.Y

	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

// 返回当前矢量的标准化矢量
func (v Vec2) Normalize() Vec2 {
	mag := v.X*v.X + v.Y*v.Y
	if mag > 0 {
		oneOverMag := 1 / float32(math.Sqrt(float64(mag)))
		return Vec2{v.X * oneOverMag, v.Y * oneOverMag}
	}

	return Vec2{0, 0}
}

// 实现玩家对象
type Player struct {
	currPos   Vec2    //当前位置
	targetPos Vec2    //目标位置
	speed     float32 //移动速度
}

// 设置玩家移动的目标位置
func (p *Player) MoveTo(v Vec2) {
	p.targetPos = v
}

// 获取当前的位置
func (p *Player) Pos() Vec2 {
	return p.currPos
}

// 判断是否到达目的地
func (p *Player) IsArrived() bool {
	return p.currPos.DistanceTo(p.targetPos) < p.speed
}

// 更新玩家的位置
func (p *Player) Update() {
	if !p.IsArrived() {
		// 计算出当前位置指向目标的朝向
		dir := p.targetPos.Sub(p.currPos).Normalize()

		// 添加速度矢量生产新的位置
		newPos := p.currPos.Add(dir.Scale(p.speed))

		// 移动完成后，更新当前位置
		p.currPos = newPos
	}
}

// 创建玩家
func NewPlayer(speed float32) *Player {
	return &Player{
		speed: speed,
	}
}

func main() {
	// 实例化玩家对象，并设置速度为0.5
	p := NewPlayer(0.5)

	// 让玩家移动到(3,1) 点
	p.MoveTo(Vec2{3, 1})

	// 如果没有到达就一直循环
	for !p.IsArrived() {
		p.Update()
		fmt.Println(p.Pos())
	}
}
