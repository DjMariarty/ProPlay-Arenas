package validation

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// TimeFormat - формат времени HH:MM
const TimeFormat = "HH:MM"

// ValidateTime проверяет и парсит время в формате HH:MM
func ValidateTime(timeStr string) (time.Time, error) {
	// Проверяем формат через regex
	matched, err := regexp.MatchString(`^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$`, timeStr)
	if err != nil || !matched {
		return time.Time{}, fmt.Errorf("неверный формат времени (ожидается HH:MM): %s", timeStr)
	}

	// Парсим строку вручную
	parts := regexp.MustCompile(`:`).Split(timeStr, -1)
	if len(parts) != 2 {
		return time.Time{}, fmt.Errorf("неверный формат времени: %s", timeStr)
	}

	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return time.Time{}, fmt.Errorf("неверный формат часов: %s", parts[0])
	}

	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("неверный формат минут: %s", parts[1])
	}

	// Валидация диапазонов
	if hours < 0 || hours > 23 {
		return time.Time{}, fmt.Errorf("часы должны быть в диапазоне 00-23: %d", hours)
	}
	if minutes < 0 || minutes > 59 {
		return time.Time{}, fmt.Errorf("минуты должны быть в диапазоне 00-59: %d", minutes)
	}

	// Создаем time.Time с текущей датой
	now := time.Now()
	t := time.Date(now.Year(), now.Month(), now.Day(), hours, minutes, 0, 0, time.UTC)

	return t, nil
}

// CompareTimes сравнивает два времени в формате HH:MM
// Возвращает true если startTime < endTime
func CompareTimes(startTime, endTime string) (bool, error) {
	start, err := ValidateTime(startTime)
	if err != nil {
		return false, fmt.Errorf("время начала: %w", err)
	}

	end, err := ValidateTime(endTime)
	if err != nil {
		return false, fmt.Errorf("время окончания: %w", err)
	}

	return start.Before(end), nil
}
