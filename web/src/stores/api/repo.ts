const REPO_URL = "https://raw.githubusercontent.com/BossHobby/fc-db/master/";

export type Pin = string;

export interface PinDevice {
  index: number;
  pin: Pin;
}

export interface ADCPin {
  pin: Pin;
  scale: number;
}

export interface UARTPort {
  index: number;
  tx_pin: Pin;
  rx_pin: Pin;
}

export interface SPIPort {
  index: number;
  sclk_pin: Pin;
  miso_pin: Pin;
  mosi_pin: Pin;
}

export interface SPIDevice {
  port: number;
  cs_pin: Pin;
}

export interface RadioDevice {
  exti_pin: Pin;
}

export interface GyroDevice {
  index: number;
  exti_pin: Pin;
}

export interface Target {
  mcu: string;
  board: string;
  manufacturer: string;
  uart_ports: UARTPort[];
  spi_ports: SPIPort[];
  beeper_pins: PinDevice[];
  motor_pins: PinDevice[];
  led_pins: PinDevice[];
  inverter_pins: PinDevice[];
  usb_detect_pin?: Pin;
  battery_pin?: ADCPin;
  current_pin?: ADCPin;
  gyros: GyroDevice[];
  osd?: SPIDevice;
  data_flash?: SPIDevice;
  sd_card?: SPIDevice;
  rx?: RadioDevice;
}

export interface TargetListEntry {
  mcu: string;
  board: string;
  manufacturer: string;
}

export class Repo {
  public static fetch(path: string): Promise<unknown> {
    return fetch(REPO_URL + path).then((r) => r.json());
  }

  public static fetchTargetList(): Promise<TargetListEntry[]> {
    return this.fetch("target/index.json").then((l) => l as TargetListEntry[]);
  }

  public static fetchTarget(mgfr: string, board: string): Promise<Target> {
    const key = (mgfr + "-" + board).toUpperCase();
    return this.fetch(`target/${key}/${key}.json`).then((t) => t as Target);
  }
}
