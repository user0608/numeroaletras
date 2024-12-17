package numeroaletras

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// NumeroALetras is a struct that provides methods to convert numbers to their Spanish word representations.
type NumeroALetras struct {
	unidades           []string
	decenas            []string
	centenas           []string
	acentosExcepciones map[string]string
	Conector           string
	apocope            bool
}

// NewNumeroALetras initializes and returns a new instance of NumeroALetras with default values.
func NewNumeroALetras() *NumeroALetras {
	return &NumeroALetras{
		unidades: []string{
			"", "UNO ", "DOS ", "TRES ", "CUATRO ", "CINCO ",
			"SEIS ", "SIETE ", "OCHO ", "NUEVE ", "DIEZ ",
			"ONCE ", "DOCE ", "TRECE ", "CATORCE ", "QUINCE ",
			"DIECISÉIS ", "DIECISIETE ", "DIECIOCHO ", "DIECINUEVE ", "VEINTE ",
		},
		decenas: []string{
			"VEINTI", "TREINTA ", "CUARENTA ", "CINCUENTA ",
			"SESENTA ", "SETENTA ", "OCHENTA ", "NOVENTA ", "CIEN ",
		},
		centenas: []string{
			"CIENTO ", "DOSCIENTOS ", "TRESCIENTOS ", "CUATROCIENTOS ",
			"QUINIENTOS ", "SEISCIENTOS ", "SETECIENTOS ", "OCHOCIENTOS ",
			"NOVECIENTOS ",
		},
		acentosExcepciones: map[string]string{
			"VEINTIDOS":  "VEINTIDÓS ",
			"VEINTITRES": "VEINTITRÉS ",
			"VEINTISEIS": "VEINTISÉIS ",
		},
		Conector: "CON",
		apocope:  false,
	}
}
func (n *NumeroALetras) redondear(numero float64, decimales int) float64 {
	factor := math.Pow(10, float64(decimales))
	return math.Round(numero*factor) / factor
}

// ToWords converts a numeric value to its word representation.
// number: The number to be converted.
// decimals: The number of decimal places to consider. Default is 2.
func (n *NumeroALetras) ToWords(number float64, decimals int) (string, error) {
	number = n.redondear(number, decimals)
	numberStr := fmt.Sprintf("%.*f", decimals, number)

	splitNumber := strings.Split(numberStr, ".")

	wholePart, err := n.wholeNumber(splitNumber[0])
	if err != nil {
		return "", err
	}

	var decimalPart string
	if len(splitNumber) > 1 && splitNumber[1] != "" {
		decimalInt, err := strconv.Atoi(splitNumber[1])
		if err != nil {
			return "", err
		}
		decimalPart = n.convertNumber(decimalInt)
	}

	return n.concat([]string{wholePart, decimalPart}), nil
}
func IsZero(s string) bool {
	if len(s) == 0 {
		return true
	}
	for _, c := range s {
		if c != '0' {
			return false
		}
	}
	return true
}

// ToMoney converts a numeric value to its money representation in words.
// number: The numeric value to be converted.
// decimals: The number of decimal places to consider. Default is 2.
// currency: The currency to append to the whole number part.
// cents: The currency to append to the decimal part.
func (n *NumeroALetras) ToMoney(number float64, decimals int, currency, cents string) (string, error) {

	numberStr := fmt.Sprintf("%.*f", decimals, number)
	splitNumber := strings.Split(numberStr, ".")

	wholePart, err := n.wholeNumber(splitNumber[0])
	if err != nil {
		return "", err
	}
	wholePart = strings.TrimSpace(wholePart) + " " + strings.ToUpper(currency)

	var decimalPart string
	if len(splitNumber) > 1 && splitNumber[1] != "" && !IsZero(splitNumber[1]) {
		decimalInt, err := strconv.Atoi(splitNumber[1])
		if err != nil {
			return "", err
		}
		decimalPart = n.convertNumber(decimalInt) + " " + strings.ToUpper(cents)
	}

	return n.concat([]string{wholePart, decimalPart}), nil
}

// ToString converts a number to its string representation with specified whole and decimal strings.
// number: The number to be converted.
// decimals: The number of decimal places to include.
// wholeStr: The string to use for the whole number part.
// decimalStr: The string to use for the decimal part.
func (n *NumeroALetras) ToString(number float64, decimals int, wholeStr, decimalStr string) (string, error) {
	return n.ToMoney(number, decimals, wholeStr, decimalStr)
}

// ToInvoice converts a number to its invoice representation in words.
// number: The number to be converted.
// decimals: The number of decimal places to consider. Default is 2.
// currency: The currency to append at the end.
func (n *NumeroALetras) ToInvoice(number float64, decimals int, currency string) (string, error) {

	numberStr := fmt.Sprintf("%.*f", decimals, number)
	splitNumber := strings.Split(numberStr, ".")

	wholeNumber, err := n.wholeNumber(splitNumber[0])
	if err != nil {
		return "", err
	}

	var decimalPart string
	if len(splitNumber) > 1 && splitNumber[1] != "" {
		decimalPart = fmt.Sprintf("%02d/100 ", mustAtoi(splitNumber[1]))
	} else {
		decimalPart = "00/100 "
	}

	invoice := fmt.Sprintf("%s %s", n.concat([]string{wholeNumber, decimalPart}), strings.ToUpper(currency))
	return invoice, nil
}

// checkApocope modifies the 'unidades' array if Apocope is set to true.
func (n *NumeroALetras) UseApocope(value bool) {
	if len(n.unidades) > 1 {
		if value {
			n.unidades[1] = "UN "

		} else {
			n.unidades[1] = "UNO "
		}
		n.apocope = value
	}
}

// wholeNumber converts the whole number part to words.
func (n *NumeroALetras) wholeNumber(number string) (string, error) {
	if number == "0" {
		return "CERO ", nil
	}
	numInt, err := strconv.Atoi(number)
	if err != nil {
		return "", err
	}
	return n.convertNumber(numInt), nil
}

// concat joins the split number parts with the connector.
func (n *NumeroALetras) concat(splitNumber []string) string {
	var parts []string
	for _, part := range splitNumber {
		part = strings.TrimSpace(part)
		if part != "" {
			parts = append(parts, part)
		}
	}
	conector := strings.ToUpper(n.Conector)
	return strings.Join(parts, fmt.Sprintf(" %s ", conector))
}

// convertNumber converts a numeric value into its Spanish words representation.
func (n *NumeroALetras) convertNumber(number int) string {
	if number < 0 || number > 999999999 {
		return "Número fuera de rango"
	}

	converted := ""

	numberStrFill := fmt.Sprintf("%09d", number)
	millonesStr := numberStrFill[0:3]
	milesStr := numberStrFill[3:6]
	cientosStr := numberStrFill[6:9]

	millones, _ := strconv.Atoi(millonesStr)
	miles, _ := strconv.Atoi(milesStr)
	cientos, _ := strconv.Atoi(cientosStr)

	if millones > 0 {
		if millones == 1 {
			converted += "UN MILLÓN "
		} else {
			converted += fmt.Sprintf("%s MILLONES ", n.convertGroup(millonesStr))
		}
	}

	if miles > 0 {
		if miles == 1 {
			converted += "MIL "
		} else {
			converted += fmt.Sprintf("%s MIL ", n.convertGroup(milesStr))
		}
	}

	if cientos > 0 {
		if cientos == 1 {
			if n.apocope {
				converted += "UN "
			} else {
				converted += "UNO "
			}
		} else {
			converted += fmt.Sprintf("%s ", n.convertGroup(cientosStr))
		}
	}

	return strings.TrimSpace(converted)
}

// convertGroup converts a three-digit group into words.
func (n *NumeroALetras) convertGroup(group string) string {
	output := ""

	// Handle hundreds
	if group == "100" {
		output = "CIEN "
	} else if group[0] != '0' {
		hundredsIndex := int(group[0]-'0') - 1
		if hundredsIndex >= 0 && hundredsIndex < len(n.centenas) {
			output = n.centenas[hundredsIndex]
		}
	}
	var unidades string
	// Handle tens and units
	k, _ := strconv.Atoi(group[1:])
	if k <= 20 {
		if k >= 0 && k < len(n.unidades) {
			unidades = n.unidades[k]
		}
	} else {
		tensIndex := (k / 10) - 2
		unitIndex := k % 10
		if tensIndex >= 0 && tensIndex < len(n.decenas) {
			if k > 30 && unitIndex != 0 {
				unidades = fmt.Sprintf("%sY %s", n.decenas[tensIndex], n.unidades[unitIndex])
			} else {
				unidades = fmt.Sprintf("%s%s", n.decenas[tensIndex], n.unidades[unitIndex])
			}
		}
	}
	// Handle accent exceptions
	trimmedUnidades := strings.TrimSpace(unidades)
	if val, exists := n.acentosExcepciones[trimmedUnidades]; exists {
		trimmedUnidades = val
	}
	if trimmedUnidades == "" {
		return strings.TrimSpace(output)
	}
	return fmt.Sprintf("%s%s", output, trimmedUnidades)
}

// mustAtoi converts a string to an integer and panics if there's an error.
// It's used internally where the input is guaranteed to be numeric.
func mustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
