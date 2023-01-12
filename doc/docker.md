# consul service

docker run -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600/udp consul
consul agent -dev -client=0.0.0.0

browser http://127.0.0.1:8500/ui/dc1/services