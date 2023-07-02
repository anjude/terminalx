package sizhi

import (
	"reflect"
	"testing"
)

func TestGetSizhiMsg(t *testing.T) {
	type args struct {
		content string
		userId  string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 error
	}{
		{name: "d", args: args{
			content: "你是谁", userId: "",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetSizhiMsg(tt.args.content, tt.args.userId)
			if got != tt.want {
				t.Errorf("GetSizhiMsg() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetSizhiMsg() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
