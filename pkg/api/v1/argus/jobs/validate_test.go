package jobs

import "testing"

func TestJob_Validate(t *testing.T) {
	tests := []struct {
		name    string
		job     Job
		wantErr bool
	}{
		{"all ok", Job{
			StaticConfigs:  []StaticConfig{{Targets: []string{"abc"}}},
			JobName:        "my-job",
			Scheme:         "http",
			ScrapeInterval: "1m",
			ScrapeTimeout:  "5s",
			MetricsPath:    "/",
		}, false},
		{"no protocol", Job{
			StaticConfigs:  []StaticConfig{{Targets: []string{"abc"}}},
			JobName:        "my-job",
			Scheme:         "",
			ScrapeInterval: "1m",
			ScrapeTimeout:  "5s",
			MetricsPath:    "/",
		}, true},
		{"bad protocol", Job{
			StaticConfigs:  []StaticConfig{{Targets: []string{"abc"}}},
			JobName:        "my-job",
			Scheme:         "abcd",
			ScrapeInterval: "1m",
			ScrapeTimeout:  "5s",
			MetricsPath:    "/",
		}, true},
		{"no targets", Job{
			StaticConfigs:  []StaticConfig{{Targets: []string{}}},
			JobName:        "my-job",
			Scheme:         "http",
			ScrapeInterval: "1m",
			ScrapeTimeout:  "5s",
			MetricsPath:    "/",
		}, true},
		{"no configs", Job{
			StaticConfigs:  nil,
			JobName:        "my-job",
			Scheme:         "http",
			ScrapeInterval: "1m",
			ScrapeTimeout:  "5s",
			MetricsPath:    "/",
		}, true},
		{"bad name", Job{
			StaticConfigs:  []StaticConfig{{Targets: []string{"abc"}}},
			JobName:        "my_job",
			Scheme:         "http",
			ScrapeInterval: "1m",
			ScrapeTimeout:  "5s",
			MetricsPath:    "/",
		}, true},
		{"no name", Job{
			StaticConfigs:  []StaticConfig{{Targets: []string{"abc"}}},
			JobName:        "",
			Scheme:         "http",
			ScrapeInterval: "1m",
			ScrapeTimeout:  "5s",
			MetricsPath:    "/",
		}, true},
		{"bad interval 1", Job{
			StaticConfigs:  []StaticConfig{{Targets: []string{"abc"}}},
			JobName:        "my-job",
			Scheme:         "http",
			ScrapeInterval: "10s",
			ScrapeTimeout:  "5s",
			MetricsPath:    "/",
		}, true},
		{"bad interval 2", Job{
			StaticConfigs:  []StaticConfig{{Targets: []string{"abc"}}},
			JobName:        "my-job",
			Scheme:         "http",
			ScrapeInterval: "1m",
			ScrapeTimeout:  "2m",
			MetricsPath:    "/",
		}, true},
		{"bad interval 3", Job{
			StaticConfigs:  []StaticConfig{{Targets: []string{"abc"}}},
			JobName:        "my-job",
			Scheme:         "http",
			ScrapeInterval: "abcd",
			ScrapeTimeout:  "2m",
			MetricsPath:    "/",
		}, true},
		{"bad interval 4", Job{
			StaticConfigs:  []StaticConfig{{Targets: []string{"abc"}}},
			JobName:        "my-job",
			Scheme:         "http",
			ScrapeInterval: "1m",
			ScrapeTimeout:  "abcd",
			MetricsPath:    "/",
		}, true},
		{"bad interval 5", Job{
			StaticConfigs:  []StaticConfig{{Targets: []string{"abc"}}},
			JobName:        "my-job",
			Scheme:         "http",
			ScrapeInterval: "1m",
			ScrapeTimeout:  "",
			MetricsPath:    "/",
		}, true},
		{"bad interval 6", Job{
			StaticConfigs:  []StaticConfig{{Targets: []string{"abc"}}},
			JobName:        "my-job",
			Scheme:         "http",
			ScrapeInterval: "",
			ScrapeTimeout:  "",
			MetricsPath:    "/",
		}, true},
		{"no mp", Job{
			StaticConfigs:  []StaticConfig{{Targets: []string{"abc"}}},
			JobName:        "my-job",
			Scheme:         "http",
			ScrapeInterval: "1m",
			ScrapeTimeout:  "5s",
			MetricsPath:    "",
		}, true},
		{"too long mp", Job{
			StaticConfigs:  []StaticConfig{{Targets: []string{"abc"}}},
			JobName:        "my-job",
			Scheme:         "http",
			ScrapeInterval: "1m",
			ScrapeTimeout:  "5s",
			MetricsPath:    "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.job.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Job.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
