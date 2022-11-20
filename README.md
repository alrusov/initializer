## Пример типового использования

```
package module1

import (
	"myapp/config"

	"github.com/alrusov/initializer"
	"github.com/alrusov/log"
	"github.com/alrusov/stdhttp"
)

func init() {
	initializer.RegisterModuleInitializer(initModule)
}

func initModule(cfg interface{}, h *stdhttp.HTTP) (err error) {
	appCfg := cfg.(*config.Config)
	// do something

	log.Message(log.INFO, "module1 initialized")
	return
}
```

```
package http

import (
	"app/config"

	"github.com/alrusov/initializer"
	"github.com/alrusov/stdhttp"
)

type HTTP struct {
	cfg *config.Config
	h   *stdhttp.HTTP
}

func NewHTTP(cfg *config.Config) (hh *stdhttp.HTTP, err error) {
	h := &HTTP{
		cfg: cfg,
	}

	h.h, err = stdhttp.NewListener(&cfg.HTTP.Listener, h)
	if err != nil {
		return nil, err
	}

	// ...
	// ...
	// ...

	err = initializer.Do(cfg, h.h)
	if err != nil {
		return nil, err
	}

	return h.h, nil
}
```
