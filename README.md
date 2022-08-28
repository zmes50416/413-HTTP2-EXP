# 413-HTTP2-EXP

**Experiment on how HTTP client &amp; server handle of duplex communication.**

## TL;DR
Recently, I found OKHttp client will receive a HTTP/2 ResetStreamException from Traefik ingress while a 413 (Entity Too Large) response 
is expected. I belived this is the issue from OKHTTP client misimplement the duplex feature of HTTP procotal(Not yet verified)
This repo will contain codes to demostrate simplified sceneraio.

## Background

Traefik is configured with [Buffering](https://doc.traefik.io/traefik/middlewares/http/buffering/) and `maxRequestBodyBytes` is set. so any request contains body larger than the limit will see be reject with 413 response.

This is tested with most browser working correctly. However, our another java process using HTTP/2 protocal communicate with the server. Instead of receving response, it just receive a RST(RESET) frame with error code: NO ERROR.
