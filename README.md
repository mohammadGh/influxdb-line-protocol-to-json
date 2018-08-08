# InfluxDB Line Protocol to JSON Converter
This library is useful for decoding `InfluxDB Line Protocol` to a `Map` instance or `JSON` string.
The InfluxDB Line Protocol is as follow:
```
<measurement>[,<tag_key>=<tag_value>[,<tag_key>=<tag_value>]] <field_key>=<field_value>[,<field_key>=<field_value>] [<timestamp>] 
```    
for more information refer to [InfluxDB Line Protocol documentation](https://docs.influxdata.com/influxdb/v1.6/write_protocols/line_protocol_reference)

The follwoing example show how library converts the `Line Protocol` data to the JSON structure:
```json
[
   {
      "measurement":"win_cpu",
        "fields":[
         {
            "key":"Percent_DPC_Time",
            "value":"0"
         },
         {
            "key":"Percent_Idle_Time",
            "value":"99.50546264648438"
         },
         {
            "key":"Percent_Interrupt_Time",
            "value":"0"
         },
         {
            "key":"Percent_Privileged_Time",
            "value":"0"
         },
         {
            "key":"Percent_Processor_Time",
            "value":"2.4173243045806885"
         },
         {
            "key":"Percent_User_Time",
            "value":"0"
         }
      ],
      "tags":[
         {
            "key":"host",
            "value":"DESKTOP-RLUGATJ"
         },
         {
            "key":"instance",
            "value":"3"
         },
         {
            "key":"objectname",
            "value":"Processor"
         }
      ],
      "timestamp":"1533206781000000000"
   },
   {
      "fields":[...],
      "measurement":"win_disk",
      "tags":[...],
      "timestamp":"1533208293000000000"
   },
   .
   .
   .
]
```
## Install
Just add the `github/mohammadGh/influxdb-line-protocol-to-json` to the import section of your `Go` file. Then run the follwing command:
```
go get -d .
```
## Usage
Just add the package name into the import section and the related methods:
```
package main

import (
	"fmt"
	inlfuxjson "github/mohammadGh/influxdb-line-protocol-to-json"
)

func main() {
	var lines = `
win_cpu,host=DESKTOP-RLUGATJ,instance=3,objectname=Processor Percent_DPC_Time=0,Percent_Idle_Time=99.50546264648438,Percent_Interrupt_Time=0,Percent_Privileged_Time=0,Percent_Processor_Time=2.4173243045806885,Percent_User_Time=0 1533206781000000000
win_cpu,host=DESKTOP-RLUGATJ,instance=4,objectname=Processor Percent_DPC_Time=0,Percent_Idle_Time=98.90975952148438,Percent_Interrupt_Time=0,Percent_Privileged_Time=0,Percent_Processor_Time=0.8683928847312927,Percent_User_Time=0 1533206781000000000
win_cpu,host=DESKTOP-RLUGATJ,instance=7,objectname=Processor Percent_DPC_Time=0,Percent_Idle_Time=99.34410858154297,Percent_Interrupt_Time=0,Percent_Privileged_Time=0,Percent_Processor_Time=0.8683928847312927,Percent_User_Time=0 1533206781000000000
win_system,host=DESKTOP-RLUGATJ,objectname=System Context_Switches_persec=1388.9049072265625,Processor_Queue_Length=0,System_Calls_persec=2906.779541015625,System_Up_Time=8072.9501953125 1533206781000000000
win_mem,host=DESKTOP-RLUGATJ,objectname=Memory Available_Bytes=14458130432,Cache_Faults_persec=1.9841499328613281,Demand_Zero_Faults_persec=15.873199462890625,Page_Faults_persec=17.857349395751953,Pages_persec=0.9920749664306641,Pool_Nonpaged_Bytes=222699520,Pool_Paged_Bytes=517771264,Standby_Cache_Core_Bytes=264531968,Standby_Cache_Normal_Priority_Bytes=1566867456,Standby_Cache_Reserve_Bytes=7005278208,Transition_Faults_persec=0 1533206781000000000
win_net,host=DESKTOP-RLUGATJ,instance=D-Link\ DWA-525\ Wireless\ N\ 150\ Desktop\ Adapter[rev.A2],objectname=Network\ Interface Bytes_Received_persec=0,Bytes_Sent_persec=0,Packets_Outbound_Discarded=0,Packets_Outbound_Errors=0,Packets_Received_Discarded=0,Packets_Received_Errors=0,Packets_Received_persec=0,Packets_Sent_persec=0 1533206781000000000
win_net,host=DESKTOP-RLUGATJ,instance=Teredo\ Tunneling\ Pseudo-Interface,objectname=Network\ Interface Bytes_Received_persec=0,Bytes_Sent_persec=0,Packets_Outbound_Discarded=0,Packets_Outbound_Errors=0,Packets_Received_Discarded=0,Packets_Received_Errors=0,Packets_Received_persec=0,Packets_Sent_persec=0 1533206781000000000`
	
	fmt.Println(inlfuxjson.LinesToJson(line))
}

```
