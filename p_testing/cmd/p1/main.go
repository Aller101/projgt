package main

import (
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	// profCPU()
	profMEM()
}

func profCPU() {
	// Создание файла для записи профиля
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("Не удалось создать файл профиля:", err)
	}
	defer f.Close()

	// Запуск профилирования CPU
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("Не удалось запустить профилирование:", err)
	}
	defer pprof.StopCPUProfile() // Остановка профилирования

	// Симуляция нагрузки
	for i := 0; i < 1_000_000_000; i++ {
		_ = i * i
	}
}

func profMEM() {
	// Создание файла для записи профиля памяти
	f, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal("Не удалось создать файл профиля:", err)
	}
	defer f.Close()

	// Симуляция выделения памяти
	slice := make([]int, 0)
	for i := 0; i < 1_000_000; i++ {
		slice = append(slice, i)
	}

	// Запись текущего состояния памяти
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("Не удалось записать профиль памяти:", err)
	}
}
