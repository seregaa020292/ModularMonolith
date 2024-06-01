package closer

var globalCloser = New()

// Add adds `func() error` callback to the globalCloser
func Add(f ...func() error) {
	globalCloser.Add(f...)
}

// Wait waiting close
func Wait() {
	globalCloser.Wait()
}

// CloseAll close all functions
func CloseAll() {
	globalCloser.CloseAll()
}
