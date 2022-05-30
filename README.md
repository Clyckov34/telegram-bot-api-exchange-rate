# Telegram Bot "Exchange Rate Center Bank RU"
Telegram Bot - Курсы валют ЦБ банка РФ.

### Перечень валют:
- `AUD` - Австралийский доллар
- `AZN` - Азербайджанский манат
- `GBP` - Фунт стерлингов
- `AMD` - Армянский драм
- `BYN` - Белорусский рубль
- `BGN` - Болгарский Лев
- `BRL` - Бразильский реал
- `HUF` - Венгерский форинт
- `HKD` - Гонконгский доллар
- `DKK` - Датская крона
- `USD` - Доллар США
- `EUR` - Евро
- `INR` - Индийская рупия
- `KZT` - Казахстанский тенге
- `CAD` - Канадский доллар
- `KGS` - Киргизский сом
- `CNY` - Китайский юань
- `MDL` - Молдавский лей
- `NOK` - Норвежская крона
- `PLN` - Польский злотый
- `RON` - Румынский лей
- `XDR` - Международный Валютный Фонд
- `SGD` - Сингапурский доллар
- `TJS` - Таджикский сомони
- `TRY` - Турецкая лира
- `TMT` - Туркменский манат
- `UZS` - Узбекский сум
- `UAH` - Украинская гривна
- `CZK` - Чешская крона
- `SEK` - Шведская крона
- `CHF` - Швейцарский франк
- `ZAR` - Южноафриканский рэнд
- `KRW` - Южнокорейская вона
- `JPY` - Японская иена
- `RUB` - Российский рубль

### Запуск утилиты:
- `go run cmd/telegramBot/main.go -t=TOKEN_TELEGRAM_BOT`

### Запуск утилиты c подробными флагами:
- `go run cmd/telegramBot/main.go --help`

### DockerFile:
- Build - `docker build -t tg .`

- Run -   `docker run tg ./app -t=TOKEN_TELEGRAM_BOT`

### Linux: Автозапуск Telegram Bot systemctl
1. Скомпилируйте приложения ` go build cmd/telegramBot/main.go `
2. Разместите скомпилированный файл ` main ` в каталог ` /usr/local/bin/MY_PROJECT`
3. Откройте файл ` my_telegram.service ` Укажите путь к приложению, и токен Telegram Bot
4. Сохраните, и закройте файл
5. Разместите файл ` my_telegram.service ` в каталоге /etc/systemd/system/
6. Запустите скрипт как службу:
- ` sudo systemctl enable my_telegram `

- ` sudo systemctl start my_telegram `
