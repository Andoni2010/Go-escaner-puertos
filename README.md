# Documentación del Código: Escáner de Puertos en Go

Este documento explica paso a paso lo que realiza el código presentado, resaltando su funcionalidad principal y los puntos clave en su ejecución. El programa es un escáner de puertos TCP sencillo que permite analizar un rango específico de puertos en una dirección IP o dominio y guardar los resultados en un archivo de texto.

---

## 1. Objetivo Principal
El programa realiza un escaneo de puertos TCP en un rango definido por el usuario para un host determinado (IP o dominio). Informa cuáles puertos están abiertos o cerrados y guarda los resultados en un archivo llamado `resultados.txt`.

---

## 2. Explicación General del Código

### Importación de Librerías
El código importa las siguientes librerías estándar de Go:

- `fmt`: Para imprimir y escanear datos en la consola.
- `net`: Para manejar conexiones de red.
- `os`: Para manejar la creación y escritura de archivos.
- `time`: Para controlar los tiempos de espera en las conexiones.

```
package main
import (
	"fmt"
	"net"
	"os"
	"time"
)
```

---

### Definición de la Función `scanPort`
Esta función se encarga de analizar un puerto específico en el host objetivo.

- **Parámetros:**
  - `host`: IP o dominio del host a escanear.
  - `port`: Número del puerto a analizar.
  - `results`: Archivo donde se escribirán los resultados del escaneo.

- **Funcionamiento:**
  - Crea una dirección en el formato `host:puerto`.
  - Intenta establecer una conexión TCP con el puerto utilizando `net.DialTimeout`.
  - Si la conexión falla (error), el puerto se considera **cerrado**.
  - Si la conexión es exitosa, el puerto se considera **abierto**.
  - Los resultados del escaneo (abierto/cerrado) se escriben en el archivo `resultados.txt`.

```
func scanPort(host string, port int, results *os.File) {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		results.WriteString(fmt.Sprintf("[-] Puerto %d cerrado\n", port))
		return
	}
	conn.Close()
	results.WriteString(fmt.Sprintf("[+] Puerto %d abierto\n", port))
}
```
---

### Función `main`
Este es el punto de entrada del programa y gestiona la interacción con el usuario y el flujo general.

1. **Solicitar datos al usuario:**
   - Pide al usuario el **host** (IP o dominio objetivo).
   - Solicita el rango de puertos (`startPort` y `endPort`) que se desea analizar.

2. **Crear el archivo `resultados.txt`:**
   - Se utiliza la función `os.Create` para crear el archivo donde se almacenarán los resultados. Si hay un error al crear el archivo, el programa se detiene.
   - El archivo se cierra automáticamente al final del programa usando `defer`.

3. **Escaneo de puertos:**
   - Itera a través del rango de puertos definido por el usuario.
   - Llama a la función `scanPort` para cada puerto, escribiendo los resultados en el archivo.

4. **Finalización del programa:**
   - Al completar el escaneo, informa al usuario que los resultados se han guardado.

```
func main() {
	var host string
	var startPort, endPort int
	fmt.Println("Ingresa la IP o dominio objetivo (e.g, 192.168.1.1): ")
	fmt.Scan(&host)
	fmt.Println("Ingresa el puerto inicial: ")
	fmt.Scan(&startPort)
	fmt.Println("Ingresa lel puerto final: ")
	fmt.Scan(&endPort)
	fmt.Printf("Escaneando %s desde el puerto %d al %d...\n ", host, startPort, endPort)
	results, err := os.Create("resultados.txt")
	if err != nil {
		fmt.Println("Error al crear archivo: ", err)
		return
	}
	defer results.Close()
	for port := startPort; port <= endPort; port++ {
		scanPort(host, port, results)
	}
	fmt.Println("Escaneo completado. Resultados guardados en 'resultados.txt'.")
}
```

---

## 3. Detalles Clave del Código

### Conexión TCP con Timeout
El uso de `net.DialTimeout` permite intentar establecer una conexión con un tiempo límite definido de 1 segundo. Esto asegura que el programa no se quede bloqueado indefinidamente en caso de que un puerto no responda.

---

### Gestión de Archivos con `os.Create`
El archivo `resultados.txt` se crea al inicio del programa.

- Se utiliza `defer results.Close()` para garantizar que el archivo se cierre correctamente una vez finalizado el programa o en caso de errores.

---

### Uso de `fmt.Sprintf` para Formatear Salidas
En múltiples partes del código, como en la función `scanPort`, se usa `fmt.Sprintf` para generar cadenas formateadas con los valores de los puertos y el estado (abierto/cerrado).

---

### Bucle para Escaneo de Puertos
El programa utiliza un bucle `for` que itera desde el puerto inicial (`startPort`) hasta el puerto final (`endPort`) definido por el usuario. En cada iteración:

- Llama a la función `scanPort`.
- Escribe los resultados en el archivo.

```
for port := startPort; port <= endPort; port++ {
		scanPort(host, port, results)
	}
	fmt.Println("Escaneo completado. Resultados guardados en 'resultados.txt'.")
```

---

## 4. Funcionamiento del Programa

1. **Inicio:**
   - El usuario ejecuta el programa y proporciona los datos necesarios (IP/host y rango de puertos).

2. **Escaneo:**
   - El programa analiza cada puerto dentro del rango especificado y verifica si está abierto o cerrado.

3. **Resultados:**
   - Los resultados se imprimen en consola y se guardan en el archivo `resultados.txt`.

---

## 5. Mejoras Potenciales

1. **Validación de Entrada del Usuario:**
   - Agregar validaciones para asegurar que los datos ingresados por el usuario sean válidos (e.g., verificar que los puertos estén en el rango 1-65535).

2. **Paralelismo:**
   - Implementar Goroutines para que los puertos sean escaneados en paralelo, lo que acelerará el proceso.

3. **Interfaz de Usuario:**
   - Crear una interfaz gráfica o una página web simple para interactuar con el escáner de puertos.

4. **Soporte para Protocolo UDP:**
   - Extender el programa para analizar puertos UDP además de TCP.

---

## 6. Conclusión
El programa es un escáner de puertos funcional y sencillo que utiliza técnicas básicas de Go. Es ideal como proyecto introductorio para entender:

- Conexiones TCP.
- Manejo de archivos.
- Interacción con el usuario.
- Implementación de loops y manejo de errores.
