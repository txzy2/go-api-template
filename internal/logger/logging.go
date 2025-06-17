package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Logger struct {
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
}

// Сокращенные методы
func (l *Logger) Debug(msg string) {
	l.debug.Println(msg)
}

func (l *Logger) Info(msg string) {
	l.info.Println(msg)
}

func (l *Logger) Warn(msg string) {
	l.warn.Println(msg)
}

func (l *Logger) Error(msg string) {
	l.error.Println(msg)
}

var AppLogger Logger

func Init() {
	// Создаем структуру директорий для логов
	now := time.Now()
	logPath := filepath.Join("logs",
		fmt.Sprintf("%d", now.Year()),
		fmt.Sprintf("%02d", now.Month()),
		fmt.Sprintf("%02d", now.Day()))

	// Создаем директории если их нет
	if err := os.MkdirAll(logPath, 0755); err != nil {
		log.Fatal("Error creating log directory:", err)
	}

	// Создаем файл лога
	logFile, err := os.OpenFile(
		filepath.Join(logPath, "app.log"),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}

	// Инициализируем логгеры
	AppLogger = Logger{
		debug: log.New(logFile, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		info:  log.New(logFile, "INFO: ", log.Ldate|log.Ltime),
		warn:  log.New(logFile, "WARN: ", log.Ldate|log.Ltime),
		error: log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
