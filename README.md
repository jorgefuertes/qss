# qss
## Queru Simple Server

A simple http file server made in Go.

### Build
```bash
$> git clone git@github.com:jorgefuertes/qss.git
$> cd qss
$> go build
```

### Usage
```bash
$> ./qss
```

That starts the server listening at port 8000 and servig the actual dir. If you want to change it:

```bash
$> ./qss -p 8080 -r /home/queru/files
```

That's all folks!

