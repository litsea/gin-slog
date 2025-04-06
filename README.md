# gin-slog

# Usage

```golang
import (
	"context"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	ginslog "github.com/litsea/gin-slog"
	sloggin "github.com/samber/slog-gin"
)

r := gin.New()

l := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}),
gsl := ginslog.New(
	l,
	ginslog.WithRequestID(true),
	ginslog.WithRequestHeader(true),
	ginslog.WithFilters(
		sloggin.IgnorePath("/v1/health"),
	),
	ginslog.WithExtraAttrs(map[string]string{ ... }),
)

r.Use(gsl)
```

* Default options: [New()](slog.go)
* Options handle functions: [Option](option.go)

### Runtime Add Custom Attributes

```golang
import (
	"context"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/litsea/gin-api/log"
	ginslog "github.com/litsea/gin-slog"
	sloggin "github.com/samber/slog-gin"
)

r := gin.New()

l := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}),
gl := log.New(l, ginslog.AddCustomAttributes)

gsl := ginslog.New(
	l,
	ginslog.WithRequestID(true),
	ginslog.WithRequestHeader(true),
	ginslog.WithFilters(
		sloggin.IgnorePath("/v1/health"),
	),
	ginslog.WithExtraAttrs(map[string]string{ ... }),
)

r.Use(log.Middleware(gl), gsl)
```

> * See:
>   * [log](https://github.com/litsea/log)
>   * [gin-api](https://github.com/litsea/gin-api)
