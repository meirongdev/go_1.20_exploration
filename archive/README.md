# unsafe archive

## TAR

1. Create unsecure tar file: `tar -cvf  example.tar example ../readme.md`

   - If you see `tar: Removing leading '../' from member names../readme.md`, run `tar cvfP example.tar example ../readme.md`.

2. Compile binary: cd tar && go build .
3. Run binary with and without GODEBUG=tarinsecurepath=0

   - `cd tar && go run main.go`

## ZIP

1. Create unsecure tar file: `zip -v -r example.zip example/ ../readme.md`
2. Compile binary: cd zip && go build .
3. Run binary with and without GODEBUG=zipinsecurepath=0

   - `./zip`
   - `GODEBUG=zipinsecurepath=0 ./zip` -> 2023/11/30 10:44:15 NewReader failed zip: insecure file path
