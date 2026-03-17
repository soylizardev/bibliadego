package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	// --- Lógica para el SUMMARY.md ---
	var sbSummary strings.Builder
	sbSummary.WriteString("# Summary\n\n[Introducción](README.md)\n\n")

	// --- Lógica para el README.md ---
	var sbReadme strings.Builder
	sbReadme.WriteString("# 🐹 Biblia de Go (Golang)\n")
	sbReadme.WriteString("> Un recurso exhaustivo para dominar el lenguaje Go, desde los fundamentos hasta el alto rendimiento.\n\n")
	sbReadme.WriteString("Este repositorio se actualiza automáticamente. ¡No edites el índice manualmente!\n\n")
	sbReadme.WriteString("## 📚 Índice del Libro\n\n")

	modules, _ := filepath.Glob("Modulo*")
	sort.Strings(modules)

	for _, mod := range modules {
		modName := strings.ReplaceAll(mod, "_", " ")
		line := fmt.Sprintf("### 📂 [%s]()\n", modName)
		sbSummary.WriteString("- " + line)
		sbReadme.WriteString(line)

		subDirs, _ := os.ReadDir(mod)
		for _, sub := range subDirs {
			if sub.IsDir() {
				subPath := filepath.Join(mod, sub.Name())
				cleanSub := sub.Name()
				if idx := strings.Index(cleanSub, "_"); idx != -1 {
					cleanSub = cleanSub[idx+1:]
				}

				sbSummary.WriteString(fmt.Sprintf("    - [%s]()\n", cleanSub))
				sbReadme.WriteString(fmt.Sprintf("* **%s**\n", cleanSub))

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

						sbSummary.WriteString(fmt.Sprintf("        - [%s](%s)\n", fileName, webPath))
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

	// Footer del README
	sbReadme.WriteString("\n---\n## 🌐 Ver como Libro Web\n")
	sbReadme.WriteString("👉 **[https://soylizardev.github.io/bibliadego/](https://soylizardev.github.io/bibliadego/)**\n")

	// Escribir ambos archivos
	os.WriteFile("SUMMARY.md", []byte(sbSummary.String()), 0644)
	os.WriteFile("README.md", []byte(sbReadme.String()), 0644)

	fmt.Println("::notice title=Generador::README.md y SUMMARY.md actualizados 🐹")
}
