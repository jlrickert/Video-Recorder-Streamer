#!/usr/bin/sh

openssl req -x509 -newkey rsa:4096 -keyout video_streamer.key -out video_streamer.crt -nodes -days 7
