# Deploying

On windows:

```
ampy -p COM3 -d 1 put .\main.py
```

# Debugging

```
# Replace COM3 with whatever serial port shows up when plugging in.
putty -serial COM3 -sercfg 115200,8,n,1,N
```