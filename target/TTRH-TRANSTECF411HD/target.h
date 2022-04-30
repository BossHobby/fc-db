
#include "config.h"
#include "config_helper.h"

#define TRANSTECF411HD

//PORTS
#define SPI_PORTS \
  SPI1_PA5PA6PA7  \

#define USART_PORTS \
	USART1_PA10PA9 \
	USART2_PA3PA2

//LEDS
#define LED_NUMBER 1
#define LED1PIN PIN_A14
#define LED1_INVERT

#define BUZZER_PIN PIN_B6

//GYRO
#define GYRO_SPI_PORT SPI_PORT1
#define GYRO_NSS PIN_A4
#define GYRO_INT PIN_A1
#define GYRO_ORIENTATION GYRO_ROTATE_90_CCW

//RADIO


#ifdef SERIAL_RX
#define RX_USART USART_PORT2
#endif

// OSD
#define ENABLE_OSD
#define MAX7456_SPI_PORT SPI_PORT0
#define MAX7456_NSS PIN_B12





//VOLTAGE DIVIDER
#define VBAT_PIN PIN_A0
#define VBAT_DIVIDER_R1 10000
#define VBAT_DIVIDER_R2 1000

#define IBAT_PIN PIN_B4
#define IBAT_SCALE 179

// MOTOR PINS
//S3_OUT
#define MOTOR_PIN0 MOTOR_PIN_PB10
//S4_OUT
#define MOTOR_PIN1 MOTOR_PIN_PB9
//S1_OUT
#define MOTOR_PIN2 MOTOR_PIN_PB0
//S2_OUT
#define MOTOR_PIN3 MOTOR_PIN_PB1
