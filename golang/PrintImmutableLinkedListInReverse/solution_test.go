package PrintImmutableLinkedListInReverse

import "testing"

func Test_printLinkedListInReverse(t *testing.T) {
	type args struct {
		head ImmutableListNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "default test",
			args: args{
				head: &immutableListNode{
					value: 1,
					next: &immutableListNode{
						value: 2,
						next: &immutableListNode{
							value: 3,
							next: &immutableListNode{
								value: 4,
								next:  nil,
							},
						},
					},
				},
			},
		},
		{
			name: "default test",
			args: args{
				head: &immutableListNode{
					value: 0,
					next: &immutableListNode{
						value: -4,
						next: &immutableListNode{
							value: -1,
							next: &immutableListNode{
								value: 3,
								next: &immutableListNode{
									value: -5,
									next:  &immutableListNode{},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "default test",
			args: args{
				head: &immutableListNode{
					value: -2,
					next: &immutableListNode{
						value: 0,
						next: &immutableListNode{
							value: 6,
							next: &immutableListNode{
								value: 4,
								next: &immutableListNode{
									value: 4,
									next: &immutableListNode{
										value: -6,
										next:  &immutableListNode{},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printLinkedListInReverse(tt.args.head)
		})
	}
}
