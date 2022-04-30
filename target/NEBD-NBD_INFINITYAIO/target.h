
#include "config.h"
#include "config_helper.h"

#define NBD_INFINITYAIO

//PORTS
#define SPI_PORTS \
  SPI1_PA5PA6PA7  \
  SPI2_PB13PB14PB15  \

#define USART_PORTS \
	USART2_PA3PA2 \
	USART3_PB11PB10 \
	USART1_PA10 \
	USART5_PD2

//LEDS
#define LED_NUMBER 1
#define LED1PIN PIN_A15
#define LED1_INVERT

#define BUZZER_PIN PIN_C15

//GYRO
#define GYRO_SPI_PORT SPI_PORT2
#define GYRO_NSS PIN_B12
#define GYRO_INT PIN_C4
#define GYRO_ORIENTATION GYRO_ROTATE_90_CCW

//RADIO


#ifdef SERIAL_RX
#define RX_USART USART_PORT2
#endif



#define USE_M25P16
#define M25P16_SPI_PORT SPI_PORT1
#define M25P16_NSS_PIN PIN_A4



//VOLTAGE DIVIDER
#define VBAT_PIN PIN_C0
#define VBAT_DIVIDER_R1 10000
#define VBAT_DIVIDER_R2 1000

#define IBAT_PIN PIN_C1
#define IBAT_SCALE 230

// MOTOR PINS
//S3_OUT
#define MOTOR_PIN0 MOTOR_PIN_PB3
//S4_OUT
#define MOTOR_PIN1 MOTOR_PIN_PB0
//S1_OUT
#define MOTOR_PIN2 MOTOR_PIN_PB4
//S2_OUT
#define MOTOR_PIN3 MOTOR_PIN_PB1
