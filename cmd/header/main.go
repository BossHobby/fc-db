package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"

	"github.com/NotFastEnuf/fc-db/pkg/fc"
)

const headerTemplate = `
#include "config.h"
#include "config_helper.h"

#define {{ .Board }}

//PORTS
#define SPI_PORTS \
{{- range $i, $v := .SPIPorts }}
  SPI{{ $v.Index }}_{{ $v.SCLKPin }}{{ $v.MISOPin }}{{ $v.MOSIPin }}  \
{{- end }}

#define USART_PORTS \
{{- range $i, $v := .UARTPorts }}
	USART{{ $v.Index }}_{{ $v.RXPin }}{{ $v.TXPin }}{{ if not (lastIndex $i $.UARTPorts) }} \{{ end }}
{{- end }}

//LEDS
#define LED_NUMBER {{ len .LEDPins }}
{{- range $i, $v := .LEDPins }}
#define LED{{ $v.Index }}PIN {{ $v.Pin | pinEnum }}
#define LED{{ $v.Index }}_INVERT
{{- end }}

{{ if len .BeeperPins -}}
#define BUZZER_PIN {{ (index .BeeperPins 0).Pin | pinEnum }}
{{- end }}

//GYRO
#define GYRO_TYPE MPU6XXX
#define GYRO_SPI_PORT SPI_PORT{{ (index .Gyros 0).Port }}
#define GYRO_NSS {{ (index .Gyros 0).CSPin | pinEnum }}
#define GYRO_INT {{ (index .Gyros 0).EXTI | pinEnum }}
#define SENSOR_ROTATE_90_CCW
#define GYRO_ID_1 0x68
#define GYRO_ID_2 0x73
#define GYRO_ID_3 0x78
#define GYRO_ID_4 0x71

//RADIO
{{ if .RX -}}
#ifdef RX_FRSKY
#define USE_CC2500
#define CC2500_SPI_PORT SPI_PORT{{ .RX.Port }}
#define CC2500_NSS {{ .RX.CSPin | pinEnum }}
#define CC2500_GDO0_PIN {{ .RX.EXTI | pinEnum }}
// #define CC2500_TX_EN_PIN
// #define CC2500_LNA_EN_PIN
// #define CC2500_ANT_SEL_PIN
#endif
{{- end }}

#ifdef SERIAL_RX
#define RX_USART USART_PORT2
#endif

{{ if .OSD -}}
// OSD
#define ENABLE_OSD
#define MAX7456_SPI_PORT SPI_PORT{{ .OSD.Port }}
#define MAX7456_NSS {{ .OSD.CSPin | pinEnum }}
{{- end }}

{{ if .DataFlash -}}
#define USE_M25P16
#define M25P16_SPI_PORT SPI_PORT{{ .DataFlash.Port }}
#define M25P16_NSS_PIN {{ .DataFlash.CSPin | pinEnum }}
{{- end }}

{{ if .BatteryPin -}}
//VOLTAGE DIVIDER
#define VBAT_PIN {{ .BatteryPin.Pin | pinEnum }}
#define VBAT_DIVIDER_R1 10000
#define VBAT_DIVIDER_R2 1000
{{- end }}

{{ if .CurrentPin -}}
#define IBAT_PIN {{ .CurrentPin.Pin | pinEnum }}
#define IBAT_SCALE {{ .CurrentPin.Scale }}
{{- end }}

// MOTOR PINS
{{- range $i, $v := .MotorPins }}
//S{{ $v.Index }}_OUT
#define MOTOR_PIN{{ $i }} MOTOR_PIN_{{ $v.Pin }}
{{- end }}
`

func readTarget(filename string) (*fc.Target, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	target := new(fc.Target)
	if err := json.NewDecoder(f).Decode(target); err != nil {
		return nil, err
	}

	return target, nil
}

func writeHeader(filename string) error {
	target, err := readTarget(filename)
	if err != nil {
		return err
	}

	funcMap := template.FuncMap{
		"mcuShort": func(mcu string) string {
			return strings.TrimPrefix(mcu, "STM32")
		},
		"pinEnum": func(pin fc.Pin) string {
			return fmt.Sprintf("PIN_%s%d", pin.Port, pin.Num)
		},
		"lastIndex": func(index int, m interface{}) bool {
			return index == reflect.ValueOf(m).Len()-1
		},
	}

	// reorder motor pins to match QS
	target.MotorPins = []fc.PinDevice{
		target.MotorPins[2],
		target.MotorPins[3],
		target.MotorPins[0],
		target.MotorPins[1],
	}

	targetDir := strings.ToLower(target.Board)
	if err := os.Mkdir(targetDir, 0755); err != nil && !os.IsExist(err) {
		return err
	}

	{
		f, err := os.Create(filepath.Join(targetDir, "target.h"))
		if err != nil {
			return err
		}
		defer f.Close()

		t := template.Must(template.New("header").Funcs(funcMap).Parse(headerTemplate))
		if err := t.Execute(f, target); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		log.Println("missing <filename>")
		os.Exit(1)
	}

	filename := flag.Arg(0)

	s, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}

	if s.IsDir() {
		files, err := ioutil.ReadDir(filename)
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			path := filepath.Join(filename, f.Name())
			if f.IsDir() {
				continue
			}

			log.Printf("processing %s\n", path)
			if err := writeHeader(path); err != nil {
				log.Fatal(err)
			}
		}
	} else {
		if err := writeHeader(filename); err != nil {
			log.Fatal(err)
		}
	}

}
