#!/bin/bash

# Caminho para o arquivo de log
LOG_FILE="./app_logs"

# Observa mudanças no arquivo de log e printa a última linha escrita
tail -F "$LOG_FILE" | while read -r line; do
    echo "$line"
done