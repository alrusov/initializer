package initializer

import (
	"github.com/alrusov/stdhttp"
)

//----------------------------------------------------------------------------------------------------------------------------//

type (
	// Инициализатор модуля
	ModuleInitializer func(cfg interface{}, h *stdhttp.HTTP) error
)

var (
	// Инициализаторы модулей
	initializers = []ModuleInitializer{}
)

//----------------------------------------------------------------------------------------------------------------------------//

// Добавить инициализатор модуля
func RegisterModuleInitializer(f ModuleInitializer) {
	initializers = append(initializers, f)
}

//----------------------------------------------------------------------------------------------------------------------------//

// Инициализируем модули
func Do(cfg interface{}, h *stdhttp.HTTP) (err error) {
	for _, f := range initializers {
		err = f(cfg, h)
		if err != nil {
			return
		}
	}

	return
}

//----------------------------------------------------------------------------------------------------------------------------//
