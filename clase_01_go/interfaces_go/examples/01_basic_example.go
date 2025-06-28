package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// ===========================================
// FUNCIONES BÁSICAS EN GO
// ===========================================

// 1. Función simple sin parámetros ni retorno
func saludar() {
	fmt.Println("¡Hola desde Go!")
}

// 2. Función con parámetros
func saludarPersona(nombre string) {
	fmt.Printf("¡Hola, %s!\n", nombre)
}

// 3. Función con múltiples parámetros del mismo tipo
func sumar(a, b int) int {
	return a + b
}

// 4. Función con múltiples parámetros de diferentes tipos
func crearMensaje(nombre string, edad int, activo bool) string {
	estado := "inactivo"
	if activo {
		estado = "activo"
	}
	return fmt.Sprintf("Usuario: %s, Edad: %d, Estado: %s", nombre, edad, estado)
}

// 5. Función con múltiples valores de retorno
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("división por cero no permitida")
	}
	return a / b, nil
}

// 6. Función con valores de retorno nombrados (named return values)
func analizarTexto(texto string) (palabras int, caracteres int, lineas int) {
	caracteres = len(texto)
	lineas = strings.Count(texto, "\n") + 1
	palabras = len(strings.Fields(texto))
	return // return implícito cuando los valores están nombrados
}

// 7. Función que puede fallar con manejo de errores
func procesarEdad(edad int) (string, error) {
	if edad < 0 {
		return "", errors.New("la edad no puede ser negativa")
	}
	if edad < 18 {
		return "menor de edad", nil
	}
	if edad < 65 {
		return "adulto", nil
	}
	return "adulto mayor", nil
}

// 8. Función con lógica de negocio más compleja
func calcularDescuento(precio float64, categoria string, esMiembro bool) (float64, string, error) {
	if precio <= 0 {
		return 0, "", errors.New("el precio debe ser mayor a cero")
	}

	var descuento float64
	var razon string

	// Aplicar descuento por categoría
	switch strings.ToLower(categoria) {
	case "premium":
		descuento = 0.20
		razon = "Descuento Premium 20%"
	case "regular":
		descuento = 0.10
		razon = "Descuento Regular 10%"
	case "basico":
		descuento = 0.05
		razon = "Descuento Básico 5%"
	default:
		descuento = 0
		razon = "Sin descuento"
	}

	// Descuento adicional para miembros
	if esMiembro {
		descuento += 0.05
		razon += " + Miembro 5%"
	}

	// Calcular precio final
	precioFinal := precio * (1 - descuento)
	return precioFinal, razon, nil
}

// 9. Función con validaciones múltiples
func validarUsuario(nombre, email string, edad int) error {
	if strings.TrimSpace(nombre) == "" {
		return errors.New("el nombre no puede estar vacío")
	}

	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return errors.New("email inválido")
	}

	if edad < 0 || edad > 120 {
		return errors.New("edad debe estar entre 0 y 120 años")
	}

	return nil
}

// 10. Función que demuestra el uso de defer
func operacionConRecursos(archivo string) error {
	fmt.Printf("📂 Abriendo archivo: %s\n", archivo)
	
	// Simular apertura de archivo
	defer func() {
		fmt.Printf("🔒 Cerrando archivo: %s\n", archivo)
	}()

	// Simular operaciones
	fmt.Printf("📝 Procesando contenido de: %s\n", archivo)
	time.Sleep(100 * time.Millisecond)
	
	// Simular un posible error
	if archivo == "error.txt" {
		return errors.New("error al procesar el archivo")
	}

	fmt.Printf("✅ Archivo procesado exitosamente: %s\n", archivo)
	return nil
}

// ===========================================
// FUNCIÓN PRINCIPAL PARA DEMOSTRAR EJEMPLOS
// ===========================================

func main() {
	fmt.Println("🚀 EJEMPLOS DE FUNCIONES BÁSICAS EN GO")
	fmt.Println("=" + strings.Repeat("=", 45))

	// 1. Función simple
	fmt.Println("\n1️⃣ Función simple:")
	saludar()

	// 2. Función con parámetros
	fmt.Println("\n2️⃣ Función con parámetros:")
	saludarPersona("Ana")
	saludarPersona("Carlos")

	// 3. Función con retorno
	fmt.Println("\n3️⃣ Función con retorno:")
	resultado := sumar(15, 25)
	fmt.Printf("15 + 25 = %d\n", resultado)

	// 4. Función con múltiples parámetros
	fmt.Println("\n4️⃣ Función con múltiples parámetros:")
	mensaje := crearMensaje("María", 28, true)
	fmt.Println(mensaje)

	// 5. Función con múltiples retornos y manejo de errores
	fmt.Println("\n5️⃣ Función con múltiples retornos:")
	division, err := dividir(10, 3)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Printf("10 / 3 = %.2f\n", division)
	}

	// Ejemplo de error
	_, err = dividir(10, 0)
	if err != nil {
		fmt.Printf("❌ Error esperado: %v\n", err)
	}

	// 6. Named return values
	fmt.Println("\n6️⃣ Named return values:")
	texto := "Hola mundo\nEste es un ejemplo\nDe múltiples líneas"
	p, c, l := analizarTexto(texto)
	fmt.Printf("Texto analizado: %d palabras, %d caracteres, %d líneas\n", p, c, l)

	// 7. Función con validación
	fmt.Println("\n7️⃣ Función con validación:")
	edades := []int{-5, 10, 25, 70}
	for _, edad := range edades {
		categoria, err := procesarEdad(edad)
		if err != nil {
			fmt.Printf("Edad %d: ❌ %v\n", edad, err)
		} else {
			fmt.Printf("Edad %d: %s\n", edad, categoria)
		}
	}

	// 8. Función con lógica de negocio compleja
	fmt.Println("\n8️⃣ Cálculo de descuentos:")
	productos := []struct {
		precio    float64
		categoria string
		miembro   bool
	}{
		{100.0, "premium", true},
		{50.0, "regular", false},
		{25.0, "basico", true},
		{200.0, "vip", false},
	}

	for i, prod := range productos {
		precioFinal, razon, err := calcularDescuento(prod.precio, prod.categoria, prod.miembro)
		if err != nil {
			fmt.Printf("Producto %d: ❌ %v\n", i+1, err)
		} else {
			fmt.Printf("Producto %d: $%.2f → $%.2f (%s)\n", i+1, prod.precio, precioFinal, razon)
		}
	}

	// 9. Validación de usuarios
	fmt.Println("\n9️⃣ Validación de usuarios:")
	usuarios := []struct {
		nombre string
		email  string
		edad   int
	}{
		{"Juan Pérez", "juan@email.com", 30},
		{"", "invalido@", 25},
		{"María García", "maria@empresa.com", -5},
		{"Pedro López", "pedro@test.com.ar", 45},
	}

	for i, usuario := range usuarios {
		err := validarUsuario(usuario.nombre, usuario.email, usuario.edad)
		if err != nil {
			fmt.Printf("Usuario %d: ❌ %v\n", i+1, err)
		} else {
			fmt.Printf("Usuario %d: ✅ Válido - %s\n", i+1, usuario.nombre)
		}
	}

	// 10. Función con defer
	fmt.Println("\n🔟 Función con defer (manejo de recursos):")
	archivos := []string{"datos.txt", "config.json", "error.txt", "backup.sql"}
	
	for _, archivo := range archivos {
		fmt.Printf("\n📋 Procesando: %s\n", archivo)
		err := operacionConRecursos(archivo)
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
		}
	}

	fmt.Println("\n🎯 CONCEPTOS DEMOSTRADOS:")
	fmt.Println("✅ Declaración de funciones con diferentes signaturas")
	fmt.Println("✅ Parámetros múltiples del mismo y diferentes tipos")
	fmt.Println("✅ Valores de retorno simples y múltiples")
	fmt.Println("✅ Named return values para claridad")
	fmt.Println("✅ Manejo apropiado de errores")
	fmt.Println("✅ Validaciones y lógica de negocio")
	fmt.Println("✅ Uso de defer para limpieza de recursos")
}