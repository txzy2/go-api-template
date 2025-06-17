package status

// ErrorMessages map с сообщениями об ошибках
var ErrorMessages = map[int]string{
	400: "Invalid data provided",
	401: "Authentication required",
	403: "Access forbidden",
	404: "Not found",
	500: "Internal server error",
}

// GetErrorMessage получает сообщение по коду ошибки
func GetErrorMessage(code int) string {
	if message, ok := ErrorMessages[code]; ok {
		return message
	}
	return "Произошла ошибка"
}
