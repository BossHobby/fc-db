
#include "config.h"
#include "config_helper.h"

#define COLIBRI

#define F4
#define F405

//PORTS
#define SPI_PORTS \
  SPI1_PA5PA6PA7  \
  SPI2_PB13PC2PC3  \

#define USART_PORTS \
	USART1_PB7PB6 \
	USART2_PA3PA2 \
	USART3_PB11PB10

//LEDS
#define LED_NUMBER 2
#define LED1PIN Pin_PIN_C14
#define LED2PIN Pin_PIN_C13

#define BUZZER_PIN PIN_C5

//GYRO
#define MPU6XXX_SPI_PORT SPI_PORT1
#define MPU6XXX_NSS PIN_C4
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





// MOTOR PINS
#define MOTOR_PIN0 MOTOR_PIN_PB0
#define MOTOR_PIN1 MOTOR_PIN_PB4
#define MOTOR_PIN2 MOTOR_PIN_PB1
#define MOTOR_PIN3 MOTOR_PIN_PB15
#define MOTOR_PIN4 MOTOR_PIN_PB5
#define MOTOR_PIN5 MOTOR_PIN_PB14
#define MOTOR_PIN6 MOTOR_PIN_PB8
#define MOTOR_PIN7 MOTOR_PIN_PB9
