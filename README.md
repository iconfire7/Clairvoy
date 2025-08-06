# vault-cli

**Vault CLI** — безопасный менеджер секретов в терминале с поддержкой множественных пользователей и категориями (`account`, `api_key`, `ssh`, `gpg`, `note`).

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
   git clone https://github.com/yourusername/vault-cli.git
   cd vault-cli

2. Соберите и установите бинарь:

```bash
go install github.com/iconfire7/Clairvoy/cmd/clairvoy@latest
source ~/.bashrc
```

# Быстрый старт

## Зарегистрировать пользователя user
```bash
vault-cli register $user
```
## Добавить пароль
```bash
vault-cli add

→ Type: account
→ Label: github
→ Login: user
→ Password: ****
```
## Вывести список
```bash
vault-cli list
```

## Расшифровать секрет
```bash
vault-cli get github
```

## Удалить секрет
```bash
vault-cli remove github
```