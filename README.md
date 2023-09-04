# peplink-client
The API client for the Peplink devices
Information could be queried from the device via the HTTP API or via SNMP

## References
- [Peplink API Documentation](https://download.peplink.com/resources/Peplink-Router-API-Documentation-for-Firmware-8.1.1.pdf)
- [Peplink SNMP MIB](https://download.peplink.com/resources/balance_max_snmp_mib-8.2.0.zip)

## supported HTTP endpoints
- [x] /api/auth.token.grant
- [x] /api/info.frw.version
- [x] /api/status.wan.connection

## supported SNMP OIDs
- [ ] serial number