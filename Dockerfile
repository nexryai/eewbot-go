FROM ubuntu:latest as builder
WORKDIR /build

COPY . ./

RUN apt update -y && apt install -y golang ca-certificates \
 && go build -ldflags="-s -w" -trimpath -o eewbot main.go

FROM ubuntu:latest
WORKDIR /app


RUN sed -i.bak -r 's!(deb|deb-src) \S+!\1 http://ftp.riken.jp/Linux/ubuntu/!' /etc/apt/sources.list\
 && apt update \
 && apt install -y curl tini xvfb graphicsmagick-imagemagick-compat wget unzip ca-certificates \
 && groupadd -g 987 app \
 && useradd -d /app -s /bin/sh -u 987 -g app app \
 && wget https://github.com/ingen084/KyoshinEewViewerIngen/releases/latest/download/KyoshinEewViewer-ubuntu-latest.zip \
 && unzip KyoshinEewViewer-ubuntu-latest.zip \
 && rm -f KyoshinEewViewer-ubuntu-latest.zip \
 && apt purge -y wget unzip \
 && apt autoremove --purge -y \
 && apt clean \
 && mkdir "hooks" \
 && chown -R app:app /app

COPY --from=builder /build/eewbot /app/hooks/eewbot
COPY --chown=app:app . .

RUN chmod +x /app/hooks/*

USER app
CMD ["tini", "--", "xvfb-run", "-s", "-ac -screen 0 1280x800x24", "./KyoshinEewViewer"]
