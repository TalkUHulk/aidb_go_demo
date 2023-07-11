package aidb

/*
#cgo CFLAGS : -I${SRCDIR}/include
#cgo LDFLAGS: -L${SRCDIR}/lib -lAiDB_C -Wl,-rpath,lib
#include <stdio.h>
#include <stdlib.h>
#include "aidb_c_api.h"
*/
import "C"
import (
	"encoding/json"
	"fmt"
	"unsafe"
)

type GoAiDB struct {
	aidb C.AiDB
}

var G GoAiDB

func AiDBRegister(flow_uuid string, model []string, backend []string, zoo string) int {

	aidb_input := AIDBInput{
		FlowUUID: flow_uuid,
		Backend:  backend,
		Model:    model,
		Zoo:      zoo,
	}

	aidb_input_str, err := json.Marshal(&aidb_input)

	if err != nil {
		fmt.Printf("序列号错误 err = %v\n", err)
	}
	fmt.Printf("序列化后= %v\n", string(aidb_input_str))

	c_aidb_input_str := C.CString(string(aidb_input_str))
	ret := C.AiDBRegister(G.aidb, c_aidb_input_str)
	C.free(unsafe.Pointer(c_aidb_input_str))

	// 	ret := C.AiDBRegister(A.aidb, (*C.char)(unsafe.Pointer(&aidb_input_str[0])))

	return int(ret)
}

func AiDBCreate() {
	G.aidb = C.AiDBCreate()
}

func AiDBFree() {
	C.AiDBFree(G.aidb)
	fmt.Printf("Free\n")
}

func AiDBForward(flow_uuid string, image_raw string, result *AiDBOutput) int {
	c_flow_uuid := C.CString(flow_uuid)
	c_image_raw := C.CString(image_raw)
	c_size_out := (C.int)(0)
	size_in := 1024 * 1024

	c_size_in := (C.int)(size_in)
	c_result := (*C.char)(C.malloc((C.ulong)(size_in)))

	// aidb_output := AiDBOutput{}

	ret := C.AiDBForward(G.aidb, c_flow_uuid, c_image_raw, c_result, c_size_in, &c_size_out)

	if ret != 0 {
		fmt.Println("forward failed:", ret)
	} else {
		err := json.Unmarshal([]byte(C.GoString(c_result)), result)
		if err != nil {
			fmt.Println("反序列化失败", err)
			ret = -100
		}
	}

	// fmt.Println("GO!!!!", aidb_output)
	// if len(aidb_output.Tddfa) != 0 {
	// 	fmt.Println("Tddfa")
	// 	decodedBytes, err := base64.StdEncoding.DecodeString(aidb_output.Tddfa)
	// 	if err != nil {
	// 		fmt.Println("Error decoding string:", err)
	// 		return
	// 	}

	// 	fo, err := os.Create("./tddfa.jpg")
	// 	if err != nil {
	// 		fmt.Println("文件创建失败", err.Error())
	// 		return
	// 	}

	// 	defer func() {
	// 		if err := fo.Close(); err != nil {
	// 			panic(err)
	// 		}
	// 	}()

	// 	if _, err := fo.Write(decodedBytes); err != nil {
	// 		panic(err)
	// 	}
	// }

	defer C.free(unsafe.Pointer(c_flow_uuid))
	defer C.free(unsafe.Pointer(c_result))
	defer C.free(unsafe.Pointer(c_image_raw))
	return int(ret)
}

func AiDBUnRegister(flow_uuid string) int {

	c_flow_uuid := C.CString(flow_uuid)
	ret := C.AiDBUnRegister(G.aidb, c_flow_uuid)
	defer C.free(unsafe.Pointer(c_flow_uuid))
	return int(ret)
}
