#!/bin/bash
docker build -t calculator-slim .
docker run -it calculator-slim
