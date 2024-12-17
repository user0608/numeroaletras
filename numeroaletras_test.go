package numeroaletras

import (
	"strings"
	"testing"
)

// Helper function to compare strings case-insensitively and ignore trailing spaces
func equalsIgnoreCaseAndTrim(a, b string) bool {
	return strings.EqualFold(a, b)
}
func TestToWordsCombined(t *testing.T) {
	tests := map[string]struct {
		number   float64
		decimals int
		expected string
	}{
		// Casos de TestToWords
		"Cien": {
			number:   100,
			decimals: 0,
			expected: "CIEN",
		},
		"Cien con noventa y nueve": {
			number:   100.99,
			decimals: 2,
			expected: "CIEN CON NOVENTA Y NUEVE",
		},
		"Ciento uno con decimales": {
			number:   100.9999,
			decimals: 2,
			expected: "CIENTO UNO",
		},
		"Dieciséis": {
			number:   16,
			decimals: 0,
			expected: "DIECISÉIS",
		},
		"Mil dieciséis": {
			number:   1016,
			decimals: 0,
			expected: "MIL DIECISÉIS",
		},
		"Ochenta y cuatro": {
			number:   84,
			decimals: 0,
			expected: "OCHENTA Y CUATRO",
		},
		"Ochenta y cuatro con decimales": {
			number:   84,
			decimals: 4,
			expected: "OCHENTA Y CUATRO",
		},
		"Ochenta y cuatro con veinte": {
			number:   84.2,
			decimals: 2,
			expected: "OCHENTA Y CUATRO CON VEINTE",
		},

		// Casos de TestToWordsThousands
		"Mil cien": {
			number:   1100,
			decimals: 0,
			expected: "MIL CIEN",
		},
		"Mil doscientos treinta y cuatro": {
			number:   1234,
			decimals: 0,
			expected: "MIL DOSCIENTOS TREINTA Y CUATRO",
		},
		"Dos mil quinientos": {
			number:   2500,
			decimals: 0,
			expected: "DOS MIL QUINIENTOS",
		},
		"Cinco mil con decimales": {
			number:   5000.75,
			decimals: 2,
			expected: "CINCO MIL CON SETENTA Y CINCO",
		},
		"Diez mil novecientos noventa y nueve": {
			number:   10999,
			decimals: 0,
			expected: "DIEZ MIL NOVECIENTOS NOVENTA Y NUEVE",
		},

		// Casos de TestToWordsMillions
		"Un millón": {
			number:   1000000,
			decimals: 0,
			expected: "UN MILLÓN",
		},
		"Dos millones quinientos mil": {
			number:   2500000,
			decimals: 0,
			expected: "DOS MILLONES QUINIENTOS MIL",
		},
		"Cinco millones con setecientos cincuenta mil": {
			number:   5750000,
			decimals: 0,
			expected: "CINCO MILLONES SETECIENTOS CINCUENTA MIL",
		},
		"Diez millones novecientos noventa y nueve mil novecientos noventa y nueve": {
			number:   10999999,
			decimals: 0,
			expected: "DIEZ MILLONES NOVECIENTOS NOVENTA Y NUEVE MIL NOVECIENTOS NOVENTA Y NUEVE",
		},
		"Un millón con cincuenta centavos": {
			number:   1000000.50,
			decimals: 2,
			expected: "UN MILLÓN CON CINCUENTA",
		},
		"Doscientos millones con treinta y cinco mil doscientos": {
			number:   200035200,
			decimals: 0,
			expected: "DOSCIENTOS MILLONES TREINTA Y CINCO MIL DOSCIENTOS",
		},
		"Quince millones con noventa y nueve mil novecientos noventa y nueve": {
			number:   15999999,
			decimals: 0,
			expected: "QUINCE MILLONES NOVECIENTOS NOVENTA Y NUEVE MIL NOVECIENTOS NOVENTA Y NUEVE",
		},
		"Cero con cero millones": {
			number:   0,
			decimals: 0,
			expected: "CERO",
		},
	}

	formatter := NewNumeroALetras()

	for name, tt := range tests {
		tt := tt // Capturar la variable para evitar problemas con clausuras
		t.Run(name, func(t *testing.T) {
			words, err := formatter.ToWords(tt.number, tt.decimals)
			if err != nil {
				t.Errorf("ToWords(%v, %v) retornó error: %v", tt.number, tt.decimals, err)
				return
			}
			if words != tt.expected {
				t.Errorf("ToWords(%v, %v) = %v; se esperaba %v", tt.number, tt.decimals, words, tt.expected)
			}
		})
	}
}

func TestToWordsCombined_useApocope(t *testing.T) {
	tests := map[string]struct {
		number   float64
		decimals int
		expected string
	}{
		// Casos donde apócope afecta el resultado
		"Ciento uno con decimales": {
			number:   100.9999,
			decimals: 2,
			expected: "CIENTO UN",
		},
		"Un millón": {
			number:   1000000,
			decimals: 0,
			expected: "UN MILLÓN",
		},
		"Un millón con cincuenta centavos": {
			number:   1000000.50,
			decimals: 2,
			expected: "UN MILLÓN CON CINCUENTA",
		},
		"Doscientos uno": {
			number:   201,
			decimals: 0,
			expected: "DOSCIENTOS UN",
		},
		"Mil uno": {
			number:   1001,
			decimals: 0,
			expected: "MIL UN",
		},
		"Un mil doscientos uno": {
			number:   1201,
			decimals: 0,
			expected: "MIL DOSCIENTOS UN",
		},
		"Un millón con un centavo": {
			number:   1000000.01,
			decimals: 2,
			expected: "UN MILLÓN CON UN",
		},
		// Puedes agregar más casos aquí que requieran apócope
	}

	formatter := NewNumeroALetras()
	formatter.UseApocope(true) // Habilitar apócope

	for name, tt := range tests {
		tt := tt // Capturar la variable para evitar problemas con clausuras
		t.Run(name, func(t *testing.T) {
			words, err := formatter.ToWords(tt.number, tt.decimals)
			if err != nil {
				t.Errorf("ToWords(%v, %v) retornó error: %v", tt.number, tt.decimals, err)
				return
			}
			if words != tt.expected {
				t.Errorf("ToWords(%v, %v) = %v; se esperaba %v", tt.number, tt.decimals, words, tt.expected)
			}
		})
	}
}

func TestToMoney(t *testing.T) {
	tests := map[string]struct {
		number      float64
		decimals    int
		currency    string
		subCurrency string
		expected    string
	}{
		"Mil cien soles": {
			number:      1100,
			decimals:    2,
			currency:    "SOLES",
			subCurrency: "CENTIMOS",
			expected:    "MIL CIEN SOLES",
		},
		"Mil cien con decimales": {
			number:      1100.50,
			decimals:    2,
			currency:    "SOLES",
			subCurrency: "CENTIMOS",
			expected:    "MIL CIEN SOLES CON CINCUENTA CENTIMOS",
		},
		"Dos mil quinientos soles": {
			number:      2500,
			decimals:    2,
			currency:    "SOLES",
			subCurrency: "CENTIMOS",
			expected:    "DOS MIL QUINIENTOS SOLES",
		},
		"Cinco mil con decimales": {
			number:      5000.75,
			decimals:    2,
			currency:    "SOLES",
			subCurrency: "CENTIMOS",
			expected:    "CINCO MIL SOLES CON SETENTA Y CINCO CENTIMOS",
		},
		"Diez mil novecientos noventa y nueve soles": {
			number:      10999,
			decimals:    2,
			currency:    "SOLES",
			subCurrency: "CENTIMOS",
			expected:    "DIEZ MIL NOVECIENTOS NOVENTA Y NUEVE SOLES",
		},
		"Un sol con cero centimos": {
			number:      1,
			decimals:    2,
			currency:    "SOLES",
			subCurrency: "CENTIMOS",
			expected:    "UNO SOLES",
		},
		"Cero soles con cincuenta centimos": {
			number:      0.50,
			decimals:    2,
			currency:    "SOLES",
			subCurrency: "CENTIMOS",
			expected:    "CERO SOLES CON CINCUENTA CENTIMOS",
		},
		"Miles y unidades con decimales": {
			number:      1234.56,
			decimals:    2,
			currency:    "SOLES",
			subCurrency: "CENTIMOS",
			expected:    "MIL DOSCIENTOS TREINTA Y CUATRO SOLES CON CINCUENTA Y SEIS CENTIMOS",
		},
		// Puedes agregar más casos aquí según sea necesario
	}

	formatter := NewNumeroALetras()

	for name, tt := range tests {
		tt := tt // Capturar la variable para evitar problemas con clausuras
		t.Run(name, func(t *testing.T) {
			money, err := formatter.ToMoney(tt.number, tt.decimals, tt.currency, tt.subCurrency)
			if err != nil {
				t.Errorf("ToMoney(%v, %v, '%v', '%v') returned error: %v", tt.number, tt.decimals, tt.currency, tt.subCurrency, err)
				return
			}
			if money != tt.expected {
				t.Errorf("ToMoney(%v, %v, '%v', '%v') = %v; want %v", tt.number, tt.decimals, tt.currency, tt.subCurrency, money, tt.expected)
			}
		})
	}
}

func TestToMoneyFloat(t *testing.T) {
	tests := map[string]struct {
		number      float64
		decimals    int
		currency    string
		subCurrency string
		expected    string
	}{
		"Diez con diez centimos": {
			number:      10.10,
			decimals:    2,
			currency:    "SOLES",
			subCurrency: "CENTIMOS",
			expected:    "DIEZ SOLES CON DIEZ CENTIMOS",
		},
		"Doscientos con cincuenta centimos": {
			number:      200.50,
			decimals:    2,
			currency:    "SOLES",
			subCurrency: "CENTIMOS",
			expected:    "DOSCIENTOS SOLES CON CINCUENTA CENTIMOS",
		},
		"Un sol con cero centimos": {
			number:      1.00,
			decimals:    2,
			currency:    "SOLES",
			subCurrency: "CENTIMOS",
			expected:    "UNO SOLES",
		},
		"Cero soles con cincuenta centimos": {
			number:      0.50,
			decimals:    2,
			currency:    "SOLES",
			subCurrency: "CENTIMOS",
			expected:    "CERO SOLES CON CINCUENTA CENTIMOS",
		},
		"Miles y unidades con decimales": {
			number:      1234.56,
			decimals:    2,
			currency:    "SOLES",
			subCurrency: "CENTIMOS",
			expected:    "MIL DOSCIENTOS TREINTA Y CUATRO SOLES CON CINCUENTA Y SEIS CENTIMOS",
		},
	}

	formatter := NewNumeroALetras()

	for name, tt := range tests {
		tt := tt // Capturar la variable para evitar problemas con clausuras
		t.Run(name, func(t *testing.T) {
			money, err := formatter.ToMoney(tt.number, tt.decimals, tt.currency, tt.subCurrency)
			if err != nil {
				t.Errorf("ToMoney(%v, %v, '%v', '%v') returned error: %v", tt.number, tt.decimals, tt.currency, tt.subCurrency, err)
				return
			}
			if money != tt.expected {
				t.Errorf("ToMoney(%v, %v, '%v', '%v') = %v; want %v", tt.number, tt.decimals, tt.currency, tt.subCurrency, money, tt.expected)
			}
		})
	}
}

func TestToInvoice(t *testing.T) {
	tests := map[string]struct {
		number   float64
		decimals int
		currency string
		expected string
	}{
		"CIEN CON CERO SOLES": {
			number:   100,
			decimals: 2,
			currency: "soles",
			expected: "CIEN CON 00/100 SOLES",
		},
		"CIENTO VEINTITRÉS CON CINCUENTA SOLES": {
			number:   123.50,
			decimals: 2,
			currency: "soles",
			expected: "CIENTO VEINTITRÉS CON 50/100 SOLES",
		},
		"UN CON VEINTE SOLES": {
			number:   1.20,
			decimals: 2,
			currency: "soles",
			expected: "UNO CON 20/100 SOLES",
		},
		"CERO CON CERO SOLES": {
			number:   0,
			decimals: 2,
			currency: "soles",
			expected: "CERO CON 00/100 SOLES",
		},
		"QUINIENTOS CON NOVENTA Y NUEVE SOLES": {
			number:   599.99,
			decimals: 2,
			currency: "soles",
			expected: "QUINIENTOS NOVENTA Y NUEVE CON 99/100 SOLES",
		},
		"QUINIENTOS CON NOVECIENTOS NOVENTA Y NUEVE SOLES": {
			number:   599.999,
			decimals: 2,
			currency: "soles",
			expected: "SEISCIENTOS CON 00/100 SOLES",
		},
	}

	formatter := NewNumeroALetras()

	for name, tt := range tests {
		tt := tt // Capturar la variable para evitar problemas con clausuras
		t.Run(name, func(t *testing.T) {
			invoice, err := formatter.ToInvoice(tt.number, tt.decimals, tt.currency)
			if err != nil {
				t.Errorf("ToInvoice(%v, %v, '%v') returned error: %v", tt.number, tt.decimals, tt.currency, err)
				return
			}
			if invoice != tt.expected {
				t.Errorf("ToInvoice(%v, %v, '%v') = %v; want %v", tt.number, tt.decimals, tt.currency, invoice, tt.expected)
			}
		})
	}
}

func TestToInvoiceFloat(t *testing.T) {
	tests := map[string]struct {
		number   float64
		decimals int
		currency string
		expected string
	}{
		"Mil setecientos con cincuenta soles": {
			number:   1700.50,
			decimals: 2,
			currency: "SOLES",
			expected: "MIL SETECIENTOS CON 50/100 SOLES",
		},
		"Mil setecientos con novecientos noventa y nueve": {
			number:   1700.999,
			decimals: 2,
			currency: "SOLES",
			expected: "MIL SETECIENTOS UNO CON 00/100 SOLES",
		},
		"Diecisiete con cero soles": {
			number:   17.00,
			decimals: 2,
			currency: "SOLES",
			expected: "DIECISIETE CON 00/100 SOLES",
		},
		"Un mil con veinticinco soles": {
			number:   1000.25,
			decimals: 2,
			currency: "SOLES",
			expected: "MIL CON 25/100 SOLES",
		},
		"Cero con cincuenta soles": {
			number:   0.50,
			decimals: 2,
			currency: "SOLES",
			expected: "CERO CON 50/100 SOLES",
		},
		"Dos mil trescientos cuarenta y cinco con setenta y ocho soles": {
			number:   2345.78,
			decimals: 2,
			currency: "SOLES",
			expected: "DOS MIL TRESCIENTOS CUARENTA Y CINCO CON 78/100 SOLES",
		},
	}

	formatter := NewNumeroALetras()

	for name, tt := range tests {
		tt := tt // Capturar la variable para evitar problemas con clausuras
		t.Run(name, func(t *testing.T) {
			invoice, err := formatter.ToInvoice(tt.number, tt.decimals, tt.currency)
			if err != nil {
				t.Errorf("ToInvoice(%v, %v, '%v') returned error: %v", tt.number, tt.decimals, tt.currency, err)
				return
			}
			if invoice != tt.expected {
				t.Errorf("ToInvoice(%v, %v, '%v') = %v; want %v", tt.number, tt.decimals, tt.currency, invoice, tt.expected)
			}
		})
	}
}
func TestApocope(t *testing.T) {
	tests := map[string]struct {
		number   float64
		decimals int
		apocope  bool
		expected string
	}{
		"101 con apocope=true": {
			number:   101,
			decimals: 0,
			apocope:  true,
			expected: "CIENTO UN",
		},
		"101 con apocope=false": {
			number:   101,
			decimals: 0,
			apocope:  false,
			expected: "CIENTO UNO",
		},
		"201 con apocope=true": {
			number:   201,
			decimals: 0,
			apocope:  true,
			expected: "DOSCIENTOS UN",
		},
		"201 con apocope=false": {
			number:   201,
			decimals: 0,
			apocope:  false,
			expected: "DOSCIENTOS UNO",
		},
		// Puedes agregar más casos aquí según sea necesario
	}

	formatter := NewNumeroALetras()

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			formatter.UseApocope(tt.apocope)
			words, err := formatter.ToWords(tt.number, tt.decimals)
			if err != nil {
				t.Errorf("ToWords(%v, %v) con apocope=%v retornó error: %v", tt.number, tt.decimals, tt.apocope, err)
				return
			}
			if words != tt.expected {
				t.Errorf("ToWords(%v, %v) con apocope=%v = %v; se esperaba %v", tt.number, tt.decimals, tt.apocope, words, tt.expected)
			}
		})
	}
}

func TestToString(t *testing.T) {
	tests := map[string]struct {
		number   float64
		decimals int
		unit     string
		subUnit  string
		expected string
	}{
		"CINCO AÑOS CON DOS MESES": {
			number:   5.2,
			decimals: 1,
			unit:     "años",
			subUnit:  "meses",
			expected: "CINCO AÑOS CON DOS MESES",
		},
		"UN AÑO CON UN MES": {
			number:   1.1,
			decimals: 1,
			unit:     "años",
			subUnit:  "meses",
			expected: "UNO AÑOS CON UNO MESES",
		},
		"CERO AÑOS CON CINCO MESES": {
			number:   0.5,
			decimals: 1,
			unit:     "años",
			subUnit:  "meses",
			expected: "CERO AÑOS CON CINCO MESES",
		},
		"DOCE AÑOS CON DIEZ MESES": {
			number:   12.10,
			decimals: 2,
			unit:     "años",
			subUnit:  "meses",
			expected: "DOCE AÑOS CON DIEZ MESES",
		},
		"VEINTICINCO AÑOS CON CERO MESES": {
			number:   25.0,
			decimals: 2,
			unit:     "años",
			subUnit:  "meses",
			expected: "VEINTICINCO AÑOS",
		},
		// Puedes agregar más casos aquí según sea necesario
	}

	formatter := NewNumeroALetras()

	for name, tt := range tests {
		tt := tt // Capturar la variable para evitar problemas con clausuras
		t.Run(name, func(t *testing.T) {
			str, err := formatter.ToString(tt.number, tt.decimals, tt.unit, tt.subUnit)
			if err != nil {
				t.Errorf("ToString(%v, %v, '%v', '%v') returned error: %v", tt.number, tt.decimals, tt.unit, tt.subUnit, err)
				return
			}
			if !equalsIgnoreCaseAndTrim(str, tt.expected) {
				t.Errorf("ToString(%v, %v, '%v', '%v') = %v; want %v", tt.number, tt.decimals, tt.unit, tt.subUnit, str, tt.expected)
			}
		})
	}
}
