build: install
	buf generate 
 

install:
	go install .\tools\protoc-gen-cc-state
	go install .\tools\protoc-gen-cckey
	
