
#include "config.h"
#include "config_helper.h"

#define F4BY

#define F4
#define F405

//PORTS
#define SPI_PORTS \
  SPI1_PA5PA6PA7  \
  SPI2_PB13PB14PB15  \
  SPI3_PB3PB4PB5  \

#define USART_PORTS \
	USART1_PB7PB6 \
	USART2_PD6PD5 \
	USART3_PD9PD8 \
	USART4_PC11PC10 \
	USART6_PC7PC6

//LEDS
#define LED_NUMBER 3
#define LED1PIN GPIO_Pin_3
#define LED1PORT GPIOE
#define LED2PIN GPIO_Pin_2
#define LED2PORT GPIOE
#define LED3PIN GPIO_Pin_1
#define LED3PORT GPIOE

#define BUZZER_PIN GPIO_Pin_5
#define BUZZER_PIN_PORT GPIOE

//GYRO
#define MPU6XXX_SPI_PORT SPI_PORT1
#define MPU6XXX_NSS PIN_A4
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



//VOLTAGE DIVIDER
#define BATTERYPIN GPIO_Pin_3
#define BATTERYPORT GPIOC
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
#define MOTOR_PIN0 MOTOR_PIN_PA0
#define MOTOR_PIN1 MOTOR_PIN_PA1
#define MOTOR_PIN2 MOTOR_PIN_PA2
#define MOTOR_PIN3 MOTOR_PIN_PA3
#define MOTOR_PIN4 MOTOR_PIN_PE9
#define MOTOR_PIN5 MOTOR_PIN_PE11
#define MOTOR_PIN6 MOTOR_PIN_PE13
#define MOTOR_PIN7 MOTOR_PIN_PE14