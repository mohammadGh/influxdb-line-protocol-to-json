# influxdb-line-protocol-to-json
This library is useful for decoding InfluxDB line protocol to a Go map instance; also the line protocol can be converted to JSON
The InfluxDB Line Protocol is as follow:
    
    <measurement>[,<tag_key>=<tag_value>[,<tag_key>=<tag_value>]] <field_key>=<field_value>[,<field_key>=<field_value>] [<timestamp>] 
    

for more information refer to [InfluxDB Line Protocol documentation](https://docs.influxdata.com/influxdb/v1.6/write_protocols/line_protocol_reference)
