FROM golang:1.11

LABEL MAINTAINER="xieWei(wuxiaoshen@shu.edu.cn)"

WORKDIR /home/src

RUN echo "Hello World"

COPY . .

CMD ["bash","-c", "exec.sh"]