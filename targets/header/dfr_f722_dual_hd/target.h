
#include "config.h"
#include "config_helper.h"

#define DFR_F722_DUAL_HD

#define F4
#define F7X2

//PORTS
#define SPI_PORTS \
  SPI1_PA5PA6PA7  \
  SPI2_PB13PB14PB15  \
  SPI3_PB3PB4PB5  \

#define USART_PORTS \
	USART1_PA10PA9 \
	USART2_PA3PA2 \
	USART3_PB11PB10 \
	USART4_PC11PC10 \
	USART5_PD2PC12

//LEDS
#define LED_NUMBER 1
#define LED1PIN Pin_PIN_C4

#define BUZZER_PIN PIN_C15

//GYRO
#define MPU6XXX_SPI_PORT SPI_PORT1
#define MPU6XXX_NSS PIN_A15
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

// OSD
#define ENABLE_OSD
#define MAX7456_SPI_PORT SPI_PORT2
#define MAX7456_NSS PIN_B12

//VOLTAGE DIVIDER
#define BATTERYPIN GPIO_Pin_PIN_C1
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

// MOTOR PINS
#define MOTOR_PIN0 MOTOR_PIN_PC8
#define MOTOR_PIN1 MOTOR_PIN_PC6
#define MOTOR_PIN2 MOTOR_PIN_PC9
#define MOTOR_PIN3 MOTOR_PIN_PC7
#define MOTOR_PIN4 MOTOR_PIN_PB6
#define MOTOR_PIN5 MOTOR_PIN_PB7
#define MOTOR_PIN6 MOTOR_PIN_PB1
#define MOTOR_PIN7 MOTOR_PIN_PB0
