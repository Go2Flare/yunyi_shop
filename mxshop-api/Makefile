default: build-goods-web build-order-web build-user-web build-oss-web build-userop-web run

build-goods-web:
	go build -o ./goods-web ./goods-web/.


build-order-web:
	go build -o ./order-web ./order-web/.


build-user-web:
	go build -o ./user-web ./user-web/.


build-oss-web:
	go build -o ./oss-web ./oss-web/.


build-userop-web:
	go build -o ./userop-web ./userop-web/.

run:
	nohup ./goods-web/goods-web &
	nohup ./order-web/order-web &
	nohup ./oss-web/oss-web &
	nohup ./user-web/user-web &
	nohup ./userop-web/userop-web &