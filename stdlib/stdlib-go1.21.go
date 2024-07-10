//go:build go1.21

package stdlib

//go:generate ../internal/cmd/extract/extract cmp
//go:generate ../internal/cmd/extract/extract maps
//go:generate ../internal/cmd/extract/extract slices
//go:generate ../internal/cmd/extract/extract log/syslog log/slog
//go:generate ../internal/cmd/extract/extract testing/slogtest
