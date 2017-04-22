# TheQuestion Speaker

Daemon that reads all new questions from [TheQuestion](https://thequestion.ru)

## Usage
1. Get project 
```bash
go get -u https://github.com/ngalayko/theq_speak
```
2. Add Yandex SpeechKei API key to [config.yaml](config.yaml). You can get it [here](https://developer.tech.yandex.ru)
3. Build it
```bash
go build main.go
```
4. Run it
```bash
chmod +x ./main
./main
```