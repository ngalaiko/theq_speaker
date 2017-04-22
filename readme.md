# TheQuestion Speaker

Daemon that reads all new questions from [TheQuestion](https://thequestion.ru)

## Usage
1. Get project 
```bash
go get -u https://github.com/ngalayko/theq_speak
```
2. Install `mpg123` via `brew` or `apt`
3. Add Yandex SpeechKei API key to [config.yaml](config.yaml). You can get it [here](https://developer.tech.yandex.ru)
4. Build it
```bash
go build main.go
```
5. Run it
```bash
chmod +x ./main
./main
```