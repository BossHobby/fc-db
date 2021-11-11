package main

import (
	"encoding/json"
	"flag"
	"fmt"
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

#define F4
#define {{ .MCU | mcuShort }}

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
#define LED{{ $v.Index }}PIN Pin_{{ $v.Pin | pinEnum }}
{{- end }}

{{ if len .BeeperPins -}}
#define BUZZER_PIN {{ (index .BeeperPins 0).Pin | pinEnum }}
{{- end }}

//GYRO
#define MPU6XXX_SPI_PORT SPI_PORT{{ (index .Gyros 0).Port }}
#define MPU6XXX_NSS {{ (index .Gyros 0).CSPin | pinEnum }}
// #define MPU6XXX_INT 
#define USE_DUMMY_I2C
#define SENSOR_ROTATE_90_CCW
#define GYRO_ID_1 0x68
#define GYRO_ID_2 0x73
#define GYRO_ID_3 0x78
#define GYRO_ID_4 0x71

//RADIO
#define RX_USART USART_PORT2
#define SOFTSPI_NONE

{{ if .OSD -}}
// OSD
#define ENABLE_OSD
#define MAX7456_SPI_PORT SPI_PORT{{ .OSD.Port }}
#define MAX7456_NSS {{ .OSD.CSPin | pinEnum }}
{{- end }}

{{ if .BatteryPin -}}
//VOLTAGE DIVIDER
#define BATTERYPIN GPIO_Pin_{{ .BatteryPin | pinEnum }}
#define BATTERY_ADC_CHANNEL ADC_Channel_8

#ifndef VOLTAGE_DIVIDER_R1
#define VOLTAGE_DIVIDER_R1 10000
#endif

#ifndef VOLTAGE_DIVIDER_R2
#define VOLTAGE_DIVIDER_R2 1000
#endif

#ifndef ADC_REF_VOLTAGE
#define ADC_REF_VOLTAGE 3.3
#endif
{{- end }}

// MOTOR PINS
{{- range $i, $v := .MotorPins }}
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

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		log.Println("missing <filename>")
		os.Exit(1)
	}

	filename := flag.Arg(0)
	log.Printf("processing %s", filename)

	target, err := readTarget(filename)
	if err != nil {
		log.Panic(err)
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

	targetDir := strings.ToLower(target.Board)
	if err := os.Mkdir(targetDir, 0755); err != nil && !os.IsExist(err) {
		log.Panic(err)
	}

	{
		f, err := os.Create(filepath.Join(targetDir, "target.h"))
		if err != nil {
			log.Panic(err)
		}
		defer f.Close()

		t := template.Must(template.New("header").Funcs(funcMap).Parse(headerTemplate))
		if err := t.Execute(f, target); err != nil {
			log.Panic(err)
		}
	}
}
