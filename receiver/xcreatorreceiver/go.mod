module go.opentelemetry.io/collector/receiver/xcreatorreceiver

go 1.22.0

require (
	github.com/stretchr/testify v1.10.0
	go.opentelemetry.io/collector/component v0.116.0
	go.opentelemetry.io/collector/component/componenttest v0.116.0
	go.opentelemetry.io/collector/confmap v1.22.0
	go.opentelemetry.io/collector/consumer v1.22.0
	go.opentelemetry.io/collector/consumer/consumertest v0.116.0
	go.opentelemetry.io/collector/pipeline v0.116.0
	go.opentelemetry.io/collector/receiver v0.116.0
	go.opentelemetry.io/collector/receiver/otlpreceiver v0.116.0
	go.opentelemetry.io/collector/receiver/receivertest v0.116.0
	go.uber.org/goleak v1.3.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.2 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mostynb/go-grpc-compression v1.2.3 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rs/cors v1.11.1 // indirect
	go.opentelemetry.io/collector v0.116.0 // indirect
	go.opentelemetry.io/collector/client v1.22.0 // indirect
	go.opentelemetry.io/collector/component/componentstatus v0.116.0 // indirect
	go.opentelemetry.io/collector/config/configauth v0.116.0 // indirect
	go.opentelemetry.io/collector/config/configcompression v1.22.0 // indirect
	go.opentelemetry.io/collector/config/configgrpc v0.116.0 // indirect
	go.opentelemetry.io/collector/config/confighttp v0.116.0 // indirect
	go.opentelemetry.io/collector/config/confignet v1.22.0 // indirect
	go.opentelemetry.io/collector/config/configopaque v1.22.0 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.116.0 // indirect
	go.opentelemetry.io/collector/config/configtls v1.22.0 // indirect
	go.opentelemetry.io/collector/config/internal v0.116.0 // indirect
	go.opentelemetry.io/collector/consumer/consumererror v0.116.0 // indirect
	go.opentelemetry.io/collector/consumer/xconsumer v0.116.0 // indirect
	go.opentelemetry.io/collector/extension v0.116.0 // indirect
	go.opentelemetry.io/collector/extension/auth v0.116.0 // indirect
	go.opentelemetry.io/collector/internal/sharedcomponent v0.116.0 // indirect
	go.opentelemetry.io/collector/pdata v1.22.0 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.116.0 // indirect
	go.opentelemetry.io/collector/receiver/xreceiver v0.116.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.56.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.56.0 // indirect
	go.opentelemetry.io/otel v1.32.0 // indirect
	go.opentelemetry.io/otel/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/sdk v1.32.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/trace v1.32.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/net v0.32.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241007155032-5fefd90f89a9 // indirect
	google.golang.org/grpc v1.68.1 // indirect
	google.golang.org/protobuf v1.35.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace go.opentelemetry.io/collector/component => ../../component

replace go.opentelemetry.io/collector/component/componenttest => ../../component/componenttest

replace go.opentelemetry.io/collector/confmap => ../../confmap

replace go.opentelemetry.io/collector/consumer => ../../consumer

replace go.opentelemetry.io/collector/receiver => ../

replace go.opentelemetry.io/collector/pdata => ../../pdata

replace go.opentelemetry.io/collector/pdata/testdata => ../../pdata/testdata

replace go.opentelemetry.io/collector/config/configtelemetry => ../../config/configtelemetry

replace go.opentelemetry.io/collector/pdata/pprofile => ../../pdata/pprofile

replace go.opentelemetry.io/collector/consumer/xconsumer => ../../consumer/xconsumer

replace go.opentelemetry.io/collector/consumer/consumertest => ../../consumer/consumertest

replace go.opentelemetry.io/collector/receiver/xreceiver => ../xreceiver

replace go.opentelemetry.io/collector/receiver/otlpreceiver => ../otlpreceiver

replace go.opentelemetry.io/collector/receiver/receivertest => ../receivertest

replace go.opentelemetry.io/collector/pipeline => ../../pipeline

replace go.opentelemetry.io/collector/consumer/consumererror => ../../consumer/consumererror

replace go.opentelemetry.io/collector/scraper => ../../scraper