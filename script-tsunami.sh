#!/bin/bash

java -cp "/home/c-tsunami/tsunami/tsunami-main-0.0.15-SNAPSHOT-cli.jar:/home/c-tsunami/tsunami/plugins/*" com.google.tsunami.main.cli.TsunamiCli --ip-v4-target=127.0.0.1 --scan-results-local-output-format=JSON --scan-results-local-output-filename=/home/c-tsunami/tsunami/tsunami-output.json