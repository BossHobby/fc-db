package betaflight

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/NotFastEnuf/fc-db/pkg/fc"
	"github.com/NotFastEnuf/fc-db/pkg/util"
)

func ensureSpi(d **fc.SPIDevice) *fc.SPIDevice {
	if *d == nil {
		*d = &fc.SPIDevice{}
	}
	return *d
}

func ensureIndex(a interface{}, index int, def interface{}) int {
	valuePtr := reflect.ValueOf(a)
	value := valuePtr.Elem()
	for i := 0; i < value.Len(); i++ {
		item := value.Index(i).FieldByName("Index")
		if item.Int() == int64(index) {
			return i
		}
	}
	value.Set(reflect.Append(value, reflect.ValueOf(def)))
	return value.Len() - 1
}

func mapResource(t *fc.Target, resource string, index int, pin fc.Pin) {
	switch resource {
	case "BEEPER":
		t.BeeperPins = append(t.BeeperPins, fc.PinDevice{
			Index: index,
			Pin:   pin,
		})

	case "MOTOR":
		t.MotorPins = append(t.MotorPins, fc.PinDevice{
			Index: index,
			Pin:   pin,
		})

	case "LED":
		t.LEDPins = append(t.LEDPins, fc.PinDevice{
			Index: index,
			Pin:   pin,
		})

	case "INVERTER":
		t.InverterPins = append(t.InverterPins, fc.PinDevice{
			Index: index,
			Pin:   pin,
		})

	case "USB_DETECT":
		t.USBDetectPin = &pin

	case "ADC_BATT":
		t.BatteryPin = &fc.ADCPin{
			Pin:        pin,
			ADCChannel: getAdcChannel(pin, t.MCU),
		}

	case "SERIAL_TX":
		i := ensureIndex(&t.UARTPorts, index, fc.UARTPort{
			Index: index,
		})
		t.UARTPorts[i].TXPin = pin

	case "SERIAL_RX":
		i := ensureIndex(&t.UARTPorts, index, fc.UARTPort{
			Index: index,
		})
		t.UARTPorts[i].RXPin = pin

	case "SPI_SCK":
		i := ensureIndex(&t.SPIPorts, index, fc.SPIPort{
			Index: index,
		})
		t.SPIPorts[i].SCLKPin = pin

	case "SPI_MISO":
		i := ensureIndex(&t.SPIPorts, index, fc.SPIPort{
			Index: index,
		})
		t.SPIPorts[i].MISOPin = pin

	case "SPI_MOSI":
		i := ensureIndex(&t.SPIPorts, index, fc.SPIPort{
			Index: index,
		})
		t.SPIPorts[i].MOSIPin = pin

	case "GYRO_CS":
		i := ensureIndex(&t.Gyros, index, fc.GyroDevice{
			Index: index,
		})
		t.Gyros[i].CSPin = pin

	case "GYRO_EXTI":
		i := ensureIndex(&t.Gyros, index, fc.GyroDevice{
			Index: index,
		})
		t.Gyros[i].EXTI = pin

	case "FLASH_CS":
		ensureSpi(&t.DataFlash).CSPin = pin

	case "OSD_CS":
		ensureSpi(&t.OSD).CSPin = pin

	case "SDCARD_CS":
		ensureSpi(&t.SDCard).CSPin = pin

	case "RX_SPI_CS":
		if t.RX == nil {
			t.RX = &fc.RadioDevice{}
		}
		t.RX.CSPin = pin

	case "RX_SPI_EXTI":
		if t.RX == nil {
			t.RX = &fc.RadioDevice{}
		}
		t.RX.EXTI = pin
	}
}

func mapSet(t *fc.Target, key, value string) {
	switch key {
	case "gyro_1_spibus":
		i := ensureIndex(&t.Gyros, 1, fc.GyroDevice{
			Index: 1,
		})
		t.Gyros[i].Port = util.MustParseInt(value)

	case "flash_spi_bus":
		ensureSpi(&t.DataFlash).Port = util.MustParseInt(value)

	case "sdcard_spi_bus":
		ensureSpi(&t.SDCard).Port = util.MustParseInt(value)

	case "max7456_spi_bus":
		ensureSpi(&t.OSD).Port = util.MustParseInt(value)

	case "rx_spi_bus":
		if t.RX == nil {
			t.RX = &fc.RadioDevice{}
		}
		t.RX.Port = util.MustParseInt(value)

	}
}

func ParseConfig(path string) (*fc.Target, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	t := fc.NewTarget()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 || line[0] == '#' {
			if strings.HasPrefix(line, "# Betaflight") {
				parts := strings.Split(line, " ")
				t.MCU = parts[3]
			}
			continue
		}

		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "board_name":
			t.Board = parts[1]

		case "manufacturer_id":
			t.Manufacturer = parts[1]

		case "resource":
			resource := parts[1]
			if len(parts) == 3 {
				if parts[2] != "NONE" {
					mapResource(t, resource, 1, fc.ParsePin(parts[2]))
				}
			} else if len(parts) == 4 {
				if parts[3] != "NONE" {
					mapResource(t, resource, util.MustParseInt(parts[2]), fc.ParsePin(parts[3]))
				}
			} else {
				log.Fatal("invalid resource " + line)
			}

		case "set":
			mapSet(t, parts[1], parts[3])
		}

	}

	return t, scanner.Err()
}
