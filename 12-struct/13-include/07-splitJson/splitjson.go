package main

import (
	"encoding/json"
	"fmt"
)

type Screen struct {
	Size       float32
	PesX, PesY int
}

type Battery struct {
	Capacity int
}

// 生成 json 数据
func genJsonDate() []byte {
	// 完整数据结构
	raw := &struct {
		Screen
		Battery
		HasTouchID bool
	}{
		// 屏幕参数
		Screen: Screen{
			Size: 5.5,
			PesX: 1920,
			PesY: 1080,
		},
		// 电池参数
		Battery: Battery{
			2910,
		},
		//是否有指纹
		HasTouchID: true,
	}

	// 将数据序列化为 json
	jsonData, _ := json.Marshal(raw)

	return jsonData
}

func main() {
	// 生成一段json
	jsonData := genJsonDate()

	fmt.Println(string(jsonData))

	// 只需要屏幕和指纹识别信息的结构和实例
	screenAndTouch := struct {
		Screen
		HasTouchID bool
	}{}

	// 反序列化到screenAndTouch中
	json.Unmarshal(jsonData, &screenAndTouch)

	fmt.Printf("%+v\n", screenAndTouch)

	// 只需要电池和指纹识别信息
	batteryAndTouch := struct {
		Battery
		HasTouchID bool
	}{}

	// 反序列化到 batteryAndTouch
	json.Unmarshal(jsonData, &batteryAndTouch)

	fmt.Printf("%+v\n", batteryAndTouch)
}
