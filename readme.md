# Email Sender CLI 工具

中文 | [English](https://github.com/jerry12122/mail-sender/Readme_en.md)

這是一個使用 SMTP 發送電子郵件的命令行工具。它支持從配置文件中讀取 SMTP 設置。

## 使用方法

1. 從 [Releases 頁面](https://github.com/jerry12122/mail-sender/releases)下載適用於您作業系統的最新版本。
2. 查看幫助:
   ```
   ./mail-sender --help
   ```
3. 使用配置文件指定 SMTP 設置:
   
   `config.yaml` 文件的內容應該是:
   ```yaml
   smtp:
     host: smtp.example.com
     port: 587
     user: your-smtp-username
     password: your-smtp-password
     use_starttls: true
   ```
4. 使用以下命令發送電子郵件:
   ```
   ./email-sender --config config.yaml --from sender@example.com --to recipient@example.com --subject "測試郵件" --body "這是一封測試郵件。"
   ```



## 從原始碼構建

1. 安裝 Go Runtime。
2. Clone Repository 並 cd 到該目錄。
3. 構建程序:
   ```
   go build -o mail-sender
   ```

4. 編譯後的二進位制文件 `mail-sender` 將創建在當前目錄中。您可以使用上述"使用方法"部分中的相同命令運行它。

## 命令行參數

| 參數 | 縮寫 | 描述 |
| --- | --- | --- |
| `--config` | `-c` | 配置文件路徑 |
| `--host` | `-H` | SMTP 主機地址 |
| `--port` | `-P` | SMTP 埠號 |
| `--user` | `-u` | SMTP 使用者名稱 |
| `--password` | `-p` | SMTP 密碼 |
| `--from` | `-f` | 發件人電子郵件地址 |
| `--to` | `-t` | 收件人電子郵件地址 |
| `--subject` | `-s` | 電子郵件主題 |
| `--body` | `-b` | 電子郵件正文 |
| `--use-starttls` | - | 使用 STARTTLS 加密 |