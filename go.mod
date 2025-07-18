module github.com/tsi4456/gator

go 1.24.5

require internal/config v1.0.0

require internal/database v1.0.0

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
)

replace internal/config => ./internal/config

replace internal/database => ./internal/database
