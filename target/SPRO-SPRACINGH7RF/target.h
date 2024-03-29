
#include "config.h"

//PORTS
#define SPI_PORTS \

#define USART_PORTS \
	USART2_PD6PD5 \
	USART3_PD9PD8 \
	USART4_PD0PD1 \
	USART8_PE0PE1 \
	USART5_PB13

//LEDS
#define LED_NUMBER 2
#define LED1PIN PIN_E5
#define LED1_INVERT
#define LED2PIN PIN_E6
#define LED2_INVERT

#define BUZZER_PIN PIN_E4

//GYRO
#define GYRO_SPI_PORT SPI_PORT6
#define GYRO_NSS PIN_A15
#define GYRO_INT PIN_D15
#define GYRO_ORIENTATION GYRO_ROTATE_90_CCW

//RADIO
#define USE_CC2500
#define CC2500_SPI_PORT SPI_PORT2
#define CC2500_NSS_PIN PIN_B12
#define CC2500_GDO0_PIN PIN_C6










//VOLTAGE DIVIDER
#define VBAT_PIN PIN_C3
#define VBAT_DIVIDER_R1 10000
#define VBAT_DIVIDER_R2 1000

#define IBAT_PIN PIN_C1
#define IBAT_SCALE 179

// MOTOR PINS
//S3_OUT
#define MOTOR_PIN0 MOTOR_PIN_PA6
//S4_OUT
#define MOTOR_PIN1 MOTOR_PIN_PA7
//S1_OUT
#define MOTOR_PIN2 MOTOR_PIN_PB0
//S2_OUT
#define MOTOR_PIN3 MOTOR_PIN_PB1
