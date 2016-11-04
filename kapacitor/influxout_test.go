package kapacitor

import "testing"
import "github.com/influxdata/chronograf"

func TestInfluxOut(t *testing.T) {
	tests := []struct {
		name string
		want chronograf.TICKScript
	}{
		{
			name: "Test influxDBOut kapacitor node",
			want: `trigger
    |eval(lambda: field)
        .as('value')
    |influxDBOut()
        .create()
        .database(output_db)
        .retentionPolicy(output_rp)
        .measurement(output_mt)
        .tag('alertName', name)
        .tag('triggerType', triggerType)
`,
		},
	}
	for _, tt := range tests {
		got := InfluxOut(chronograf.AlertRule{
			Name:    "name",
			Trigger: "deadman",
		})
		formatted, err := formatTick(got)
		if err != nil {
			t.Errorf("%q. formatTick() error = %v", tt.name, err)
			continue
		}
		if formatted != tt.want {
			t.Errorf("%q. InfluxOut() = %v, want %v", tt.name, formatted, tt.want)
		}
	}
}
