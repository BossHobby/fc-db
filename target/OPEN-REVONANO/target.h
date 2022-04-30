
#include "config.h"
#include "config_helper.h"

#define REVONANO

//PORTS
#define SPI_PORTS \
  SPI2_PB13PB14PB15  \

#define USART_PORTS \
	USART1_PB7PB6 \
	USART2_PA3PA2

//LEDS
#define LED_NUMBER 2
#define LED1PIN PIN_C14
#define LED1_INVERT
#define LED2PIN PIN_C13
#define LED2_INVERT

#define BUZZER_PIN PIN_C13

//GYRO
#define GYRO_SPI_PORT SPI_PORT2
#define GYRO_NSS PIN_B12
#define GYRO_INT PIN_A15
#define GYRO_ORIENTATION GYRO_ROTATE_90_CCW

//RADIO


#ifdef SERIAL_RX
#define RX_USART USART_PORT2
#endif







//VOLTAGE DIVIDER
#define VBAT_PIN PIN_A6
#define VBAT_DIVIDER_R1 10000
#define VBAT_DIVIDER_R2 1000

#define IBAT_PIN PIN_A7
#define IBAT_SCALE 179

// MOTOR PINS
//S3_OUT
#define MOTOR_PIN0 MOTOR_PIN_PB8
//S4_OUT
#define MOTOR_PIN1 MOTOR_PIN_PB9
//S1_OUT
#define MOTOR_PIN2 MOTOR_PIN_PA10
//S2_OUT
#define MOTOR_PIN3 MOTOR_PIN_PB3
