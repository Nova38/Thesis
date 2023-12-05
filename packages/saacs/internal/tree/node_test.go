package tree

// type Holder struct {
// 	name  string
// 	valid bool
// }

// func TestNode_AddPath(t *testing.T) {

//     type args[T any] struct {
//         path  string
//         value *T
//     }
//     type testCase[T any] struct {
//         name    string
//         n       Node[T]
//         args    args[T]
//         wantState bool
//         wantErr bool
//     }
//     tests := []testCase[ Holder ]{
//         // TODO: Add test cases.
//         {
//         	name:      "",
//         	n:         NewTree[Holder](),
//         	args:      args{},
//         	wantState: false,
//         	wantErr:   false,
//         },
//     }
//     for _, tt := range tests {

//         t.Run(tt.name, func(t *testing.T) {
//             if err := tt.n.AddPath(tt.args.path, tt.args.value); (err != nil) != tt.wantErr {
//                 t.Errorf("AddPath() error = %v, wantErr %v", err, tt.wantErr)
//             }
//         })
//     }
// }

// func TestNode_GetPath(t *testing.T) {
//     type args struct {
//         path string
//     }
//     type testCase[T any] struct {
//         name     string
//         n        Node[T]
//         args     args
//         wantNode *Node[T]
//     }
//     tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//         // TODO: Add test cases.
//     }
//     for _, tt := range tests {
//         t.Run(tt.name, func(t *testing.T) {
//             if gotNode := tt.n.GetPath(tt.args.path); !reflect.DeepEqual(gotNode, tt.wantNode) {
//                 t.Errorf("GetPath() = %v, want %v", gotNode, tt.wantNode)
//             }
//         })
//     }
// }

// func TestNode_Merge(t *testing.T) {
//     type args[T any] struct {
//         trees []*Node[T]
//     }
//     type testCase[T any] struct {
//         name string
//         n    Node[T]
//         args args[T]
//     }
//     tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//         // TODO: Add test cases.
//     }
//     for _, tt := range tests {
//         t.Run(tt.name, func(t *testing.T) {
//             tt.n.Merge(tt.args.trees...)
//         })
//     }
// }

// func TestNode_Visit(t *testing.T) {
//     type args[T any] struct {
//         options VisitorOptions
//         visitFn func(node *Node[T])
//     }
//     type testCase[T any] struct {
//         name string
//         n    Node[T]
//         args args[T]
//     }
//     tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//         // TODO: Add test cases.
//     }
//     for _, tt := range tests {
//         t.Run(tt.name, func(t *testing.T) {
//             tt.n.Visit(tt.args.options, tt.args.visitFn)
//         })
//     }
// }

// func TestNode_VisitPaths(t *testing.T) {
//     type args[T any] struct {
//         paths   []string
//         resolve bool
//         visitor func(node *Node[T])
//     }
//     type testCase[T any] struct {
//         name string
//         n    Node[T]
//         args args[T]
//     }
//     tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//         // TODO: Add test cases.
//     }
//     for _, tt := range tests {
//         t.Run(tt.name, func(t *testing.T) {
//             tt.n.VisitPaths(tt.args.paths, tt.args.resolve, tt.args.visitor)
//         })
//     }
// }

// func TestNode_addPath(t *testing.T) {
//     type args[T any] struct {
//         split []string
//         value *T
//     }
//     type testCase[T any] struct {
//         name    string
//         n       Node[T]
//         args    args[T]
//         wantErr bool
//     }
//     tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//         // TODO: Add test cases.
//     }
//     for _, tt := range tests {
//         t.Run(tt.name, func(t *testing.T) {
//             if err := tt.n.addPath(tt.args.split, tt.args.value); (err != nil) != tt.wantErr {
//                 t.Errorf("addPath() error = %v, wantErr %v", err, tt.wantErr)
//             }
//         })
//     }
// }

// func TestNode_error(t *testing.T) {
//     type testCase[T any] struct {
//         name    string
//         n       Node[T]
//         wantErr oops.OopsErrorBuilder
//     }
//     tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//         // TODO: Add test cases.
//     }
//     for _, tt := range tests {
//         t.Run(tt.name, func(t *testing.T) {
//             if gotErr := tt.n.error(); !reflect.DeepEqual(gotErr, tt.wantErr) {
//                 t.Errorf("error() = %v, want %v", gotErr, tt.wantErr)
//             }
//         })
//     }
// }

// func TestNode_getPath(t *testing.T) {
//     type args[T any] struct {
//         p     *Node[T]
//         split []string
//     }
//     type testCase[T any] struct {
//         name     string
//         n        Node[T]
//         args     args[T]
//         wantNode *Node[T]
//     }
//     tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//         // TODO: Add test cases.
//     }
//     for _, tt := range tests {
//         t.Run(tt.name, func(t *testing.T) {
//             if gotNode := tt.n.getPath(tt.args.p, tt.args.split); !reflect.DeepEqual(gotNode, tt.wantNode) {
//                 t.Errorf("getPath() = %v, want %v", gotNode, tt.wantNode)
//             }
//         })
//     }
// }

// func TestNode_merge(t *testing.T) {
//     type args[T any] struct {
//         tree *Node[T]
//     }
//     type testCase[T any] struct {
//         name string
//         n    Node[T]
//         args args[T]
//     }
//     tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//         // TODO: Add test cases.
//     }
//     for _, tt := range tests {
//         t.Run(tt.name, func(t *testing.T) {
//             tt.n.merge(tt.args.tree)
//         })
//     }
// }

// func TestNode_visit(t *testing.T) {
//     type args[T any] struct {
//         options VisitorOptions
//         visitFn func(node *Node[T])
//     }
//     type testCase[T any] struct {
//         name string
//         n    Node[T]
//         args args[T]
//     }
//     tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//         // TODO: Add test cases.
//     }
//     for _, tt := range tests {
//         t.Run(tt.name, func(t *testing.T) {
//             tt.n.visit(tt.args.options, tt.args.visitFn)
//         })
//     }
// }
