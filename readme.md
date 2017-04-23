# TheQuestion Speaker

Daemon that reads all new questions from [TheQuestion](https://thequestion.ru)

## Usage
1. Get project 
```bash
go get -u https://github.com/ngalayko/theq_speak
```
2. Install `mpg123` via `brew` or `apt`
3. Build it
```bash
go build main.go 
```
4. Run it
```bash
chmod +x ./main
./main -key=yandex-api-key
```
You can get Yandex SpeechKit key [here](https://developer.tech.yandex.ru)