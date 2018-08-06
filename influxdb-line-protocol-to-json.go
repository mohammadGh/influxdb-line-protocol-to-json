package influxdb_line_protocol_to_json

import (
	"encoding/json"
	"strings"
)

func LinesToJsonBytes(linesString string) []byte {
	jsonStr, _ := json.Marshal(LinesToMapList(linesString))
	return jsonStr
}

func LinesToJson(linesString string) string {
	jsonString, _ := json.Marshal(LinesToMapList(linesString))
	return string(jsonString)
}

func LineToJsonBytes(lineStr string) []byte {
	jsonStr, _ := json.Marshal(LineToMap(lineStr))
	return jsonStr
}

func LineToJson(lineStr string) string {
	jsonString, _ := json.Marshal(LineToMap(lineStr))
	return string(jsonString)
}

func LinesToMapList(linesStr string) []map[string]interface{} {
	lines := strings.Split(linesStr, "\n")
	var result []map[string]interface{}
	for _, line := range lines {
		if line != "" {
			result = append(result, LineToMap(line))
		}
	}
	return result
}

func LineToMap(lineStr string) map[string]interface{} {
	var jsonMap map[string]interface{}
	jsonMap = make(map[string]interface{})

	//decode three main section of Line Protocol
	tags, fields, timestamp := extractMainSections(lineStr)

	//decode metric name and tags
	if strings.Contains(tags, ",") {
		jsonMap["measurement"] = tags[:strings.Index(tags, ",")]
		jsonMap["tags"] = decodeAsKeyValueList(tags[strings.Index(tags, ",")+1:], ",", "=")
	} else {
		jsonMap["measurement"] = tags
	}
	//decode fields and set time stamp
	jsonMap["fields"] = decodeAsKeyValueList(fields, ",", "=")
	jsonMap["timestamp"] = timestamp
	return jsonMap
}

func decodeAsKeyValueList(txt string, itemSplitter string, keyValueSplitter string) []map[string]interface{} {
	var result []map[string]interface{}
	items := strings.Split(txt, itemSplitter)
	for _, item := range items {
		var decodedItem = make(map[string]interface{})
		decodedItem["key"] = item[:strings.Index(item, keyValueSplitter)]
		decodedItem["value"] = item[strings.Index(item, keyValueSplitter)+1:]
		result = append(result, decodedItem)
	}
	return result
}

func extractMainSections(line string) (string, string, string) {
	var endOfTagSection = 0
	var tagSection, fieldSection, timestamp string

	//extract tag section
	for index, chr := range line {
		if chr == ' ' && line[index-1] != '\\' {
			if tagSection == "" {
				tagSection = line[:index]
				endOfTagSection = index
				continue
			}
			fieldSection = line[endOfTagSection+1 : index]
			timestamp = line[index+1:]
		}
	}
	return tagSection, fieldSection, timestamp
}
