package initializer

import (
	"fmt"
	"reflect"
	"runtime"
)

//----------------------------------------------------------------------------------------------------------------------------//

type (
	// Инициализатор модуля
	ModuleInitializer func(cfg any, h any) error
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
func Do(cfg any, h any) (err error) {
	for _, f := range initializers {
		err = f(cfg, h)
		if err != nil {
			err = fmt.Errorf("[%s] %s", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), err)
			return
		}
	}

	return
}

//----------------------------------------------------------------------------------------------------------------------------//
