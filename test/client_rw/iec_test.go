package client_rw

import (
	"log"
	"testing"

	"github.com/wendy512/iec61850"
)

func Test_read618501(t *testing.T) {
	// 1. 配置连接参数
	settings := iec61850.NewSettings()
	settings.Host = "192.168.3.225" // 替换为你的IED设备IP
	settings.Port = 102             // IEC61850标准端口

	// 2. 建立连接
	client, err := iec61850.NewClient(settings)
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer client.Close()
	log.Println("连接成功")

	float, _ := client.ReadFloat("ZE13CTMP01/MMXU1.yc1.mag.f", iec61850.MX)
	log.Printf("float: %v", float)

	float, _ = client.ReadFloat("ZE13CVOL01/MMXU1.yc2689.mag.f", iec61850.MX)
	log.Printf("float: %v", float)

	float, _ = client.ReadFloat("ZE13RACK01/MMXU1.yc7681.mag.f", iec61850.MX)
	log.Printf("float: %v", float)

	bo, _ := client.ReadBool("ZE13RACK01/GGIO1.yx1.stVal", iec61850.ST)
	log.Printf("bo: %v", bo)
}

func Test_write618501(t *testing.T) {
	// 1. 配置连接参数
	settings := iec61850.NewSettings()
	settings.Host = "192.168.3.225" // 替换为你的IED设备IP
	settings.Port = 102             // IEC61850标准端口

	// 2. 建立连接
	client, err := iec61850.NewClient(settings)
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer client.Close()
	log.Println("连接成功")

	//mode, _ := client.Read("ZE13STCK01/CSWI0.Mod.stVal", iec61850.ST)
	//if mode != "on" {
	//	log.Printf("IED不处于可操作模式，当前模式: %v", mode)
	//}
	//
	//if err := client.Write("ZE13STCK01/CSWI0.Mod.stVal", iec61850.SP, 1); err != nil {
	//	log.Printf("写入异常: %v", err)
	//}

	//ctlModel, _ := client.Read("ZE13STCK01/CSWI0.yk1.ctlModel", iec61850.CF)
	//log.Printf("ctlModel: %v", ctlModel)

	//if err := client.Write("ZE13STCK01/CSWI0.yk1.Oper.ctlVal", iec61850.CO, 1); err != nil {
	//	log.Printf("写入异常: %v", err)
	//}
	//if err := client.Write("ZE13STCK01/CSWI0.yk1.stVal", iec61850.ST, 1); err != nil {
	//	log.Printf("写入异常: %v", err)
	//}

	bo, _ := client.Read("ZE13STCK01/CSWI0.yk1.stVal", iec61850.ST)
	log.Printf("bo: %v", bo)

	//err = client.ControlForDirectWithNormalSecurity("ZE13STCK01/CSWI0.yk1", false)

	inc := iec61850.NewControlObjectParamINC(0)
	inc.OrCat = 2
	err = client.ControlByControlModelINC("ZE13STCK01/CSWI0.yk1",
		iec61850.CONTROL_MODEL_DIRECT_NORMAL,
		inc)
	if err != nil {
		log.Fatalf("Control 操作失败: %v", err)
	}
	log.Println("Control 操作成功，已发送")

	bo, _ = client.Read("ZE13STCK01/CSWI0.yk1.stVal", iec61850.ST)
	log.Printf("bo: %v", bo)
}
