package mycalc

import (
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{
			name: "valid operation",
			args: args{args: []string{"10", "2"}},
			wantSum: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := Add(tt.args.args); gotSum != tt.wantSum {
				t.Errorf("Add() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name     string
		args     args
		wantDiff int
	}{
		{
			name: "valid operation",
			args: args{args: []string{"10", "2"}},
			wantDiff: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDiff := Subtract(tt.args.args); gotDiff != tt.wantDiff {
				t.Errorf("Subtract() = %v, want %v", gotDiff, tt.wantDiff)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name        string
		args        args
		wantProduct int
	}{
		{
			name: "valid operation",
			args: args{args: []string{"10", "2"}},
			wantProduct: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotProduct := Multiply(tt.args.args); gotProduct != tt.wantProduct {
				t.Errorf("Multiply() = %v, want %v", gotProduct, tt.wantProduct)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name         string
		args         args
		wantDividend int
	}{
		{
			name: "valid operation",
			args: args{args: []string{"10", "2"}},
			wantDividend: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDividend := Divide(tt.args.args); gotDividend != tt.wantDividend {
				t.Errorf("Divide() = %v, want %v", gotDividend, tt.wantDividend)
			}
		})
	}
}
