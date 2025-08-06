# clairvoy

**Clairvoy** — безопасный менеджер секретов в терминале с поддержкой множественных пользователей и категориями (`account`, `api_key`, `ssh`, `gpg`, `note`).

## Особенности

- Хранит данные в `~/.vault/<username>/…` с отдельной солью для каждого пользователя
- Категории: пароли, API-ключи, SSH-ключи, GPG-ключи и произвольные заметки
- Мастер-пароль для шифрования/дешифрования через Argon2id + AES-GCM
- Команды:
    - `vault-cli register <username>` — инициализация хранилища для пользователя
    - `vault-cli add` — добавить секрет
    - `vault-cli list` — список всех метаданных
    - `vault-cli get <id|label>` — вывести и расшифровать секрет
    - `vault-cli remove <id|label>` — удалить секрет
- Автодополнение для Bash через `vault-cli completion bash`

## Установка

1. Склонируйте и перейдите в репозиторий:
   ```bash
   git clone https://github.com/iconfire7/clairvoy.git
   cd clairvoy

2. Соберите и установите бинарь:

```bash
go install github.com/iconfire7/Clairvoy/cmd/clairvoy@latest
source ~/.bashrc
```

# Быстрый старт

## Зарегистрировать пользователя user
```bash
clairvoy register $user
```
## Добавить пароль
```bash
clairvoy add

→ Type: account
→ Label: github
→ Login: user
→ Password: ****
```
## Вывести список
```bash
clairvoy list
```

## Расшифровать секрет
```bash
clairvoy get github
```

## Удалить секрет
```bash
clairvoy remove github
```