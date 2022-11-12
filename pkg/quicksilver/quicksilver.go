package quicksilver

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/template"

	"github.com/BossHobby/fc-db/pkg/fc"
)

const headerTemplate = `
#include "config.h"

//PORTS
#define SPI_PORTS \
{{- range $i, $v := .SPIPorts }}
  SPI{{ $v.Index }}_{{ $v.SCLKPin }}{{ $v.MISOPin }}{{ $v.MOSIPin }}{{ if not (lastIndex $i $.SPIPorts) }} \{{ end }}
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
#define GYRO_SPI_PORT SPI_PORT{{ (index .Gyros 0).Port }}
#define GYRO_NSS {{ (index .Gyros 0).CSPin | pinEnum }}
#define GYRO_INT {{ (index .Gyros 0).EXTI | pinEnum }}
#define GYRO_ORIENTATION GYRO_ROTATE_90_CCW

{{ if .CC2500 -}}
//RADIO
#define USE_CC2500
#define CC2500_SPI_PORT SPI_PORT{{ .CC2500.Port }}
#define CC2500_NSS_PIN {{ .CC2500.CSPin | pinEnum }}
#define CC2500_GDO0_PIN {{ .CC2500.EXTI | pinEnum }}
{{ if .CC2500.TXEnPin -}}#define CC2500_TX_EN_PIN {{ .CC2500.TXEnPin | pinEnum }}{{- end }}
{{ if .CC2500.LNAEnPin -}}#define CC2500_LNA_EN_PIN {{ .CC2500.LNAEnPin | pinEnum }}{{- end }}
{{ if .CC2500.ANTSelPin -}}#define CC2500_ANT_SEL_PIN {{ .CC2500.ANTSelPin | pinEnum }}{{- end }}
{{- end }}

{{ if .OSD -}}
// OSD
#define USE_MAX7456
#define MAX7456_SPI_PORT SPI_PORT{{ .OSD.Port }}
#define MAX7456_NSS {{ .OSD.CSPin | pinEnum }}
{{- end }}

{{ if .DataFlash -}}
#define USE_M25P16
#define M25P16_SPI_PORT SPI_PORT{{ .DataFlash.Port }}
#define M25P16_NSS_PIN {{ .DataFlash.CSPin | pinEnum }}
{{- end }}

{{ if .SDCard -}}
#define USE_SDCARD
#define SDCARD_SPI_PORT SPI_PORT{{ .SDCard.Port }}
#define SDCARD_NSS_PIN {{ .SDCard.CSPin | pinEnum }}
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

func WriteHeader(target fc.Target, filename string) error {
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

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	t := template.Must(template.New("header").Funcs(funcMap).Parse(headerTemplate))
	if err := t.Execute(f, target); err != nil {
		return err
	}
	return nil
}
