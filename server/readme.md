# TheQuestion Speaker

Daemon that reads all new questions from [TheQuestion](https://thequestion.ru)

## Usage
1. Get project 
```bash
go get -u https://github.com/ngalayko/theq_speak
```
2. Install dependencies via [dep](https://github.com/golang/dep)
```bash
dep ensure
```
3. Build it
```bash
go build main.go 
```
4. Make config file
```bash
cp config_example.yaml config.yaml
```
You can get Yandex SpeechKit key [here](https://developer.tech.yandex.ru)
5. Run it
```bash
chmod +x ./main
./main 
```