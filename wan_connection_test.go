package peplink

import (
	"context"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
)

func TestClient_StatusWanConnection(t *testing.T) {
	tests := []struct {
		name     string
		response string
		want     []WanStatus
		wantErr  bool
	}{
		{"happy",
			`{
				"stat": "ok",
				"response": {
				  "1": {
					"name": "WAN 1",
					"enable": true,
					"locked": false,
					"managementOnly": false,
					"statusLed": "red",
					"message": "No Cable Detected",
					"uptime": 0,
					"type": "ethernet",
					"virtualType": "ethernet",
					"virtual": false,
					"priority": 1,
					"bandwidthAllowanceMonitor": {
					  "enable": false,
					  "hasSmtp": false
					},
					"method": "dhcp",
					"routingMode": "NAT",
					"mtu": 1440
				  },
				  "2": {
					"name": "WAN 2",
					"enable": true,
					"locked": false,
					"managementOnly": false,
					"statusLed": "red",
					"message": "No Cable Detected",
					"uptime": 0,
					"type": "ethernet",
					"virtualType": "ethernet",
					"virtual": false,
					"priority": 1,
					"bandwidthAllowanceMonitor": {
					  "enable": false,
					  "hasSmtp": false
					},
					"method": "dhcp",
					"routingMode": "NAT",
					"mtu": 1440
				  },
				  "3": {
					"name": "Cellular 1",
					"enable": true,
					"locked": false,
					"managementOnly": false,
					"ip": "10.10.10.10",
					"statusLed": "green",
					"message": "Connected to Carrier1",
					"uptime": 3314261,
					"type": "cellular",
					"virtualType": "cellular",
					"virtual": false,
					"priority": 2,
					"cellular": {
					  "network": "LTE",
					  "mobileType": "LTE",
					  "modulePowerOn": true,
					  "sim": {
						"1": {
						  "active": true,
						  "simCardDetected": true,
						  "imsi": "111111111111",
						  "iccid": "1111111111111",
						  "autoApn": true,
						  "apn": "web.carrier1.de",
						  "bandwidthAllowanceMonitor": {
							"enable": false,
							"hasSmtp": false
						  }
						},
						"2": {
						  "active": false,
						  "simCardDetected": false,
						  "autoApn": true,
						  "bandwidthAllowanceMonitor": {
							"enable": false,
							"hasSmtp": false
						  }
						},
						"order": [
						  1,
						  2
						]
					  },
					  "speedfusionConnect5gLte": {
						"active": false,
						"iccid": "1111111111111"
					  },
					  "carrier": {
						"name": "Carrier1",
						"country": "Germany"
					  },
					  "carrierAggregation": false,
					  "signalLevel": 5,
					  "rat": [
						{
						  "name": "",
						  "band": [
							{
							  "name": "LTE Band 1 (2100 MHz)",
							  "channel": 100,
							  "channelWidth": "20 MHz",
							  "signal": {
								"rssi": -63,
								"sinr": 19.399999999999999,
								"rsrp": -90,
								"rsrq": -8.0
							  }
							}
						  ]
						}
					  ],
					  "imei": "1111111111111",
					  "meid": {
						"hex": "",
						"dec": ""
					  },
					  "esn": "",
					  "dataTechnology": "LTE",
					  "mcc": "262",
					  "mnc": "02",
					  "cellTower": {
						"cellId": 11111,
						"cellPlmn": 22222,
						"cellUtranId": 333333,
						"tac": 33333
					  },
					  "model": "Model1",
					  "firmware": "VF.008"
					},
					"dns": [
					  "10.10.1.1",
					  "10.10.1.2"
					],
					"mask": 31,
					"gateway": "10.10.1.1",
					"bandwidthAllowanceMonitor": {
					  "enable": false,
					  "hasSmtp": false
					},
					"method": "dhcp",
					"routingMode": "NAT",
					"mtu": 1428
				  },
				  "4": {
					"name": "Cellular 2",
					"enable": true,
					"locked": false,
					"managementOnly": false,
					"ip": "10.10.10.11",
					"statusLed": "green",
					"message": "Connected to Carrier2",
					"uptime": 330759,
					"type": "cellular",
					"virtualType": "cellular",
					"virtual": false,
					"priority": 2,
					"cellular": {
					  "network": "LTE",
					  "mobileType": "LTE",
					  "modulePowerOn": true,
					  "sim": {
						"1": {
						  "active": true,
						  "simCardDetected": true,
						  "imsi": "1111111111",
						  "iccid": "22222222",
						  "mtn": "3333333333",
						  "autoApn": true,
						  "apn": "wap.carreir2.de",
						  "username": "user",
						  "password": "pass",
						  "bandwidthAllowanceMonitor": {
							"enable": false,
							"hasSmtp": false
						  }
						},
						"2": {
						  "active": false,
						  "simCardDetected": false,
						  "autoApn": true,
						  "bandwidthAllowanceMonitor": {
							"enable": false,
							"hasSmtp": false
						  }
						},
						"order": [
						  1,
						  2
						]
					  },
					  "speedfusionConnect5gLte": {
						"active": false,
						"iccid": "1111111111"
					  },
					  "carrier": {
						"name": "Carrier2",
						"country": "Germany"
					  },
					  "carrierAggregation": false,
					  "signalLevel": 5,
					  "rat": [
						{
						  "name": "",
						  "band": [
							{
							  "name": "LTE Band 3 (1800 MHz)",
							  "channel": 1300,
							  "channelWidth": "20 MHz",
							  "signal": {
								"rssi": -43,
								"sinr": 12.199999999999999,
								"rsrp": -76,
								"rsrq": -12.4
							  }
							}
						  ]
						}
					  ],
					  "imei": "11111111111",
					  "meid": {
						"hex": "",
						"dec": ""
					  },
					  "esn": "",
					  "dataTechnology": "LTE",
					  "mcc": "262",
					  "mnc": "01",
					  "cellTower": {
						"cellId": 11111,
						"cellPlmn": 22222,
						"cellUtranId": 333333,
						"tac": 1111
					  },
					  "manufacturer": "Example",
					  "model": "Model2"
					},
					"dns": [
					  "10.10.12.13",
					  "10.10.12.11"
					],
					"mask": 31,
					"gateway": "10.10.12.13",
					"bandwidthAllowanceMonitor": {
					  "enable": false,
					  "hasSmtp": false
					},
					"method": "dhcp",
					"routingMode": "NAT",
					"mtu": 1428
				  },
				  "5": {
					"name": "USB",
					"enable": true,
					"locked": false,
					"managementOnly": false,
					"statusLed": "empty",
					"message": "No Device Detected",
					"uptime": 0,
					"type": "modem",
					"virtualType": "modem",
					"virtual": false,
					"priority": 2,
					"bandwidthAllowanceMonitor": {
					  "enable": false,
					  "hasSmtp": false
					},
					"method": "ppp",
					"routingMode": "NAT",
					"mtu": 1428
				  },
				  "6": {
					"name": "Wi-Fi WAN on 2.4 GHz",
					"enable": false,
					"locked": false,
					"managementOnly": false,
					"statusLed": "gray",
					"message": "Disabled",
					"uptime": 0,
					"type": "wifi",
					"virtualType": "wifi",
					"virtual": false,
					"bandwidthAllowanceMonitor": {
					  "enable": false,
					  "hasSmtp": false
					},
					"method": "unknown",
					"routingMode": "NAT",
					"mtu": 1500
				  },
				  "7": {
					"name": "Wi-Fi WAN on 5 GHz",
					"enable": false,
					"locked": false,
					"managementOnly": false,
					"statusLed": "gray",
					"message": "Disabled",
					"uptime": 0,
					"type": "wifi",
					"virtualType": "wifi",
					"virtual": false,
					"bandwidthAllowanceMonitor": {
					  "enable": false,
					  "hasSmtp": false
					},
					"method": "unknown",
					"routingMode": "NAT",
					"mtu": 1500
				  },
				  "8": {
					"name": "VLAN WAN 1",
					"enable": false,
					"locked": false,
					"managementOnly": false,
					"statusLed": "gray",
					"message": "Disabled",
					"uptime": 0,
					"type": "wovlan",
					"virtualType": "wovlan",
					"virtual": false,
					"bandwidthAllowanceMonitor": {
					  "enable": false,
					  "hasSmtp": false
					},
					"method": "dhcp",
					"routingMode": "NAT"
				  },
				  "order": [
					1,
					2,
					3,
					4,
					5,
					6,
					7,
					8
				  ],
				  "timestamp": 1693680864,
				  "reportTimestamp": 1693680864,
				  "supportGatewayProxy": true
				}
			  }`,

			[]WanStatus{
				{Name: "WAN 1", StatusLed: "red", AsLan: false, Enable: true, Locked: false, ScheduledOff: false, Message: "No Cable Detected", Uptime: 0, Type: "ethernet", VirtualType: "ethernet", Priority: 1, Groupset: 0, Ip: "", Mask: 0, Gateway: "", Method: "dhcp", Mode: "", Dns: []string(nil), AditionalIp: []string(nil), MTU: 1440, MSS: 0, Mac: "", Wireless: WifiInfo{SSID: "", BSSID: "", Signal: Signal{RSSI: 0, SINR: 0, SNR: 0, ECIO: 0, RSRP: 0, RSRQ: 0, Strength: 0}}, Modem: ModemObj{Name: "", VendorID: 0, ProductID: 0, Manufacturer: "", Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, Network: "", MobileType: "", IMSIStr: "", ICCID: []string(nil), ESN: []string(nil), MTN: []string(nil), APN: "", Username: "", Password: "", DialNumber: "", Band: []BandObj(nil), Gobi: GobiObj{RoamingStatus: RoamingObj{Code: 0, Message: ""}, Network: "", MobileType: "", SIM: SIMGroupObj{Order: []int(nil)}, RemoteSIM: RemoteSIMObj{IMSI: "", SerialNumber: "", Slot: 0, AutoApp: false, APN: "", Username: "", Password: ""}, Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, MEID: MEIDObj{Hex: "", Dec: ""}, IMEI: "", ESN: "", Mode: "", Band: []BandObj(nil), MCC: "", MNC: "", CellTower: CellTowerObj{CellID: 0}}}, Cellular: GobiObj{RoamingStatus: RoamingObj{Code: 0, Message: ""}, Network: "", MobileType: "", SIM: SIMGroupObj{Order: []int(nil)}, RemoteSIM: RemoteSIMObj{IMSI: "", SerialNumber: "", Slot: 0, AutoApp: false, APN: "", Username: "", Password: ""}, Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, MEID: MEIDObj{Hex: "", Dec: ""}, IMEI: "", ESN: "", Mode: "", Band: []BandObj(nil), MCC: "", MNC: "", CellTower: CellTowerObj{CellID: 0}}, Gobi: GobiObj{RoamingStatus: RoamingObj{Code: 0, Message: ""}, Network: "", MobileType: "", SIM: SIMGroupObj{Order: []int(nil)}, RemoteSIM: RemoteSIMObj{IMSI: "", SerialNumber: "", Slot: 0, AutoApp: false, APN: "", Username: "", Password: ""}, Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, MEID: MEIDObj{Hex: "", Dec: ""}, IMEI: "", ESN: "", Mode: "", Band: []BandObj(nil), MCC: "", MNC: "", CellTower: CellTowerObj{CellID: 0}}},
				{Name: "WAN 2", StatusLed: "red", AsLan: false, Enable: true, Locked: false, ScheduledOff: false, Message: "No Cable Detected", Uptime: 0, Type: "ethernet", VirtualType: "ethernet", Priority: 1, Groupset: 0, Ip: "", Mask: 0, Gateway: "", Method: "dhcp", Mode: "", Dns: []string(nil), AditionalIp: []string(nil), MTU: 1440, MSS: 0, Mac: "", Wireless: WifiInfo{SSID: "", BSSID: "", Signal: Signal{RSSI: 0, SINR: 0, SNR: 0, ECIO: 0, RSRP: 0, RSRQ: 0, Strength: 0}}, Modem: ModemObj{Name: "", VendorID: 0, ProductID: 0, Manufacturer: "", Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, Network: "", MobileType: "", IMSIStr: "", ICCID: []string(nil), ESN: []string(nil), MTN: []string(nil), APN: "", Username: "", Password: "", DialNumber: "", Band: []BandObj(nil), Gobi: GobiObj{RoamingStatus: RoamingObj{Code: 0, Message: ""}, Network: "", MobileType: "", SIM: SIMGroupObj{Order: []int(nil)}, RemoteSIM: RemoteSIMObj{IMSI: "", SerialNumber: "", Slot: 0, AutoApp: false, APN: "", Username: "", Password: ""}, Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, MEID: MEIDObj{Hex: "", Dec: ""}, IMEI: "", ESN: "", Mode: "", Band: []BandObj(nil), MCC: "", MNC: "", CellTower: CellTowerObj{CellID: 0}}}, Cellular: GobiObj{RoamingStatus: RoamingObj{Code: 0, Message: ""}, Network: "", MobileType: "", SIM: SIMGroupObj{Order: []int(nil)}, RemoteSIM: RemoteSIMObj{IMSI: "", SerialNumber: "", Slot: 0, AutoApp: false, APN: "", Username: "", Password: ""}, Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, MEID: MEIDObj{Hex: "", Dec: ""}, IMEI: "", ESN: "", Mode: "", Band: []BandObj(nil), MCC: "", MNC: "", CellTower: CellTowerObj{CellID: 0}}, Gobi: GobiObj{RoamingStatus: RoamingObj{Code: 0, Message: ""}, Network: "", MobileType: "", SIM: SIMGroupObj{Order: []int(nil)}, RemoteSIM: RemoteSIMObj{IMSI: "", SerialNumber: "", Slot: 0, AutoApp: false, APN: "", Username: "", Password: ""}, Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, MEID: MEIDObj{Hex: "", Dec: ""}, IMEI: "", ESN: "", Mode: "", Band: []BandObj(nil), MCC: "", MNC: "", CellTower: CellTowerObj{CellID: 0}}},
				{Name: "Cellular 1", StatusLed: "green", AsLan: false, Enable: true, Locked: false, ScheduledOff: false, Message: "Connected to Carrier1", Uptime: 3314261, Type: "cellular", VirtualType: "cellular", Priority: 2, Groupset: 0, Ip: "10.10.10.10", Mask: 31, Gateway: "10.10.1.1", Method: "dhcp", Mode: "", Dns: []string{"10.10.1.1", "10.10.1.2"}, AditionalIp: []string(nil), MTU: 1428, MSS: 0, Mac: "", Wireless: WifiInfo{SSID: "", BSSID: "", Signal: Signal{RSSI: 0, SINR: 0, SNR: 0, ECIO: 0, RSRP: 0, RSRQ: 0, Strength: 0}}, Modem: ModemObj{Name: "", VendorID: 0, ProductID: 0, Manufacturer: "", Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, Network: "", MobileType: "", IMSIStr: "", ICCID: []string(nil), ESN: []string(nil), MTN: []string(nil), APN: "", Username: "", Password: "", DialNumber: "", Band: []BandObj(nil), Gobi: GobiObj{RoamingStatus: RoamingObj{Code: 0, Message: ""}, Network: "", MobileType: "", SIM: SIMGroupObj{Order: []int(nil)}, RemoteSIM: RemoteSIMObj{IMSI: "", SerialNumber: "", Slot: 0, AutoApp: false, APN: "", Username: "", Password: ""}, Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, MEID: MEIDObj{Hex: "", Dec: ""}, IMEI: "", ESN: "", Mode: "", Band: []BandObj(nil), MCC: "", MNC: "", CellTower: CellTowerObj{CellID: 0}}}, Cellular: GobiObj{RoamingStatus: RoamingObj{Code: 0, Message: ""}, Network: "LTE", MobileType: "LTE", SIM: SIMGroupObj{Order: []int{1, 2}}, RemoteSIM: RemoteSIMObj{IMSI: "", SerialNumber: "", Slot: 0, AutoApp: false, APN: "", Username: "", Password: ""}, Carrier: CarrierObj{Name: "Carrier1", Country: "Germany"}, SignalLevel: 5, MEID: MEIDObj{Hex: "", Dec: ""}, IMEI: "1111111111111", ESN: "", Mode: "", Band: []BandObj(nil), MCC: "262", MNC: "02", CellTower: CellTowerObj{CellID: 11111}}, Gobi: GobiObj{RoamingStatus: RoamingObj{Code: 0, Message: ""}, Network: "", MobileType: "", SIM: SIMGroupObj{Order: []int(nil)}, RemoteSIM: RemoteSIMObj{IMSI: "", SerialNumber: "", Slot: 0, AutoApp: false, APN: "", Username: "", Password: ""}, Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, MEID: MEIDObj{Hex: "", Dec: ""}, IMEI: "", ESN: "", Mode: "", Band: []BandObj(nil), MCC: "", MNC: "", CellTower: CellTowerObj{CellID: 0}}},
				{Name: "Cellular 2", StatusLed: "green", AsLan: false, Enable: true, Locked: false, ScheduledOff: false, Message: "Connected to Carrier2", Uptime: 330759, Type: "cellular", VirtualType: "cellular", Priority: 2, Groupset: 0, Ip: "10.10.10.11", Mask: 31, Gateway: "10.10.12.13", Method: "dhcp", Mode: "", Dns: []string{"10.10.12.13", "10.10.12.11"}, AditionalIp: []string(nil), MTU: 1428, MSS: 0, Mac: "", Wireless: WifiInfo{SSID: "", BSSID: "", Signal: Signal{RSSI: 0, SINR: 0, SNR: 0, ECIO: 0, RSRP: 0, RSRQ: 0, Strength: 0}}, Modem: ModemObj{Name: "", VendorID: 0, ProductID: 0, Manufacturer: "", Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, Network: "", MobileType: "", IMSIStr: "", ICCID: []string(nil), ESN: []string(nil), MTN: []string(nil), APN: "", Username: "", Password: "", DialNumber: "", Band: []BandObj(nil), Gobi: GobiObj{RoamingStatus: RoamingObj{Code: 0, Message: ""}, Network: "", MobileType: "", SIM: SIMGroupObj{Order: []int(nil)}, RemoteSIM: RemoteSIMObj{IMSI: "", SerialNumber: "", Slot: 0, AutoApp: false, APN: "", Username: "", Password: ""}, Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, MEID: MEIDObj{Hex: "", Dec: ""}, IMEI: "", ESN: "", Mode: "", Band: []BandObj(nil), MCC: "", MNC: "", CellTower: CellTowerObj{CellID: 0}}}, Cellular: GobiObj{RoamingStatus: RoamingObj{Code: 0, Message: ""}, Network: "LTE", MobileType: "LTE", SIM: SIMGroupObj{Order: []int{1, 2}}, RemoteSIM: RemoteSIMObj{IMSI: "", SerialNumber: "", Slot: 0, AutoApp: false, APN: "", Username: "", Password: ""}, Carrier: CarrierObj{Name: "Carrier2", Country: "Germany"}, SignalLevel: 5, MEID: MEIDObj{Hex: "", Dec: ""}, IMEI: "11111111111", ESN: "", Mode: "", Band: []BandObj(nil), MCC: "262", MNC: "01", CellTower: CellTowerObj{CellID: 11111}}, Gobi: GobiObj{RoamingStatus: RoamingObj{Code: 0, Message: ""}, Network: "", MobileType: "", SIM: SIMGroupObj{Order: []int(nil)}, RemoteSIM: RemoteSIMObj{IMSI: "", SerialNumber: "", Slot: 0, AutoApp: false, APN: "", Username: "", Password: ""}, Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, MEID: MEIDObj{Hex: "", Dec: ""}, IMEI: "", ESN: "", Mode: "", Band: []BandObj(nil), MCC: "", MNC: "", CellTower: CellTowerObj{CellID: 0}}},
				{Name: "USB", StatusLed: "empty", AsLan: false, Enable: true, Locked: false, ScheduledOff: false, Message: "No Device Detected", Uptime: 0, Type: "modem", VirtualType: "modem", Priority: 2, Groupset: 0, Ip: "", Mask: 0, Gateway: "", Method: "ppp", Mode: "", Dns: []string(nil), AditionalIp: []string(nil), MTU: 1428, MSS: 0, Mac: "", Wireless: WifiInfo{SSID: "", BSSID: "", Signal: Signal{RSSI: 0, SINR: 0, SNR: 0, ECIO: 0, RSRP: 0, RSRQ: 0, Strength: 0}}, Modem: ModemObj{Name: "", VendorID: 0, ProductID: 0, Manufacturer: "", Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, Network: "", MobileType: "", IMSIStr: "", ICCID: []string(nil), ESN: []string(nil), MTN: []string(nil), APN: "", Username: "", Password: "", DialNumber: "", Band: []BandObj(nil), Gobi: GobiObj{RoamingStatus: RoamingObj{Code: 0, Message: ""}, Network: "", MobileType: "", SIM: SIMGroupObj{Order: []int(nil)}, RemoteSIM: RemoteSIMObj{IMSI: "", SerialNumber: "", Slot: 0, AutoApp: false, APN: "", Username: "", Password: ""}, Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, MEID: MEIDObj{Hex: "", Dec: ""}, IMEI: "", ESN: "", Mode: "", Band: []BandObj(nil), MCC: "", MNC: "", CellTower: CellTowerObj{CellID: 0}}}, Cellular: GobiObj{RoamingStatus: RoamingObj{Code: 0, Message: ""}, Network: "", MobileType: "", SIM: SIMGroupObj{Order: []int(nil)}, RemoteSIM: RemoteSIMObj{IMSI: "", SerialNumber: "", Slot: 0, AutoApp: false, APN: "", Username: "", Password: ""}, Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, MEID: MEIDObj{Hex: "", Dec: ""}, IMEI: "", ESN: "", Mode: "", Band: []BandObj(nil), MCC: "", MNC: "", CellTower: CellTowerObj{CellID: 0}}, Gobi: GobiObj{RoamingStatus: RoamingObj{Code: 0, Message: ""}, Network: "", MobileType: "", SIM: SIMGroupObj{Order: []int(nil)}, RemoteSIM: RemoteSIMObj{IMSI: "", SerialNumber: "", Slot: 0, AutoApp: false, APN: "", Username: "", Password: ""}, Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, MEID: MEIDObj{Hex: "", Dec: ""}, IMEI: "", ESN: "", Mode: "", Band: []BandObj(nil), MCC: "", MNC: "", CellTower: CellTowerObj{CellID: 0}}},
				{Name: "Wi-Fi WAN on 2.4 GHz", StatusLed: "gray", AsLan: false, Enable: false, Locked: false, ScheduledOff: false, Message: "Disabled", Uptime: 0, Type: "wifi", VirtualType: "wifi", Priority: 0, Groupset: 0, Ip: "", Mask: 0, Gateway: "", Method: "unknown", Mode: "", Dns: []string(nil), AditionalIp: []string(nil), MTU: 1500, MSS: 0, Mac: "", Wireless: WifiInfo{SSID: "", BSSID: "", Signal: Signal{RSSI: 0, SINR: 0, SNR: 0, ECIO: 0, RSRP: 0, RSRQ: 0, Strength: 0}}, Modem: ModemObj{Name: "", VendorID: 0, ProductID: 0, Manufacturer: "", Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, Network: "", MobileType: "", IMSIStr: "", ICCID: []string(nil), ESN: []string(nil), MTN: []string(nil), APN: "", Username: "", Password: "", DialNumber: "", Band: []BandObj(nil), Gobi: GobiObj{RoamingStatus: RoamingObj{Code: 0, Message: ""}, Network: "", MobileType: "", SIM: SIMGroupObj{Order: []int(nil)}, RemoteSIM: RemoteSIMObj{IMSI: "", SerialNumber: "", Slot: 0, AutoApp: false, APN: "", Username: "", Password: ""}, Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, MEID: MEIDObj{Hex: "", Dec: ""}, IMEI: "", ESN: "", Mode: "", Band: []BandObj(nil), MCC: "", MNC: "", CellTower: CellTowerObj{CellID: 0}}}, Cellular: GobiObj{RoamingStatus: RoamingObj{Code: 0, Message: ""}, Network: "", MobileType: "", SIM: SIMGroupObj{Order: []int(nil)}, RemoteSIM: RemoteSIMObj{IMSI: "", SerialNumber: "", Slot: 0, AutoApp: false, APN: "", Username: "", Password: ""}, Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, MEID: MEIDObj{Hex: "", Dec: ""}, IMEI: "", ESN: "", Mode: "", Band: []BandObj(nil), MCC: "", MNC: "", CellTower: CellTowerObj{CellID: 0}}, Gobi: GobiObj{RoamingStatus: RoamingObj{Code: 0, Message: ""}, Network: "", MobileType: "", SIM: SIMGroupObj{Order: []int(nil)}, RemoteSIM: RemoteSIMObj{IMSI: "", SerialNumber: "", Slot: 0, AutoApp: false, APN: "", Username: "", Password: ""}, Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, MEID: MEIDObj{Hex: "", Dec: ""}, IMEI: "", ESN: "", Mode: "", Band: []BandObj(nil), MCC: "", MNC: "", CellTower: CellTowerObj{CellID: 0}}},
				{Name: "Wi-Fi WAN on 5 GHz", StatusLed: "gray", AsLan: false, Enable: false, Locked: false, ScheduledOff: false, Message: "Disabled", Uptime: 0, Type: "wifi", VirtualType: "wifi", Priority: 0, Groupset: 0, Ip: "", Mask: 0, Gateway: "", Method: "unknown", Mode: "", Dns: []string(nil), AditionalIp: []string(nil), MTU: 1500, MSS: 0, Mac: "", Wireless: WifiInfo{SSID: "", BSSID: "", Signal: Signal{RSSI: 0, SINR: 0, SNR: 0, ECIO: 0, RSRP: 0, RSRQ: 0, Strength: 0}}, Modem: ModemObj{Name: "", VendorID: 0, ProductID: 0, Manufacturer: "", Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, Network: "", MobileType: "", IMSIStr: "", ICCID: []string(nil), ESN: []string(nil), MTN: []string(nil), APN: "", Username: "", Password: "", DialNumber: "", Band: []BandObj(nil), Gobi: GobiObj{RoamingStatus: RoamingObj{Code: 0, Message: ""}, Network: "", MobileType: "", SIM: SIMGroupObj{Order: []int(nil)}, RemoteSIM: RemoteSIMObj{IMSI: "", SerialNumber: "", Slot: 0, AutoApp: false, APN: "", Username: "", Password: ""}, Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, MEID: MEIDObj{Hex: "", Dec: ""}, IMEI: "", ESN: "", Mode: "", Band: []BandObj(nil), MCC: "", MNC: "", CellTower: CellTowerObj{CellID: 0}}}, Cellular: GobiObj{RoamingStatus: RoamingObj{Code: 0, Message: ""}, Network: "", MobileType: "", SIM: SIMGroupObj{Order: []int(nil)}, RemoteSIM: RemoteSIMObj{IMSI: "", SerialNumber: "", Slot: 0, AutoApp: false, APN: "", Username: "", Password: ""}, Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, MEID: MEIDObj{Hex: "", Dec: ""}, IMEI: "", ESN: "", Mode: "", Band: []BandObj(nil), MCC: "", MNC: "", CellTower: CellTowerObj{CellID: 0}}, Gobi: GobiObj{RoamingStatus: RoamingObj{Code: 0, Message: ""}, Network: "", MobileType: "", SIM: SIMGroupObj{Order: []int(nil)}, RemoteSIM: RemoteSIMObj{IMSI: "", SerialNumber: "", Slot: 0, AutoApp: false, APN: "", Username: "", Password: ""}, Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, MEID: MEIDObj{Hex: "", Dec: ""}, IMEI: "", ESN: "", Mode: "", Band: []BandObj(nil), MCC: "", MNC: "", CellTower: CellTowerObj{CellID: 0}}},
				{Name: "VLAN WAN 1", StatusLed: "gray", AsLan: false, Enable: false, Locked: false, ScheduledOff: false, Message: "Disabled", Uptime: 0, Type: "wovlan", VirtualType: "wovlan", Priority: 0, Groupset: 0, Ip: "", Mask: 0, Gateway: "", Method: "dhcp", Mode: "", Dns: []string(nil), AditionalIp: []string(nil), MTU: 0, MSS: 0, Mac: "", Wireless: WifiInfo{SSID: "", BSSID: "", Signal: Signal{RSSI: 0, SINR: 0, SNR: 0, ECIO: 0, RSRP: 0, RSRQ: 0, Strength: 0}}, Modem: ModemObj{Name: "", VendorID: 0, ProductID: 0, Manufacturer: "", Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, Network: "", MobileType: "", IMSIStr: "", ICCID: []string(nil), ESN: []string(nil), MTN: []string(nil), APN: "", Username: "", Password: "", DialNumber: "", Band: []BandObj(nil), Gobi: GobiObj{RoamingStatus: RoamingObj{Code: 0, Message: ""}, Network: "", MobileType: "", SIM: SIMGroupObj{Order: []int(nil)}, RemoteSIM: RemoteSIMObj{IMSI: "", SerialNumber: "", Slot: 0, AutoApp: false, APN: "", Username: "", Password: ""}, Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, MEID: MEIDObj{Hex: "", Dec: ""}, IMEI: "", ESN: "", Mode: "", Band: []BandObj(nil), MCC: "", MNC: "", CellTower: CellTowerObj{CellID: 0}}}, Cellular: GobiObj{RoamingStatus: RoamingObj{Code: 0, Message: ""}, Network: "", MobileType: "", SIM: SIMGroupObj{Order: []int(nil)}, RemoteSIM: RemoteSIMObj{IMSI: "", SerialNumber: "", Slot: 0, AutoApp: false, APN: "", Username: "", Password: ""}, Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, MEID: MEIDObj{Hex: "", Dec: ""}, IMEI: "", ESN: "", Mode: "", Band: []BandObj(nil), MCC: "", MNC: "", CellTower: CellTowerObj{CellID: 0}}, Gobi: GobiObj{RoamingStatus: RoamingObj{Code: 0, Message: ""}, Network: "", MobileType: "", SIM: SIMGroupObj{Order: []int(nil)}, RemoteSIM: RemoteSIMObj{IMSI: "", SerialNumber: "", Slot: 0, AutoApp: false, APN: "", Username: "", Password: ""}, Carrier: CarrierObj{Name: "", Country: ""}, SignalLevel: 0, MEID: MEIDObj{Hex: "", Dec: ""}, IMEI: "", ESN: "", Mode: "", Band: []BandObj(nil), MCC: "", MNC: "", CellTower: CellTowerObj{CellID: 0}}},
			},
			false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					require.Equal(t, "/api/status.wan.connection", r.URL.Path)
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(tt.response))
				}),
			)

			defer srv.Close()

			c := Client{
				httpClient: resty.New().
					SetBaseURL(srv.URL).
					SetHeader("Content-Type", "application/json").
					SetHeader("Accept", "application/json"),
				log: slog.Default(),
			}

			got, err := c.StatusWanConnection(context.Background())
			require.Equal(t, tt.wantErr, err != nil, "FirmwareVersion() error = %v, wantErr %v", err, tt.wantErr)
			require.Equal(t, tt.want, got)
		})
	}
}
