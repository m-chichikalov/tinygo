// Automatically generated file. DO NOT EDIT.
// Generated by gen-device.py from ATmega162.atdf, see http://packs.download.atmel.com/

// +build avr,atmega162

// Device information for the ATmega162.
//
package avr

// Magic type name for the compiler.
type __reg uint8

// Export this magic type name.
type RegValue = __reg

// Some information about this device.
const (
	DEVICE = "ATmega162"
	ARCH   = "AVR8"
	FAMILY = "megaAVR"
)

// Interrupts
const (
	IRQ_RESET        = 0  // External Pin,Power-on Reset,Brown-out Reset,Watchdog Reset,and JTAG AVR Reset. See Datasheet.
	IRQ_INT0         = 1  // External Interrupt Request 0
	IRQ_INT1         = 2  // External Interrupt Request 1
	IRQ_INT2         = 3  // External Interrupt Request 2
	IRQ_PCINT0       = 4  // Pin Change Interrupt Request 0
	IRQ_PCINT1       = 5  // Pin Change Interrupt Request 1
	IRQ_TIMER3_CAPT  = 6  // Timer/Counter3 Capture Event
	IRQ_TIMER3_COMPA = 7  // Timer/Counter3 Compare Match A
	IRQ_TIMER3_COMPB = 8  // Timer/Counter3 Compare Match B
	IRQ_TIMER3_OVF   = 9  // Timer/Counter3 Overflow
	IRQ_TIMER2_COMP  = 10 // Timer/Counter2 Compare Match
	IRQ_TIMER2_OVF   = 11 // Timer/Counter2 Overflow
	IRQ_TIMER1_CAPT  = 12 // Timer/Counter1 Capture Event
	IRQ_TIMER1_COMPA = 13 // Timer/Counter1 Compare Match A
	IRQ_TIMER1_COMPB = 14 // Timer/Counter Compare Match B
	IRQ_TIMER1_OVF   = 15 // Timer/Counter1 Overflow
	IRQ_TIMER0_COMP  = 16 // Timer/Counter0 Compare Match
	IRQ_TIMER0_OVF   = 17 // Timer/Counter0 Overflow
	IRQ_SPI_STC      = 18 // SPI Serial Transfer Complete
	IRQ_USART0_RXC   = 19 // USART0, Rx Complete
	IRQ_USART1_RXC   = 20 // USART1, Rx Complete
	IRQ_USART0_UDRE  = 21 // USART0 Data register Empty
	IRQ_USART1_UDRE  = 22 // USART1, Data register Empty
	IRQ_USART0_TXC   = 23 // USART0, Tx Complete
	IRQ_USART1_TXC   = 24 // USART1, Tx Complete
	IRQ_EE_RDY       = 25 // EEPROM Ready
	IRQ_ANA_COMP     = 26 // Analog Comparator
	IRQ_SPM_RDY      = 27 // Store Program Memory Read
	IRQ_max          = 27 // Highest interrupt number on this device.
)

// Peripherals
var (
	// Fuses
	FUSE = struct {
		EXTENDED __reg
		HIGH     __reg
		LOW      __reg
	}{
		EXTENDED: 0x2,
		HIGH:     0x1,
		LOW:      0x0,
	}

	// Lockbits
	LOCKBIT = struct {
		LOCKBIT __reg
	}{
		LOCKBIT: 0x0,
	}

	// Timer/Counter, 16-bit
	TC16 = struct {
		TCCR1A __reg
		TCCR1B __reg
		TCNT1L __reg
		TCNT1H __reg
		OCR1AL __reg
		OCR1AH __reg
		OCR1BL __reg
		OCR1BH __reg
		ICR1L  __reg
		ICR1H  __reg
		ETIMSK __reg
		ETIFR  __reg
		TCCR3A __reg
		TCCR3B __reg
		TCNT3L __reg
		TCNT3H __reg
		OCR3AL __reg
		OCR3AH __reg
		OCR3BL __reg
		OCR3BH __reg
		ICR3L  __reg
		ICR3H  __reg
	}{
		TCCR1A: 0x4f, // Timer/Counter1 Control Register A
		TCCR1B: 0x4e, // Timer/Counter1 Control Register B
		TCNT1L: 0x4c, // Timer/Counter1  Bytes
		TCNT1H: 0x4c, // Timer/Counter1  Bytes
		OCR1AL: 0x4a, // Timer/Counter1 Output Compare Register A  Bytes
		OCR1AH: 0x4a, // Timer/Counter1 Output Compare Register A  Bytes
		OCR1BL: 0x48, // Timer/Counter1 Output Compare Register B  Bytes
		OCR1BH: 0x48, // Timer/Counter1 Output Compare Register B  Bytes
		ICR1L:  0x44, // Timer/Counter1 Input Capture Register  Bytes
		ICR1H:  0x44, // Timer/Counter1 Input Capture Register  Bytes
		ETIMSK: 0x7d, // Extended Timer/Counter Interrupt Mask Register
		ETIFR:  0x7c, // Extended Timer/Counter Interrupt Flag register
		TCCR3A: 0x8b, // Timer/Counter3 Control Register A
		TCCR3B: 0x8a, // Timer/Counter3 Control Register B
		TCNT3L: 0x88, // Timer/Counter3  Bytes
		TCNT3H: 0x88, // Timer/Counter3  Bytes
		OCR3AL: 0x86, // Timer/Counter3 Output Compare Register A  Bytes
		OCR3AH: 0x86, // Timer/Counter3 Output Compare Register A  Bytes
		OCR3BL: 0x84, // Timer/Counte3 Output Compare Register B  Bytes
		OCR3BH: 0x84, // Timer/Counte3 Output Compare Register B  Bytes
		ICR3L:  0x80, // Timer/Counter3 Input Capture Register  Bytes
		ICR3H:  0x80, // Timer/Counter3 Input Capture Register  Bytes
	}

	// Timer/Counter, 8-bit Async
	TC8_ASYNC = struct {
		TCCR2 __reg
		TCNT2 __reg
		OCR2  __reg
		ASSR  __reg
	}{
		TCCR2: 0x47, // Timer/Counter Control Register
		TCNT2: 0x43, // Timer/Counter Register
		OCR2:  0x42, // Output Compare Register
		ASSR:  0x46, // Asynchronous Status Register
	}

	// Analog Comparator
	AC = struct {
		ACSR __reg
	}{
		ACSR: 0x28, // Analog Comparator Control And Status Register
	}

	// USART
	USART = struct {
		UDR0   __reg
		UCSR0A __reg
		UCSR0B __reg
		UCSR0C __reg
		UBRR0H __reg
		UBRR0L __reg
		UDR1   __reg
		UCSR1A __reg
		UCSR1B __reg
		UCSR1C __reg
		UBRR1H __reg
		UBRR1L __reg
	}{
		UDR0:   0x2c, // USART I/O Data Register
		UCSR0A: 0x2b, // USART Control and Status Register A
		UCSR0B: 0x2a, // USART Control and Status Register B
		UCSR0C: 0x40, // USART Control and Status Register C
		UBRR0H: 0x40, // USART Baud Rate Register High Byte
		UBRR0L: 0x29, // USART Baud Rate Register Low Byte
		UDR1:   0x23, // USART I/O Data Register
		UCSR1A: 0x22, // USART Control and Status Register A
		UCSR1B: 0x21, // USART Control and Status Register B
		UCSR1C: 0x5c, // USART Control and Status Register C
		UBRR1H: 0x5c, // USART Baud Rate Register Highg Byte
		UBRR1L: 0x20, // USART Baud Rate Register Low Byte
	}

	// Serial Peripheral Interface
	SPI = struct {
		SPCR __reg
		SPSR __reg
		SPDR __reg
	}{
		SPCR: 0x2d, // SPI Control Register
		SPSR: 0x2e, // SPI Status Register
		SPDR: 0x2f, // SPI Data Register
	}

	// CPU Registers
	CPU = struct {
		SREG   __reg
		SPL    __reg
		SPH    __reg
		OSCCAL __reg
		CLKPR  __reg
		SFIOR  __reg
	}{
		SREG:   0x5f, // Status Register
		SPL:    0x5d, // Stack Pointer
		SPH:    0x5d, // Stack Pointer
		OSCCAL: 0x24, // Oscillator Calibration Value
		CLKPR:  0x61, // Clock prescale register
		SFIOR:  0x50, // Special Function IO Register
	}

	// JTAG Interface
	JTAG = struct {
		OCDR __reg
	}{
		OCDR: 0x24, // On-Chip Debug Related Register in I/O Memory
	}

	// Bootloader
	BOOT_LOAD = struct {
		SPMCR __reg
	}{
		SPMCR: 0x57, // Store Program Memory Control Register
	}

	// EEPROM
	EEPROM = struct {
		EEARL __reg
		EEARH __reg
		EEDR  __reg
		EECR  __reg
	}{
		EEARL: 0x3e, // EEPROM Address Register Bytes
		EEARH: 0x3e, // EEPROM Address Register Bytes
		EEDR:  0x3d, // EEPROM Data Register
		EECR:  0x3c, // EEPROM Control Register
	}

	// I/O Port
	PORT = struct {
		PORTA __reg
		DDRA  __reg
		PINA  __reg
		PORTB __reg
		DDRB  __reg
		PINB  __reg
		PORTC __reg
		DDRC  __reg
		PINC  __reg
		PORTD __reg
		DDRD  __reg
		PIND  __reg
		PORTE __reg
		DDRE  __reg
		PINE  __reg
	}{
		PORTA: 0x3b, // Port A Data Register
		DDRA:  0x3a, // Port A Data Direction Register
		PINA:  0x39, // Port A Input Pins
		PORTB: 0x38, // Port B Data Register
		DDRB:  0x37, // Port B Data Direction Register
		PINB:  0x36, // Port B Input Pins
		PORTC: 0x35, // Port C Data Register
		DDRC:  0x34, // Port C Data Direction Register
		PINC:  0x33, // Port C Input Pins
		PORTD: 0x32, // Port D Data Register
		DDRD:  0x31, // Port D Data Direction Register
		PIND:  0x30, // Port D Input Pins
		PORTE: 0x27, // Data Register, Port E
		DDRE:  0x26, // Data Direction Register, Port E
		PINE:  0x25, // Input Pins, Port E
	}

	// Timer/Counter, 8-bit
	TC8 = struct {
		TCCR0 __reg
		TCNT0 __reg
		OCR0  __reg
	}{
		TCCR0: 0x53, // Timer/Counter 0 Control Register
		TCNT0: 0x52, // Timer/Counter 0 Register
		OCR0:  0x51, // Timer/Counter 0 Output Compare Register
	}

	// Watchdog Timer
	WDT = struct {
		WDTCR __reg
	}{
		WDTCR: 0x41, // Watchdog Timer Control Register
	}

	// External Interrupts
	EXINT = struct {
		GICR   __reg
		GIFR   __reg
		PCMSK1 __reg
		PCMSK0 __reg
	}{
		GICR:   0x5b, // General Interrupt Control Register
		GIFR:   0x5a, // General Interrupt Flag Register
		PCMSK1: 0x6c, // Pin Change Mask Register 1
		PCMSK0: 0x6b, // Pin Change Enable Mask
	}
)

// Bitfields for FUSE: Fuses
const (
	// EXTENDED
	EXTENDED_M161C    = 0x10 // ATmega161 compability mode
	EXTENDED_BODLEVEL = 0xe  // Brownout detector trigger level

	// HIGH
	HIGH_OCDEN   = 0x80 // On-Chip Debug Enabled
	HIGH_JTAGEN  = 0x40 // JTAG Interface Enabled
	HIGH_SPIEN   = 0x20 // Serial program downloading (SPI) enabled
	HIGH_WDTON   = 0x10 // Watchdog timer always on
	HIGH_EESAVE  = 0x8  // Preserve EEPROM through the Chip Erase cycle
	HIGH_BOOTSZ  = 0x6  // Select Boot Size
	HIGH_BOOTRST = 0x1  // Boot Reset vector Enabled

	// LOW
	LOW_CKDIV8    = 0x80 // Divide clock by 8 internally
	LOW_CKOUT     = 0x40 // Clock output on PORTB0
	LOW_SUT_CKSEL = 0x3f // Select Clock Source
)

// Bitfields for LOCKBIT: Lockbits
const (
	// LOCKBIT
	LOCKBIT_LB   = 0x3  // Memory Lock
	LOCKBIT_BLB0 = 0xc  // Boot Loader Protection Mode
	LOCKBIT_BLB1 = 0x30 // Boot Loader Protection Mode
)

// Bitfields for TC16: Timer/Counter, 16-bit
const (
	// TCCR1A: Timer/Counter1 Control Register A
	TCCR1A_COM1A = 0xc0 // Compare Output Mode 1A, bits
	TCCR1A_COM1B = 0x30 // Compare Output Mode 1B, bits
	TCCR1A_FOC1A = 0x8  // Force Output Compare for Channel A
	TCCR1A_FOC1B = 0x4  // Force Output Compare for Channel B
	TCCR1A_WGM1  = 0x3  // Pulse Width Modulator Select Bits

	// TCCR1B: Timer/Counter1 Control Register B
	TCCR1B_ICNC1 = 0x80 // Input Capture 1 Noise Canceler
	TCCR1B_ICES1 = 0x40 // Input Capture 1 Edge Select
	TCCR1B_WGM1  = 0x18 // Pulse Width Modulator Select Bits
	TCCR1B_CS1   = 0x7  // Clock Select1 bits

	// ETIMSK: Extended Timer/Counter Interrupt Mask Register
	ETIMSK_TICIE3 = 0x20 // Timer/Counter3 Input Capture Interrupt Enable
	ETIMSK_OCIE3A = 0x10 // Timer/Counter3 Output CompareA Match Interrupt Enable
	ETIMSK_OCIE3B = 0x8  // Timer/Counter3 Output CompareB Match Interrupt Enable
	ETIMSK_TOIE3  = 0x4  // Timer/Counter3 Overflow Interrupt Enable

	// ETIFR: Extended Timer/Counter Interrupt Flag register
	ETIFR_ICF3  = 0x20 // Input Capture Flag 3
	ETIFR_OCF3A = 0x10 // Output Compare Flag 3A
	ETIFR_OCF3B = 0x8  // Output Compare Flag 3B
	ETIFR_TOV3  = 0x4  // Timer/Counter3 Overflow Flag

	// TCCR3A: Timer/Counter3 Control Register A
	TCCR3A_COM3A = 0xc0 // Compare Output Mode 3A, bits
	TCCR3A_COM3B = 0x30 // Compare Output Mode 3B, bits
	TCCR3A_FOC3A = 0x8  // Force Output Compare for Channel A
	TCCR3A_FOC3B = 0x4  // Force Output Compare for Channel B
	TCCR3A_WGM3  = 0x3  // Pulse Width Modulator Select Bits

	// TCCR3B: Timer/Counter3 Control Register B
	TCCR3B_ICNC3 = 0x80 // Input Capture 3 Noise Canceler
	TCCR3B_ICES3 = 0x40 // Input Capture 3 Edge Select
	TCCR3B_WGM3  = 0x18 // Pulse Width Modulator Select Bits
	TCCR3B_CS3   = 0x7  // Clock Select3 bits
)

// Bitfields for TC8_ASYNC: Timer/Counter, 8-bit Async
const (
	// TCCR2: Timer/Counter Control Register
	TCCR2_FOC2  = 0x80 // Forde Output Compare
	TCCR2_WGM20 = 0x40 // Pulse Width Modulator Select Bit 0
	TCCR2_COM2  = 0x30 // Compare Match Output Mode
	TCCR2_WGM21 = 0x8  // Pulse Width Modulator Select Bit 1
	TCCR2_CS2   = 0x7  // Clock Select

	// ASSR: Asynchronous Status Register
	ASSR_AS2    = 0x8 // Asynchronous Timer 2
	ASSR_TCN2UB = 0x4 // Timer/Counter2 Update Busy
	ASSR_OCR2UB = 0x2 // Output Compare Register2 Update Busy
	ASSR_TCR2UB = 0x1 // Timer/Counter Control Register2 Update Busy
)

// Bitfields for AC: Analog Comparator
const (
	// ACSR: Analog Comparator Control And Status Register
	ACSR_ACD  = 0x80 // Analog Comparator Disable
	ACSR_ACBG = 0x40 // Analog Comparator Bandgap Select
	ACSR_ACO  = 0x20 // Analog Compare Output
	ACSR_ACI  = 0x10 // Analog Comparator Interrupt Flag
	ACSR_ACIE = 0x8  // Analog Comparator Interrupt Enable
	ACSR_ACIC = 0x4  // Analog Comparator Input Capture Enable
	ACSR_ACIS = 0x3  // Analog Comparator Interrupt Mode Select bits
)

// Bitfields for USART: USART
const (
	// UCSR0A: USART Control and Status Register A
	UCSR0A_RXC0  = 0x80 // USART Receive Complete
	UCSR0A_TXC0  = 0x40 // USART Transmitt Complete
	UCSR0A_UDRE0 = 0x20 // USART Data Register Empty
	UCSR0A_FE0   = 0x10 // Framing Error
	UCSR0A_DOR0  = 0x8  // Data overRun
	UCSR0A_UPE0  = 0x4  // Parity Error
	UCSR0A_U2X0  = 0x2  // Double the USART transmission speed
	UCSR0A_MPCM0 = 0x1  // Multi-processor Communication Mode

	// UCSR0B: USART Control and Status Register B
	UCSR0B_RXCIE0 = 0x80 // RX Complete Interrupt Enable
	UCSR0B_TXCIE0 = 0x40 // TX Complete Interrupt Enable
	UCSR0B_UDRIE0 = 0x20 // USART Data register Empty Interrupt Enable
	UCSR0B_RXEN0  = 0x10 // Receiver Enable
	UCSR0B_TXEN0  = 0x8  // Transmitter Enable
	UCSR0B_UCSZ02 = 0x4  // Character Size
	UCSR0B_RXB80  = 0x2  // Receive Data Bit 8
	UCSR0B_TXB80  = 0x1  // Transmit Data Bit 8

	// UCSR0C: USART Control and Status Register C
	UCSR0C_URSEL0 = 0x80 // Register Select
	UCSR0C_UMSEL0 = 0x40 // USART Mode Select
	UCSR0C_UPM0   = 0x30 // Parity Mode Bits
	UCSR0C_USBS0  = 0x8  // Stop Bit Select
	UCSR0C_UCSZ0  = 0x6  // Character Size
	UCSR0C_UCPOL0 = 0x1  // Clock Polarity

	// UBRR0H: USART Baud Rate Register High Byte
	UBRR0H_URSEL0 = 0x80 // Register Select
	UBRR0H_UBRR0  = 0xf  // USART Baud Rate Register High bits

	// UBRR0L: USART Baud Rate Register Low Byte
	UBRR0L_UBRR0 = 0xff // USART Baud Rate Register Low bits

	// UCSR1A: USART Control and Status Register A
	UCSR1A_RXC1  = 0x80 // USART Receive Complete
	UCSR1A_TXC1  = 0x40 // USART Transmitt Complete
	UCSR1A_UDRE1 = 0x20 // USART Data Register Empty
	UCSR1A_FE1   = 0x10 // Framing Error
	UCSR1A_DOR1  = 0x8  // Data overRun
	UCSR1A_UPE1  = 0x4  // Parity Error
	UCSR1A_U2X1  = 0x2  // Double the USART transmission speed
	UCSR1A_MPCM1 = 0x1  // Multi-processor Communication Mode

	// UCSR1B: USART Control and Status Register B
	UCSR1B_RXCIE1 = 0x80 // RX Complete Interrupt Enable
	UCSR1B_TXCIE1 = 0x40 // TX Complete Interrupt Enable
	UCSR1B_UDRIE1 = 0x20 // USART Data register Empty Interrupt Enable
	UCSR1B_RXEN1  = 0x10 // Receiver Enable
	UCSR1B_TXEN1  = 0x8  // Transmitter Enable
	UCSR1B_UCSZ12 = 0x4  // Character Size
	UCSR1B_RXB81  = 0x2  // Receive Data Bit 8
	UCSR1B_TXB81  = 0x1  // Transmit Data Bit 8

	// UCSR1C: USART Control and Status Register C
	UCSR1C_URSEL1 = 0x80 // Register Select
	UCSR1C_UMSEL1 = 0x40 // USART Mode Select
	UCSR1C_UPM1   = 0x30 // Parity Mode Bits
	UCSR1C_USBS1  = 0x8  // Stop Bit Select
	UCSR1C_UCSZ1  = 0x6  // Character Size
	UCSR1C_UCPOL1 = 0x1  // Clock Polarity

	// UBRR1H: USART Baud Rate Register Highg Byte
	UBRR1H_URSEL0 = 0x80 // Register Select
	UBRR1H_UBRR1  = 0xf  // USART Baud Rate Register High bits

	// UBRR1L: USART Baud Rate Register Low Byte
	UBRR1L_UBRR1 = 0xff // USART Baud Rate Register Low bits
)

// Bitfields for SPI: Serial Peripheral Interface
const (
	// SPCR: SPI Control Register
	SPCR_SPIE = 0x80 // SPI Interrupt Enable
	SPCR_SPE  = 0x40 // SPI Enable
	SPCR_DORD = 0x20 // Data Order
	SPCR_MSTR = 0x10 // Master/Slave Select
	SPCR_CPOL = 0x8  // Clock polarity
	SPCR_CPHA = 0x4  // Clock Phase
	SPCR_SPR  = 0x3  // SPI Clock Rate Selects

	// SPSR: SPI Status Register
	SPSR_SPIF  = 0x80 // SPI Interrupt Flag
	SPSR_WCOL  = 0x40 // Write Collision Flag
	SPSR_SPI2X = 0x1  // Double SPI Speed Bit
)

// Bitfields for CPU: CPU Registers
const (
	// SREG: Status Register
	SREG_I = 0x80 // Global Interrupt Enable
	SREG_T = 0x40 // Bit Copy Storage
	SREG_H = 0x20 // Half Carry Flag
	SREG_S = 0x10 // Sign Bit
	SREG_V = 0x8  // Two's Complement Overflow Flag
	SREG_N = 0x4  // Negative Flag
	SREG_Z = 0x2  // Zero Flag
	SREG_C = 0x1  // Carry Flag

	// OSCCAL: Oscillator Calibration Value
	OSCCAL_OSCCAL = 0xff // Oscillator Calibration

	// CLKPR: Clock prescale register
	CLKPR_CLKPCE = 0x80 // Clock Prescaler Change Enable
	CLKPR_CLKPS  = 0xf  // Clock Prescaler Select Bits

	// SFIOR: Special Function IO Register
	SFIOR_TSM    = 0x80 // Timer/Counter Synchronization Mode
	SFIOR_XMBK   = 0x40 // External Memory Bus Keeper Enable
	SFIOR_XMM    = 0x38 // External Memory High Mask Bits
	SFIOR_PUD    = 0x4  // Pull-up Disable
	SFIOR_PSR2   = 0x2  // Prescaler Reset Timer/Counter2
	SFIOR_PSR310 = 0x1  // Prescaler Reset Timer/Counter3, Timer/Counter1 and Timer/Counter0
)

// Bitfields for JTAG: JTAG Interface
const (
	// OCDR: On-Chip Debug Related Register in I/O Memory
	OCDR_OCDR = 0xff // On-Chip Debug Register Bits
)

// Bitfields for BOOT_LOAD: Bootloader
const (
	// SPMCR: Store Program Memory Control Register
	SPMCR_SPMIE  = 0x80 // SPM Interrupt Enable
	SPMCR_RWWSB  = 0x40 // Read While Write Section Busy
	SPMCR_RWWSRE = 0x10 // Read While Write secion read enable
	SPMCR_BLBSET = 0x8  // Boot Lock Bit Set
	SPMCR_PGWRT  = 0x4  // Page Write
	SPMCR_PGERS  = 0x2  // Page Erase
	SPMCR_SPMEN  = 0x1  // Store Program Memory Enable
)

// Bitfields for EEPROM: EEPROM
const (
	// EEARL: EEPROM Address Register Bytes

	// EEARH: EEPROM Address Register Bytes
	EEAR_EEAR = 0x1ff // EEPROM Address Register bits

	// EEDR: EEPROM Data Register
	EEDR_EEDR = 0xff // EEPROM Data Register bits

	// EECR: EEPROM Control Register
	EECR_EERIE = 0x8 // EEPROM Ready Interrupt Enable
	EECR_EEMWE = 0x4 // EEPROM Master Write Enable
	EECR_EEWE  = 0x2 // EEPROM Write Enable
	EECR_EERE  = 0x1 // EEPROM Read Enable
)

// Bitfields for TC8: Timer/Counter, 8-bit
const (
	// TCCR0: Timer/Counter 0 Control Register
	TCCR0_FOC0  = 0x80 // Force Output Compare
	TCCR0_WGM00 = 0x40 // Waveform Generation Mode 0
	TCCR0_COM0  = 0x30 // Compare Match Output Modes
	TCCR0_WGM01 = 0x8  // Waveform Generation Mode 1
	TCCR0_CS0   = 0x7  // Clock Selects
)

// Bitfields for WDT: Watchdog Timer
const (
	// WDTCR: Watchdog Timer Control Register
	WDTCR_WDCE = 0x10 // Watchdog Change Enable
	WDTCR_WDE  = 0x8  // Watch Dog Enable
	WDTCR_WDP  = 0x7  // Watch Dog Timer Prescaler bits
)

// Bitfields for EXINT: External Interrupts
const (
	// GICR: General Interrupt Control Register
	GICR_INT0  = 0x40 // External Interrupt Request 0 Enable
	GICR_INT1  = 0x80 // External Interrupt Request 1 Enable
	GICR_INT2  = 0x20 // External Interrupt Request 2 Enable
	GICR_PCIE  = 0x18 // Pin Change Interrupt Enables
	GICR_IVSEL = 0x2  // Interrupt Vector Select
	GICR_IVCE  = 0x1  // Interrupt Vector Change Enable

	// GIFR: General Interrupt Flag Register
	GIFR_INTF  = 0xc0 // External Interrupt Flags
	GIFR_INTF2 = 0x20 // External Interrupt Flag 2
	GIFR_PCIF  = 0x18 // Pin Change Interrupt Flags

	// PCMSK1: Pin Change Mask Register 1
	PCMSK1_PCINT = 0xff // Pin Change Interrupt mask bits

	// PCMSK0: Pin Change Enable Mask
	PCMSK0_PCINT = 0xff // Pin Change Interrupt mask bits
)