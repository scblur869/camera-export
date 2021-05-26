
# what does this do?
- this utility exports the users from the temper-sure camera / scanner
- creates a folder called 'exported-photos'
- creates a CSV with the UUID of the device in the name of the CSV
- goal of this is to allow the user to import / restore this user data back to the device(s)

###  To Run

## from the src root run 
```bash
go build
```

## next run the application
```bash
./camera-export -deviceip=192.168.1.47 -user=admin -pass=admin
```