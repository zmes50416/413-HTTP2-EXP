# 413-HTTP2-EXP

**Experiment on how HTTP client &amp; server handle of duplex communication.**

## TL;DR
Recently, I found OKHttp client will receive a HTTP/2 ResetStreamException from the Traefik ingress while a 413 (Entity Too Large) response 
is expected. I believed this is the issue from the OkHttp client implementation to the duplex feature of HTTP protocol(Not yet verified).

This repo will contain codes to demonstrate simplified scenarios.

## Background

Traefik is configured with [Buffering](https://doc.traefik.io/traefik/middlewares/http/buffering/) and `maxRequestBodyBytes` is set. so any request contains body larger than the limit will see be reject with 413 response.

This is tested with most browser working correctly. However, our another java process using HTTP/2 protocol communicate with the server. Instead of receiving response, it just receive a RST(RESET) frame with error code: NO ERROR.
