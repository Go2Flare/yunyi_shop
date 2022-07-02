default: build_goods_srv build_order_srv build_user_srv build_inventory_srv build_userop_srv run

build_goods_srv:
	go build -o ./goods_srv ./goods_srv/.


build_order_srv:
	go build -o ./order_srv ./order_srv/.


build_user_srv:
	go build -o ./user_srv ./user_srv/.


build_inventory_srv:
	go build -o ./inventory_srv ./inventory_srv/.


build_userop_srv:
	go build -o ./userop_srv ./userop_srv/.

run:
	nohup ./goods_srv/goods_srv &
	nohup ./order_srv/order_srv &
	nohup ./inventory_srv/inventory_srv &
	nohup ./user_srv/user_srv &
	nohup ./userop_srv/userop_srv &