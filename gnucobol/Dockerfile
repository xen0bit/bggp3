FROM ubuntu:rolling

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get -y update && apt-get install -y \
    gcc \
    git \
    make \
    pkg-config \
	libipt-dev \
	libunwind8-dev \
	binutils-dev \
    wget \
    libdb-dev \
    libgmp-dev \
&& rm -rf /var/lib/apt/lists/*

RUN git clone --depth=1 https://github.com/xen0bit/bggp3.git

WORKDIR /bggp3/gnucobol

RUN ./configure \
    make -j$(nproc) \
    make install \
    ldconfig \
    cobc -V \
    cobc crash.on.the.cob

