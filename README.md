# NumeroALetras para Go

**NumeroALetras** es una biblioteca escrita en Go que convierte valores numéricos a su representación literal en español. Es útil para facturación, reportes financieros, documentos legales y cualquier sistema donde se requiera expresar montos en texto.

Este proyecto está basado en la biblioteca [php-numero-a-letras](https://github.com/luecano/php-numero-a-letras) desarrollada por [@luecano](https://github.com/luecano) y licenciada bajo [MIT License](https://github.com/luecano/php-numero-a-letras/blob/master/LICENSE).

---

## ✨ Características

- Conversión de números enteros y decimales a palabras en español.
- Representación de montos con moneda y centavos.
- Formato especial para facturación electrónica SUNAT (`45.50` → `CUARENTA Y CINCO 50/100 SOLES`).
- Apócope opcional de “UNO” a “UN”.
- Corrección de acentos en casos especiales (`VEINTIDÓS`, `VEINTITRÉS`, `VEINTISÉIS`).
- Personalización del conector (por defecto: "CON").

---

## 📦 Instalación

```bash
go get github.com/user0608/numeroaletras
```

---

## 🧪 Ejemplos de uso

### Convertir número a palabras

```go
n := numeroaletras.NewNumeroALetras()
res, _ := n.ToWords(1100, 0)
fmt.Println(res)
// Salida: "MIL CIEN"
```

### Activar apócope

```go
n := numeroaletras.NewNumeroALetras()
n.UseApocope(true)
res, _ := n.ToWords(101, 0)
fmt.Println(res + " AÑOS")
// Salida: "CIENTO UN AÑOS"
```

### Representación monetaria

```go
n := numeroaletras.NewNumeroALetras()
res, _ := n.ToMoney(2500.90, 2, "DÓLARES", "CENTAVOS")
fmt.Println(res)
// Salida: "DOS MIL QUINIENTOS DÓLARES CON NOVENTA CENTAVOS"
```

### Representación con conector personalizado

```go
n := numeroaletras.NewNumeroALetras()
n.Conector = "Y"
res, _ := n.ToMoney(11.10, 2, "pesos", "centavos")
fmt.Println(res)
// Salida: "ONCE PESOS Y DIEZ CENTAVOS"
```

### Formato SUNAT (factura)

```go
n := numeroaletras.NewNumeroALetras()
res, _ := n.ToInvoice(1700.50, 2, "soles")
fmt.Println(res)
// Salida: "MIL SETECIENTOS CON 50/100 SOLES"
```

### Formato libre

```go
n := numeroaletras.NewNumeroALetras()
res, _ := n.ToString(5.2, 1, "años", "meses")
fmt.Println(res)
// Salida: "CINCO AÑOS CON DOS MESES"
```
