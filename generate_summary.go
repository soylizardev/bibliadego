package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	var sbSummary strings.Builder
	var sbReadme strings.Builder

	// Configuración inicial
	sbSummary.WriteString("# Summary\n\n[Introducción](README.md)\n\n")
	sbReadme.WriteString("# 🐹 Biblia de Go (Golang)\n")
	sbReadme.WriteString("> Un recurso de alto rendimiento. *Actualizado automáticamente por el Robot Gopher.*\n\n")
	sbReadme.WriteString("## 📚 Índice del Contenido\n\n")

	modules, _ := filepath.Glob("Modulo*")
	sort.Strings(modules)

	for _, mod := range modules {
		modClean := strings.ReplaceAll(mod, "_", " ")

		// SUMMARY: Lista simple sin ###
		sbSummary.WriteString(fmt.Sprintf("- [%s]()\n", modClean))
		// README: Aquí sí usamos ### para que se vea bien en GitHub
		sbReadme.WriteString(fmt.Sprintf("### 📂 %s\n", modClean))

		subDirs, _ := os.ReadDir(mod)
		for _, sub := range subDirs {
			if sub.IsDir() {
				subPath := filepath.Join(mod, sub.Name())
				subClean := sub.Name()
				if idx := strings.Index(subClean, "_"); idx != -1 {
					subClean = subClean[idx+1:]
				}

				// SUMMARY: Sangría de 4 espacios
				sbSummary.WriteString(fmt.Sprintf("    - [%s]()\n", subClean))
				// README: Negrita para subcategorías
				sbReadme.WriteString(fmt.Sprintf("* **%s**\n", subClean))

				files, _ := os.ReadDir(subPath)
				for _, file := range files {
					if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") && !strings.HasPrefix(file.Name(), "00_") {
						filePath := filepath.Join(subPath, file.Name())
						fileName := strings.TrimSuffix(file.Name(), ".md")
						if idx := strings.Index(fileName, "_"); idx != -1 {
							fileName = fileName[idx+1:]
						}

						webPath := strings.ReplaceAll(filePath, " ", "%20")
						webPath = strings.ReplaceAll(webPath, "\\", "/")

						// SUMMARY: Sangría de 8 espacios y enlace
						sbSummary.WriteString(fmt.Sprintf("        - [%s](%s)\n", fileName, webPath))
						// README: Lista con enlace
						sbReadme.WriteString(fmt.Sprintf("  - [%s](%s)\n", fileName, webPath))
					}
				}
				sbReadme.WriteString("\n")
			}
		}
	}

	// Roadmap
	if _, err := os.Stat("Roadmap Go Backend High Performance.md"); err == nil {
		path := strings.ReplaceAll("Roadmap Go Backend High Performance.md", " ", "%20")
		sbSummary.WriteString(fmt.Sprintf("\n- [Roadmap](%s)\n", path))
		sbReadme.WriteString(fmt.Sprintf("### 🚀 [Roadmap](%s)\n", path))
	}

	sbReadme.WriteString("\n---\n## 🌐 Ver como Libro Web\n👉 **[https://soylizardev.github.io/bibliadego/](https://soylizardev.github.io/bibliadego/)**\n")

	os.WriteFile("SUMMARY.md", []byte(sbSummary.String()), 0644)
	os.WriteFile("README.md", []byte(sbReadme.String()), 0644)
	fmt.Println("::notice title=Generador::Archivos actualizados con formato compatible.")
}
