# Email Sender CLI Tool

[中文](https://github.com/jerry12122/mail-sender) | English

This is a command-line tool for sending emails using SMTP. It supports reading SMTP settings from a configuration file.

## Usage

1. Download the latest release for your operating system from the [Releases page](https://github.com/jerry12122/mail-sender/releases).
2. Extract the downloaded archive and navigate to the extracted directory.
3. Run the program using the following command:
   ```
   ./mail-sender --help
   ```
4. You use a configuration file to specify the SMTP settings:
   
   The contents of the `config.yaml` file should be:
   ```yaml
   smtp:
     host: smtp.example.com
     port: 587
     user: your-smtp-username
     password: your-smtp-password
     use_starttls: true
   ```
5. Use the following command to send an email:
   ```
   ./mail-sender --config config.yaml --from sender@example.com --to recipient@example.com --subject "Test Email" --body "This is a test email."
   ```



## Building from Source

1. Install the Go runtime.
2. Clone this repository and navigate to the project directory.
3. Build the program:
   ```
   go build -o mail-sender
   ```

4. The compiled binary `mail-sender` will be created in the current directory. You can run it using the same commands as in the Usage section above.

## Command-line Arguments

| Argument | Short Form | Description |
| --- | --- | --- |
| `--config` | `-c` | Path to the configuration file |
| `--host` | `-H` | SMTP host address |
| `--port` | `-P` | SMTP port number |
| `--user` | `-u` | SMTP username |
| `--password` | `-p` | SMTP password |
| `--from` | `-f` | Sender's email address |
| `--to` | `-t` | Recipient's email address |
| `--subject` | `-s` | Email subject |
| `--body` | `-b` | Email body |
| `--use-starttls` | - | Use STARTTLS encryption |