Packettik
=========
Really simple utility to check tcp connectivity at interval and print success/fail counts

to run:
```
go run packettik.go -d google.com -p 443 -i 1 -t 1 -l google_com_443.log
```

```
Help:
  -d string
        destination host <mandatory>
  -i int
        check interval (sec) <mandatory>
  -l string
        log to file [filename] <optional>
  -p int
        destination port number <mandatory>
  -t int
        session timeout (sec) <mandatory>
``` 