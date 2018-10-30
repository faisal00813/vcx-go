# GO Bindings for LibVCX

### Structure
- Folder `vcx` contains the go-bindings for LibVCX
- Folder `credex` contains wrapper for LibVCX
  

## Contibution guidelines.
### How to generate bindings on your system after modifying vcx.h header file.
- Install go 
- Install c-for-go by running `go get github.com/xlab/c-for-go`
- Goto folder vcx-go
- Run `make` command
- This will generate the go bindings