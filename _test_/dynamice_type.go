package main

import (
	"fmt"
	"math"
)

func main() {
	truckSize := 0.0
	truckSize2 := 0.0
	// 声明空接口类型变量materials，存放各种不同体积的家具
	var materials []interface{}
	materials = append(materials, cube{12.5})
	materials = append(materials, cuboid{25, 13, 60})
	materials = append(materials, cylinder{5, 25.3})
	// 遍历materials切片，依次计算每个家具的体积，并相加求和
	for _, singleMaterial := range materials {
		truckSize += calcSize(singleMaterial)
		calcSize2(singleMaterial)
	}
	fmt.Println(truckSize, truckSize2)
}

// 计算某个物体的体积
func calcSize(material interface{}) float64 {
	cubeMaterial, cubeOk := material.(cube)
	cuboidMaterial, cuboidOk := material.(cuboid)
	cylinderMaterial, cylinderOk := material.(cylinder)
	if cubeOk {
		return cubeMaterial.cubeVolume()
	} else if cuboidOk {
		return cuboidMaterial.cuboidVolume()
	} else if cylinderOk {
		return cylinderMaterial.cylinderVolume()
	} else {
		return 0
	}
}

func calcSize2(material interface{}) float64 {
	var result float64 = 0.0
	switch data := material.(type) {
	case cube:
		fmt.Println("cube", data)
	case cuboid:
		fmt.Println("cuboid", data)
	case cylinder:
		fmt.Println("cylinder", data)
	}
	return result
}

// 正方体
type cube struct {
	// 边长
	length float64
}

// 正方体的体积计算
func (c *cube) cubeVolume() float64 {
	return c.length * c.length * c.length
}

// 长方体
type cuboid struct {
	// 长
	length float64
	// 宽
	width float64
	// 高
	height float64
}

// 长方体的体积计算
func (c *cuboid) cuboidVolume() float64 {
	return c.length * c.width * c.height
}

// 圆柱体
type cylinder struct {
	// 直径
	diameter float64
	// 高度
	height float64
}

// 圆柱体的体积计算
func (c *cylinder) cylinderVolume() float64 {
	return math.Pi * (c.diameter / 2) * (c.diameter / 2) * c.height
}
