package digibyte_addrdec

import "testing"

func TestCheck(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "valid", args: args{address: "D8qHfnugKAgavULzVjQKyjqxMD7wBETN4s"}, want: true, wantErr: false},
		{name: "valid", args: args{address: "DKzQUa2z9G1pzhhUcGzR2qJUxVzMo4vgeq"}, want: true, wantErr: false},
		{name: "invalid", args: args{address: "BAo2HCZgvPTmFxFANdhznfHAX5Pjpg4Kg6"}, want: false, wantErr: true},
		{name: "invalid", args: args{address: "DAo2HsZgvPTmFxFANdhznfHAX5Pjpg4Kgs"}, want: false, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Check(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
}
