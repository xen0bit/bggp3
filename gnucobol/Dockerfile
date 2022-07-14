FROM ubuntu:rolling

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get -y update && apt-get install -y \
    gcc \
    git \
    make \
    wget \
    libdb-dev \
    libgmp-dev \
&& rm -rf /var/lib/apt/lists/*

RUN git clone --depth=1 https://github.com/xen0bit/bggp3.git

WORKDIR /bggp3/gnucobol/gnucobol-3.1.2

RUN ./configure
RUN make -j$(nproc)
RUN make install
RUN ldconfig

WORKDIR /bggp3/gnucobol/

RUN cobc -V
RUN cobc crash.on.the.cob