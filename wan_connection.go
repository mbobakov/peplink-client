package peplink

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmespath/go-jmespath"
)

type WanStatus struct {
	// Name of the WAN connection
	Name string `json:"name"`
	// LED color for UI { empty, gray, red, yellow, green, flash }
	StatusLed string `json:"statusLed"`
	// WAN port is performing WAN as LAN or not
	AsLan bool `json:"asLan"`
	//WAN is enabled or not
	Enable bool `json:"enable"`
	// WAN is locked or not
	Locked bool `json:"locked"`
	// Only appear if Connection is scheduled and currently off
	ScheduledOff bool `json:"scheduledOff"`
	// WAN status message
	Message string `json:"message"`
	// WAN connection uptime in seconds
	Uptime int `json:"uptime"`
	// WAN connection type
	// For cellular WAN
	// In fw8.0.1 or later, it will return “cellular”.
	// Before fw8.0.1, it will return “gobi”
	// { modem, wireless, gobi, cellular, ipsec, adsl, ethernet }
	Type string `json:"type"`

	// For cellular WAN
	// In fw8.0.1 or later, it will return “cellular”.
	// Before fw8.0.1, it will return “gobi”
	// { modem, wireless, gobi, cellular, ipsec, adsl, ethernet }
	VirtualType string `json:"virtualType"`
	// Priority of the WAN. The field will not appear if the WAN is disabled
	Priority int `json:"priority"`
	//Group set of the WAN connection
	Groupset int `json:"groupset"`
	// IP address
	Ip string `json:"ip"`
	// Subnet mask. The field will not appear if ip is not exist or lite=yes
	Mask int `json:"mask"`
	// Gateway. The field will not appear if ip is not exist or lite=yes
	Gateway string `json:"gateway"`
	// Connection method, DHCP or Static IP. The field will not appear if lite=yes
	//{ dhcp static }
	Method string `json:"method"`
	// Connection mode. The field will not appear if lite=yes
	// { NAT, IP Forwarding }
	Mode string `json:"mode"`
	// DNS Server list. The field will not appear if lite=yes
	Dns []string `json:"dns"`
	// Additional IP address list. The field will not appear if lite=yes
	AditionalIp []string `json:"aditionalIp"`
	// MTU value. The field will not appear if auto or lite=yes
	// [576, 9000]
	MTU int `json:"mtu"`
	// MSS value. The field will not appear if auto or lite=yes
	// [536, 8960]
	MSS int `json:"mss"`
	// MAC address. The field will not appear if lite=yes
	Mac string `json:"mac"`
	// WAN connection detail for wireless. The field will only appear if type is wifi
	Wireless WifiInfo `json:"wireless"`
	// WAN connection detail for modem. The field will only appear if type is modem
	Modem ModemObj `json:"modem"`
	// WAN connection detail for gobi. The field will only appear if type is cellular
	Cellular GobiObj `json:"cellular"`
	//WAN connection detail for gobi. The field will only appear if type is gob NOTE: This object is deprecated in firmware 8.0.1.
	Gobi GobiObj `json:"gobi"`
}

// Signal represents the signal information.
type Signal struct {
	RSSI     int     `json:"rssi,omitempty"`     // Received Signal Strength Indicator (RSSI), only appear in Gobi and Modem
	SINR     float64 `json:"sinr,omitempty"`     // Signal to Interference plus Noise Ratio (SINR), only appear in Gobi and Modem
	SNR      float64 `json:"snr,omitempty"`      // Signal-to-noise ratio (SNR), only appear in Gobi and has value
	ECIO     float64 `json:"ecio,omitempty"`     // Energy to Interference Ratio (Ec/Io), only appear in Gobi and has value
	RSRP     float64 `json:"rsrp,omitempty"`     // Reference Signal Received Power (RSRP), only appear in Gobi and Modem
	RSRQ     float64 `json:"rsrq,omitempty"`     // Reference Signal Received Quality (RSRQ), only appear in Gobi
	Strength float64 `json:"strength,omitempty"` // Wi-Fi signal strength, only appear in Wifi
}

// WifiInfo represents information about Wi-Fi networks.
type WifiInfo struct {
	SSID   string `json:"ssid"`   // SSID of the Wifi. The field will not appear if lite=yes
	BSSID  string `json:"bssid"`  // BSSID. The field will not appear if lite=yes
	Signal Signal `json:"signal"` // Signal information
}

// ModemObj represents modem adaptor information.
type ModemObj struct {
	Name         string     `json:"name"`                 // Modem adaptor name
	VendorID     int        `json:"vendorId"`             // Modem adaptor vendor ID
	ProductID    int        `json:"productId"`            // Modem adaptor product ID
	Manufacturer string     `json:"manufacturer"`         // Modem adaptor manufacturer
	Carrier      CarrierObj `json:"carrier"`              // Carrier Information
	SignalLevel  int        `json:"signalLevel"`          // Signal level [0,5]
	Network      string     `json:"network"`              // Network name (deprecated in fw8.0.1)
	MobileType   string     `json:"mobileType"`           // Network name (use "mobileType" in fw8.0.1 or later)
	IMSIStr      string     `json:"imsi,omitempty"`       // International Mobile Subscriber Identity (IMSI)
	ICCID        []string   `json:"iccid,omitempty"`      // Integrated Circuit Card Identity (ICCID)
	ESN          []string   `json:"esn,omitempty"`        // Electronic Serial Number (ESN)
	MTN          []string   `json:"mtn,omitempty"`        // Mobile Telecommunications Network (MTN)
	APN          string     `json:"apn,omitempty"`        // APN
	Username     string     `json:"username,omitempty"`   // Username for APN
	Password     string     `json:"password,omitempty"`   // Password for APN
	DialNumber   string     `json:"dialNumber,omitempty"` // Dial number for APN
	Band         []BandObj  `json:"band"`                 // Cellular band information
	Gobi         GobiObj    `json:"gobi"`                 // Gobi network information
}

// GobiObj represents Gobi network information.
type GobiObj struct {
	RoamingStatus RoamingObj   `json:"roamingStatus"`       // Roaming status information
	Network       string       `json:"network"`             // Network name (deprecated in fw8.0.1)
	MobileType    string       `json:"mobileType"`          // Network name (use "mobileType" in fw8.0.1 or later)
	SIM           SIMGroupObj  `json:"sim"`                 // SIM information
	RemoteSIM     RemoteSIMObj `json:"remoteSim,omitempty"` // Remote SIM information (only when remote SIM is enabled)
	Carrier       CarrierObj   `json:"carrier"`             // Carrier information
	SignalLevel   int          `json:"signalLevel"`         // Signal Level [0,5]
	MEID          MEIDObj      `json:"meid,omitempty"`      // Hex and Dec value of Mobile Equipment Identifier (MEID)
	IMEI          string       `json:"imei,omitempty"`      // International Mobile Equipment Identity (IMEI)
	ESN           string       `json:"esn,omitempty"`       // Electronic Serial Number (ESN)
	Mode          string       `json:"mode,omitempty"`      // Gobi network mode
	Band          []BandObj    `json:"band"`                // Gobi band information
	MCC           string       `json:"mcc,omitempty"`       // Mobile Country Code (MCC)
	MNC           string       `json:"mnc,omitempty"`       // Mobile Network Code (MNC)
	CellTower     CellTowerObj `json:"cellTower"`           // Cell Tower information
}

// BandObj represents cellular band information.
type BandObj struct {
	Name   string    `json:"name"`   // Band Name
	Signal SignalObj `json:"signal"` // Signal information
}

// SignalObj represents signal information.
type SignalObj struct {
	RSSI     int     `json:"rssi,omitempty"`     // Received Signal Strength Indicator (RSSI)
	SINR     float64 `json:"sinr,omitempty"`     // Signal to Interference plus Noise Ratio (SINR)
	SNR      float64 `json:"snr,omitempty"`      // Signal-to-noise ratio (SNR)
	ECIO     float64 `json:"ecio,omitempty"`     // Energy to Interference Ratio (Ec/Io)
	RSRP     float64 `json:"rsrp,omitempty"`     // Reference Signal Received Power (RSRP)
	RSRQ     float64 `json:"rsrq,omitempty"`     // Reference Signal Received Quality (RSRQ)
	Strength float64 `json:"strength,omitempty"` // Wi-Fi signal strength
}

// SIMGroupObj represents a group of SIM cards.
type SIMGroupObj struct {
	Order []int `json:"order"` // List of SIM IDs
}

// RemoteSIMObj represents remote SIM information.
type RemoteSIMObj struct {
	IMSI         string `json:"imsi"`               // IMSI
	SerialNumber string `json:"serialNumber"`       // Serial Number
	Slot         int    `json:"slot"`               // Number of slot
	AutoApp      bool   `json:"autoApp,omitempty"`  // Indicates if APN, Username, and Password fields are auto-detect
	APN          string `json:"apn,omitempty"`      // APN (only available in fw8.1.1 or later)
	Username     string `json:"username,omitempty"` // Username for APN (only available in fw8.1.1 or later)
	Password     string `json:"password,omitempty"` // Password for APN (only available in fw8.1.1 or later)
}

// CarrierObj represents carrier information.
type CarrierObj struct {
	Name    string `json:"name"`    // Carrier name
	Country string `json:"country"` // Carrier country (field does not appear if lite = yes)
}

// MEIDObj represents Mobile Equipment Identifier (MEID) information.
type MEIDObj struct {
	Hex string `json:"hex"` // MEID value in HEX
	Dec string `json:"dec"` // MEID value in DEC
}

// SIMObj represents SIM card information.
type SIMObj struct {
	Status   string `json:"status"`             // SIM card status {In Use, SIM Card Detected, No SIM Card Detected}
	Active   bool   `json:"active"`             // SIM card active status
	APN      string `json:"apn,omitempty"`      // APN
	Username string `json:"username,omitempty"` // Username for APN
	Password string `json:"password,omitempty"` // Password for APN
	IMSI     string `json:"imsi,omitempty"`     // International Mobile Subscriber Identity (IMSI)
	ICCID    string `json:"iccid,omitempty"`    // Integrated Circuit Card Identity (ICCID)
	MTN      string `json:"mtn,omitempty"`      // Mobile Telecommunications Network (MTN)
}

// RoamingObj represents roaming status information.
type RoamingObj struct {
	Code    int    `json:"code"`    // Roaming Status Code {0, 1, 2}
	Message string `json:"message"` // Readable Roaming Status Code and message relation
}

// CellTowerObj represents cell tower information.
type CellTowerObj struct {
	CellID int `json:"cellId"`
}

// StatusWanConnection returns the status of the WAN connections
func (c *Client) StatusWanConnection(ctx context.Context) ([]WanStatus, error) {
	msg, err := c.doRequest(ctx, "/api/status.wan.connection", http.MethodGet, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get wan status via http: %w", err)
	}

	var buf interface{}

	err = json.Unmarshal(msg, &buf)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	orderI, err := jmespath.Search("order", buf)
	if err != nil {
		return nil, fmt.Errorf("failed to get firmware version from json: %w", err)
	}
	statuses := []WanStatus{}

	for _, oi := range orderI.([]interface{}) {
		i, ok := oi.(float64)
		if !ok {
			return nil, fmt.Errorf("failed to get wan status: order is not a float64")
		}
		wanI, err := jmespath.Search(fmt.Sprintf("\"%d\"", int(i)), buf)
		if err != nil {
			return nil, fmt.Errorf("failed to get wan status from json: %w", err)
		}
		// Do marshal/unmarshal to workaround interface{} hussle
		buf, err := json.Marshal(wanI)
		if err != nil {
			return nil, fmt.Errorf("failed to get wan status from json: %w", err)
		}
		wan := WanStatus{}
		err = json.Unmarshal(buf, &wan)
		if err != nil {
			return nil, fmt.Errorf("failed to get wan status from json: %w", err)
		}
		statuses = append(statuses, wan)
	}

	return statuses, nil
}
