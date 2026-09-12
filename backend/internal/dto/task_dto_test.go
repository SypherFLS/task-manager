package dto

import (
	"testing"
)

func ptr[T any](v T) *T {
    return &v
}


func TestToMap(t *testing.T) {
	tests := []struct{
		Name string
		Dto UpdateTaskDTO
		Want bool
	}{	
		{
			Name : "usual",
			Dto : UpdateTaskDTO{
				Label: ptr("usual"),
				Description: ptr("something"),
				Priority: ptr(Low),
			},
			Want : true,
		},
		{
			Name : "Empty",
			Dto : UpdateTaskDTO{
				Label : ptr(""),
				Description: ptr(""),
				Priority: nil,
			},
			Want : true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T){
			data := tt.Dto.ToMap()
			_ = data
			
		})
	}
}