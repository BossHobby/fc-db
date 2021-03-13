package fc

import (
	"encoding/json"
	"strconv"

	"github.com/NotFastEnuf/fc-db/pkg/util"
)

type Pin struct {
	Port string
	Num  int
}

func ParsePin(s string) Pin {
	if s == "" {
		return Pin{}
	}

	i := 0
	if s[0] == 'P' {
		i = 1
	}
	return Pin{
		Port: string(s[i]),
		Num:  util.MustParseInt(string(s[(i + 1):])),
	}
}

func (p *Pin) String() string {
	if p.Port == "" {
		return ""
	}
	return "P" + p.Port + strconv.Itoa(p.Num)
}

func (p *Pin) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *Pin) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	*p = ParsePin(str)
	return nil
}

type PinDevice struct {
	Index int `json:"index"`
	Pin   Pin `json:"pin"`
}

type UARTPort struct {
	Index int `json:"index"`
	TXPin Pin `json:"tx_pin"`
	RXPin Pin `json:"rx_pin"`
}

type SPIPort struct {
	Index   int `json:"index"`
	SCLKPin Pin `json:"sclk_pin"`
	MISOPin Pin `json:"miso_pin"`
	MOSIPin Pin `json:"mosi_pin"`
}

type SPIDevice struct {
	Port  int `json:"port"`
	CSPin Pin `json:"cs_pin"`
}

type GyroDevice struct {
	Index  int `json:"index"`
	INTPin Pin `json:"int_pin"`
	SPIDevice
}

type Target struct {
	MCU          string `json:"mcu"`
	Board        string `json:"board"`
	Manufacturer string `json:"manufacturer"`

	UARTPorts []UARTPort `json:"uart_ports"`
	SPIPorts  []SPIPort  `json:"spi_ports"`

	BeeperPins   []PinDevice `json:"beeper_pins"`
	MotorPins    []PinDevice `json:"motor_pins"`
	LEDPins      []PinDevice `json:"led_pins"`
	InverterPins []PinDevice `json:"inverter_pins"`
	USBDetectPin *Pin        `json:"usb_detect_pin,omitempty"`
	BatteryPin   *Pin        `json:"battery_pin,omitempty"`

	Gyros     []GyroDevice `json:"gyros"`
	OSD       *SPIDevice   `json:"osd,omitempty"`
	DataFlash *SPIDevice   `json:"data_flash,omitempty"`
	SDCard    *SPIDevice   `json:"sd_card,omitempty"`
	RX        *SPIDevice   `json:"rx,omitempty"`
}

func NewTarget() *Target {
	return &Target{
		UARTPorts: make([]UARTPort, 0),
		SPIPorts:  make([]SPIPort, 0),

		BeeperPins:   make([]PinDevice, 0),
		MotorPins:    make([]PinDevice, 0),
		LEDPins:      make([]PinDevice, 0),
		InverterPins: make([]PinDevice, 0),

		Gyros: make([]GyroDevice, 0),
	}
}
