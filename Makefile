build: install
	buf generate 
 

install:
	go install .\tools\cmd\protoc-gen-cckey
	
