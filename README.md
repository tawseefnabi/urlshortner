# URL Shortener

Created url shortener using golang that will minimize url length and built us an optimize url.


##### 1) Clone project on your system and run command mentioned below.
```
go mod download
```

##### 2) Run project using command below.
```
 go run .
``` 

##### 3) Use below api to generate tiny url using our system

###### Request-Type: POST
```
  http://localhost:8080/generateTinyUrl/
```

###### Request-Body
```
{
  "url":"https://www.google.com/"
}
```

###### Response-Body
```
{
  "url": "localhost:8080/ulnYa"
}
```

