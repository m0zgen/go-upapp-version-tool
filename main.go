package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {

	// Версия приложения
	vv := "v1.1.1"

	// Зелёный цвет и жирный шрифт
	greenBold := "\033[1;32m"
	// Сброс форматирования
	reset := "\033[0m"

	// Вывод сплэш-экрана с версией
	fmt.Println("──────────────────────────────")
	fmt.Printf(" %s%s%-1s%s %s      \n", greenBold, "Update Releaser Tool", "", reset, vv)
	fmt.Println("──────────────────────────────")

	versionFile := "version.txt"

	data, err := ioutil.ReadFile(versionFile)
	if err != nil {
		log.Fatalf("Ошибка при чтении версии: %v", err)
	}

	version := string(data)
	fmt.Printf("Текущая версия: %s\n", version)

	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		log.Fatalf("Неверный формат версии: %s", version)
	}

	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		log.Fatalf("Ошибка при преобразовании патча версии: %v", err)
	}
	patch++
	parts[2] = strconv.Itoa(patch)

	newVersion := strings.Join(parts, ".")
	err = ioutil.WriteFile(versionFile, []byte(newVersion), 0644)
	if err != nil {
		log.Fatalf("Ошибка при записи новой версии: %v", err)
	}

	fmt.Printf("Новая версия: %s\n", newVersion)
}
