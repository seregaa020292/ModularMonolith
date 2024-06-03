package closer

var globalCloser = New()

// Add добавляет обратный вызов `func() error` в globalCloser
func Add(f ...func() error) {
	globalCloser.Add(f...)
}

// Wait ожидает завершения работы
func Wait() {
	globalCloser.Wait()
}

// CloseAll закрывает все функции
func CloseAll() {
	globalCloser.CloseAll()
}
