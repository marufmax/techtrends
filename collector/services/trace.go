package techtrends

import "github.com/uptrace/uptrace-go/uptrace"

uptrace.ConfigureOpentelemetry(
    uptrace.WithDSN("https://8_sBTg3xwW9sa2HbEo9sMg@uptrace.dev/332"),

    uptrace.WithServiceName("collector"),
    uptrace.WithServiceVersion("v1.0.0"),
)