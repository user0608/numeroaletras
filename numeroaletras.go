package numeroaletras

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type NumeroALetras struct {
	unidades           []string
	decenas            []string
	centenas           []string
	acentosExcepciones map[string]string
	Conector           string
	apocope            bool
}

func NewNumeroALetras() *NumeroALetras {
	return &NumeroALetras{
		unidades:           []string{"", "UNO ", "DOS ", "TRES ", "CUATRO ", "CINCO ", "SEIS ", "SIETE ", "OCHO ", "NUEVE ", "DIEZ ", "ONCE ", "DOCE ", "TRECE ", "CATORCE ", "QUINCE ", "DIECISÉIS ", "DIECISIETE ", "DIECIOCHO ", "DIECINUEVE ", "VEINTE "},
		decenas:            []string{"VEINTI", "TREINTA ", "CUARENTA ", "CINCUENTA ", "SESENTA ", "SETENTA ", "OCHENTA ", "NOVENTA ", "CIEN "},
		centenas:           []string{"CIENTO ", "DOSCIENTOS ", "TRESCIENTOS ", "CUATROCIENTOS ", "QUINIENTOS ", "SEISCIENTOS ", "SETECIENTOS ", "OCHOCIENTOS ", "NOVECIENTOS "},
		acentosExcepciones: map[string]string{"VEINTIDOS": "VEINTIDÓS ", "VEINTITRES": "VEINTITRÉS ", "VEINTISEIS": "VEINTISÉIS "},
		Conector:           "CON",
		apocope:            false,
	}
}

func (n *NumeroALetras) redondear(numero float64, decimales int) float64 {
	factor := math.Pow(10, float64(decimales))
	return math.Round(numero*factor) / factor
}

func (n *NumeroALetras) ToWords(number float64, decimals int) (string, error) {
	number = n.redondear(number, decimals)
	numberStr := fmt.Sprintf("%.*f", decimals, number)
	parts := strings.Split(numberStr, ".")

	whole, err := n.wholeNumber(parts[0])
	if err != nil {
		return "", err
	}

	var decimal string
	if len(parts) > 1 && parts[1] != "" {
		decimalInt, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", err
		}
		decimal = n.convertNumber(decimalInt)
	}

	return n.concat([]string{whole, decimal}), nil
}

func (n *NumeroALetras) ToMoney(number float64, decimals int, currency, cents string) (string, error) {
	numberStr := fmt.Sprintf("%.*f", decimals, number)
	parts := strings.Split(numberStr, ".")

	whole, err := n.wholeNumber(parts[0])
	if err != nil {
		return "", err
	}
	whole = strings.TrimSpace(whole) + " " + strings.ToUpper(currency)

	var decimal string
	if len(parts) > 1 && parts[1] != "" && !isZero(parts[1]) {
		decimalInt, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", err
		}
		decimal = n.convertNumber(decimalInt) + " " + strings.ToUpper(cents)
	}

	return n.concat([]string{whole, decimal}), nil
}

func (n *NumeroALetras) ToString(number float64, decimals int, wholeStr, decimalStr string) (string, error) {
	return n.ToMoney(number, decimals, wholeStr, decimalStr)
}

func (n *NumeroALetras) ToInvoice(number float64, decimals int, currency string) (string, error) {
	numberStr := fmt.Sprintf("%.*f", decimals, number)
	parts := strings.Split(numberStr, ".")

	whole, err := n.wholeNumber(parts[0])
	if err != nil {
		return "", err
	}

	decimal := "00/100 "
	if len(parts) > 1 && parts[1] != "" {
		d := mustAtoi(parts[1])
		decimal = fmt.Sprintf("%02d/100 ", d)
	}

	return fmt.Sprintf("%s %s", n.concat([]string{whole, decimal}), strings.ToUpper(currency)), nil
}

func (n *NumeroALetras) UseApocope(value bool) {
	if len(n.unidades) > 1 {
		n.unidades[1] = map[bool]string{true: "UN ", false: "UNO "}[value]
		n.apocope = value
	}
}

func (n *NumeroALetras) wholeNumber(number string) (string, error) {
	if number == "0" {
		return "CERO ", nil
	}
	num, err := strconv.Atoi(number)
	if err != nil {
		return "", err
	}
	return n.convertNumber(num), nil
}

func (n *NumeroALetras) concat(parts []string) string {
	var clean []string
	for _, part := range parts {
		if p := strings.TrimSpace(part); p != "" {
			clean = append(clean, p)
		}
	}
	var results = strings.Join(clean, fmt.Sprintf(" %s ", strings.ToUpper(n.Conector)))
	return strings.ReplaceAll(results, "  ", " ")
}

func (n *NumeroALetras) convertNumber(num int) string {
	if num < 0 || num > 999999999 {
		return "Número fuera de rango"
	}

	s := fmt.Sprintf("%09d", num)
	mill, _ := strconv.Atoi(s[0:3])
	thou, _ := strconv.Atoi(s[3:6])
	hund, _ := strconv.Atoi(s[6:9])

	var res strings.Builder
	if mill > 0 {
		if mill == 1 {
			res.WriteString("UN MILLÓN ")
		} else {
			res.WriteString(fmt.Sprintf("%s MILLONES ", n.convertGroup(s[0:3])))
		}
	}
	if thou > 0 {
		if thou == 1 {
			res.WriteString("MIL ")
		} else {
			res.WriteString(fmt.Sprintf("%s MIL ", n.convertGroup(s[3:6])))
		}
	}
	if hund > 0 {
		if hund == 1 {
			res.WriteString(map[bool]string{true: "UN ", false: "UNO "}[n.apocope])
		} else {
			res.WriteString(fmt.Sprintf("%s ", n.convertGroup(s[6:9])))
		}
	}
	return strings.TrimSpace(res.String())
}

func (n *NumeroALetras) convertGroup(group string) string {
	if group == "100" {
		return "CIEN "
	}

	h := int(group[0] - '0')
	t := int(group[1] - '0')
	u := int(group[2] - '0')
	lastTwo := t*10 + u

	var res strings.Builder
	if h > 0 {
		res.WriteString(n.centenas[h-1])
	}
	var unit string
	if lastTwo <= 20 {
		unit = n.unidades[lastTwo]
	} else {
		if t-2 >= 0 && t-2 < len(n.decenas) {
			if lastTwo > 30 && u != 0 {
				unit = fmt.Sprintf("%sY %s", n.decenas[t-2], n.unidades[u])
			} else {
				unit = fmt.Sprintf("%s%s", n.decenas[t-2], n.unidades[u])
			}
		}
	}
	unit = strings.TrimSpace(unit)
	if val, ok := n.acentosExcepciones[strings.ToUpper(unit)]; ok {
		unit = val
	}
	res.WriteString(unit)
	return res.String()
}

func isZero(s string) bool {
	return strings.TrimLeft(s, "0") == ""
}

func mustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
