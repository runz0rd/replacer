package replacer

import (
	"testing"
)

func TestReplace(t *testing.T) {
	type args struct {
		inputFile  string
		outputFile string
		from       string
		to         string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"phoney", args{"test-input.json", "test-output.json", "dev", "prod"}, false},
	}
	c, _ := LoadConfig("test.yaml")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Replace(tt.args.inputFile, tt.args.outputFile, tt.args.from, tt.args.to, c.Rules); (err != nil) != tt.wantErr {
				t.Errorf("Replace() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
