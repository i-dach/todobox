package domain

import (
	"reflect"
	"testing"
)

func TestNewTask(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Task
	}{
		{
			name: "Create Sample task",
			args: args{name: "Sample"},
			want: &Task{
				Name:  "Sample",
				state: Backlog,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTask(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_GetState(t *testing.T) {
	type fields struct {
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		want   TaskState
	}{
		{
			name: "Get state Backlog",
			fields: fields{
				Name: "Sample",
			},
			want: Backlog,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := NewTask(tt.fields.Name)
			if got := tr.GetState(); got != tt.want {
				t.Errorf("Task.GetState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_NextState(t *testing.T) {
	type fields struct {
		Name string
	}
	type args struct {
		next TaskState
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   TaskState
	}{
		{
			name: "Move state Backlog to Done",
			fields: fields{
				Name: "Sample",
			},
			args: args{Done},
			want: Done,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := NewTask(tt.fields.Name)
			tr.NextState(tt.args.next)
			if got := tr.GetState(); got != tt.want {
				t.Errorf("Task.NextState() = %v, want %v", got, tt.want)
			}
		})
	}
}
