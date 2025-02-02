# alert-from-db
Tool for get metric from db and visualise  threshold for this metric.


* Generate go files form topo file

from root project directory run this command

```
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  proto/backend/backend.proto
```