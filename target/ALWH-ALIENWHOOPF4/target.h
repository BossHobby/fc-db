
#include "config.h"
#include "config_helper.h"

#define ALIENWHOOPF4

//PORTS
#define SPI_PORTS \
  SPI1_PA5PA6PA7  \
  SPI2_PB13PB14PB15  \
  SPI3_PB3PB4PB5  \

#define USART_PORTS \
	USART1_PA10PA9 \
	USART2_PA3PA2 \
	USART3_PC11PC10 \
	USART4_PA1PA0

//LEDS
#define LED_NUMBER 2
#define LED1PIN PIN_C12
#define LED1_INVERT
#define LED2PIN PIN_D2
#define LED2_INVERT

#define BUZZER_PIN PIN_A2

//GYRO
#define GYRO_SPI_PORT SPI_PORT1
#define GYRO_NSS PIN_A4
#define GYRO_INT PIN_C14
#define GYRO_ORIENTATION GYRO_ROTATE_90_CCW

//RADIO


#ifdef SERIAL_RX
#define RX_USART USART_PORT2
#endif

// OSD
#define ENABLE_OSD
#define MAX7456_SPI_PORT SPI_PORT2
#define MAX7456_NSS PIN_B12

#define USE_M25P16
#define M25P16_SPI_PORT SPI_PORT3
#define M25P16_NSS_PIN PIN_A15







// MOTOR PINS
//S3_OUT
#define MOTOR_PIN0 MOTOR_PIN_PC7
//S4_OUT
#define MOTOR_PIN1 MOTOR_PIN_PC6
//S1_OUT
#define MOTOR_PIN2 MOTOR_PIN_PC9
//S2_OUT
#define MOTOR_PIN3 MOTOR_PIN_PC8
