// Copyright 2014 The GiterLab Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// example for ots2
package main

import (
	"fmt"
	"os"
	"time"

	ots2 "github.com/GiterLab/goots"
	"github.com/GiterLab/goots/log"
	. "github.com/GiterLab/goots/otstype"
)

// modify it to yours
const (
	ENDPOINT     = "http://127.0.0.1:8800"
	ACCESSID     = "OTSMultiUser177_accessid"
	ACCESSKEY    = "OTSMultiUser177_accesskey"
	INSTANCENAME = "TestInstance177"
)

func main() {
	// set running environment
	ots2.OTSDebugEnable = true
	ots2.OTSLoggerEnable = true
	log.OTSErrorPanicMode = true // 默认为开启，如果不喜欢panic则设置此为false

	fmt.Println("Test goots start ...")

	ots_client, err := ots2.New(ENDPOINT, ACCESSID, ACCESSKEY, INSTANCENAME)
	if err != nil {
		fmt.Println(err)
	}

	// delete_table
	ots_err := ots_client.DeleteTable("myTable")
	if ots_err != nil {
		fmt.Println(ots_err)
		// os.Exit(1)
	}
	fmt.Println("表已删除")

	// create_table
	table_meta := &OTSTableMeta{
		TableName: "myTable",
		SchemaOfPrimaryKey: OTSSchemaOfPrimaryKey{
			"gid": "INTEGER",
			"uid": "INTEGER",
		},
	}

	reserved_throughput := &OTSReservedThroughput{
		OTSCapacityUnit{100, 100},
	}

	ots_err = ots_client.CreateTable(table_meta, reserved_throughput)
	if ots_err != nil {
		fmt.Println(ots_err)
		os.Exit(1)
	}
	fmt.Println("表已创建")

	// list_table
	list_tables, ots_err := ots_client.ListTable()
	if ots_err != nil {
		fmt.Println(ots_err)
		os.Exit(1)
	}
	fmt.Println("表的列表如下：")
	fmt.Println("list_tables:", list_tables.TableNames)

	// update_table
	//
	// 每次调整操作的间隔应大于10分钟
	// 如果是刚创建表，需要10分钟之后才能调整表的预留读写吞吐量。
	update_reserved_throughput := &OTSReservedThroughput{
		OTSCapacityUnit{5000, 5000},
	}

	fmt.Println("Need to sleep 12 Minute, be patient...")
	time.Sleep(12 * time.Minute)
	update_response, ots_err := ots_client.UpdateTable("myTable", update_reserved_throughput)
	if ots_err != nil {
		fmt.Println(ots_err)
		os.Exit(1)
	}
	fmt.Println("表的预留读吞吐量:", update_response.ReservedThroughputDetails.CapacityUnit.Read)
	fmt.Println("表的预留写吞吐量:", update_response.ReservedThroughputDetails.CapacityUnit.Write)
	fmt.Println("最后一次上调预留读写吞吐量时间:", update_response.ReservedThroughputDetails.LastIncreaseTime)
	fmt.Println("最后一次下调预留读写吞吐量时间:", update_response.ReservedThroughputDetails.LastDecreaseTime)
	fmt.Println("UTC自然日内总的下调预留读写吞吐量次数:", update_response.ReservedThroughputDetails.NumberOfDecreasesToday)

	// describe_table
	describe_response, ots_err := ots_client.DescribeTable("myTable")
	if ots_err != nil {
		fmt.Println(ots_err)
		os.Exit(1)
	}
	fmt.Println("表的名称:", describe_response.TableMeta.TableName)
	fmt.Println("表的主键:", describe_response.TableMeta.SchemaOfPrimaryKey)
	fmt.Println("表的预留读吞吐量:", describe_response.ReservedThroughputDetails.CapacityUnit.Read)
	fmt.Println("表的预留写吞吐量:", describe_response.ReservedThroughputDetails.CapacityUnit.Write)
	fmt.Println("最后一次上调预留读写吞吐量时间:", describe_response.ReservedThroughputDetails.LastIncreaseTime)
	fmt.Println("最后一次下调预留读写吞吐量时间:", describe_response.ReservedThroughputDetails.LastDecreaseTime)
	fmt.Println("UTC自然日内总的下调预留读写吞吐量次数:", describe_response.ReservedThroughputDetails.NumberOfDecreasesToday)
}
