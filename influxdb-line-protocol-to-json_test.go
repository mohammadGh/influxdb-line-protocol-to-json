package influxdb_line_protocol_to_json

import (
	"fmt"
	"testing"
)

func TestLineToJson(t *testing.T) {
	var line = `win_cpu,host=DESKTOP-RLUGATJ,instance=3,objectname=Processor Percent_DPC_Time=0,Percent_Idle_Time=99.50546264648438,Percent_Interrupt_Time=0,Percent_Privileged_Time=0,Percent_Processor_Time=2.4173243045806885,Percent_User_Time=0 1533206781000000000`
	fmt.Println(LineToJson(line))
}

func TestLinesToJson(t *testing.T) {
	var line = `
win_cpu,host=DESKTOP-RLUGATJ,instance=3,objectname=Processor Percent_DPC_Time=0,Percent_Idle_Time=99.50546264648438,Percent_Interrupt_Time=0,Percent_Privileged_Time=0,Percent_Processor_Time=2.4173243045806885,Percent_User_Time=0 1533206781000000000
win_cpu,host=DESKTOP-RLUGATJ,instance=4,objectname=Processor Percent_DPC_Time=0,Percent_Idle_Time=98.90975952148438,Percent_Interrupt_Time=0,Percent_Privileged_Time=0,Percent_Processor_Time=0.8683928847312927,Percent_User_Time=0 1533206781000000000
win_cpu,host=DESKTOP-RLUGATJ,instance=7,objectname=Processor Percent_DPC_Time=0,Percent_Idle_Time=99.34410858154297,Percent_Interrupt_Time=0,Percent_Privileged_Time=0,Percent_Processor_Time=0.8683928847312927,Percent_User_Time=0 1533206781000000000
win_system,host=DESKTOP-RLUGATJ,objectname=System Context_Switches_persec=1388.9049072265625,Processor_Queue_Length=0,System_Calls_persec=2906.779541015625,System_Up_Time=8072.9501953125 1533206781000000000
win_mem,host=DESKTOP-RLUGATJ,objectname=Memory Available_Bytes=14458130432,Cache_Faults_persec=1.9841499328613281,Demand_Zero_Faults_persec=15.873199462890625,Page_Faults_persec=17.857349395751953,Pages_persec=0.9920749664306641,Pool_Nonpaged_Bytes=222699520,Pool_Paged_Bytes=517771264,Standby_Cache_Core_Bytes=264531968,Standby_Cache_Normal_Priority_Bytes=1566867456,Standby_Cache_Reserve_Bytes=7005278208,Transition_Faults_persec=0 1533206781000000000
win_net,host=DESKTOP-RLUGATJ,instance=D-Link\ DWA-525\ Wireless\ N\ 150\ Desktop\ Adapter[rev.A2],objectname=Network\ Interface Bytes_Received_persec=0,Bytes_Sent_persec=0,Packets_Outbound_Discarded=0,Packets_Outbound_Errors=0,Packets_Received_Discarded=0,Packets_Received_Errors=0,Packets_Received_persec=0,Packets_Sent_persec=0 1533206781000000000
win_net,host=DESKTOP-RLUGATJ,instance=Teredo\ Tunneling\ Pseudo-Interface,objectname=Network\ Interface Bytes_Received_persec=0,Bytes_Sent_persec=0,Packets_Outbound_Discarded=0,Packets_Outbound_Errors=0,Packets_Received_Discarded=0,Packets_Received_Errors=0,Packets_Received_persec=0,Packets_Sent_persec=0 1533206781000000000`
	fmt.Println(LinesToJson(line))
}

func TestExtractPorotcolSection(t *testing.T) {
	var line = `win_cpu,host=DESKTOP-RLUGATJ,instance=3,objectname=Processor Percent_DPC_Time=0,Percent_Idle_Time=99.50546264648438,Percent_Interrupt_Time=0,Percent_Privileged_Time=0,Percent_Processor_Time=2.4173243045806885,Percent_User_Time=0 1533206781000000000`
	a, b, c := extractMainSections(line)
	fmt.Printf("%s\n%s\n%s\n", a, b, c)
}

func TestExtractPorotcolSectionWithEscapeCharacters(t *testing.T) {
	var line = `win_perf_counters,host=DESKTOP-RLUGATJ,instance=D-Link\ DWA-525\ Wireless\ N\ 150\ Desktop\ Adapter[rev.A2],objectname=Network\ Interface Bytes_Received_persec=0,Bytes_Sent_persec=0,Packets_Outbound_Discarded=0,Packets_Outbound_Errors=0,Packets_Received_Discarded=0,Packets_Received_Errors=0,Packets_Received_persec=0,Packets_Sent_persec=0 1533206781000000000`
	a, b, c := extractMainSections(line)
	fmt.Printf("%s\n%s\n%s\n", a, b, c)
}
