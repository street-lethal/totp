package src

import (
	"testing"
	"time"
)

func TestTOTP_Calc(t1 *testing.T) {
	type fields struct {
		Time time.Time
	}
	type args struct {
		secret string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "1",
			fields: fields{
				Time: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
			},
			args: args{"AAAZ===="},
			want: "281526",
		},
		{
			name: "2",
			fields: fields{
				Time: time.Date(2030, time.December, 30, 15, 45, 39, 0, time.UTC),
			},
			args: args{"3HW6QABXP5DSLRVYKTOMIG2N4ZJFC7UE3MS57EUFCPT6ABZWOV2NQDLJGK4XYIHR"},
			want: "861638",
		},
		{
			name: "3",
			fields: fields{
				Time: time.Date(2040, time.July, 15, 14, 45, 59, 0, time.UTC),
			},
			args: args{"PSEJDQLVKZ4RH2FMU7XOW56GCIANBT3YKTIWO7M3ZH6LEPFQNXUR2ASJYVD5B4GCIKJW4FYPMGCSV3572HZXDOBUENTA6QLR4QKM3HZVEICYU2SBO7DJ6GNR5AXLPWFTPRMJULD32HKICS4O6TXQW7EN5YZABVGF"},
			want: "956780",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := TOTP{
				Time:     tt.fields.Time,
				digits:   6,
				timeStep: 30,
			}

			secret := Secret(tt.args.secret)
			bin, err := secret.Base32()
			if err != nil {
				t1.Error(err)
				return
			}

			if got := t.Calc(bin); got != tt.want {
				t1.Errorf("Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
