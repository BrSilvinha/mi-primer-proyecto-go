package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// ===========================================
// FUNCIONES VARIÁDICAS EN GO
// ===========================================

// 1. Función variádica básica - suma números
func sumar(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}

// 2. Función variádica con parámetros fijos y variables
func saludarConMensajes(nombre string, mensajes ...string) {
	fmt.Printf("👋 Hola %s!\n", nombre)
	
	if len(mensajes) == 0 {
		fmt.Println("   (Sin mensajes adicionales)")
		return
	}
	
	fmt.Println("   Mensajes:")
	for i, mensaje := range mensajes {
		fmt.Printf("   %d. %s\n", i+1, mensaje)
	}
}

// 3. Función variádica para logging con niveles
func log(nivel string, formato string, args ...interface{}) {
	timestamp := time.Now().Format("15:04:05")
	prefijo := fmt.Sprintf("[%s] [%s]", timestamp, strings.ToUpper(nivel))
	
	mensaje := fmt.Sprintf(formato, args...)
	fmt.Printf("%s %s\n", prefijo, mensaje)
}

// 4. Función variádica para cálculos estadísticos
func calcularEstadisticas(numeros ...float64) (min, max, promedio float64, count int) {
	if len(numeros) == 0 {
		return 0, 0, 0, 0
	}
	
	min = numeros[0]
	max = numeros[0]
	suma := 0.0
	
	for _, num := range numeros {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		suma += num
	}
	
	promedio = suma / float64(len(numeros))
	count = len(numeros)
	return
}

// 5. Función variádica para concatenar strings con separador
func concatenar(separador string, elementos ...string) string {
	if len(elementos) == 0 {
		return ""
	}
	
	return strings.Join(elementos, separador)
}

// 6. Función variádica avanzada - validador de campos
func validarCampos(requeridos bool, campos ...string) (bool, []string) {
	var camposVacios []string
	
	for _, campo := range campos {
		if strings.TrimSpace(campo) == "" {
			camposVacios = append(camposVacios, "campo vacío")
		}
	}
	
	if requeridos && len(camposVacios) > 0 {
		return false, camposVacios
	}
	
	return len(camposVacios) == 0, camposVacios
}

// 7. Función variádica con diferentes tipos - builder de consultas
func construirSQL(tabla string, operacion string, condiciones ...string) string {
	var query strings.Builder
	
	switch strings.ToUpper(operacion) {
	case "SELECT":
		query.WriteString("SELECT * FROM " + tabla)
	case "DELETE":
		query.WriteString("DELETE FROM " + tabla)
	case "UPDATE":
		query.WriteString("UPDATE " + tabla + " SET ")
	default:
		return "Operación no soportada"
	}
	
	if len(condiciones) > 0 {
		query.WriteString(" WHERE ")
		query.WriteString(strings.Join(condiciones, " AND "))
	}
	
	return query.String() + ";"
}

// 8. Función variádica para operaciones matemáticas
func operar(operacion string, numeros ...float64) (float64, error) {
	if len(numeros) == 0 {
		return 0, fmt.Errorf("se requiere al menos un número")
	}
	
	resultado := numeros[0]
	
	switch strings.ToLower(operacion) {
	case "suma":
		for _, num := range numeros[1:] {
			resultado += num
		}
	case "multiplicacion":
		for _, num := range numeros[1:] {
			resultado *= num
		}
	case "maximo":
		for _, num := range numeros[1:] {
			if num > resultado {
				resultado = num
			}
		}
	case "minimo":
		for _, num := range numeros[1:] {
			if num < resultado {
				resultado = num
			}
		}
	default:
		return 0, fmt.Errorf("operación '%s' no soportada", operacion)
	}
	
	return resultado, nil
}

// 9. Función variádica que acepta funciones - pipeline de procesamiento
func procesarDatos(datos []int, funciones ...func(int) int) []int {
	resultado := make([]int, len(datos))
	copy(resultado, datos)
	
	// Aplicar cada función en secuencia
	for _, fn := range funciones {
		for i, valor := range resultado {
			resultado[i] = fn(valor)
		}
	}
	
	return resultado
}

// 10. Función variádica para reporting - generar reportes
func generarReporte(titulo string, secciones ...func() string) string {
	var reporte strings.Builder
	
	// Encabezado
	reporte.WriteString("=" + strings.Repeat("=", len(titulo)+2) + "=\n")
	reporte.WriteString("| " + titulo + " |\n")
	reporte.WriteString("=" + strings.Repeat("=", len(titulo)+2) + "=\n\n")
	
	// Agregar cada sección
	for i, seccion := range secciones {
		reporte.WriteString(fmt.Sprintf("--- Sección %d ---\n", i+1))
		reporte.WriteString(seccion())
		reporte.WriteString("\n\n")
	}
	
	reporte.WriteString("Generado: " + time.Now().Format("02/01/2006 15:04:05"))
	return reporte.String()
}

// ===========================================
// FUNCIÓN PRINCIPAL PARA DEMOSTRAR EJEMPLOS
// ===========================================

func main() {
	fmt.Println("🔢 EJEMPLOS DE FUNCIONES VARIÁDICAS EN GO")
	fmt.Println("=" + strings.Repeat("=", 45))

	// 1. Función variádica básica
	fmt.Println("\n1️⃣ Suma variádica:")
	fmt.Printf("Sin números: %d\n", sumar())
	fmt.Printf("Un número: %d\n", sumar(5))
	fmt.Printf("Varios números: %d\n", sumar(1, 2, 3, 4, 5))
	
	// Usando slice con operador ...
	numeros := []int{10, 20, 30, 40, 50}
	fmt.Printf("Desde slice: %d\n", sumar(numeros...))

	// 2. Parámetros fijos + variádicos
	fmt.Println("\n2️⃣ Saludos con mensajes:")
	saludarConMensajes("Ana")
	saludarConMensajes("Carlos", "¡Bienvenido!")
	saludarConMensajes("María", "Feliz cumpleaños", "Que tengas un gran día", "🎉")

	// 3. Logging variádico
	fmt.Println("\n3️⃣ Sistema de logging:")
	log("info", "Usuario %s ha iniciado sesión", "juan.perez")
	log("warning", "Quedan %d intentos para el usuario %s", 2, "ana.garcia")
	log("error", "Falló la conexión a la base de datos: %s", "timeout")

	// 4. Cálculos estadísticos
	fmt.Println("\n4️⃣ Estadísticas:")
	datos := []float64{1.5, 2.8, 3.2, 1.9, 4.1, 2.3, 3.7}
	min, max, prom, count := calcularEstadisticas(datos...)
	fmt.Printf("Datos: %v\n", datos)
	fmt.Printf("Count: %d, Min: %.2f, Max: %.2f, Promedio: %.2f\n", count, min, max, prom)

	// 5. Concatenación flexible
	fmt.Println("\n5️⃣ Concatenación:")
	fmt.Printf("Con comas: %s\n", concatenar(", ", "manzana", "banana", "naranja"))
	fmt.Printf("Con guiones: %s\n", concatenar(" - ", "Uno", "Dos", "Tres"))
	fmt.Printf("Sin elementos: '%s'\n", concatenar(", "))

	// 6. Validación de campos
	fmt.Println("\n6️⃣ Validación de campos:")
	campos1 := []string{"Juan", "juan@email.com", "123456789"}
	campos2 := []string{"María", "", "987654321"}
	
	valido1, errores1 := validarCampos(true, campos1...)
	valido2, errores2 := validarCampos(true, campos2...)
	
	fmt.Printf("Campos 1 válidos: %t, errores: %v\n", valido1, errores1)
	fmt.Printf("Campos 2 válidos: %t, errores: %v\n", valido2, errores2)

	// 7. Constructor de SQL
	fmt.Println("\n7️⃣ Constructor de consultas SQL:")
	fmt.Println(construirSQL("usuarios", "SELECT"))
	fmt.Println(construirSQL("usuarios", "SELECT", "edad > 18", "activo = true"))
	fmt.Println(construirSQL("productos", "DELETE", "stock = 0"))

	// 8. Operaciones matemáticas
	fmt.Println("\n8️⃣ Operaciones matemáticas:")
	valores := []float64{5, 10, 15, 20, 25}
	
	operaciones := []string{"suma", "multiplicacion", "maximo", "minimo"}
	for _, op := range operaciones {
		resultado, err := operar(op, valores...)
		if err != nil {
			fmt.Printf("%s: ❌ %v\n", op, err)
		} else {
			fmt.Printf("%s de %v = %.2f\n", op, valores, resultado)
		}
	}

	// 9. Pipeline de procesamiento
	fmt.Println("\n9️⃣ Pipeline de procesamiento:")
	datosOriginales := []int{1, 2, 3, 4, 5}
	
	// Definir funciones de transformación
	duplicar := func(x int) int { return x * 2 }
	sumarUno := func(x int) int { return x + 1 }
	elevarCuadrado := func(x int) int { return x * x }
	
	fmt.Printf("Datos originales: %v\n", datosOriginales)
	
	resultado1 := procesarDatos(datosOriginales, duplicar)
	fmt.Printf("Después de duplicar: %v\n", resultado1)
	
	resultado2 := procesarDatos(datosOriginales, duplicar, sumarUno)
	fmt.Printf("Duplicar → +1: %v\n", resultado2)
	
	resultado3 := procesarDatos(datosOriginales, duplicar, sumarUno, elevarCuadrado)
	fmt.Printf("Duplicar → +1 → ²: %v\n", resultado3)

	// 10. Generador de reportes
	fmt.Println("\n🔟 Generador de reportes:")
	
	seccionUsuarios := func() string {
		return "Usuarios activos: 1,234\nUsuarios nuevos hoy: 56"
	}
	
	seccionVentas := func() string {
		return fmt.Sprintf("Ventas del día: $%.2f\nTransacciones: %d", 15789.50, 89)
	}
	
	seccionEstadisticas := func() string {
		return "Tiempo promedio de sesión: 12 min\nPáginas más visitadas: /home, /productos"
	}
	
	reporte := generarReporte("Reporte Diario", seccionUsuarios, seccionVentas, seccionEstadisticas)
	fmt.Println(reporte)

	fmt.Println("\n🎯 CONCEPTOS DEMOSTRADOS:")
	fmt.Println("✅ Sintaxis de funciones variádicas (...tipo)")
	fmt.Println("✅ Combinación de parámetros fijos y variádicos")
	fmt.Println("✅ Uso del operador ... para expandir slices")
	fmt.Println("✅ Funciones variádicas con diferentes tipos")
	fmt.Println("✅ Aplicaciones prácticas: logging, validación, SQL")
	fmt.Println("✅ Procesamiento de datos con pipelines")
	fmt.Println("✅ Generación dinámica de contenido")
}