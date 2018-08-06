# influxdb-line-protocol-to-json
This library is useful for decoding InfluxDB line protocol to a Go map instance; also the line protocol can be converted to JSON

<code> <measurement>[,<tag_key>=<tag_value>[,<tag_key>=<tag_value>]] <field_key>=<field_value>[,<field_key>=<field_value>] [<timestamp>] </code>
