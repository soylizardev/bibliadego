package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	var sb strings.Builder
	sb.WriteString("# Summary\n\n[Introducción](README.md)\n\n")

	// 1. Obtener y ordenar los módulos (Carpetas que empiezan con "Modulo")
	modules, _ := filepath.Glob("Modulo*")
	sort.Strings(modules)

	for _, mod := range modules {
		modName := strings.ReplaceAll(mod, "_", " ")
		sb.WriteString(fmt.Sprintf("- [%s]()\n", modName))

		// 2. Obtener y ordenar subcarpetas (ej: 01_Fundamentos)
		subDirs, _ := os.ReadDir(mod)
		for _, sub := range subDirs {
			if sub.IsDir() {
				subPath := filepath.Join(mod, sub.Name())
				cleanSub := sub.Name()
				if idx := strings.Index(cleanSub, "_"); idx != -1 {
					cleanSub = cleanSub[idx+1:]
				}
				sb.WriteString(fmt.Sprintf("    - [%s]()\n", cleanSub))

				// 3. Obtener y ordenar archivos .md
				files, _ := os.ReadDir(subPath)
				for _, file := range files {
					if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") && !strings.HasPrefix(file.Name(), "00_") {
						filePath := filepath.Join(subPath, file.Name())

						// Limpiar nombre para el menú
						fileName := strings.TrimSuffix(file.Name(), ".md")
						if idx := strings.Index(fileName, "_"); idx != -1 {
							fileName = fileName[idx+1:]
						}

						// Codificar ruta para URL (reemplazar espacios)
						webPath := strings.ReplaceAll(filePath, " ", "%20")
						webPath = strings.ReplaceAll(webPath, "\\", "/")

						sb.WriteString(fmt.Sprintf("        - [%s](%s)\n", fileName, webPath))
					}
				}
			}
		}
	}

	// 4. Agregar Roadmap si existe
	if _, err := os.Stat("Roadmap Go Backend High Performance.md"); err == nil {
		path := strings.ReplaceAll("Roadmap Go Backend High Performance.md", " ", "%20")
		sb.WriteString(fmt.Sprintf("\n- [Roadmap](%s)\n", path))
	}

	// Escribir el archivo
	err := os.WriteFile("SUMMARY.md", []byte(sb.String()), 0644)
	if err != nil {
		fmt.Println("Error escribiendo SUMMARY.md:", err)
		return
	}
	fmt.Println("SUMMARY.md generado exitosamente con Go.")
}
